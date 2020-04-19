package bootstrap

// Events returns the channel of progress messages
func (c *Client) Events() chan interface{} {
	return c.events
}
