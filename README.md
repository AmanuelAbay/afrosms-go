# github.com/amanuelabay/afro-sms

# Usage

## Initialize your module

`go get init example.com/my-afrosms-demo`

## Get the afrosms-go module

Note that you need to include the *v* in the version tag.

`go get github.com/amanuelabay/afro-sms@v1.0.0`

```go
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
```

Check out the [afro-sms documentation](https://www.afromessage.com/developers) for more information