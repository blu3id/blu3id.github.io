---
layout: default
title: "Journal: Right-half-plane zeros"
permalink: "/posts/journal-right-half-plane-zeros"
date: 2022-06-23
tag: "journal"
---

# Journal: Right-half-plane zeros

> - [You need to know what right-half-plane zeros are](https://jbconsulting.substack.com/p/you-need-to-know-what-right-half) by [Jacob Bayless](https://www.jbaylessconsulting.ca/)

This is the introductory article in a series using systems-based theory as a lens to view and interpret macro economic data and narratives (which is not itself [new](https://en.wikipedia.org/wiki/Stock-flow_consistent_model)). The remainder of the series looking an economics is less interesting than this rather good introduction/overview of _control theory in dynamic systems_.

Understanding _control theory in dynamic systems_ provides important insights needed to safely work with tightly coupled complex adaptive system. And as I have said before the NHS and health and care in general are [tightly coupled complex adaptive systems](complex-system-dangers).

It is not possible to properly summarise the already heavily summarised introduction to modelling in control theory in this article so go read it. Particularly the footnote that explains the differential equation _dy/dt = dx/dt + Ax_ as an example of system modelling (and how calculus has a real world use outside the classroom). However the titular concept is worth picking out:

> To put it as simply as possible, when subjected to an input that ought to make it move up, a system with an RHP \[right-half-plane\] zero will move down first, before moving up. This is known as an inverse response.

alongside this wonderful quote:

> A system with a right-half-plane zero is one type of non-minimum-phase system, which is a code-word for “trouble”.

and the explanation of how this can cause harm:

> The danger of the right-half-plane zero is that it lures you into reacting to it, but that is precisely the wrong thing to do. Attempting to apply a new control input to cancel the inverse response only sets off an even worse chain of events, where the resulting secondary inverse reaction becomes even more severe, requiring even more corrective action, until finally you’ve slammed into the ground.

### Right-half-plane zeros in health and care

Digital transformation - particularly when not done well - is the perfect example. E.g the "big bang" introduction of a mega suite BigEHR™ will likely initially result in negative effects and costs rather than the theoretically promised benefits and savings.

The risk that right-half-plane zeros model is twofold:

1. As the article states - that there may be temptation to react and change things that need to settle out ("wait and see" in the case of a mega suite BigEHR™ is probably the wrong course of action) and/or that the rate at which these corrective changes can be made is limited to the timescale of the right-half-plane zero. Which is a limitation of the system itself - effectively baking in:

2. The subsequent drop - negative effects and costs - which will likely be greater than the system can cope with (i.e harm will come to patients, staff etc.) before it starts to recover. Expanding on this point would be series of posts. 

In most digital change projects there is an unrecognised[^1] poor understanding of risk and some deep-rooted (false) preconceptions about the differences in risk with e.g:
 - a "big bang" vs. gradual roll out of change.
 - an established supplier / product vs. a new market entrant without existing deployments.

These biases frequently taint specifications, procurement processes, and project planning leading to increased risk, opportunity cost, and sub-par outcomes.

[^1]: See the Dunning–Kruger cognitive bias, unknown unknowns, and why you should employ a Human Factors professional rather than someone with experience or training in Human Factors.