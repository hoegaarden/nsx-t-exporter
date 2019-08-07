package exporter

import (
	"crypto/tls"
	"net/url"

	// "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	nsxv3config "github.com/sapcc/nsx-t-exporter/config"
)

/*
TODO
- Implement rate limit NSXv3Configuration.RequestsPerSecond
*/

const pathCreateSession = "/api/session/create"
const httpHeaderAcceptJSON = "application/json"
const httpHarderContentTypeJSON = "application/json"

// Nsxv3Client representes connection to NSXc3 Manger
type Nsxv3Client struct {
	config nsxv3config.NSXv3Configuration
	client http.Client
	cookie string
	token  string
}

// Nsxv3Response http.Response wrapper including the error and extracted response content bytes
type Nsxv3Response struct {
	path     string
	response *http.Response
	body     []byte
	err      error
}

// GetClient initialize NSXv3 http client
func GetClient(c nsxv3config.NSXv3Configuration) Nsxv3Client {
	if c.SuppressSslWornings {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	return Nsxv3Client{
		config: c,
		client: http.Client{
			Timeout: time.Second * 10,
		},
	}
}

// login to NSXv3 manager
func (c *Nsxv3Client) login(force bool) error {

	if !force && (c.cookie != "" || c.token != "") {
		return nil
	}

	requestBody := url.Values{}
	requestBody.Set("j_username", c.config.LoginUser)
	requestBody.Set("j_password", c.config.LoginPassword)

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("https://%s%s", c.config.LoginHost, pathCreateSession),
		strings.NewReader(requestBody.Encode()))

	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.client.Do(req)

	if err != nil {
		return err
	}

	c.cookie = resp.Header.Get("Set-Cookie")
	c.token = resp.Header.Get("X-XSRF-TOKEN")
	return nil
}

// AsyncGetRequest executes http get requests in an asych mode
func (c *Nsxv3Client) AsyncGetRequest(path string, ch chan<- *Nsxv3Response) error {
	c.login(false)

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("https://%s%s", c.config.LoginHost, path), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", httpHeaderAcceptJSON)
	req.Header.Set("Content-Type", httpHarderContentTypeJSON)
	req.Header.Set("Cookie", c.cookie)
	req.Header.Set("X-XSRF-TOKEN", c.token)

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	ch <- &Nsxv3Response{path, resp, bodyBytes, err}

	return err

}
