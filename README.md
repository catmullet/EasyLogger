![EasyLogger](https://raw.githubusercontent.com/catmullet/EasyLogger/master/EasyLogger_Banner.png)
[![Go Report Card](https://goreportcard.com/badge/github.com/catmullet/EasyLogger)](https://goreportcard.com/report/github.com/catmullet/EasyLogger)

# _Simple, useful Golang logs on ElasticSearch_
## Features
* Automatic capture of useful debug info
* Captures HTTP Request and Response Body
* Captures the Caller Func
* Posts to ElasticSearch
* Easily search logs in Kibana

## Getting Started
#### Run Go Get
```json
go get github.com/catmullet/EasyLogger
```
#### Add your environment variables
##### Set The Following Environment Variables
```
ES_HOST={{ElasticSearch Instance URL}}
APP_NAME={{Applications Identity (ex.EasyLogger)}}
```

##### Important 
Set your ElasticSearch instance to only accept connections from your Apps IP or modify to meet your security needs.
