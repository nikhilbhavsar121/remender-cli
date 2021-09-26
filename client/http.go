package client

import (
	"net/http"
	"strings"
	"time"
)

type HTTPClient struct {
	Client     *http.Client
	BackendURI string
}

func NewHTTPClient(uri string) HTTPClient {
	return HTTPClient{
		BackendURI: uri,
		Client:     &http.Client{},
	}
}

// Create calls the create API endpoint
func (c HTTPClient) Create(title, message string, duration time.Duration) ([]byte, error) {
	res := []byte{}
	return res, nil
}

// Edit calls the edit API endpoint
func (c HTTPClient) Edit(id string, title, message string, duration time.Duration) ([]byte, error) {
	res := []byte{`response for edit`}
	return res, nil
}

// Fetch calls the fetch API endpoint
func (c HTTPClient) Fetch(ids []string) ([]byte, error) {
	idsSet := strings.Join(ids, ",")
	return c.apiCall(
		http.MethodGet,
		"/reminders/"+idsSet,
		nil,
		http.StatusOK,
	)
}

// Delete calls the delete API endpoint
func (c HTTPClient) Delete(ids []string) error {
	idsSet := strings.Join(ids, ",")
	_, err := c.apiCall(
		http.MethodDelete,
		"/reminders/"+idsSet,
		nil,
		http.StatusNoContent,
	)
	return err
}

// Healthy checks whether a given host is up and running
func (c HTTPClient) Healthy(host string) bool {
	res, err := http.Get(host + "/health")
	if err != nil || res.StatusCode != http.StatusOK {
		return false
	}
	return true
}
