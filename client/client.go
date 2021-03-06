package client

import (
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
)

// Client is the eventory service client.
type Client struct {
	*goaclient.Client
	KeySigner goaclient.Signer
	Encoder   *goa.HTTPEncoder
	Decoder   *goa.HTTPDecoder
}

// New instantiates the client.
func New(c goaclient.Doer) *Client {
	client := &Client{
		Client:  goaclient.New(c),
		Encoder: goa.NewHTTPEncoder(),
		Decoder: goa.NewHTTPDecoder(),
	}

	// Setup encoders and decoders
	client.Encoder.Register(goa.NewJSONEncoder, "application/json")
	client.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	client.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	client.Decoder.Register(goa.NewJSONDecoder, "application/json")
	client.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	client.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	client.Encoder.Register(goa.NewJSONEncoder, "*/*")
	client.Decoder.Register(goa.NewJSONDecoder, "*/*")

	return client
}

// SetKeySigner sets the request signer for the key security scheme.
func (c *Client) SetKeySigner(signer goaclient.Signer) {
	c.KeySigner = signer
}
