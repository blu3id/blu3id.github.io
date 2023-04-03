---
layout: default
title: "The unappreciated risks of IT deployment within Complex Systems"
permalink: "/posts/complex-system-dangers"
date: 2017-06-18
---

# The unappreciated risks of deploying IT within Complex Systems
{:.no_toc}

Put another way: _The real-world patient safety implications of undertaking large-scale change (deploying large IT projects) within a tightly coupled complex adaptive system (the National Health Service)._

> "Complex systems are not responsive to warnings of unimaginable or highly unlikely accidents"  
> &mdash; Charles Perrow 1981[^1]

_I should prefix this with a brief warning: I am not &mdash; and do not pretend to be &mdash; an expert on sociology, organisational theory or system safety. However I do have a fairly good understanding of the above and plenty of personal professional experience within the specific 'change field' of IT deployment, I also have domain experience within the NHS._

- Table of Contents
{:toc}

## The Problems

### Normal Accidents and Complex Systems

There is a naïve view within the NHS (and other healthcare organisations around the world) that the introduction of IT systems will do among other things one or many of the following:

1. Improve productivity & "patient flow".[^2][^3] _"For every hour physicians provide direct clinical face time to patients, nearly 2 additional hours is spent on EHR"_
2. Improve patient outcomes.[^4] _"Mortality rate ... increased from 2.80% ... to 6.57% (36 of 548) after ... implementation"_
3. Decrease/Eliminate common prescribing errors.[^4][^5][^6] _"...that many of these errors are not caught by CPOE systems in use today"_
4. Never have to be used in a major incident/ eliminate business continuity concerns.[^7] _"...problems with health IT can disrupt care delivery and harm patients"_
5. Be loved by users, especially after they have had training and so "actually know how to use the system properly and better".

