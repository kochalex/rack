package client

import (
	"fmt"
	"time"
)

type Certificate struct {
	Id         string    `json:"id"`
	Domain     string    `json:"domain"`
	Expiration time.Time `json:"expiration"`
}

type Certificates []Certificate

func (c *Client) CreateCertificate(pub, key, chain string) (*Certificate, error) {
	var cert Certificate

	params := Params{
		"pub":   pub,
		"key":   key,
		"chain": chain,
	}

	err := c.Post("/certificates", params, &cert)

	if err != nil {
		return nil, err
	}

	return &cert, nil
}

func (c *Client) DeleteCertificate(id string) error {
	return c.Delete(fmt.Sprintf("/certificates/%s", id), nil)
}

func (c *Client) ListCertificates() (Certificates, error) {
	var certs Certificates

	err := c.Get("/certificates", &certs)

	if err != nil {
		return nil, err
	}

	return certs, nil
}
