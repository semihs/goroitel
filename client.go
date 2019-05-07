package goroitel

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RoitelClient struct {
	userName string
	password string
}

type request struct {
	Header   string  `json:"header"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Action   string  `json:"action"`
	Text     string  `json:"text"`
	Phones   []Phone `json:"phones"`
}

type Phone struct {
	Phone string `json:"phone"`
}

type response struct {
	Success bool `json:"success"`
}

func NewRoitelClient(userName, password string) *RoitelClient {
	return &RoitelClient{
		userName: userName,
		password: password,
	}
}

func (roitel *RoitelClient) SendSms(header, to, message string) error {
	return roitel.request(request{
		Header:   header,
		Username: roitel.userName,
		Password: roitel.password,
		Phones:   []Phone{{Phone: to}},
		Text:     message,
		Action:   "send_sms",
	})
}

func (roitel *RoitelClient) request(request request) error {
	body, err := json.Marshal(request)
	if err != nil {
		return err
	}
	fmt.Printf("roitel request body %s\n", string(body))

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("POST", "https://portal.roitel.com.tr/sms/api", bytes.NewBuffer([]byte(body)))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	responseBodyStr := string(responseBody)
	fmt.Printf("roitel response body %s\n", responseBodyStr)

	if resp.StatusCode != 200 {
		return fmt.Errorf("response status not ok, expected 200 given %d", resp.StatusCode)
	}

	response := response{}
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return err
	}

	return nil
}
