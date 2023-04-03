---
layout: default
title: "Snippet: Filling a PDF form using Python"
permalink: "/posts/snippet-python-pdf-form-filling"
date: 2019-07-02
tag: "snippet"
---

# Snippet: Filling a PDF form using Python

This is just a short note/documentation on something I "wrote" for a project (that didn't end up being presented) at [NHS Hackday #22 in London](https://nhshackday.com/events/2019/06/london).

TL;DR find the code here: <https://github.com/blu3id/python-pdf-form-filling>

## The What

This Snippet was to enable the [HMRC Starter Checklist](https://www.gov.uk/government/publications/paye-starter-checklist) [PDF](https://web.archive.org/web/20190725170132/https://assets.publishing.service.gov.uk/government/uploads/system/uploads/attachment_data/file/783186/Starter_checklist_for_2019_to_2020.pdf){:data-originalurl="https://assets.publishing.service.gov.uk/government/uploads/system/uploads/attachment_data/file/783186/Starter_checklist_for_2019_to_2020.pdf" data-versiondate="2019-07-25"} to be pragmatically filled by a python app.

## The How

Standing on the shoulders of others. A quick search of the internet showed a wonderful starting point from `Jan Chęć` in the post "[Filling PDF Forms In Python — The Right Way](https://medium.com/@zwinny/filling-pdf-forms-in-python-the-right-way-eb9592e03dba)" which lead to "[How to Populate Fillable PDF's with Python](https://web.archive.org/web/20190830215626/https://bostata.com/how-to-populate-fillable-pdfs-with-python/){:data-originalurl="https://bostata.com/how-to-populate-fillable-pdfs-with-python/" data-versiondate="2019-08-30"}"

## The Detail

The heavy lifting is done by the [`pdfrw`](https://github.com/pmaupin/pdfrw) library which provides a convenient way to read and write PDFs. The sticking point with the helpful instruction from the above posts was that the neatest way of filling a PDF is to populate a native PDF form (this also provides the benifit of producing a PDF that can then be further edited).

Unfortunately HMRC no longer provide a PDF form instead opting for a convoluted online form that spits out a rather ugly PDF. So the first step was to create a PDF form from the provided Starter Checklist (any good PDF editor).

The next step was to stich it all together and solve the issue of filled form fields not having an associated rendered state. Trial and error and delving through the [PDF specification](https://web.archive.org/web/20190806160841/https://www.adobe.com/content/dam/acom/en/devnet/pdf/pdfs/pdf_reference_archives/PDFReference.pdf){:data-originalurl="https://www.adobe.com/content/dam/acom/en/devnet/pdf/pdfs/pdf_reference_archives/PDFReference.pdf" data-versiondate="2019-08-06"} lead to the realisation that simply setting the `/ap dictionary` to an empty value made most PDF viewers re-render the set value.

## The Code

As written for the original project see: <https://github.com/blu3id/python-pdf-form-filling>. The simplified transferable version (just add template PDF form with corresponding field names/values):

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

def fill_pdf(input_pdf_path, output_pdf_path, data_dict):
    template_pdf = pdfrw.PdfReader(input_pdf_path)
    for page in template_pdf.pages:
        annotations = page[ANNOT_KEY]
        for annotation in annotations:
            if annotation[SUBTYPE_KEY] == WIDGET_SUBTYPE_KEY:
                if annotation[ANNOT_FIELD_KEY]:
                    key = annotation[ANNOT_FIELD_KEY][1:-1]
                    if key in data_dict.keys():
                        if type(data_dict[key]) == bool:
                            if data_dict[key] == True:
                                annotation.update(pdfrw.PdfDict(
                                    AS=pdfrw.PdfName('Yes')))
                        else:
                            annotation.update(
                                pdfrw.PdfDict(V='{}'.format(data_dict[key]))
                            )
                            annotation.update(pdfrw.PdfDict(AP=''))
    pdfrw.PdfWriter().write(output_pdf_path, template_pdf)


fill_pdf(TEMPLATE_PATH, OUTPUT_PATH, data_dict)

```