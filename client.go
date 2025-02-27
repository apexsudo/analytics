package analytics

import (
	"github.com/rudderlabs/analytics-go/v4"
)

type Client interface {
	Identify(userID string, traits map[string]any) error
	Track(userID string, event string, properties map[string]any) error
	Close() error
}

type analyticsClient struct {
	rudderStackClient analytics.Client
}

func (a *analyticsClient) Close() error {
	return a.rudderStackClient.Close()
}

func (a *analyticsClient) Identify(userID string, traits map[string]any) error {
	return a.rudderStackClient.Enqueue(analytics.Identify{
		UserId: userID,
		Traits: traits,
	})
}

func (a *analyticsClient) Track(userID string, event string, properties map[string]any) error {
	return a.rudderStackClient.Enqueue(analytics.Track{
		UserId:     userID,
		Event:      event,
		Properties: properties,
	})
}

func New(writeKey string, dataPlaneURL string) Client {
	client := analytics.New(writeKey, dataPlaneURL)

	return &analyticsClient{rudderStackClient: client}
}
