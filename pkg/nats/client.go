package nats

// Client is a client for the NATS messaging system.
type Client struct {
	// ...
}

// NewClient creates a new Client.
func NewClient() *Client {
	return &Client{
		// ...
	}
}

// Publish publishes a message.
func (c *Client) Publish(subject string, body []byte) error {
	// ...
	return nil
}
