package pipedrive

var (
	// Client is the Pipedrive client being tested.
	client *Client
)

func init() {
	config := &Config{
		ApiKey: "bc5b30cb07ac9572597b427c1767ab650eef03ef",
	}

	client = New(config)
}
