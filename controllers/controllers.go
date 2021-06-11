package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	model "github.com/mktsy/go-webhook/models"
)

func HandlerMessenger(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		secretKey := os.Getenv("FACEBOOK_ACCESS_TOKEN")
		u, err := url.Parse(r.RequestURI)
		if err != nil {
			log.Printf("Failed parsing URL: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("An error occured: %s", err)))
			return
		}

		values, err := url.ParseQuery(u.RawQuery)
		if err != nil {
			log.Printf("Failed parsing value URL: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("An error occured: %s", err)))
			return
		}

		token := values.Get("hub.verify_token")
		if token == secretKey {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(values.Get("hub.challenge")))
			return
		}

		log.Printf("VALUES: %#v", values)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Token not found`))
		return
	}

	// Anything that reaches here in POST.
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Failed parsing body %s", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("An error occured"))
		return
	}

	// Parse message into the Message struct
	var message model.InputMessage
	err = json.Unmarshal(body, &message)
	if err != nil {
		log.Printf("Failed unmarshalling message: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("An error occured"))
		return
	}

	// Find messages
	for _, entry := range message.Entry {
		if len(entry.Messaging) == 0 {
			log.Printf("No messages")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("An error occured"))
			return
		}

		event := entry.Messaging[0]

		err = handleMessage(event.Sender.ID, event.Message.Text)
		if err != nil {
			log.Printf("Failed sending message: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("An error occurred"))
			return
		}
	}
}

func handleMessage(senderId, message string) error {
	if len(message) == 0 {
		return errors.New("No message found.")
	}
	response := model.ResponseMessage{
		Recipient: model.Recipient{
			ID: senderId,
		},
		Message: model.Message{
			Text: "Hello",
		},
	}
	data, err := json.Marshal(response)
	if err != nil {
		log.Printf("Marshal error: %s", err)
		return err
	}
	uri := "https://graph.facebook.com/v11.0/me/messages"
	uri = fmt.Sprintf("%s?access_token=%s", uri, os.Getenv("FACEBOOK_ACCESS_TOKEN"))
	log.Printf("URI: %s", uri)
	req, err := http.NewRequest(
		"POST",
		uri,
		bytes.NewBuffer(data),
	)
	if err != nil {
		log.Printf("Failed making request: %s", err)
		return err
	}
	req.Header.Add("Content-type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Failed doing request: %s", err)
		return err
	}
	log.Printf("MESSAGE SENT?\n%#v", res)
	return nil
}

// func handleAttachment(senderId, message string) error {
// 	if len(message) == 0 {
// 		return errors.New("No message found.")
// 	}

// 	response := model.ResponseAttachment{
// 		Recipient: model.Recipient{
// 			ID: senderId,
// 		},
// 		Message: model.Attachment{},
// 	}

// 	elements := []model.Element{
// 		model.Element{
// 			Title:    "Contact us.",
// 			ImageURL: "https://static.wixstatic.com/media/0f18c2_90117d17ff6c4d74901ba19be63dfcde~mv2.png/v1/fill/w_1424,h_729,al_c,q_90,usm_0.66_1.00_0.01/0f18c2_90117d17ff6c4d74901ba19be63dfcde~mv2.webp",
// 			Subtitle: "Botio service",
// 			DefaultAction: model.DefaultAction{
// 				Type:                "web_url",
// 				URL:                 "https://www.botio.services/",
// 				WebViewHeightRation: "tall",
// 			},
// 			Buttons: []model.Button{
// 				model.Button{
// 					Type:  "web_url",
// 					URL:   "https://www.botio.services/",
// 					Title: "Join us",
// 				},
// 			},
// 		},
// 	}
// 	response.Message.Attachment.Type = "template"
// 	response.Message.Attachment.Payload.TemplateType = "generic"
// 	response.Message.Attachment.Payload.Elements = elements

// 	data, err := json.Marshal(response)
// 	if err != nil {
// 		log.Printf("Marshal error %s", err)
// 		return err
// 	}
// 	log.Printf("DATA: %s", string(data))
// 	return sendRequest(data)
// }
