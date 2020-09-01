package gofor

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

var GoFor *Gofor

type Gofor struct {
	Hosts []string
	ListenPort string
}

func New(h []string, l string) *Gofor {
	GoFor = &Gofor{
		Hosts: h,
		ListenPort: l,
	}
	return GoFor
}

func (g *Gofor) Get(u *url.URL) ([]byte, error) {
	url, err := url.Parse(u.String())
	if err != nil {
		return nil, err
	}
	r := make(chan []byte)
	for _, h := range g.Hosts {
		host := h
		go func() {
			parsedHost, err := url.Parse(host)
			if err != nil {
				return
			}
			url.Scheme = parsedHost.Scheme
			url.Host = parsedHost.Host
			resp, err := http.Get(url.String())
			defer resp.Body.Close()
			contents, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return
			}
			r <- contents
		}()

	}
	i := <- r
	return i, nil
}
