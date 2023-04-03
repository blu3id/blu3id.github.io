---
layout: default
title: "Kitchen Sink"
permalink: "/kitchen-sink"
unlisted: true
---

# Kitchen Sink
{:.no_toc}

First paragraph text. This is a kitchen sink page showing all the styled elements on the site in one go. This needs to be long enough to wrap at all viewport widths. So there are extra words here.

Second paragraph text. This is a paragraph with some *emphasized text*, **strongly emphasized text**, ~~struck-through text~~, `inline code`, a <abbr title="Three Letter Acronym">TLA</abbr>, a [link to the homepage](/) and this is an [external link](https://en.wikipedia.org/).

- Table of Contents
{:toc}

## Headings

## Heading 2
{:.no_toc}
### Heading 3
{:.no_toc}
#### Heading 4
{:.no_toc}
##### Heading 5
{:.no_toc}

## Links

[Internal Link](/) and this is an [external link](https://en.wikipedia.org/) and one that is [broken](https://en.wikipedia.org/){:class="broken"}, and one that is [archived](https://en.wikipedia.org/){:data-originalurl="https://en.wikipedia.org/"}.

[Link to a .pdf](fake.pdf)  
[Link to a .xlsx](fake.xlsx)

## Footnotes

This is a paragraph with some footnotes[^1] in it. We have multiple[^2][^3] notes and reuse[^1] some of the notes as references.

[^1]: This is the footnote text
[^2]: Second footnote with some other text. This needs to be long enough to wrap at all viewport widths. So there are extra words here.
[^3]: Another note

## Unordered list

- unordered
- list. This needs to be long enough to wrap at all viewport widths. So there are extra words here.
- items
- Paragraphs.

  This needs to be long enough to wrap at all viewport widths. So there are extra words here.

- Line break  
  in paragraph.

  And a second paragraph. This needs to be long enough to wrap at all viewport widths. So there are extra words here.

## Ordered list

1. ordered
2. list. This needs to be long enough to wrap at all viewport widths. So there are extra words here.
3. items
4. Paragraphs.

   This needs to be long enough to wrap at all viewport widths. So there are extra words here.

5. Line break  
   in paragraph.

   And a second paragraph. This needs to be long enough to wrap at all viewport widths. So there are extra words here.

## Blockquote

> Blockquote. [Internal Link](/) and this is an [external link](https://en.wikipedia.org/), one that is [broken](https://en.wikipedia.org/){:class="broken"}, and one that is [archived](https://en.wikipedia.org/){:data-originalurl="https://en.wikipedia.org/"}.
>
> — Quote Source

## Tables

| Left columns  | Center columns | Right columns |
| :------------ |:--------------:|--------------:|
| left foo      | center foo     | right foo     |
| left bar      | center bar     | right bar     |
| left baz      | center baz     | right baz     |

## Code

```python
import pdfrw

TEMPLATE_PATH = 'template.pdf'
OUTPUT_PATH = 'output.pdf'

ANNOT_KEY = '/Annots'
ANNOT_FIELD_KEY = '/T'
ANNOT_VAL_KEY = '/V'
ANNOT_RECT_KEY = '/Rect'
SUBTYPE_KEY = '/Subtype'
WIDGET_SUBTYPE_KEY = '/Widget'

data_dict = {
    'field_name' : 'Value'
}
```

## Info Box

{: class="info"}
> Information in a box. [Internal Link](/) and this is an [external link](https://en.wikipedia.org/) and one that is [broken](https://en.wikipedia.org/){:class="broken"}. This needs to be long enough to wrap at all viewport widths. So there are extra words here.

## Warning Box

{: class="warning"}
> Warning in a box. [Internal Link](/) and this is an [external link](https://en.wikipedia.org/) and one that is [broken](https://en.wikipedia.org/){:class="broken"}. This needs to be long enough to wrap at all viewport widths. So there are extra words here.

## Box

{: class="box"}
> Text in a box. [Internal Link](/) and this is an [external link](https://en.wikipedia.org/) and one that is [broken](https://en.wikipedia.org/){:class="broken"}. This needs to be long enough to wrap at all viewport widths. So there are extra words here.

## Timeline

{: class="timeline"}
- {: data-timeline="1999"}
    **Title**

    Paragraph of text in the timeline entry. This needs to be long enough to wrap at all viewport widths. So there are extra words here.

- {: data-timeline="2007"}
    **Standalone Title**

- {: data-timeline="2009"}
    **Title**

    With pragraph

## Figure

<figure>
    <img src="data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='2560' height='1440'%3E%3Crect width='100%25' height='100%25' fill='gray'/%3E%3Cpath stroke='%23dcdcdc' d='m0 0 2560 1440M2560 0 0 1440'/%3E%3Ctext x='1280' y='725' fill='%23fff' font-size='80' text-anchor='middle'%3E2560 x 1440%3C/text%3E%3C/svg%3E" alt="2560 x 1440 Placeholder">
    <img src="data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='2560' height='1440'%3E%3Crect width='100%25' height='100%25' fill='gray'/%3E%3Cpath stroke='%23dcdcdc' d='m0 0 2560 1440M2560 0 0 1440'/%3E%3Ctext x='1280' y='725' fill='%23fff' font-size='80' text-anchor='middle'%3E2560 x 1440%3C/text%3E%3C/svg%3E" alt="2560 x 1440 Placeholder">
    <figcaption>Figure caption goes here</figcaption>
</figure>

## Twitter Embed

<blockquote class="twitter-embed">
  <header>
    <img class="profile" alt="@handle profile picture" src="data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='2560' height='1440'%3E%3Crect width='100%25' height='100%25' fill='gray'/%3E%3Cpath stroke='%23dcdcdc' d='m0 0 2560 1440M2560 0 0 1440'/%3E%3Ctext x='1280' y='725' fill='%23fff' font-size='80' text-anchor='middle'%3E2560 x 1440%3C/text%3E%3C/svg%3E"/>
    <div class="name">
      Display Name
      <div class="handle">@handle</div>
    </div>
  </header>
  <a class="reply-to" href="https://en.wikipedia.org/" target="_blank">
    Replying to @handle
  </a>
  <p>
    Tweet text goes here. This needs to be long enough to wrap at all viewport widths. So there are extra words here. Also a link to <a href="https://en.wikipedia.org/" target="_blank">en.wikipedia.org</a>
  </p>
  <img alt="Image embeded in Tweet: 2560 x 1440 Placeholder" src="data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='2560' height='1440'%3E%3Crect width='100%25' height='100%25' fill='gray'/%3E%3Cpath stroke='%23dcdcdc' d='m0 0 2560 1440M2560 0 0 1440'/%3E%3Ctext x='1280' y='725' fill='%23fff' font-size='80' text-anchor='middle'%3E2560 x 1440%3C/text%3E%3C/svg%3E" />
  <a class="permalink" href="https://en.wikipedia.org/" target="_blank">
    <time datetime="2018-08-09T19:34:15.000Z">8:34 pm · August 9, 2018</time>
  </a>
</blockquote>
