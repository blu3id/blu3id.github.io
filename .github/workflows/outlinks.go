package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	LogInfo  *log.Logger
	LogError *log.Logger
	LogDebug *log.Logger
)

func init() {
	var debug string

	LogInfo = log.New(os.Stderr, "[Info]  ", log.LstdFlags|log.Lmsgprefix)
	LogError = log.New(os.Stderr, "[Error] ", log.LstdFlags|log.Lmsgprefix)
	LogDebug = log.New(io.Discard, "[Debug] ", log.LstdFlags|log.Lmsgprefix)

	debug = os.Getenv("DEBUG")
	debug = strings.ToLower(debug)
	if debug == "true" || debug == "1" || debug == "y" {
		LogDebug.SetOutput(os.Stderr)
	}
}

// spnAPI models the SPN2 JSON API response
type spnAPI struct {
	Status      string   `json:"status"`
	JobId       string   `json:"job_id"`
	OriginalUrl string   `json:"original_url"`
	Screenshot  string   `json:"screenshot"`
	Timestamp   string   `json:"timestamp"`
	Duration    float32  `json:"duration_sec"`
	HTTPStatus  int      `json:"http_status"`
	Resources   []string `json:"resources"`
	Outlinks    []string `json:"outlinks"`

	Exception string `json:"exception"`
	StatusExt string `json:"status_ext"`
	Message   string `json:"message"`
}

// spnStatus models the SPN2 User Status JSON API response
type spnStatus struct {
	Available          int `json:"available"`
	DailyCaptures      int `json:"daily_captures"`
	DailyCapturesLimit int `json:"daily_captures_limit"`
	Processing         int `json:"processing"`
}

// waybackAPI models the Archive.org Wayback Machine JSON API response
// (https://archive.org/help/wayback_api.php)
type waybackAPI struct {
	Url               string                      `json:"url"`
	ArchivedSnapshots waybackAPIArchivedSnapshots `json:"archived_snapshots"`
}

// waybackAPIArchivedSnapshots models a nested type of waybackAPI
type waybackAPIArchivedSnapshots struct {
	Closest waybackAPIClosest `json:"closest"`
}

