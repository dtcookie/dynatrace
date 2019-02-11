package users

import (
	"encoding/json"
	"errors"

	resterrors "github.com/dtcookie/dynatrace/apis/errors"
	"github.com/dtcookie/dynatrace/rest"
)

// API is able to make REST API Calls to the Users API of an
// OnPremise Dynatrace Cluster
type API struct {
	client *rest.Client
}

// NewAPI creates a preconfigured API for accessing the Users API
// of an OnPremise Dynatrace Cluster
func NewAPI(config *rest.Config, credentials *rest.Credentials) *API {
	return &API{client: rest.NewClient(config, credentials)}
}

// GetUsers queries for the currently configured users
func (api *API) GetUsers() ([]UserConfig, error) {
	var err error
	var bytes []byte

	if bytes, err = api.client.GET("/api/v1.0/onpremise/users", 200); err != nil {
		return nil, resterrors.Resolve(bytes, err)
	}
	var response []UserConfig
	if err = json.Unmarshal(bytes, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// Create TODO: documentation
func (api *API) Create(config *UserConfig) error {
	return errors.New("Not Implemented")
}
