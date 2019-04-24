package ping

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

// WebHookConfig stores information for sending WebHooks
type WebHookConfig struct {
	WebhookURL string
	Username   string
	IconEmoji  string
	IconURL    string
}

// Alert will send the given message
func (wh *WebHookConfig) Alert(message string) (err error) {
	jsonStr := fmt.Sprintf(`{"text":"%s","username":"%s","icon_emoji":"%s"}`, message, wh.Username, wh.IconEmoji)
	log.Printf("Webhook alert fired: %v", string(jsonStr))
	request, err := http.NewRequest("POST", wh.WebhookURL, bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		panic(err)
	}
	// defer resp.Body.Close()
	return nil
}
