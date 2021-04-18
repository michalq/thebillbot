package coinbase

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const BaseUrl = "https://api.coinbase.com"

type Client struct {
	apiKey    string
	apiSecret string
}

func NewClient(apiKey, apiSecret string) *Client {
	return &Client{apiKey, apiSecret}
}

func (c *Client) request(ctx context.Context, method, endpoint string, responseBucket interface{}) error {

	url := fmt.Sprintf("%s%s", BaseUrl, endpoint)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}
	now := strconv.FormatInt(time.Now().Unix(), 10)
	req = req.WithContext(ctx)
	req.Header.Add("CB-ACCESS-KEY", c.apiKey)
	req.Header.Add("CB-ACCESS-SIGN", c.createAccessSign(now, method, url, ""))
	req.Header.Add("CB-ACCESS-TIMESTAMP", now)
	httpClient := http.DefaultClient
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	if responseBucket == nil {
		return nil
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bodyBytes, responseBucket); err != nil {
		return err
	}
	return nil
}

func (c *Client) createAccessSign(timestamp, method, requestPath, body string) string {

	key := []byte(c.apiSecret)
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(timestamp + method + requestPath + body))
	sign := mac.Sum(nil)
	signBase64 := make([]byte, base64.StdEncoding.EncodedLen(len(sign)))
	base64.StdEncoding.Encode(signBase64, sign)
	fmt.Println(string(signBase64))
	return string(signBase64)
}
