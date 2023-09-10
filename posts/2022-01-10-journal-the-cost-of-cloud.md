---
layout: default
title: "Journal: The Cost of Cloud"
permalink: "/posts/journal-the-cost-of-cloud"
date: 2022-01-10
tag: "journal"
---

# Journal: The Cost of Cloud

> - [The cost of cloud](https://ptribble.blogspot.com/2021/12/the-cost-of-cloud.html) by [Peter Tribble](http://www.petertribble.co.uk/)

A very sensible - and refreshing - article weighing the trade-offs of moving infrastructure[^1] into <abbr title="AKA someone else's computer">"the cloud"</abbr>.

- Roughly x3 unit cost in "the cloud" vs a competent in-house IT organisation
- [Egregious network costs for egress](https://blog.cloudflare.com/aws-egregious-egress/) traffic
- Need to carefully capacity plan (as with in-house infrastructure)
- Benefit to small scale where no existing team to make in-house operations worthwhile exists
- Variable workloads or a need for deployment globally that lend themselves to "the cloud".

It doesn't get into the detail of capital vs. operational budgets but highlights that weighing the above many organisations still save due to inefficient and overpriced traditional legacy infrastructure (making the argument for in-house infrastructure reform).

> If you're using commercial virtualization platforms and/or SAN storage, then you're overpaying by as much as a factor of 10, and getting an inferior product into the bargain - so that while many organizations could save a huge amount of money by moving to the cloud, they could save even more by running their internal operations better.

### Relevance to health and care

However the real take home message from the article is trade-offs aside using "the cloud" _enables organisations to reduce the time spent outside their core competencies_. This is where the fundamental ideological augment needs to be made in health and care - technology is a core competency.

This can be hard to understand surely the core competency is providing health and care? Supermarkets are a good example of a similar perception problem - their core competency is selling food surely? Wrong - they are specialised logistics companies (and like logistics companies entirely dependent on in-house technology to run their operations).

In-house capacity to build and maintain the digital tools needed to run a modern health and care organisation are essential for both the cost and agility benefits that are simply not possible within the current entrenched market. We must [fund teams not projects](https://defradigital.blog.gov.uk/2018/10/29/funding-product-teams-not-projects/) to maintain and grow this capacity.

[^1]: A distinction needs to be made here between "moving infrastructure to the cloud" and buying so-called "cloud hosted" software-as-a-service and calling this "moving to the cloud". This article discusses the former not the later.