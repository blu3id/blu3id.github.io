---
layout: default-date
title: "NHS Digital's War Against Innovation"
permalink: "/posts/nhs-digital-war-against-innovation"
date: 2021-11-19
---

# NHS Digital's War Against Innovation

A recent[^1] [post](https://discourse.digitalhealth.net/t/ers-api-anyone-with-any-experience-with-this/20124) on the [Digital Health Networks](https://discourse.digitalhealth.net/) forum asking about how to use Care Identity Service SmartCard authentication for the [eRS](https://digital.nhs.uk/services/e-referral-service) reminded me of the challenges I started to uncover as I went about trying to realise my idea of a stand alone virtual prescription pad for [EPS](https://digital.nhs.uk/services/electronic-prescription-service) at the start of the pandemic (something being widely sought by all manner of services that were used to using FP10s but that is another story for another time).

As I went about digging up memories of what I had gleaned through hour of research, experimentation and use of human sources I though perhaps it is worth writing out ways that NHS Digital stifles and suppresses innovation on it's platform of digital services - the so-called [Spine](https://digital.nhs.uk/services/spine) - that underpin large swathes of the day to day business of the NHS in England.

My aim here is to:
1. To publicly encourage NHS Digital to do better.
2. To shame NHSx who proclaim [responsibility](https://www.nhsx.nhs.uk/about-us/what-we-do/) for this under so-called "Mission 2" & "Mission 3" by:
   > **Driving implementation**  
   > Delivering APIs and documentation to empower developers and data analysts across the NHS and the health tech industry

   > **Radical innovation**  
   > Supporting the use of new, emergent and effective technologies by the NHS, both by working with industry and through its own prototyping and development

2. To provide a road-map of exactly what **NOT** to do for the NHS in Wales and Scotland as they start to contemplate building "open" platforms.

## Identity Agent Authentication / Care Identity Service / SmartCards

As a quick background for the uninitiated - the majority of national services (Spine) in England hosted by NHS Digital have required (and for the most part still do) a Smart Card to use. By national services I mean among others the: Electronic Prescription Service, Summary Care Record, e-Referral Service, Child Protection Information System and Personal Demographic Service (PDS).[^2] If you want to innovate in the Health & Care IT space you will almost certainly require access to one or more of these.

The first argument that will be made about the lack of investment/poor documentation are that all the authentication APIs are being replaced with newer APIs (so called [NHS Care Identity Service 2](https://digital.nhs.uk/services/identity-and-access-management)) based on modern standards such as OpenID Connect. This is flawed on multiple levels: 
- The "legacy" systems will still require a SmartCard are some of those most in need of innovation and have no [road map](https://digital.nhs.uk/services/identity-and-access-management/identity-and-access-management-roadmap) to being replaced or adopting CIS2.
- This doesn't excuse the the multiple years of (and ongoing) lack of documentation.
- And even if there was a timeline for the "legacy" systems why should innovation and new market entrants have to wait.

The only public [documentation](https://developer.nhs.uk/apis/spine-core/smartcards.html) as it stands - which is relatively recent (~2019) - is ok at outlining the overall architecture. The fact it is buried amongst FIHR documentation is just confusing. There is a set of much more detailed documentation and code samples for accessing the Identity Agent (IA) API that is available if you have access to the HSCN (formerly N3) here: http://nww.hscic.gov.uk/dir/downloads/ alongside the download links for the IA and ancillary tools (Chrome/Edge browser extension for use where the Java Applet bridge is not possible).

## Bureaucratic

So as an SME if you want to get into providing innovative services to the NHS just to get access to to all the documentation about the various APIs and how to access them you need to be on the [HSCN](https://digital.nhs.uk/services/health-and-social-care-network). To get access to the HSCN you need to:
- Get an ODS code
- Sign a HSCN Connection Agreement
- Buy access to the HSCN from a Network Provider

Then once you have read the documentation and decide you want to build something you need to go through a multi layer bureaucratic process to:
- Make sure your ODS code has an associate Application Service Provider (ASP) Code for Spine access
- Have completed the Data Security and Protection Toolkit
- Been issued with some Testing SmartCards and a Reader
- Undergone and paid for the bespoke Assurance process for your application depending on the APIs you wish to access
- Have a sponsoring Health and/or Care organisation for your application (with the product you don't have, that isn't assured, that you haven't even been able to access the documentation for the APIs that you might be able to use to build it - you can see the problem right?)

Congratulations after working you way through all this you can finally actually offer a product to the NHS. I can only imagine the financial cost and uncertainty this presents to an SME. So I don't think it is unreasonable to apportion a sizeable amount of blame to NHS Digital and NHSx for stifling innovation.

## Conclusion

I haven't even touched on the documentation (or lack there of) for the various "Legacy" Spine services. Suffice to say it is in a similar state of disorganised at best, partially hidden on the HSCN and non-existent at worse. There are also - as mentioned above - varying process of Assurance and Conformance testing that are poorly documented beyond "Contact Us" to discuss and we will Work Out Costs&trade;.

So to conclude no-wonder the health IT marketplace is strewn with substandard products from a limited range of suppliers charged at eyewatering prices. NHS Digital simply makes it inhospitable and actively difficult for innovation both from SMEs, in-house teams & well meaning citizens/CICs as part of e.g [NHS Hack Days](https://nhshackday.com/). Some work is underway to improve things - there is now an "internet first" attitude and better [documentation](https://digital.nhs.uk/developer/api-catalogue) - but this is more focused on placating existing suppliers rather than attracting new ones.

### Update:

So since April it appears that things have started to improve slightly using my original use case of EPS there is now a Beta production RESTful API that provides a limited interface to EPS, and it is reasonably [documented](https://digital.nhs.uk/developer/api-catalogue/electronic-prescription-service-fhir). There are however fundamental limitations around authentication and signing that have still not been resolved (you still need HSCN, a SmartCard and and Assured EPS solution to do the  prescription signing). But things are moving forward perhaps?

---

[^1]: Well it was recent when I started writing this post back in April 2021... graduation from my drafts folder however appears to have taken rather some time.

[^2]: NHS Digital have actually provided Smart Card and HSCN network free access to a limited subset of the PDS API for some time (~2018) via the Spine Mini Service Provider.