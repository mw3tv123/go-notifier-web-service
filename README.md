# go-notifier-web-service
> An all-round notifier build in Go as a web hooker for other tools in terms of 3rd party notifier.

[![Build Status](https://app.travis-ci.com/mw3tv123/go-notifier-web-service.svg?branch=main)](https://app.travis-ci.com/mw3tv123/go-notifier-web-service)
[![codebeat badge](https://codebeat.co/badges/8ad64015-9459-4c8e-ad69-a5531972a966)](https://codebeat.co/projects/github-com-mw3tv123-go-notifier-web-service-main)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/d5b8c91c34594512aff383129239d4d4)](https://www.codacy.com/gh/mw3tv123/go-notifier-web-service/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=mw3tv123/go-notifier-web-service&amp;utm_campaign=Badge_Grade)
[![Maintainability](https://api.codeclimate.com/v1/badges/531ffd1ac5852a2bfe56/maintainability)](https://codeclimate.com/github/mw3tv123/go-notifier-web-service/maintainability)
    ![Go Version](https://img.shields.io/badge/Go%20version-v1.16-blue)
![License](https://img.shields.io/badge/License-GLP%203.0-blue)

## Installation

```shell
# Pull dependencies
$ make deps
$ go build 
$ go run .
2021/09/11 21:36:50 Listening and serving HTTP on 0.0.0.0:80
```

## Configuration
All configurations are stored in `config/app.env`. Besides, configurations can be overrided from export environment.

Beside default message struct, you can change the template used to send message. Reference to following for each channel template:
- [MS Team](https://docs.microsoft.com/en-us/microsoftteams/platform/webhooks-and-connectors/how-to/connectors-using?tabs=cURL#example-of-connector-message): all message send to MS Teams must use legacy `MessageCard`, detail can be found in link;

## Usage example
Current developed API:
- `/health`
- `/notify/ms_teams`
- `/alert/ms_teams`

```shell
$ # For simple notify message
$ curl http://localhost:8090/notify/ms_teams -d '{ "title": "test-title", "content": "test content" }' -H 'Content-Type: application/json

$ # For alert notify message
$ curl http://localhost:8090/alert/ms_teams -d '{ "title": "test", "priority": 1, "service_name": "monitor a", "description": "Alert test a", "create_date": "2018-09-22T12:42:31+07:00" }' -H 'Content-Type: application/json'
```
- ***Health*** --- `/health`, return HTTP code 200 when application ready to serve API


- ***Notify*** --- Group API `/notyfy/`, request body with 2 params:
  + `title` - Title of message (`required,min=3,max=100`)
  + `content` - Content of message (`required,min=3,max=200`)


- ***Alert*** --- Group API `/alert/`, request body with 5 params:
  + `title` - Title of alert (`required,min=3,max=100`)
  + `service_name` - Service trigger alert (`required,min=3,max=100`)
  + `description` - Alert description (`can empty,max=100`)
  + `priority` - Or in another word: `critical level` (`required,min=0,max=10`)
  + `create_date` - Date the alert create, which is control by request client (`can empty`)
