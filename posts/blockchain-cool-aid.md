---
layout: default-date
title: "The Blockchain Cool-Aid and Snake Oil"
permalink: "/posts/blockchain-cool-aid"
date: 2018-08-11
---

# The Blockchain Cool-Aid and Snake Oil

I have been driven to write this _rapid response_ post by the outpouring of support for a "Blockchain Based Solution" to the problem of electronic staff identity within the NHS that has been plaguing Twitter for the past 48 hours (see following thread). 

<blockquote class="twitter-tweet" data-lang="en"><p lang="en" dir="ltr">So the future is coming now - fab session today on the collaboration of NHS, vendors, medical directors and RAs on how we can use Blockchain for staff identity... <a href="https://twitter.com/hashtag/CCIO7">#CCIO7</a> <a href="https://twitter.com/NHSCCIO">@NHSCCIO</a> <a href="https://twitter.com/PhilipGrahamNHS">@PhilipGrahamNHS</a> <a href="https://twitter.com/SarahFWilkinson">@SarahFWilkinson</a> <a href="https://twitter.com/NHSDigital">@NHSDigital</a> <a href="https://t.co/YCVJcurW22">pic.twitter.com/YCVJcurW22</a></p>&mdash; indi singh (@mr_indisingh) <a href="https://twitter.com/mr_indisingh/status/1027639212173127680">August 9, 2018</a></blockquote>


## What is a Blockchain/Distributed Ledger?

First it is vitally important that we clarify and explain just exactly what a "Blockchain" or "Distributed Ledger" is that way (hopefully) it will be abundantly clear why there is no logical need for one within the NHS. Least of all to "solve" staff identity and access to electronic systems.

