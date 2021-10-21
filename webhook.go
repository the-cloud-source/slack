package slack

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type WebHook struct {
	hookURL string

	hook  string
	proxy string
	host  string
}

type WebHookPostPayload struct {
	Text        string        `json:"text,omitempty"`
	Channel     string        `json:"channel,omitempty"`
	Username    string        `json:"username,omitempty"`
	IconUrl     string        `json:"icon_url,omitempty"`
	IconEmoji   string        `json:"icon_emoji,omitempty"`
	UnfurlLinks bool          `json:"unfurl_links,omitempty"`
	LinkNames   string        `json:"link_names,omitempty"`
	Attachments []*Attachment `json:"attachments,omitempty"`
}

func NewWebHook(hookURL string) *WebHook {
	return &WebHook{hookURL}
}

func NewWebHookProxy(hookUrl, proxy string) (*WebHook, error) {

	h, err := url.Parse(hookUrl)
	if err != nil {
		return nil, err
	}

	p, err := url.Parse(hookUrl)
	if err != nil {
		return nil, err
	}

	proxied = h
	proxied.Scheme = p.Scheme
	proxied.Host = p.Host

	wh := &WebHook{
		hookURL: proxied.String(),
		hook:    hookUrl,
		proxy:   proxy,
		host:    h.Hostname(),
	}

	wh.hookURL = proxied.String()
	return wh, nil
}

func (hk *WebHook) PostMessage(payload *WebHookPostPayload) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err

	}

	req, err := http.NewRequest("POST", URL, bytes.NewReader(body))
	req.Host = hk.Host
	req.Header.Add("User-Agent", "go-slack/v1")
	req.Header.Add("Content-Type", "application/json")

	resp, err := httpClientWH.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t, _ := ioutil.ReadAll(resp.Body)
		return errors.New(string(t))
	}

	return nil
}

var httpClientWH = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	},
	Timeout: 60 * time.Second,
}
