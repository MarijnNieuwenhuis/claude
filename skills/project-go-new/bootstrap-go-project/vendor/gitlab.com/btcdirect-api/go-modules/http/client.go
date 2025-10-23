package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"go.uber.org/zap"
)

const (
	DefaultAuthenticateEndpoint = "/token/authenticate"
	DefaultTokenExpireTime      = time.Hour - 20*time.Second
)

type AuthenticatedClient interface {
	BearerToken() (string, error)
	AddAuthorizationHeader(r *http.Request) error
	DoRequest(rc RequestConfig) error
}

type AuthenticatedClientConfig struct {
	BaseUrl              string
	AuthenticateEndpoint string
	Username             string
	Password             string
	TokenExpireTime      time.Duration
	Logger               *zap.SugaredLogger
}

type authenticatedClient struct {
	AuthenticatedClientConfig
	token bearerToken
}

type bearerToken struct {
	Token     string
	ExpiresAt time.Time
}

type RequestConfig struct {
	Method             string
	URL                string
	Data               any
	ExpectedStatusCode int
	Reader             io.Reader
}

func NewAuthenticatedClient(c AuthenticatedClientConfig) AuthenticatedClient {
	if c.AuthenticateEndpoint == "" {
		c.AuthenticateEndpoint = DefaultAuthenticateEndpoint
	}
	if c.TokenExpireTime == 0 {
		c.TokenExpireTime = DefaultTokenExpireTime
	}

	return &authenticatedClient{
		AuthenticatedClientConfig: c,
	}
}

func (c *authenticatedClient) BearerToken() (string, error) {
	if !c.token.Valid() {
		if err := c.authenticate(); err != nil {
			c.Logger.Errorw("Failed to obtain an authorization token", "error", err)
			return "", err
		}
	}

	return c.token.Token, nil
}

func (c *authenticatedClient) AddAuthorizationHeader(r *http.Request) error {
	token, err := c.BearerToken()
	if err != nil {
		return err
	}

	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	return nil
}

func (t bearerToken) Valid() bool {
	if t.Token == "" {
		return false
	}

	return t.ExpiresAt.After(time.Now())
}

func (c *authenticatedClient) authenticate() error {
	c.Logger.Info("Requesting an authorization token")

	body := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{c.Username, c.Password}

	js, err := json.Marshal(body)
	if err != nil {
		return err
	}

	r, err := http.NewRequest(http.MethodPost, c.BaseUrl+c.AuthenticateEndpoint, bytes.NewBuffer(js))
	if err != nil {
		return err
	}
	r.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("authentication failed: %s", res.Status)
	}

	defer res.Body.Close()

	token := struct {
		Token string `json:"token"`
	}{}
	if err := json.NewDecoder(res.Body).Decode(&token); err != nil {
		return err
	}

	c.Logger.Info("Successfully obtained an authorization token")

	c.token.Token = token.Token
	c.token.ExpiresAt = time.Now().Add(c.TokenExpireTime)

	return nil
}

func (c *authenticatedClient) DoRequest(rc RequestConfig) error {
	if rc.ExpectedStatusCode == 0 {
		if rc.Method == http.MethodPost || rc.Method == http.MethodPut {
			rc.ExpectedStatusCode = http.StatusCreated
		} else {
			rc.ExpectedStatusCode = http.StatusOK
		}
	}

	r, err := http.NewRequest(http.MethodGet, rc.URL, rc.Reader)
	if err != nil {
		return err
	}
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Accept", "application/json")

	err = c.AddAuthorizationHeader(r)
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return err
	}

	if res.StatusCode != rc.ExpectedStatusCode {
		return fmt.Errorf("request failed: %s", res.Status)
	}

	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(rc.Data); err != nil {
		return err
	}

	return nil
}
