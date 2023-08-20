package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

const (
	EnvSlackWebhook   = "SLACK_WEBHOOK"
	EnvSlackChannel   = "SLACK_CHANNEL"
	EnvSlackUserName  = "SLACK_USERNAME"
	EnvSlackIcon      = "SLACK_ICON"
)

type Webhook struct {
	Text        string       `json:"text,omitempty"`
	UserName    string       `json:"username,omitempty"`
	IconURL     string       `json:"icon_url,omitempty"`
	IconEmoji   string       `json:"icon_emoji,omitempty"`
	Channel     string       `json:"channel,omitempty"`
	LinkNames   string       `json:"link_names,omitempty"`
	UnfurlLinks bool         `json:"unfurl_links"`
}

func main() {
	endpoint := os.Getenv(EnvSlackWebhook)
	if endpoint == "" {
		fmt.Fprintln(os.Stderr, "URL is required")
		os.Exit(1)
	}
	channel := os.Getenv(EnvSlackChannel)
	if channel == "" {
		fmt.Fprintln(os.Stderr, "Channel is required")
		os.Exit(1)
	}

	text := os.Getenv(EnvSlackIcon)
	text += "<https://github.com/"
	text += os.Getenv("GITHUB_REPOSITORY")
	text += "/commit/"
	text += os.Getenv("GITHUB_SHA")
	text += "/checks|["
	text += strings.Split(os.Getenv("GITHUB_REPOSITORY"), "/")[1]
	text += "|" + os.Getenv("GITHUB_REF_NAME")
	text += "|" + os.Getenv("GITHUB_SHA")[0:6]
	text += "]>"
	text += " - "
	text += os.Getenv("GITHUB_ACTOR")
	text += " (" + os.Getenv("GITHUB_WORKFLOW") + ")"
	text += " - "
	text += os.Getenv("COMMIT_MESSAGE")

	msg := Webhook{
		Text:	  text,
		UserName:  os.Getenv(EnvSlackUserName),
		Channel:   os.Getenv(EnvSlackChannel),
	}

	if err := send(endpoint, msg); err != nil {
		fmt.Fprintf(os.Stderr, "Error sending message: %s\n", err)
		os.Exit(2)
	}
}

func envOr(name, def string) string {
	if d, ok := os.LookupEnv(name); ok {
		return d
	}
	return def
}

func send(endpoint string, msg Webhook) error {
	enc, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	b := bytes.NewBuffer(enc)
	res, err := http.Post(endpoint, "application/json", b)
	if err != nil {
		return err
	}

	if res.StatusCode >= 299 {
		return fmt.Errorf("Error on message: %s\n", res.Status)
	}
	fmt.Println(res.Status)
	return nil
}
