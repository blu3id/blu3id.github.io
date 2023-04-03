---
layout: default
title: "How-to: Gain Work/Life Balance on your Android Phone"
permalink: "/posts/how-to-work-life-phone"
date: 2021-03-28
tag: "how-to"
---

# How-to: Gain Work/Life Balance on your Android Phone

In an ideal world your employer would provide you with a "Work Phone" so you can easily and physically separate your work from your personal life by turning off/physically discarding the device. Unfortunately one of the largest employers[^1] in the UK doesn't routinely offer this for a large proportion of their workforce. This results in work inevitably encroaching on personal time either through unwanted phone calls, group messages, pleading e-mails and other distractions.

Recent (and old) discussions on Twitter have once again highlighted the different 'hacks' that people have developed to gain some sort of separation. This has finally prompted me to finally writeup the different components/options/evolution of my solution to this problem over the years.

Fair warning: The easiest and simplest solution with the best experience is to just get a separate "Work Phone" what follows is a 'hack' to avoid this you have been warned.

## How-to Guide
{:.no_toc}

{: class="warning"}
> **First and most importantly a disclaimer:**
>
> This is not an idiot's-guide with detailed step-by-step instructions and as such requires some savvy, intuition and ability to seek out manuals / guides / documentation. _**No support or help beyond this post will be provided.**_
>
> Obviously no liability is accepted for damage, expense, data-loss, missed calls etc. etc. caused by following these ideas. **Do so entirely at your own risk**.

- Table of Contents
{:toc}

The requirements we are trying to achieve are to effectively merge two phones into one so we have the following functionality:
- Two phone numbers (so the second can be used for work)
- The ability to make and receive calls from both numbers
- The ability to send and receive SMS text messages from both numbers
- The ability to use OTT messaging Apps such as WhatsApp with from both numbers
- The ability to "Turn Off" the second number and not receive calls, texts or messages from Apps associated with it


### Step 1. Acquire another Phone Number

There are two options here each with different pros. and cons. and unfortunately only one provides "full" functionality as described above.

#### Option A: Dual SIM

