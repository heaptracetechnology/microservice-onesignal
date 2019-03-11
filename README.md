# OneSignal as a microservice
An OMG service for OneSignal, it allows to send message to the players(users) who have subscribed for the application on onesignal.

[![Open Microservice Guide](https://img.shields.io/badge/OMG-enabled-brightgreen.svg?style=for-the-badge)](https://microservice.guide)
[![Build Status](https://travis-ci.org/heaptracetechnology/microservice-onesignal.svg?branch=master)](https://travis-ci.org/heaptracetechnology/microservice-onesignal)
[![codecov](https://codecov.io/gh/heaptracetechnology/microservice-onesignal/branch/master/graph/badge.svg)](https://codecov.io/gh/heaptracetechnology/microservice-onesignal)

## [OMG](hhttps://microservice.guide) CLI

### OMG

* omg validate
```
omg validate
```
* omg build
```
omg build
```
### Test Service

* Test the service by following OMG commands

### CLI

##### List Application
```sh
$ omg run list -e USER_KEY=<USER_AUTH_KEY> -e APP_KEY=<APP_KEY>
```

##### Send Message
```sh
$ omg run send -a app_Id=<APPLICATION_ID> -a headings=<HEADING> -a contents=<CONTENTS> -a isAnyWeb=<BOOLEAN> -a include_player_ids=<PLAYERS_ID>
```
## License
### [MIT](https://choosealicense.com/licenses/mit/)

## Docker
### Build
```
docker build -t microservice-onesignal
```
### RUN
```
docker run -p 3000:3000 microservice-onesignal
```
