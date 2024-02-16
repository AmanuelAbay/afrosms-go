![Logo](https://www.afromessage.com/assets/images/logo.png) 

# AfroSms

This library allows you to quickly and easily use the afro-sms Web API via Go.

## Badges  
[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)  

# github.com/amanuelabay/afro-sms


# Usage

## Initialize your module

~~~bash
go mod init example.com/my-afrosms-demo
~~~

## Environment Variables  

To run this project, you will need to add the following environment variables to your .env file  

`AFRO_SMS_API_KEY`  

`AFRO_SMS_SENDER_NAME` 

`AFRO_SMS_RECEIVER_PHONE_NUMBER` 

## Get the afrosms-go module

Note that you need to include the **v** in the version tag.

~~~bash
go get github.com/amanuelabay/afrosms-go
~~~


~~~go
package main

import (
	"fmt"
	afro "github.com/amanuelabay/afrosms-go"
	"os"
        "fmt"
)


func main(){
    
    AFRO_SMS_API_KEY := os.Getenv("AFRO_SMS_API_KEY")
    AFRO_SMS_SENDER_NAME := os.Getenv("AFRO_SMS_SENDER_NAME")
    AFRO_SMS_RECEIVER_PHONE_NUMBER := os.Getenv("AFRO_SMS_RECEIVER_PHONE_NUMBER")
    HOST := "https://api.afromessage.com"
    END_POINT := "/api/send"
    
    request := afro.GetRequest(API_KEY, END_POINT, HOST)
    request.Method = "GET"
    request.Sender(AFRO_SMS_SENDER_NAME)
    Message := "Hey, from afro sms"
    request.To(AFRO_SMS_RECEIVER_PHONE_NUMBER, Message)

	response, err := afro.Api(request)

    if response["acknowledge"] == "success" {
        fmt.Println("message_successfully_sent!")
        fmt.Println(response["response"])
    }else {
        fmt.Println("unable_to_send_sms")
        fmt.Println(response["response"])

    }
}
~~~


## Tech Stack  
**Client:** Go 


## Authors  
- [@Amanuel Abay](https://www.github.com/amanuelabay) 

## Documentation  
[Documentation](https://www.afromessage.com/developers)