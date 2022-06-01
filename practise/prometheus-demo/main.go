package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/twindagger/httpsig.v1"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	JmsServerURL    = "http://jump.voneyun.com"
	AccessKeyID     = "2ddbc068-50f5-4f83-b771-a03044e7f448"
	AccessKeySecret = "8c110db6-bad0-4a74-8352-5616c084648b"
)

type SigAuth struct {
	KeyID    string
	SecretID string
}

type Body struct {
	Count    int64
	Next     string
	Previous string
	Results  []Rusults
}

type Rusults struct {
	Id       string
	Hostname string
	Ip       string
}

func (auth *SigAuth) Sign(r *http.Request) error {
	headers := []string{"(request-target)", "date"}
	signer, err := httpsig.NewRequestSigner(auth.KeyID, auth.SecretID, "hmac-sha256")
	if err != nil {
		return err
	}
	return signer.SignRequest(r, headers, nil)
}

func GetUserInfo(jmsurl string, auth *SigAuth) (*Body, error) {
	url := jmsurl + "/api/v1/perms/users/nodes/3f797c52-4f8a-4548-aad8-505c65156467/assets/?cache_policy=1&offset=0&limit=2&display=1&draw=1"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-JMS-ORG", "00000000-0000-0000-0000-000000000002")
	if err != nil {
		log.Fatal(err)
	}
	if err := auth.Sign(req); err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var bodys Body
	if err := json.Unmarshal(body, &bodys); err != nil {
		return nil, err
	}
	return &bodys, nil

}

func main() {
	auth := SigAuth{
		KeyID:    AccessKeyID,
		SecretID: AccessKeySecret,
	}
	bodys, err := GetUserInfo(JmsServerURL, &auth)
	if err != nil {
		panic(err)
	}

	for _, v := range bodys.Results {
		fmt.Println(v.Id)
	}
}
