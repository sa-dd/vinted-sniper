package client

import (
	"net/url"
	"strings"
)

func ExtractOfferingID(notificationLink string) (string, error) {
	if notificationLink == "" {
		return "", nil
	}

	parsedURL, err := url.Parse(notificationLink)
	if err != nil {
		return "", err
	}

	offeringID := parsedURL.Query().Get("offering_id")
	offeringID = strings.TrimSpace(offeringID)

	return offeringID, nil
}