// waybackAPIClosest models a nested type of waybackAPIArchivedSnapshots
type waybackAPIClosest struct {
	Available bool   `json:"available"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Url       string `json:"url"`
}

// sendRequest is a helper function that wraps HTTP requests and adds the SPN2
// authorization header from the environment variable SPN2_S3_KEY if setAuth is
// true
func sendRequest(requestMethod string, requestUrl string,
	requestBody io.Reader, setAuth bool) (responseBytes []byte, err error) {
	var (
		archiveClient http.Client
		clientRequest *http.Request
		response      *http.Response
		s3Key         string
	)

	clientRequest, err = http.NewRequest(requestMethod,
		requestUrl, requestBody)
	if err != nil {
		return responseBytes, err
	}

	clientRequest.Header.Set("User-Agent", "outlink-archiver-spn2-v0.1.0")
	clientRequest.Header.Set("Accept", "application/json")
	if requestBody != nil {
		clientRequest.Header.Set("Content-Type",
			"application/x-www-form-urlencoded")
	}
	if setAuth {
		s3Key = os.Getenv("SPN2_S3_KEY")
		if s3Key == "" {
			return responseBytes,
				fmt.Errorf("SPN2_S3_KEY not found in environment")
		}
		clientRequest.Header.Set("authorization", fmt.Sprintf("LOW %s", s3Key))
	}

	archiveClient = http.Client{
		Timeout: time.Second * 20,
	}
	response, err = archiveClient.Do(clientRequest)
	if err != nil {
		return responseBytes, err
	}

	if response.StatusCode != http.StatusOK {
		return responseBytes, fmt.Errorf("HTTP Response Status: %d",
			response.StatusCode)
	}

	if response.Body != nil {
		defer response.Body.Close()
	}

	responseBytes, err = io.ReadAll(response.Body)
	if err != nil {
		return responseBytes, err
	}

	return responseBytes, nil
}

// requestSPNCapture submits a URL to the archive.org SPN2 API
// (https://docs.google.com/document/d/1Nsv52MvSjbLb2PCpHlat0gkzw0EvtSgpKHu4mk0MnrA/edit)
// and returns the jobId. Throws an error if the SPN2 API doesn't return a JobId
// (which will happen if the URL has had a snapshot in the last 7d)
func requestSPNCapture(pageUrl string) (jobId string, err error) {
	var (
		responseBytes []byte
		postBody      = url.Values{}
		apiResponse   spnAPI
	)

	postBody.Add("url", pageUrl)
	postBody.Add("skip_first_archive", "1")
	postBody.Add("if_not_archived_within", "7d")
	//postBody.Add("capture_outlinks", "1")

	responseBytes, err = sendRequest(http.MethodPost,
		"https://web.archive.org/save", strings.NewReader(postBody.Encode()),
		true)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(responseBytes, &apiResponse)
	if err != nil {
		return "", err
	}

	if apiResponse.JobId == "" {
		return "", fmt.Errorf("SPN2 Error: %s [%s]", apiResponse.Message,
			apiResponse.StatusExt)
	}

	LogDebug.Printf("requestSPNCapture: %v\n", apiResponse)
	return apiResponse.JobId, nil
}

// checkSPNJobStatus returns the snapshot timestamp if the status of the SPN2
// job status is "success" otherwise returns an empty string if the job status
// is "pending". If the status is "error" or unknown an "SPN2 Error..." error is
// thrown including the returned message and status_ext code
func checkSPNJobStatus(jobId string) (timestamp string, err error) {
	var (
		requestUrl    string
		responseBytes []byte
		apiResponse   spnAPI
	)

	requestUrl = fmt.Sprintf("https://web.archive.org/save/status/%s", jobId)
	responseBytes, err = sendRequest(http.MethodGet, requestUrl, nil, true)
	if err != nil {
		return "", err
	}

	apiResponse = spnAPI{}
	err = json.Unmarshal(responseBytes, &apiResponse)
	if err != nil {
		return "", err
	}

	if apiResponse.Status == "error" {
		return "", fmt.Errorf("SPN2 Error: %s [%s]", apiResponse.Message,
			apiResponse.StatusExt)
	}

	if apiResponse.Status == "pending" {
		LogDebug.Printf("checkSPNJobStatus: %v\n", apiResponse)
		return "", nil
	}

	if apiResponse.Status == "success" {
		LogDebug.Printf("checkSPNJobStatus: %v\n", apiResponse)
		return apiResponse.Timestamp, nil
	}

	return "", fmt.Errorf(
		"Unknown SPN Error (status: %s, status_ext: %s, message: %s)",
		apiResponse.Status, apiResponse.StatusExt, apiResponse.Message)
}

// checkSPNUserStatus retuns the User Status for the SPN2 API. This shows
// remaining dailiy captures and currently available parallel jobs.
func checkSPNUserStatus() (apiResponse spnStatus, err error) {
	var (
		responseBytes []byte
	)

	responseBytes, err = sendRequest(http.MethodGet,
		"https://web.archive.org/save/status/user", nil, true)
	if err != nil {
		return apiResponse, err
	}

	apiResponse = spnStatus{}
	err = json.Unmarshal(responseBytes, &apiResponse)
	if err != nil {
		return apiResponse, err
	}

	LogDebug.Printf("checkSPNUserStatus: %v\n", apiResponse)
	return apiResponse, nil
}

// SPNCapture wraps the archive.org SPN2 API and synchronously captures the
// supplied pageURL and retuns the URL to the archive.org snapshot. Will timout
// after 5min and throw and error.
func SPNCapture(pageUrl string) (snapshotURL string, err error) {
	var (
		status    spnStatus
		jobId     string
		timestamp string
		attempt   int
	)

	status, err = checkSPNUserStatus()
	if err != nil {
		return "", err
	}

	if status.DailyCaptures >= status.DailyCapturesLimit ||
		status.Available < 1 {
		return "", fmt.Errorf(
			"SPN Status API Error (avaiable: %d daily_captures: %d limit: %d)",
			status.Available, status.DailyCaptures, status.DailyCapturesLimit)
	}

	jobId, err = requestSPNCapture(pageUrl)
	if err != nil {
		return "", err
	}

	for {
		time.Sleep(time.Second * 5)
		if attempt > 60 {
			return "", fmt.Errorf("SPN Job (%s) Timeout (>5min)", jobId)
		}

		timestamp, err = checkSPNJobStatus(jobId)
		if err != nil {
			return "", err
		}
		if timestamp != "" {
			break
		}

		attempt++
	}

	return fmt.Sprintf("https://web.archive.org/web/%s/%s",
		timestamp, pageUrl), nil
}

// findSnapshot returns a URL to the archive.org snapshot closest to the
// supplied period. If no snapshot is available an empty string is returned.
func findSnapshot(pageUrl string, period time.Time) (
	snapshotUrl string, err error) {
	var (
		responseBytes []byte
		apiResponse   waybackAPI
	)

	pageUrl = fmt.Sprintf(
		"https://archive.org/wayback/available?url=%s&timestamp=%s",
		pageUrl, period.Format("20060102"))

	responseBytes, err = sendRequest(http.MethodGet, pageUrl, nil, false)
	if err != nil {
		return "", err
	}

	apiResponse = waybackAPI{}
	err = json.Unmarshal(responseBytes, &apiResponse)
	if err != nil {
		return "", err
	}

	//LogDebug.Printf("findSnapshot: %v\n", apiResponse)

	if apiResponse.ArchivedSnapshots.Closest.Available {
		return strings.Replace(apiResponse.ArchivedSnapshots.Closest.Url,
			"http://web.archive.org/", "https://web.archive.org/", 1), nil
	}

	return "", nil
}

// checkURL returns true if a HEAD request to the supplied pageURL returns HTTP
// Status 200. Otherwise returns false with the error (if e.g DNS) or the error
// "HTTP Response Status: xxx" with the returned Status code. Request timeout
// set to 20 seconds.
func checkURL(pageUrl string) (ok bool, err error, responseStatus string) {
	var (
		checkClient   http.Client
		clientRequest *http.Request
		response      *http.Response
		dnsError      *net.DNSError
	)

	clientRequest, err = http.NewRequest(http.MethodGet, pageUrl, nil)
	if err != nil {
		return false, err, responseStatus
	}

	clientRequest.Header.Set("Accept", "text/html")
	clientRequest.Header.Set("User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 "+
			"(KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")

	checkClient = http.Client{
		Timeout: time.Second * 20,
	}
	response, err = checkClient.Do(clientRequest)
	if err != nil {
		if errors.As(err, &dnsError) {
			return false, dnsError, "DNS"
		}
		return false, err, responseStatus
	}
	if response.Body != nil {
		defer response.Body.Close()
	}

	if response.StatusCode != http.StatusOK {
		return false, fmt.Errorf("HTTP Response Status: %d",
			response.StatusCode), strconv.Itoa(response.StatusCode)
	}

	return true, nil, strconv.Itoa(response.StatusCode)
}

// Outlink Utilities

func archiveOutlinks(outlinksPath string) (jobSummary string, err error) {
	var (
		outlinksFile *os.File
		outlinks     map[string]map[string]string
		jsonEnc      *json.Encoder
	)

	outlinks, err = readOutlinks(outlinksPath)
	if err != nil {
		LogError.Fatal(err)
	}

	outlinks = outlinksFindSnapshots(outlinks)
	outlinks, jobSummary = outlinksArchive(outlinks)

	outlinksFile, err = os.OpenFile(outlinksPath,
		os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		LogError.Fatal(err)
	}
	defer outlinksFile.Close()

	jsonEnc = json.NewEncoder(outlinksFile)
	jsonEnc.SetIndent("", "  ")
	jsonEnc.SetEscapeHTML(false)
	if err := jsonEnc.Encode(outlinks); err != nil {
		LogError.Fatal(err)
	}

	outlinksFile.Close()
	return jobSummary, nil
}

func readOutlinks(outlinksPath string) (
	outlinks map[string]map[string]string, err error) {
	var (
		outlinksBytes []byte
	)

	outlinksBytes, err = os.ReadFile(outlinksPath)
	if err != nil {
		return outlinks, err
	}

	err = json.Unmarshal(outlinksBytes, &outlinks)
	if err != nil {
		return outlinks, err
	}

	return outlinks, nil
}

func outlinksFindSnapshots(
	outlinks map[string]map[string]string) map[string]map[string]string {

	for page, links := range outlinks {
		LogDebug.Printf("Checking for existing snapshots of outlinks in: %s",
			page)

		for url, snapshot := range links {
			if strings.HasPrefix(snapshot, "http") {
				continue
			}

			period, err := time.Parse("2006-01-02", snapshot)
			if err != nil {
				continue
			}

			LogInfo.Printf(
				"Searching for snapshot near %s of: %s",
				period.Format("2006-01-02"), url)

			snapshot, err := findSnapshot(url, period)
			if err != nil {
				LogError.Printf("findSnapshot Error: %v", err)
			}
			if snapshot != "" {
				LogDebug.Printf("Found snapshot: %s", snapshot)
				outlinks[page][url] = snapshot
			}
		}
	}

	return outlinks
}

func outlinksArchive(
	outlinks map[string]map[string]string) (map[string]map[string]string,
	string) {
	var jobSummary string

	jobSummary = "### Archive.org SPN2 snapshots requested for: \n\n"

	for page, links := range outlinks {
		for url, snapshot := range links {
			if strings.HasPrefix(snapshot, "http") {
				continue
			} else {
				_, err := time.Parse("2006-01-02", snapshot)
				if err != nil {
					continue
				}
				jobSummary += fmt.Sprintf("- <%s>\n", url)
				LogInfo.Printf("Requesting SPN2 snapshot of: %s ", url)

				snapshotURL, err := SPNCapture(url)
				if err != nil {
					LogError.Printf("SPNCapture Error: %v", err)
				}

				if snapshotURL != "" {
					LogDebug.Printf("Created snapshot: %s", snapshotURL)
					outlinks[page][url] = snapshotURL
				}
			}
		}
	}

	return outlinks, jobSummary
}

func checkOutlinks(checkOutlinksPath string) (
	brokenOutlinks map[string][]string, jobSummary string) {
	var (
		checkOutlinksBytes []byte
		checkOutlinks      map[string][]string
		err                error
	)
	brokenOutlinks = map[string][]string{}

	checkOutlinksBytes, err = os.ReadFile(checkOutlinksPath)
	if err != nil {
		LogError.Fatal(err)
	}

	err = json.Unmarshal(checkOutlinksBytes, &checkOutlinks)
	if err != nil {
		LogError.Fatal(err)
	}

	jobSummary = "### Broken Outlinks \n\n"
	jobSummary += "| Page | Status | Outlink |\n"
	jobSummary += "| --- | :---: | --- |\n"

	for page, links := range checkOutlinks {
		LogDebug.Printf("Checking for broken oulinks in: %s", page)

		for _, url := range links {
			ok, err, status := checkURL(url)
			if err != nil {
				LogDebug.Printf("checkURL Error %s: %v", url, err)
			}
			if ok {
				continue
			}

			elipsis := ""
			if len(url) > 60 {
				elipsis = "..."
			}
			jobSummary += fmt.Sprintf("| `%s` | %s | [%.60s%s](%s) |\n", page,
				status, url, elipsis, url)

			LogInfo.Printf("Broken [%s] outlink in: %s (%s)",
				status, page, url)
			if status == "404" || status == "DNS" {
				brokenOutlinks[page] = append(brokenOutlinks[page], url)
			}
		}
	}

	return brokenOutlinks, jobSummary
}

func fixOutlinks(brokenOutlinks map[string][]string, outlinksPath string,
	sitePath string) {
	var (
		outlinks              map[string]map[string]string
		postBytes             []byte
		postString            string
		err                   error
		archiveTimestampRegex = regexp.MustCompile(`([1-2]\d{3}\d*)/(.*)`)
	)

	outlinks, err = readOutlinks(outlinksPath)
	if err != nil {
		LogError.Fatal(err)
	}

	for path, links := range brokenOutlinks {
		LogInfo.Printf("Fixing/Marking broken oulinks in: %s", path)

		postBytes, err = os.ReadFile(sitePath + path)
		if err != nil {
			LogError.Fatal(err)
		}

		for _, url := range links {
			urlRegex := regexp.MustCompile(`\]\((` +
				regexp.QuoteMeta(url) + `(#.*)*)\)`)

			if strings.HasPrefix(outlinks[path][url], "http") {
				LogDebug.Printf("Replacing broken oulink (%s) with: %s in: %s",
					url, outlinks[path][url], path)

				timestamp := archiveTimestampRegex.FindStringSubmatch(
					outlinks[path][url])[1]

				postString = urlRegex.ReplaceAllString(string(postBytes),
					fmt.Sprintf("](%s$2){:data-originalurl=\"$1\" "+
						"data-versiondate=\"%s\"}", outlinks[path][url],
						fmt.Sprintf("%s-%s-%s", timestamp[0:4], timestamp[4:6],
							timestamp[6:8])))
			} else {
				LogDebug.Printf("Marking oulink (%s) as broken in: %s",
					url, path)

				postString = urlRegex.ReplaceAllString(string(postBytes),
					"]($1){:class=\"broken\"}")
			}
		}

		err = os.WriteFile(sitePath+path, []byte(postString), 0)
		if err != nil {
			LogError.Fatal(err)
		}
	}
}

