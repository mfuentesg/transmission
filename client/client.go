package client

import (
	"errors"

	"github.com/mfuentesg/transmission/torrent"
)

type Option func(*options)

type client struct {
	options
}

type options struct {
	User     string
	Password string
	URL      string
}

func WithURL(url string) Option {
	return func(o *options) {
		o.URL = url
	}
}

func WithBasicAuth(user, password string) Option {
	return func(o *options) {
		o.User = user
		o.Password = password
	}
}

func New(opts ...Option) *client {
	defaults := options{}
	for _, o := range opts {
		o(&defaults)
	}
	return &client{options: defaults}
}

func (c client) Get(id int64, fields []string) (torrent.Torrent, error) {
	var empty torrent.Torrent
	if len(fields) == 0 {
		return empty, errors.New("request must includes at least one field")
	}

	response, err := torrent.DoRequest(&torrent.Request{
		URL: c.URL,
		Auth: torrent.RequestAuth{
			Username: c.User,
			Password: c.Password,
		},
		Payload: torrent.RequestPayload{
			Method:    torrent.MethodGet,
			Arguments: torrent.RequestArgs{Ids: []int64{id}, Fields: fields},
		},
	})

	if err != nil {
		return empty, err
	}

	if response.Result == torrent.ResponseResultError {
		return empty, errors.New("unable to get results from transmission service")
	}

	for _, t := range response.Args.Torrents {
		if t.Id == id {
			return t, nil
		}
	}

	return empty, errors.New("torrent not found")
}

func (c client) GetAll(fields []string) ([]torrent.Torrent, error) {
	var empty []torrent.Torrent
	if len(fields) == 0 {
		return empty, errors.New("request must includes at least one field")
	}

	response, err := torrent.DoRequest(&torrent.Request{
		URL: c.URL,
		Auth: torrent.RequestAuth{
			Username: c.User,
			Password: c.Password,
		},
		Payload: torrent.RequestPayload{
			Method:    torrent.MethodGet,
			Arguments: torrent.RequestArgs{Fields: fields},
		},
	})

	if err != nil {
		return empty, err
	}

	if response.Result == torrent.ResponseResultError {
		return empty, errors.New("unable to get results from transmission service")
	}

	return response.Args.Torrents, nil
}
