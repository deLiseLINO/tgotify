package telegram

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

// Client represents a client for interacting with the Telegram Bot API.
type Client struct {
	host     string
	basePath string
	client   http.Client
}

const (
	sendMessageMethod = "sendMessage"
	updateMethod      = "getUpdates"
)

func New(host string) *Client {
	return &Client{
		host:   host,
		client: http.Client{},
	}
}

// SendMessage sends a message to a specified chat ID.
func (c *Client) SendMessage(token string, chatID uint, message string) error {
	c.setToken(token)
	q := url.Values{}
	q.Add("chat_id", strconv.FormatUint(uint64(chatID), 10))
	q.Add("text", message)

	_, err := c.doRequest(sendMessageMethod, q)
	if err != nil {
		return fmt.Errorf("unable to send message: %w", err)
	}

	return nil
}

// Update retrieves updates from the Telegram Bot API.
// func (c *Client) Update() ([]Update, error) {
// 	body, err := c.doRequest(updateMethod, url.Values{})

// 	if err != nil {
// 		return nil, fmt.Errorf("unable to get updates: %w", err)
// 	}

// 	var res UpdatesResponse

// 	if err := json.Unmarshal(body, &res); err != nil {
// 		return nil, fmt.Errorf("unable to get updated: %w", err)
// 	}

// 	return res.Updates, nil
// }

// doRequest performs an HTTP request to the Telegram API.
func (c *Client) doRequest(method string, query url.Values) (data []byte, err error) {
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
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

func (c *Client) setToken(token string) {
	c.basePath = newBasePath(token)
}
