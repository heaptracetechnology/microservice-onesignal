# _OneSignal_ Open Microservice

> This is a OneSignal service to push notification message.

[![Open Microservice Specification Version](https://img.shields.io/badge/Open%20Microservice-1.0-477bf3.svg)](https://openmicroservices.org) [![Open Microservices Spectrum Chat](https://withspectrum.github.io/badge/badge.svg)](https://spectrum.chat/open-microservices) [![Open Microservices Code of Conduct](https://img.shields.io/badge/Contributor%20Covenant-v1.4%20adopted-ff69b4.svg)](https://github.com/oms-services/.github/blob/master/CODE_OF_CONDUCT.md) [![Open Microservices Commitzen](https://img.shields.io/badge/commitizen-friendly-brightgreen.svg)](http://commitizen.github.io/cz-cli/) [![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com) 
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

## Introduction

This project is an example implementation of the [Open Microservice Specification](https://openmicroservices.org), a standard originally created at [Storyscript](https://storyscript.io) for building highly-portable "microservices" that expose the events, actions, and APIs inside containerized software.

## Getting Started

The `oms` command-line interface allows you to interact with Open Microservices. If you're interested in creating an Open Microservice the CLI also helps validate, test, and debug your `oms.yml` implementation!

See the [oms-cli](https://github.com/microservices/oms) project to learn more!

### Installation

```
npm install -g @microservices/oms
```

## Usage

### Open Microservices CLI Usage

Once you have the [oms-cli](https://github.com/microservices/oms) installed, you can run any of the following commands from within this project's root directory:

#### Actions

##### list

> List Applications
##### Action Arguments

| Argument Name | Type | Required | Default | Description |
|:------------- |:---- |:-------- |:--------|:----------- |
| USER_KEY | `string` | `false` | None | User key for onesignal. |
| APP_KEY | `string` | `false` | None | App key for onesignal. |

``` shell
oms run list \ 
    -e USER_KEY=$USER_KEY \ 
    -e APP_KEY=$APP_KEY
```

##### send

> Send message to the subscribed players.
##### Action Arguments

| Argument Name | Type | Required | Default | Description |
|:------------- |:---- |:-------- |:--------|:----------- |
| app_Id | `string` | `true` | None | Application ID of application created on onesignal. |
| headings | `map` | `false` | None | Heading of the message,Key as "en"(english) and Value. |
| contents | `map` | `false` | None | Content of the message,Key as "en"(english) and Value. |
| isIos | `boolean` | `false` | None | Defines device type (true/false). |
| isAndroid | `boolean` | `false` | None | Defines device type (true/false). |
| isWP | `boolean` | `false` | None | Defines it is web push or not (true/false). |
| isAdm | `boolean` | `false` | None | Defines it is Adm or not (true/false). |
| isChrome | `boolean` | `false` | None | Defines it is Chrome or not (true/false). |
| isChromeWeb | `boolean` | `false` | None | Defines it is Chrome web notification or not (true/false). |
| isAdmisSafari | `boolean` | `false` | None | Defines it is Safari browser or not (true/false). |
| data | `map` | `false` | None | Data of the message,Key as "en"(english) and Value. |
| isAnyWeb | `boolean` | `false` | None | Defines for any web push notification. |
| include_player_ids | `list` | `false` | None | Players Id of onesignal subscribers to whom notification is to be sent. |
| include_ios_tokens | `list` | `false` | None | List of IOS device token. |
| include_android_reg_ids | `list` | `false` | None | List of android registration IDs. |
| include_wp_uris | `list` | `false` | None | List of web push URIs. |
| include_wp_wns_uris | `list` | `false` | None | List of wns web push URis. |
| include_amazon_reg_ids | `list` | `false` | None | List of amazon registered registration Ids. |
| include_chrome_reg_ids | `list` | `false` | None | List of chrome registered Ids. |
| include_chrome_web_reg_ids | `list` | `false` | None | List of chrome web registered Ids. |
| app_ids | `list` | `false` | None | List of application ids where notification is to be sent. |
| tags | `map` | `false` | None | Tags for the notifications. |
| included_segments | `list` | `false` | None | List of the segments to be included to send notification. |
| excluded_segments | `list` | `false` | None | List of the segments to be excluded to send notification. |
| small_icon | `string` | `false` | None | Url of the icon. |
| large_icon | `string` | `false` | None | Url of the icon. |
| big_picture | `string` | `false` | None | Url of the icon. |
| chrome_icon | `string` | `false` | None | Url of the icon. |
| chrome_web_icon | `string` | `false` | None | Url of the icon. |
| ios_badgeType | `string` | `false` | None | Badge type of the IOS device. |
| ios_badgeCount | `string` | `false` | None | Badge count of the IOS device. |
| ios_sound | `string` | `false` | None | Notification sound of IOS device. |
| android_sound | `string` | `false` | None | Notification sound of android device. |
| adm_sound | `string` | `false` | None | ADM sound. |
| wp_sound | `string` | `false` | None | Web push notification sound. |
| wp_wns_sound | `string` | `false` | None | Web push wns notification sound. |
| send_after | `string` | `false` | None | Delay time to send message. |
| delayed_option | `string` | `false` | None | Delay time to send message. |
| delivery_time_of_day | `string` | `false` | None | Delivery time to send message. |
| android_led_color | `string` | `false` | None | Android notification LED colour. |
| android_accent_color | `string` | `false` | None | Android accent colour. |
| android_visibility | `int` | `false` | None | Android visibility. |
| android_background_data | `int` | `false` | None | Android background data. |
| filters | `map` | `false` | None | User filters to be added while sending message |
| USER_KEY | `string` | `false` | None | User key for onesignal. |
| APP_KEY | `string` | `false` | None | App key for onesignal. |

``` shell
oms run send \ 
    -a app_Id='*****' \ 
    -a headings='*****' \ 
    -a contents='*****' \ 
    -a isIos='*****' \ 
    -a isAndroid='*****' \ 
    -a isWP='*****' \ 
    -a isAdm='*****' \ 
    -a isChrome='*****' \ 
    -a isChromeWeb='*****' \ 
    -a isAdmisSafari='*****' \ 
    -a data='*****' \ 
    -a isAnyWeb='*****' \ 
    -a include_player_ids='*****' \ 
    -a include_ios_tokens='*****' \ 
    -a include_android_reg_ids='*****' \ 
    -a include_wp_uris='*****' \ 
    -a include_wp_wns_uris='*****' \ 
    -a include_amazon_reg_ids='*****' \ 
    -a include_chrome_reg_ids='*****' \ 
    -a include_chrome_web_reg_ids='*****' \ 
    -a app_ids='*****' \ 
    -a tags='*****' \ 
    -a included_segments='*****' \ 
    -a excluded_segments='*****' \ 
    -a small_icon='*****' \ 
    -a large_icon='*****' \ 
    -a big_picture='*****' \ 
    -a chrome_icon='*****' \ 
    -a chrome_web_icon='*****' \ 
    -a ios_badgeType='*****' \ 
    -a ios_badgeCount='*****' \ 
    -a ios_sound='*****' \ 
    -a android_sound='*****' \ 
    -a adm_sound='*****' \ 
    -a wp_sound='*****' \ 
    -a wp_wns_sound='*****' \ 
    -a send_after='*****' \ 
    -a delayed_option='*****' \ 
    -a delivery_time_of_day='*****' \ 
    -a android_led_color='*****' \ 
    -a android_accent_color='*****' \ 
    -a android_visibility='*****' \ 
    -a android_background_data='*****' \ 
    -a filters='*****' \ 
    -e USER_KEY=$USER_KEY \ 
    -e APP_KEY=$APP_KEY
```

## Contributing

All suggestions in how to improve the specification and this guide are very welcome. Feel free share your thoughts in the Issue tracker, or even better, fork the repository to implement your own ideas and submit a pull request.

[![Edit onesignal on CodeSandbox](https://codesandbox.io/static/img/play-codesandbox.svg)](https://codesandbox.io/s/github/oms-services/onesignal)

This project is guided by [Contributor Covenant](https://github.com/oms-services/.github/blob/master/CODE_OF_CONDUCT.md). Please read out full [Contribution Guidelines](https://github.com/oms-services/.github/blob/master/CONTRIBUTING.md).

## Additional Resources

* [Install the CLI](https://github.com/microservices/oms) - The OMS CLI helps developers create, test, validate, and build microservices.
* [Example OMS Services](https://github.com/oms-services) - Examples of OMS-compliant services written in a variety of languages.
* [Example Language Implementations](https://github.com/microservices) - Find tooling & language implementations in Node, Python, Scala, Java, Clojure.
* [Storyscript Hub](https://hub.storyscript.io) - A public registry of OMS services.
* [Community Chat](https://spectrum.chat/open-microservices) - Have ideas? Questions? Join us on Spectrum.
