# _OneSignal_ OMG Microservice

[![Open Microservice Guide](https://img.shields.io/badge/OMG-enabled-brightgreen.svg?style=for-the-badge)](https://microservice.guide)
[![Build Status](https://travis-ci.org/omg-services/onesignal.svg?branch=master)](https://travis-ci.org/omg-services/onesignal)
[![codecov](https://codecov.io/gh/omg-services/onesignal/branch/master/graph/badge.svg)](https://codecov.io/gh/omg-services/onesignal)


An OMG service for OneSignal, it allows to send message to the players(users) who have subscribed for the application on onesignal.

## Direct usage in [Storyscript](https://storyscript.io/):

##### List Application
```coffee
>>> onesignal list
["List all apps"]
```
##### Send Message
```coffee
>>> onesignal send app_Id:'appId' headings:'headings' contents:'contents' isAnyWeb:'true/false' include_player_ids:'includePlayerIds'
{"success": "true","message": "Notification sent","statusCode": 200}
```

Curious to [learn more](https://docs.storyscript.io/)?

✨🍰✨

## Usage with [OMG CLI](https://www.npmjs.com/package/omg)

##### List Application
```shell
$ omg run list -e USER_KEY=<USER_AUTH_KEY> -e APP_KEY=<APP_KEY>
```

##### Send Message
```shell
$ omg run send -a app_Id=<APPLICATION_ID> -a headings=<HEADING> -a contents=<CONTENTS> -a isAnyWeb=<BOOLEAN> -a include_player_ids=<PLAYERS_ID>
```

**Note**: The OMG CLI requires [Docker](https://docs.docker.com/install/) to be installed.

## License
[MIT License](https://github.com/omg-services/onesignal/blob/master/LICENSE).
