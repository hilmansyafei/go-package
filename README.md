# Canopus Go Package
Canopu Go package is general package used by all canopus microservice with golang

## Package Dependencies

Below is the packages used by this project

* Framework : https://github.com/labstack/echo
* Logging : 
    - https://github.com/sirupsen/logrus
    - https://github.com/lestrrat-go/file-rotatelogs
* ODM (mongodb) : https://github.com/zebresel-com/mongodm

## List Package in Canopus Go Package

1. database
    - mongodb
2. middleware
    - auth
3. modules
    - client
    - logger
    - mockServer
    - notification
    - rabbitMQ
4. response
    - notificationVa
    - response
5. status
    - log
    - status

## How to use package.

to use this package in your app just import the package you want
```
import (
    "github.com/sepulsa/canopus-gope/modules"
)
```