This is the preferred option although it requires a phone that supports Dual SIM cards. Fortunately you may already have one and not realise. There are two types of Dual SIM phone: Two physical SIM cards OR [eSIM](https://en.wikipedia.org/wiki/ESIM) and SIM card (more common).

In the UK there are only a limited number of networks that provide an eSIM. This means you may need to think carefully about what network to chose (particularly as there are no networks that provide an eSIM on PAYG plans currently). You may need to switch your current plan to an eSIM so you can use the physical SIM for your new number.

{: class="box"}
> Pros.
> - Relatively simple to set-up (particularly if you have a phone with two physical SIM cards)
> - May already have a Dual SIM phone (through the magic of eSIM)
> - Ability to have a separate account for work related calls so can clearly identify expenses (may be tax deducible)
> - Full functionality of phone number (calls, SMS text messages)
> 
> Cons.
> - May have need a new phone
> - Depending on set-up may need to engage in shenanigans to switch too/from an eSIM

#### Option B: Use a SIP number (e.g Sipgate)

The alternative (and admittedly slightly more complex to set-up) is to use a VOIP provider to get a [SIP](https://en.wikipedia.org/wiki/Session_Initiation_Protocol) number you can use to make and receive calls over WiFi and/or Mobile Data.

The breakthrough here was provided to me in this wonderful [blog post](https://shkspr.mobi/blog/2020/07/adding-sip-calls-to-android-for-free/) by [Terence Eden](https://twitter.com/edent) introducing [SIPGate Basic](https://www.sipgatebasic.co.uk/) who provide a free UK PAYG SIP number (most other providers usualy charge line-rental).

{: class="box"}
> **If you decide to go down this route there a number of important points to note:**
>
> Cons.
> - Your new number will **NOT** be a "mobile" number and will **NOT** be in the form 07xxx xxx xxx
> - You will need to either be on WiFi or use Mobile Data to make and receive calls through this number
> - You won't be able to receive or send* traditional SMS Text Messages with this number (it isn't a "mobile" number)
> - Terence Eden's [blog post](https://shkspr.mobi/blog/2020/07/adding-sip-calls-to-android-for-free/) describes the steps to forward your existing number through SIP to achieve "poor mans" WiFi Calling so don't follow all the instructions
> - Creating a SIPGate account is not instant and requires physical physical receipt of a letter, you can also only have one account per postal address
> - Calling plans/costs are not as competitive as traditional providers if you are only making a few calls
> - You will have to take care when verifying the number with WhatsApp (or other OTT messaging apps) and use a phone call to verify rather than SMS
> 
> Pros.
> - Can receive calls over Wifi (may be useful in buildings with poor phone reception)
> - Voicemail messages can be emailed to you
> - Ability to have a separate account for work related calls so can clearly identify expenses (may be tax deducible)


### Step 2. The Magic (getting WhatsApp to work)

There are some options here that depend on the manufacture of your phone. There are different pros. and cons. depending on how you proceed.

#### Option A: Manufacture provided

Samsung has the [Dual Messenger](https://www.samsung.com/uk/support/mobile-devices/how-do-i-set-up-dual-messenger/) feature on some of their phones. Most OnePlus phones have the slightly more flexible [Parallel Apps](https://www.oneplus.com/uk/support/answer/detail/op98) feature.

{: class="box"}
> Pros.
> - Relatively easy to set-up
> - Supported, tested, bug-free and known to work
> 
> Cons.
> - Requires manufacture support
> - No easy way to "turn off" the app when not it work
> - Can't be used alongside a [MDM](https://en.wikipedia.org/wiki/Mobile_device_management) system that uses a _Work Profile_

#### Option B: Shelter

[Shelter](https://web.archive.org/web/20201125020741/https://play.google.com/store/apps/details?id=net.typeblog.shelter){:data-originalurl="https://play.google.com/store/apps/details?id=net.typeblog.shelter" data-versiondate="2020-11-25"} is an App that enables you to use the "Work Profile" feature of phones running Android 7 "Nougat" and above (phone released after late 2016) to create isolated copies of other Apps on your phone.

It is important to note that you can **NOT** run Shelter on a phone that already has a _Work Profile_ enabled (or on Samsung/OnePlus phones that are _using_ their manufacture provided solution as above).

{: class="box"}
> Pros.
> - Should work on most Android phones (provided they are running Android 7 or above)
> - Can be used on most Apps so you can put e.g work email, Teams etc. inside the _Work Profile_ (provided said email account isn't using an [MDM](https://en.wikipedia.org/wiki/Mobile_device_management) system that uses a _Work Profile_)
> - On most devices you literally have an "off" switch to [suspend](https://support.google.com/work/android/answer/7029561?hl=en) the _Work Profile_
> 
> Cons.
> - May not be supported/work with your phone
> - Can't be used alongside a [MDM](https://en.wikipedia.org/wiki/Mobile_device_management) system that uses a _Work Profile_
> - Not "idiot proof" to set-up 


### Step 3.

There is no step 3. Unless you count trying to erase all trace of your personal number at work and replace it with you new "work" phone number.

Other tasks to consider also include: 
- Getting you Union to lobby your employer to provide you with a "work" phone (or option of a work phone) like the majority of employers in the UK so you don't have to go through all this.
- Contemplating just buying a separate cheap "smart" phone for work rather than trying to piece together the suggestions in this post.


### Known Issues

- With all the options presented above there is no ready-built ridiculously easy "off" switch for the second phone number. With both Dual SIM and SIP you need to open phone settings and disable the SIM or SIP account. That said it is perfectly possible to use e.g [Tasker](https://play.google.com/store/apps/details?id=net.dinglisch.android.taskerm) and [plugins](https://play.google.com/store/apps/details?id=com.joaomgcd.taskersettings) to create a shortcut to do this. But I will leave that as an exercise to the reader.
- Re-read the cons. listed for each of the options you have picked.


[^1]: The NHS, which is also one of the [largest employers in the world](https://en.wikipedia.org/wiki/List_of_largest_employers). Though admittedly responsibility of "Work Phones" etc. is really that of the individual employing NHS organisations as "The NHS" isn't really an institution but that is a much larger footnote...