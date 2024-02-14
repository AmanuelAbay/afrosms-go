![Logo](https://www.afromessage.com/assets/images/logo.png) 

# AfroSms

This library allows you to quickly and easily use the afro-sms Web API via Go.

## Badges  
[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)  

# github.com/amanuelabay/afro-sms


# Usage

## Initialize your module

~~~bash
go get init example.com/my-afrosms-demo
~~~

## Environment Variables  

To run this project, you will need to add the following environment variables to your .env file  

`AFRO_SMS_API_KEY`  

`AFRO_SMS_SENDER_NAME` 

## Get the afrosms-go module

Note that you need to include the **v** in the version tag.

~~~bash
go get github.com/amanuelabay/afro-sms@v1.0.0
~~~


~~~go
package main

import (
	"fmt"
	"github.com/amanuelabay/afrosms-go"
	"os"
        "fmt"
)

AFRO_SMS_API_KEY := os.Getenv("AFRO_SMS_API_KEY")
AFRO_SMS_SENDER_NAME := os.Getenv("AFRO_SMS_SENDER_NAME")
HOST := "https://api.afromessage.com"
END_POINT := "/api/send"

func main(){

    request := afro.GetRequest(API_KEY, END_POINT, HOST)
    request.Sender(AFRO_SMS_SENDER_NAME)
    Message := "Hey, from afro sms"
    request.To(AFRO_SMS_SENDER_NAME, Message)

    data := afro.Api(request)

    if data.Acknowledgement == "error" {
        fmt.Println(data.Response.Errors)
    }else {
        fmt.Println(data.Response.Status)
        fmt.Println(data.Response.MessageId)
        fmt.Println(data.Response.Message)
        fmt.Println(data.Response.To)

    }
}
~~~


## Tech Stack  
**Client:** Go 


## Authors  
- [@Amanuel Abay](https://www.github.com/amanuelabay) 

## Documentation  
[Documentation](https://www.afromessage.com/developers)