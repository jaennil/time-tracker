package service

import (
	"encoding/json"
	"errors"
	"github.com/jaennil/time-tracker/config"
	"github.com/jaennil/time-tracker/internal/model"
	"net/http"
	"net/url"
)

type UserAPI struct {
	config *config.Config
}

func NewUserAPI(config *config.Config) *UserAPI {
	return &UserAPI{config}
}

func (a *UserAPI) UserInfo(passportSerie string, passportNumber string) (*model.User, error) {
	apiURL, err := url.Parse(a.config.UserApiUrl)
	if err != nil {
		return nil, err
	}

	apiURL = apiURL.JoinPath("info")
	query := apiURL.Query()
	query.Set("passportSerie", passportSerie)
	query.Set("passportNumber", passportNumber)
	apiURL.RawQuery = query.Encode()

	response, err := http.Get(apiURL.String())
	if err != nil {
		return nil, err
	}
	// TODO: handle close error
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("bad request")
	}

	user := new(model.User)
	if err := json.NewDecoder(response.Body).Decode(user); err != nil {
		return nil, err
	}

	return user, nil
}
