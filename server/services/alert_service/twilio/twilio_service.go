package twilio

import (
	"main/server/request"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

var TwilioClient *twilio.RestClient

func TwilioInit(password string) {
	TwilioClient = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: os.Getenv("TWILIO_ACCOUNT_SID"),
		Password: password,
	})
}

func SendOtpService(ctx *gin.Context, req request.TwilioSmsRequest) (bool, *string) {
	params := &openapi.CreateVerificationParams{}

	params.SetTo(req.Contact)

	params.SetChannel("sms")

	resp, err := TwilioClient.VerifyV2.CreateVerification(os.Getenv("VERIFY_SERVICE_SID"), params)

	if err != nil {
		return false, nil
	} else {
		return true, resp.Sid
	}
}

func VerifyOtpService(ctx *gin.Context, number string, otp string) (string, error) {
	params := &openapi.CreateVerificationCheckParams{}

	params.SetTo("+91" + number)

	params.SetCode(otp)

	resp, err := TwilioClient.VerifyV2.CreateVerificationCheck(os.Getenv("VERIFY_SERVICE_SID"), params)

	if err != nil {
		return "", err
	} else if *resp.Status == "approved" {
		return *resp.Status, err
	}

	return "", nil

}

//function to send normal sms

// func SendSms(ctx *gin.Context, contact string) {

// 	client := twilio.NewRestClient()

// 	params := &api.CreateMessageParams{}
// 	params.SetFrom("+15557771212")
// 	params.SetBody("Ahoy! This message was sent from my Twilio phone number!")
// 	params.SetTo("+15559991111")

// 	resp, err := client.Api.CreateMessage(params)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	} else {
// 		if resp.Body != nil {
// 			fmt.Println(*resp.Body)
// 		} else {
// 			fmt.Println(resp.Body)
// 		}
// 	}
// }
