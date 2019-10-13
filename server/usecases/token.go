package usecases

import (
	"fmt"
	"encoding/json"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/developer-kikikaikai/githubapi/server/data"
)

type PostToken struct {
	Code string `json:"code"`
	ID string `json:"client_id"`
	Secret string `json:"client_secret"`
}

type Result struct {
	Token string `json:"access_token"`
}

func GenerateToken(code string) (string,error) {
	body := PostToken{
		Code: code,
		ID: data.GetClientKey(),
		Secret: data.GetClientSecret(),
	}
	//fmt.Printf("Code:%s ID:%s Secret:%s\n", body.Code, body.ID, body.Secret)
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&body).
		SetResult(&Result{}).
		Post("https://github.com/login/oauth/access_token")
	if err != nil || resp.StatusCode() != http.StatusOK {
		fmt.Printf("Call https://github.com/login/oauth/access_token error\n")
		return "", err
	}

	result := Result{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		fmt.Printf("Parse response error %s\n", err.Error())
		return "", err
	}

	return result.Token, nil
} 
