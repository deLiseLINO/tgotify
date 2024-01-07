package telegram

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
	models "tgotify/storage"

	"github.com/sirupsen/logrus"
)

type TokensGetter interface {
	EnabledClients() ([]models.EnabledClientsResponse, error)
}

// Client represents a client for interacting with the Telegram Bot API.
type Client struct {
	host      string
	client    http.Client
	tgClients []models.EnabledClientsResponse
	tokensSvc TokensGetter
}

const (
	sendMessageMethod = "sendMessage"
	updateMethod      = "getUpdates"
)

func New(
	host string,
	tgClients []models.EnabledClientsResponse,
	tokensSvc TokensGetter,
) *Client {
	return &Client{
		host:      host,
		client:    http.Client{},
		tgClients: tgClients,
		tokensSvc: tokensSvc,
	}
}

// SendMessage sends a message to a specified chat ID.
func (c *Client) SendMessage(token string, chatID uint, message string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.FormatUint(uint64(chatID), 10))
	q.Add("text", message)

	basePath := c.setToken(token)
	_, err := c.doRequest(sendMessageMethod, q, basePath)
	if err != nil {
		return fmt.Errorf("unable to send message: %w", err)
	}

	return nil
}

// Update retrieves updates from the Telegram Bot API.
func (c *Client) Updates(offset int, limit int) []Update {
	var res []Update
	for _, client := range c.tgClients {
		updates, err := c.update(offset, limit, client)
		if err != nil {
			logrus.Error("Unable to get updates from token: ", client)
			continue
		}
		res = append(res, updates...)
	}
	return res
}

func (c *Client) UpdateClients() error {
	clients, err := c.tokensSvc.EnabledClients()
	c.tgClients = clients
	return err
}

func (c *Client) update(
	offset int,
	limit int,
	tgClient models.EnabledClientsResponse,
) ([]Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	basePath := c.setToken(tgClient.Token)
	body, err := c.doRequest(updateMethod, q, basePath)

	if err != nil {
		return nil, fmt.Errorf("unable to get updates: %w", err)
	}

	var res UpdatesResponse

	if err := json.Unmarshal(body, &res); err != nil {
		return nil, fmt.Errorf("unable to get updated: %w", err)
	}

	for i := range res.Updates {
		res.Updates[i].ClientToken = tgClient.Token
		res.Updates[i].ClientID = tgClient.ID
	}

	return res.Updates, nil
}

// doRequest performs an HTTP request to the Telegram API.
func (c *Client) doRequest(method string, query url.Values, basePath string) (data []byte, err error) {
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("unable to perform the request: %w", err)
	}

	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to perform the request: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to perform the request: %w", err)
	}

	return body, nil
}

// newBasePath constructs the base path for the Telegram client.
func newBasePath(token string) string {
	return "bot" + token
}

func (c *Client) setToken(token string) string {
	return newBasePath(token)
}