From [Wikipedia](https://en.wikipedia.org/wiki/Blockchain):
> Blockchain was invented by Satoshi Nakamoto in 2008 to serve as the public transaction ledger of the cryptocurrency bitcoin. The invention of the blockchain for bitcoin made it the first digital currency to solve the double-spending problem without the need of a trusted authority or central server.

A blockchain is made up of a list of records (blocks) that are linked cryptographically (usually by a [Merkle trees](https://en.wikipedia.org/wiki/Merkle_tree)) that have the singular property of making it dificult to alter older blocks. As all block created since will have to be altered acordingly to ensure the cryptographic signatures match. This property is exploited in the context of cryptocurrency and combined with consesus mechanics such a proof-of-work (mining) and multiple nodes (many users) to create the public transaction ledger.

To be explicit [Merkle trees](https://en.wikipedia.org/wiki/Merkle_tree) are indeed useful for tamper resistant audit records - see some of my [other](/posts/auditable-records-access-infrastructure) ramblings or the work on [Certificate Transparency](https://en.wikipedia.org/wiki/Certificate_Transparency). To be clear Blockchains/Distributed Ledgers are *not* Merkle trees they are the combination of a cryptographic scheme, consensus protocol (proof-of-work) and other mechanics.

## What problems are we "solving"?

There are many, many, _many_ problems with technology and IT in the the NHS. But the problems that Twitter has been abuzz with are as follows:

1. Staff Identity
2. Staff Access to Systems

<blockquote class="twitter-tweet" data-conversation="none" data-lang="en"><p lang="en" dir="ltr">The hypothesis is to test whether Distributed Ledger Technology can improve the management of staff identity and link to system access... in particular the PoC is looking at Locums.... lots of opportunities.... <a href="https://twitter.com/hashtag/WatchThisSpace">#WatchThisSpace</a></p>&mdash; Philip Graham (@PhilipGrahamNHS) <a href="https://twitter.com/PhilipGrahamNHS/status/1027886273313341442">August 10, 2018</a></blockquote>


And as secondary goals per Simon Eccles and others:

<blockquote class="twitter-tweet" data-conversation="none" data-lang="en"><p lang="en" dir="ltr">No - not at all. The GMC ledger is vital but doesn’t confirm your ID. Nor does it prove you hold the postgrad quals and competencies you claim. A GMC no (or equivalent for other professions, and the no. sets overlap) doesn’t get you paid. Multiple ledgers - for all clinicians.</p>&mdash; Simon Eccles (@NHSCCIO) <a href="https://twitter.com/NHSCCIO/status/1027977904360173568">August 10, 2018</a></blockquote>


3. Confirmation of Licence to Practice (GMC, NMC etc. Register Entry)
4. Post-graduate qualifications (Passed membership exams of Royal Colleges)
5. Occupational Health Records

## Where we are now / Historical context

For those new to the party (or for those old enough to remember but have forgotten) we need a quick recap on the short history of NHS IT initiatives borne out of the calamitous [NPfIT](https://en.wikipedia.org/wiki/NHS_Connecting_for_Health).

### SmartCards

Within NHS England since ~2005 there has been a unified central electronic identity service for staff. This has matured and improved in the intervening years and is now know as the [Care Identity Service](https://digital.nhs.uk/services/registration-authorities-and-smartcards/care-identity-service). This has formed the backbone of access to the national Spine and [Spine2](https://digital.nhs.uk/services/spine) services through one of the largest active deployments of Public Key Infrastructure with SmartCards. There are clear [guidelines and processes](https://digital.nhs.uk/services/registration-authorities-and-smartcards#registration-authorities) about identity verification before issuing cards. SmartCards are unique to the individual and specifically designed to be used across organisations to maintain a consistent digital identity for staff.

> "Your Smartcard is a national token of your identity which is not specific to a particular organisation.  You should retain your card if leaving an organisation for use in other health and social care settings.."  
&mdash; [Leaflet for new smartcard users](https://digital.nhs.uk/binaries/content/assets/legacy/word/4/k/end_user_leaflet_v2.0_120117_locked.docx)

It is important to mention the caveat that within the NHS the implementation of how SmartCards are used with clinical systems and across organisational boundaries is significantly deficient and poses many barriers.

### NHSMail

NHSMail (or Contact as it was named in 2004) is now in it's second revision as [NHSMail2](https://digital.nhs.uk/services/nhsmail) operating as a fully managed service contracted out to Accenture.

NHSMail offers a secure email service for health and care users across England and Scotland. It boasts a single directory of staff across a large number of NHS organisations (but not all as some still operate their own email systems).

### ESR

The [Electronic Staff Record](http://www.electronicstaffrecord.nhs.uk/home/) is an unwieldly monstrosity used across England and Wales to manage HR functions and other add-ons (e-learning and mandatory training). This is a single unified platform that tracks employees across their careers and through multiple organisations within the NHS. 


## What others are doing to "solve" these problems

Wales has been largely unaffected by the legacy of [NPfIT](https://en.wikipedia.org/wiki/NHS_Connecting_for_Health) and has be quietly building its own health and care infrastructure. Early on in this journey it was recognised that a consistent electronic staff identity was required. Starting in 2009 the [NADEX](http://www.wales.nhs.uk/nwis/page/52675) has been successfully rolled out across the whole of Wales providing as single set of logon credentials to any NHS Wales computer, national clinical systems, most local systems and a national email service similar to NHSMail. This single electronic identity has vastly simplified the deployment of Audit tools such as [NIIAS](http://www.wales.nhs.uk/nwis/page/87467) used to track staff access to patient records.

Again much like the SmartCard it is important to caveat that while for the most part this single national identity and credential works well there are a number of implementation issues that cause occasional problems day-to-day for staff as they move from organisation to organisation.

## Conclusion

So the original premise from Twitter (and no doubt actively being perused at significant cost both financially and in terms of time) was that the NHS needs to embrace Blockchain/Distributed ledger technology to "solve" the electronic staff identity and access problem.

However what we have learnt is NHS England already has three different central authorities that could or do already provide a consistent electronic staff identity. That Blockchain technology is suited to enforcing trust (through an immutable record) across multiple thousands of systems that are inherently untrustworthy and have no pre-existing relationships (the exact opposite of NHs organisations).

And that Wales has already "solved" the identity and access problem  (back in 2009) without the use of a Blockchain.

### Where efforts should be focused:

While it is true NHS England does have existing central authorities of staff identity unfortunately the [Care Identity Service](https://digital.nhs.uk/services/registration-authorities-and-smartcards/care-identity-service) doesn't provide for the [other problems]() listed out on twitter - and to be clear neither would a distributed ledger - and isn't quite set-up to act as a single combined identity/authentication service for the whole of NHS England... yet.

Delegating Access to individual systems will always be a problem (one of the implementation issues in Wales) as regardless of technical solution employed Access to patient record will always need to be individually granted by organisations (local or regional) when employees join them.

However there is room to improve the electronic view of NHS staff. For example looking at Doctors it would be great if there was a single federated view provided by a [EMPI](https://en.wikipedia.org/wiki/Enterprise_master_patient_index) like system that could collate in real-time:
- Professional registration (GMC Register e.g Specialty register entry, responsible officer etc.)
- Post-graduate qualifications (College Exams e.g MRCP, FRCEM)
- Previous and Current employers
- Mandatory training compliance
- Occupational Health clearance

However for the above to become a reality work would need to be done by the various organisations to provide and APIs exposing the relevant information - a brief overview of how that might work would be an entire series of posts on its own. 

**But suffice to say given the existence of central authorities (GMC, NMC, Colleges etc.) and inherent trust between organisations (NHS organisations implicitly trust the GMC and Colleges) and only single organisations updating their relevant records (An NHS trust isn't going to update the Royal College of Physicians Exam Records to show that a Doctor has passed their MRCP Part 1) there is _absolutely no need_ for a distributed, eventually consistent ledger that was originally conceived to enforce trust between inherently untrustworthy systems.**

Therefore effort would be better spent extending the role of the [Care Identity Service](https://digital.nhs.uk/services/registration-authorities-and-smartcards/care-identity-service) and ironing out some of the implementation issues to improve user experience and ease of use.


<blockquote class="twitter-tweet" data-conversation="none" data-lang="en"><p lang="en" dir="ltr">In conclusion: it should not be normal for NHS IT to be so bad. It&#39;s 2018. <br>Fix the basics and stop getting distracted by the latest fads. <a href="https://twitter.com/bengoldacre">@bengoldacre</a> <a href="https://twitter.com/MattHancock">@MattHancock</a></p>&mdash; Dan Sheldon (@sheldonline) <a href="https://twitter.com/sheldonline/status/1027825755936509952">August 10, 2018</a></blockquote>


### Conflict of Interest Statement

I work in the NHS and pay Tax in the UK. Therefore I have vested interest in: 1) money no being wasted on the latest technological fad 2) the rapid improvement of technology in the NHS.

I welcome any feedback, comments, criticism or suggestions on what I have outlined. If you can explain the need for a novel, pointless, expensive and difficult technology to me I will happily update this post with more information as it become available.


<script async src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>