Unfortunately reality has a different set of options. There are a number of complex and deeply nuanced reasons for this... but that would be a whole different series of posts. The referenced papers against each of the above 'outcomes' show that the majority do not achieve much from this list and often are left with the exact opposite. One of the often ignored causes of this harm is the well recognised pattern of system accidents that are inevitable in complex systems, frustratingly much of this harm can be mitigated in healthcare[^8] by learning from the wealth of published [theory](#theory-a-skeleton-overview).

### Lack of Evidence, Scrutiny, Monitoring and Effective Change Management

In modern medicine introduction of a new drug, treatment or procedure is highly scrutinised and evidence based. There are trials, audits and research studies that evaluate the effectiveness and safety of these changes. However when is comes to introduction of IT systems to the clinical environment there is no such scrutiny, there is no requirement for evidence, there is no on-going supervisions and reporting of adverse incidents.

IT deployments are a change management process - at risk of veering off-topic - there is a vast amount of theory on how best to manage and implement change and yet within the NHS there is either wilful ignorance of this amongst IT staff/project managers or (perhaps more worryingly) a complete lack of understanding. The emphasis is always on the project or the software never on the "product" (providing safe, efficient healthcare) or the end users who deliver it (health care professionals) which is at odds with change management theories (such as Kotter, Lewin or McKinsey) that all focus on the users (health care professionals) and how the change will improve the product (providing safe, efficient healthcare).

### Work as Imagined vs. Work as Done

<blockquote class="twitter-embed">
  <header>
    <img class="profile" alt="@MarkGraban profile picture" src="/assets/images/twitter-embed/@markgraban.jpg"/>
    <div class="name">
      Mark Graban
      <div class="handle">@MarkGraban</div>
    </div>
  </header>
  <p>
    When the MA takes your temp and weight &amp; writes it on a post it note until she gets to the EMR. Ah the digital revolution. 🤓
  </p>
  <a class="permalink" href="https://twitter.com/MarkGraban/status/854010172913713154" target="_blank">
    <time datetime="2017-04-17T16:34:26.000Z">5:34 PM · April 17, 2017</time>
  </a>
</blockquote>

This is a classic example of IT for 'Work as Imagined' vs. real world needs of IT for 'Work as Actually Done' that represents technology in the NHS in a nutshell. The focus is only ever on the problem as perceived by committee or management not the problem as it is in the real world. This is a well recognised problem in literature[^3] as a cause for failure of IT deployments to produce 'expected' benefits and gains.

## Theory (_a skeleton overview_)

_There are many theories encompassing the topic explored in this post, I will try to cover the basics and provide a starting point for further reading (at some point I may expand this section, but for now it shall be left as an exercise for the reader to research)._

### Human Factors & Swiss Cheese Model

The practice of modern healthcare within the NHS is a complex, tightly coupled and decidedly human process that carries risk.[^9] The most common group of theory in this area is Human Factors (the understanding of interactions between humans and other system elements). The Swiss Cheese Model[^10] describes the series of defensive measures inherent in a system to reduce risk and prevent accidents.

### Normal Accident theory (NAT)

Described by Perrow in 1981 in his analysis of the events surrounding the accident at Three Mile Island.[^1] He described how accidents are inevitable in extremely complex systems. That given the nature and tightly coupled complexity of such systems multiple failures which interact with each other will occur despite best efforts to avoid them.

> "In normal accidents the particular trigger is relatively insignificant; the interaction is significant"  
> &mdash; Charles Perrow 1981[^1]

### High Reliability theory (HRT)

Is the antithesis of Normal Accident theory. It attempts to describe the factors that contribute to reliability within an organisation.[^11] The premise being that accidents are avoidable in complex organisations if enough planning and preparation is performed. Briefly it describes four organisational characteristics that limit failure:

1. Prioritisation of both safety and performance.
2. A culture of reliability that is decentralised and paradoxically centralised to allow decision making authority to be available where needed.
3. Learning organisation, learning and changing based following accidents, incidents and near-misses.
4. Redundancy working beyond technology extending to behaviour.

More recently there has been work on attempting to view NAT and HRT as complementary rather than competitive.[^12] This has focused on different temporal perspectives and how reframing that analysis of NAT within an open system can address the non-falsifiability of the theory.

### Safety-I and Safety-II

Other relevant reading material includes the excellent White Paper by Professor Erik Hollnagel "From Safety-I to Safety-II"[^13] which puts into context NAT and HRT in a very approachable and healthcare specific manner and introduces the concepts of Safety-I and Safety-II. Two brief quotes for definitions as follows:

> "Most people think of safety as the absence of accidents and incidents (or as an acceptable level of risk). In this perspective, which we term Safety-I, safety is defined as a state where as few things as possible go wrong."  
> &mdash; Hollnagel et al. 2015[^13]

> "Safety management should therefore move from ensuring that ‘as few things as possible go wrong’ to ensuring that ‘as many things as possible go right’. We call this perspective Safety-II; it relates to the system’s ability to succeed under varying conditions."  
> &mdash; Hollnagel et al. 2015[^13]

### Complex Adaptive Systems

Wikipedia does a fine job defining and describing [Complex Adaptive Systems](https://en.wikipedia.org/wiki/Complex_adaptive_system) as follows:

> A complex adaptive system is a system in which a perfect understanding of the individual parts does not automatically convey a perfect understanding of the whole system's behaviour.

> They are complex in that they are dynamic networks of interactions, and their relationships are not aggregations of the individual static entities, i.e., the behaviour of the ensemble is not predicted by the behaviour of the components. They are adaptive in that the individual and collective behaviour mutate and self-organize corresponding to the change-initiating micro-event or collection of events.

## Worked Example

### A Patients discharge from hospital / Electronic test requesting and reports
{:.no_toc}

This example will examine the unintended impact of the introduction of (another) system to view results and electronic test requesting. The detailed background is beyond scope the salient points to have sufficient context for this example are:

- The new system for viewing lab test results introduced the new functionality of having insight into when a specific sample had been received by the laboratory and analysis was 'in-progress'.
- The new electronic test requesting system has a poorly designed 'Directory' of available tests not consistent with day-to-day practice or common sense.

Events unfold as follows. This scenario is fictitious but incorporates a number of real examples and situations that have resulted from the deployment of a new IT system.

1. _Decision to discharge._

   The decision to discharge the patient was taken and specific criteria were set to enable a safe discharge: serum potassium improving following changes in medication.

2. _Organisation and planning discharge._

   Hospital transport is booked and specific start times for packages of care to coincide with the arrival home of the patient are organised.

3. _Planning blood test._

   The Junior Doctor at the end of the working day (as it is not time efficient to access a computer, log-in and follow the laborious process of requesting tests) requests the necessary blood test to monitor the serum potassium.

   As they are unfamiliar with the online system and/or are tired they forget that unlike 'traditional' paper requests – where they would simply write "U/E" – the online directory lists this specific test as "Electrolyte Profile" so when they by habit select the first presented test after typing "U" the wrong test is requested – "Urea" a single electrolyte, and not the potassium needed to support the discharge.

4. _Blood sample taken._

   The blood sample is collected from the patient in the early morning so results will be available by the planned discharge time.

5. _Results reviewed._

   The team note that that laboratory have only tested serum urea rather than the full electrolyte profile. A manual "add-on" request is sent to lab.

6. _Results reviewed, again._

   There are still no results for the remaining electrolytes. Remembering that the new results system displays status information the Junior Doctor opts to use this system to check on the status of the "add-on" test. As the ‘traditional’ method of ringing the laboratory and navigating a time consuming and patronising set of interactive voice prompts would be unproductive. The results viewing system to confirm that additional tests are "in-progress".

7. _Patient missed arranged hospital transport slot._

8. _Results reviewed, for the third time._

   There are still no results for remaining electrolytes. The Junior Doctor decides to spend the time being on hold and navigating the IVR to contact the laboratory for an update. Perhaps there has been an equipment malfunction or emergency? After waiting to get through the laboratory staff report that the results should be available as they have already been released.

9. _Attempt to re-publish the results to the report system._

   This fails. The results are verbally released.

10. _Results reviewed._

    The potassium has improved following the medication changes and the patient can be safely discharged.

11. _Patient discharged._

    Taken home by next available hospital transport.

12. Due to later than expected discharge, package of care unable to meet patient on arrival so returned to hospital.

In the above example events conspired to result is a delayed discharge - patient harm, equally the circumstances could have been different: this could have simply been a set of routine blood tests that picked up a dangerously high potassium similar events could have conspired and ultimately due to lack of a timely result the patient could have died. You are probably thinking that there are other fail-safes (slices of the Swiss Cheese) to prevent this - and you would be right - but it is disturbingly becoming more routinely (?due to workload or other reasons) for abnormal results to be highlighted by the lab a considerable time (+4h) after they are "published" to reporting systems.

Let us now explore the above example and apply the theory to learn more about how this theoretical scenario unfolded with the aim of understand where different approaches to deployment could have prevented or reduced the risk of harm:

- The chain of events leading to harm begins in 3. the root cause (though it would be unfair to blame one single event - see Perrow's [NAT](#normal-accident-theory-nat)[^1]) is that it is difficult and time consuming to request blood tests on the new requesting system. This in turn led to the Junior Doctor making the decision (logical but only unwise if they had specific human factors training to be more aware of their limitations) to request all tests at the end of the day - when they are most tired, distracted and prone to make errors. This ultimately leads to the "wrong" test being requested.
- The second significant inflection point is in 6. if the virtues of the new reporting system had not been explained to the Junior Doctor they would have rung the lab at this point and the error with the lab systems would have been identified earlier. The other contributing factor is the time consuming and patronising set of interactive voice prompts - this system is in place to serve a purpose but again much like the IT deployment we are analysing it had unintended consequences.
- The final error/accident if the malfunction of the laboratory systems to publish released results. If this did not happen the scenario would have been cut short.

The above are the three separate occasions when the pieces of the Swiss Cheese[^10] could have been prevented the ultimate outcome. It is also clear that there was no single event that caused this scenario to evolve, instead it was a series of minor errors the combined to form this one event an outcome - exactly as described by Perrow's [NAT](#normal-accident-theory-nat).[^1] Attempting to apply [HRT](#high-reliability-theory-hrt) it is clear that NHS organisations aspire to the four main characteristics to limit failure, however unfortunately for a number of complex and deeply nuanced reasons (a whole other series of posts) in reality it is rare that any of the four are applied in practice. Similarly there is almost no application of [Safety-II](#safety-i-and-safety-ii) emphasis is almost exclusively within the [Safety-I](#safety-i-and-safety-ii) frame of reference.

For a specific lesson from this scenario: The introduction and design of the new test requesting and reporting system did not make a meaningful attempt to understand the system beyond 'Work as Imagined' to encompass 'Work as Actually Done'. If it had there may have been a realisation of the need to effortlessly and quickly request tests as a real requirement. Further analysis of the whole system may have highlighted a consequence of ignoring this requirement being exposure to increased risk of errors due to human factors as describe above.

Unfortunately reality it is far more complex than this example hints at - the test requesting and reporting system has many users and stakeholders each with different requirements which have influenced is design and deployment - a whole post would barely scratch the surface in exploring the significant interactions and complexities of this. This example is more of a taster to show a whole new set of underappreciated risks.

## Conclusion

I would suggest that fundamentally you should not be deploying IT projects - novel or not - in healthcare (especially not at scale) if you:

1. Don't even have insight into the inherent system risks or any of the theory touched on in the post.
2. Don't actually understand how a system works (not just how it is designed to work - see IT for work as imagined vs. real world needs of IT for work as actually done).
3. Don't have effective (and tested) systems in-pace to monitor for and measure adverse consequences.<sup>[b](#noteb)</sup>
4. Are knowingly not addressing basic system requirements in order to meet deployment deadlines or targets. Safety trumps all; as outlined in this post basic requirements can have a bigger impact that may be first apparent.

A cautionary lesson (and aside) if you believe that a well designed, validated and redundant systems in a well understood system are infallible when you remove the human component, picture this: An aircraft with multiple redundant and validated systems enters an unexpected failure mode where inbuilt flight envelop protection systems cause the plane to violently pitch-down and hurtle towards the ground completely unresponsive to pilots attempts to save the aircraft. Well this happened to [Quantas Flight 72](https://en.wikipedia.org/wiki/Qantas_Flight_72) there have been some lessons leaned but no root-cause identified.<sup>[a](#notea)</sup>

My recommendations would therefore be:

- At a minimum read: [Normal Accidents](https://www.amazon.co.uk/Normal-Accidents-Technologies-Princeton-Paperbacks/dp/0691004129) by Charles Perrow and the references on this post.
- Establish out-of-band organisational/system monitoring<sup>[b](#noteb)</sup> to analyse impact of deployment and pick-up warning signs of failure &mdash; a slightly different recommendation than that made by Meeks et al.[^14]
- Teams leading deployment and training should be experienced in:
  1. Human factors
  2. The (actual not theoretical) workflow of the system within which they are deploying.
- **TALK TO YOUR _ACTUAL_ END USERS** (not just their managers and service leads) to understand how a system works and what impact IT will have end-to-end. But for this to be meaningful keep in mind:
  <blockquote class="twitter-embed">
    <header>
      <img class="profile" alt="@dr_lungs profile picture" src="/assets/images/twitter-embed/@dr_lungs.jpg"/>
      <div class="name">
        John Watson
        <div class="handle">@dr_lungs</div>
      </div>
    </header>
    <a class="reply-to" href="https://twitter.com/dr_lungs/status/866387344726904833" target="_blank">
      Replying to @dr_lungs @DrUmeshPrabhu and @doctorcaldwell
    </a>
    <p>
      Trouble is nobody asks the FY1s and if they do they&#39;re too nervous/polite to authority to say what they really think
    </p>
    <a class="permalink" href="https://twitter.com/dr_lungs/status/866387713381019648" target="_blank">
      <time datetime="2017-05-21T20:18:22.000Z">9:18 PM · May 21, 2017</time>
    </a>
  </blockquote>

As ever I welcome any feedback, comments, criticism or suggestions on what I have outlined.

---

**Notes:**

<a name="notea">[a]</a>: Technical side note/rabbit hole for the curious: Airbus flight control law failure modes don't account for this failure - you can only enter Direct Law (no envelop protection) or Alternate Law 2 (no high AoA or high speed protection) if two ADIRU's fail not if there isn't parity across all three ADIRU (which also isn't/wasn't a cause to fail out an ADIRU).

<a name="noteb">[b]</a>: And no existing "Incident Reporting" systems such as DATIX are not the type of monitoring I am talking about - in the current climate (and for that matter culture of many organisations) in the NHS these systems record a metaphorical scraping of the tip of the iceberg. [Dr Gordon Cadwell](https://twitter.com/doctorcaldwell) has [tweeted](https://twitter.com/doctorcaldwell/status/876144163200393217) about some of the resons behind this. The type of monitoring that should be aspired to is something similar to the [Neuraxial Connector Surveillance](https://web.archive.org/web/20170504205951/http://www.wales.nhs.uk/sites3/page.cfm?orgid=457&pid=72625){:data-originalurl="http://www.wales.nhs.uk/sites3/page.cfm?orgid=457&pid=72625" data-versiondate="2017-05-04"} programme that is being run in Wales. There needs to be minimally (or ideally completely unobtrusive) baseline monitoring before deployment, during and after in order to pick up issues - again there could be a whole series of posts on the topic of effective monitoring.

> "The worst pain a man can suffer: to have insight into much and power over nothing." ― Herodotus


[^1]: [Perrow, C., 1981. Normal accident at three mile island. Society, 18(5), pp.17–26.](https://dx.doi.org/10.1007/BF02701322)

[^2]: [Villares, J.M.M., 2016. Allocation of physician time in ambulatory practice: A time and motion study in 4 specialties. Acta Pediatrica Espanola, 74(8), p.203.](https://dx.doi.org/10.7326/M16-0961)

[^3]: [Jones, S.S. et al., 2012. Unraveling the IT Productivity Paradox — Lessons for Health Care. New England Journal of Medicine, 366(24), pp.2243–2245.](https://dx.doi.org/10.1056/NEJMp1204980)

[^4]: [Han, Y.Y. et al., 2005. Unexpected Increased Mortality Arter Implementation of a Commercially Sold Computerized Physician Order Entry System. Pediatrics, 116(6), pp.1506–1512.](https://dx.doi.org/10.1542/peds.2005-1287)

[^5]: [Coiera, E., 2015. Technology, cognition and error. BMJ quality & safety, 24(7), pp.417–22.](https://dx.doi.org/10.1136/bmjqs-2014-003484)

[^6]: [Payne, T.H., 2015. Electronic health records and patient safety: should we be discouraged? BMJ Quality & Safety, 24(4), pp.239–240.](https://dx.doi.org/10.1136/bmjqs-2015-004039)

[^7]: [Kim, M.O., Coiera, E. & Magrabi, F., 2017. Problems with health information technology and their effects on care delivery and patient outcomes: A systematic review. Journal of the American Medical Informatics Association, 24(2), pp.246–260.](https://dx.doi.org/10.1093/jamia/ocw154)

[^8]: [Tamuz, M. & Harrison, M.I., 2006. Improving patient safety in hospitals: Contributions of high-reliability theory and normal accident theory. Health Services Research, 41(4 II), pp.1654–1676.](https://dx.doi.org/10.1111/j.1475-6773.2006.00570.x)

[^9]: [Reason, J.T., Carthey, J. & de Leval, M.R., 2001. Diagnosing “vulnerable system syndrome”: an essential prerequisite to effective risk management. Quality in health care, 10(Suppl 2), p.ii21-5.](https://qualitysafety.bmj.com/content/10/suppl_2/ii21)

[^10]: [Reason, J., 2000. Human error: models and management. Bmj, 320(March), pp.768–770.](https://dx.doi.org/10.1136/bmj.320.7237.768)

[^11]: [Weick, K.E., 1987. Organizational Culture as a Source of High Reliability. California Management Review, 29(2), pp.112–127.](https://dx.doi.org/10.2307/41165243)

[^12]: [Shrivastava, S., Sonpar, K. & Pazzaglia, F., 2009. Normal Accident Theory versus High Reliability Theory: A resolution and call for an open systems view of accidents. Human Relations, 62(9), pp.1357–1390.](https://dx.doi.org/10.1177/0018726709339117)

[^13]: [Hollnagel, E., Wears, R.L. & Braithwaite, J., 2015. From Safety-I to Safety-II: A White Paper, University of Southern Denmark, University of Florida, USA, and Macquarie University, Australia.](https://www.england.nhs.uk/signuptosafety/wp-content/uploads/sites/16/2015/10/safety-1-safety-2-whte-papr.pdf)

[^14]: [Meeks, D.W. et al., 2014. An analysis of electronic health record-related patient safety concerns. Journal of the American Medical Informatics Association, 21(6), pp.1053–1059.](https://dx.doi.org/10.1136/amiajnl-2013-002578)
