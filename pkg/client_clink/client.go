package clientclink

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/davesavic/clink"
)

type ClientC struct {
	client *clink.Client
}

type Body struct {
	Token string `json:"yandexPassportOauthToken"`
}

type YandexToken struct {
	Expires    string `json:"expiresAt"`
	BasicToken string `json:"iamToken"`
}

func (cl *ClientC) Clink(url string, token string) (YandexToken, error) {
	cl.client = clink.NewClient()
	jsonBody := prepareBody(token)
	resp, err := cl.client.Post(url, jsonBody)
	if err != nil {
		log.Fatalf("request error: %s", err)
		return YandexToken{}, err
	}

	var target YandexToken
	err = clink.ResponseToJson(resp, &target)
	return target, err
}

func prepareBody(token string) *bytes.Reader {
	body := &Body{
		Token: token,
	}
	marshalled, err := json.Marshal(body)
	if err != nil {
		log.Fatalf("marshalling error: %s", err)
	}
	return bytes.NewReader(marshalled)
}
