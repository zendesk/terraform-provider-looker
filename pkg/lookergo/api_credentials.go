package lookergo

import (
	"context"
	"fmt"
)

const apiCredentialsBasePath = "4.0/users"

type ApiCredential struct {
	ID           string `json:"id"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret,omitempty"`
	Type         string `json:"type"`
	IsDisabled   bool   `json:"is_disabled"`
	URL          string `json:"url"`
}

type ApiCredentialsResource interface {
	Get(ctx context.Context, userID int, credentialsID string) (*ApiCredential, *Response, error)
	Create(ctx context.Context, userID int, newCredential *ApiCredential) (*ApiCredential, *Response, error)
}

type ApiCredentialsResourceOp struct {
	client *Client
}

var _ ApiCredentialsResource = &ApiCredentialsResourceOp{}

// Get fetches a specific API credential for a user
func (s *ApiCredentialsResourceOp) Get(ctx context.Context, userID int, credentialsID string) (*ApiCredential, *Response, error) {
	path := fmt.Sprintf("%s/%d/credentials_api3/%s", apiCredentialsBasePath, userID, credentialsID)
	return doGet(ctx, s.client, path, new(ApiCredential))
}

// Create creates a new API credential for a user
func (s *ApiCredentialsResourceOp) Create(ctx context.Context, userID int, newCredential *ApiCredential) (*ApiCredential, *Response, error) {
	path := fmt.Sprintf("%s/%d/credentials_api3", apiCredentialsBasePath, userID)
	return doCreate(ctx, s.client, path, newCredential, new(ApiCredential))
}