func main() {
	var (
		archiveFlags = flag.NewFlagSet("archive", flag.ExitOnError)
		checkFlags   = flag.NewFlagSet("check", flag.ExitOnError)
	)

	if len(os.Args) < 2 {
		fmt.Println("expected 'archive' or 'check' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "archive":
		var (
			outlinksPath = archiveFlags.String("outlinks",
				"./outlinks.json", "path to outlinks.json file")
			outputJobSummary = archiveFlags.Bool("jobSummary",
				true, "output markdown Job Summary")
		)
		archiveFlags.Parse(os.Args[2:])

		jobSummary, err := archiveOutlinks(*outlinksPath)
		if err != nil {
			LogError.Printf("%v\n", err)
		}

		if *outputJobSummary {
			fmt.Print(jobSummary)
		}

	case "check":
		var (
			outlinksPath = checkFlags.String("outlinks",
				"./outlinks.json", "path to outlinks.json file")
			outlinksCheckPath = checkFlags.String("outlinksCheck",
				"./outlinks-check.json", "path to outlinks-check.json file")
			sitePath = checkFlags.String("sitePath",
				"./", "path to site root")
			outputJobSummary = checkFlags.Bool("jobSummary",
				true, "output markdown Job Summary")
		)
		checkFlags.Parse(os.Args[2:])

		brokenOutlinks, jobSummary := checkOutlinks(*outlinksCheckPath)
		fixOutlinks(brokenOutlinks, *outlinksPath, *sitePath)

		if *outputJobSummary {
			fmt.Print(jobSummary)
		}

	default:
		fmt.Println("expected 'archive' or 'check' subcommands")
		os.Exit(1)
	}
}
