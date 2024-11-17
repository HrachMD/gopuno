package fcmclient

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

type Client struct {
	*messaging.Client
}

func New(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	app, err := firebase.NewApp(ctx, nil, opts...)
	if err != nil {
		return nil, err
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		return nil, err
	}

	return &Client{Client: client}, nil
}

// Google Application Credentials (service-account-file.json)
func WithCreds(path string) option.ClientOption {
	return option.WithCredentialsFile(path)
}
