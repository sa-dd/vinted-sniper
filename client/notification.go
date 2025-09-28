package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Notification struct {
	Id            string `json:"id"`
	UserId        int64  `json:"user_id"`
	SubjectId     int64  `json:"subject_id"`
	IsRead        bool   `json:"is_read"`
	UpdatedAt     string `json:"updated_at"`
	EntryType     int    `json:"entry_type"`
	Body          string `json:"body"`
	Link          string `json:"link"`
	PhotoType     string `json:"photo_type"`
	SmallPhotoUrl string `json:"small_photo_url"`
}

type NotificationResponse struct {
	Notifications []Notification
}

func FetchNotifications(client *http.Client) ([]Notification, error) {
	const url = "https://www.vinted.co.uk/api/v2/notifications?page=1&per_page=5"

	body, err := Get(url)

	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	var response NotificationResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("JSON unmarshal failed: %w", err)
	}

	return response.Notifications, nil
}
