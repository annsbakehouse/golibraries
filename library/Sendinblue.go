package library

import (
	"context"
	"fmt"
	"os"

	emailverifier "github.com/AfterShip/email-verifier"
	sendinblue "github.com/sendinblue/APIv3-go-library/v2/lib"
)

func SIBEmailVerify(email string) bool {
	email := "example@exampledomain.org"

	ret, err := emailverifier.Verify(email)
	if err != nil {
		fmt.Println("verify email address failed, error is: ", err)
		return
	}
	if !ret.Syntax.Valid {
		fmt.Println("email address syntax is invalid")
		return
	}

	fmt.Println("email validation result", ret)
}

func SIBSendEmail() {
	var ctx context.Context
	cfg := sendinblue.NewConfiguration()
	//Configure API key authorization: api-key
	cfg.AddDefaultHeader("api-key", os.Getenv("SIB_api_key"))
	//Configure API key authorization: partner-key
	cfg.AddDefaultHeader("partner-key", os.Getenv("SIB_api_key"))

	sib := sendinblue.NewAPIClient(cfg)
	result, resp, err := sib.AccountApi.GetAccount(ctx)
	if err != nil {
		fmt.Println("Error when calling AccountApi->get_account: ", err.Error())
		return
	}
	fmt.Println("GetAccount Object:", result, " GetAccount Response: ", resp)
	return
}
