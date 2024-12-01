package utils

import (
	"bytes"
	"owlint/models"

	"encoding/json"
	"log"
	"net/http"
	"time"
)

func SendCommentNotification(comment models.Comment) {
	url := "http://tech-test-back.owlint.fr:8080/on_comment"

	message := comment.TextEn
	if message == "" {
		message = comment.TextFr
	}

	payload := map[string]string{
		"author":  comment.AuthorID,
		"message": message,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Println("error while marshaling payload:", err)
		return
	}

	// the service may be unstable, so we retry in case of failure
	maxRetries := 3
	for i := 0; i < maxRetries; i++ {
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(payloadBytes))
		if err != nil {
			log.Println("error while sending notification:", err)
			time.Sleep(2 * time.Second)
			continue
		}
		resp.Body.Close()
		if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated {
			log.Println("notification sent successfully")
			return
		} else {
			log.Println("unexpected response from service:", resp.Status)
			time.Sleep(2 * time.Second)
		}
	}
	log.Printf("failed to send notification after [%d] retries", maxRetries)
}
