package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var disableRedirectsFunc = func(req *http.Request, via []*http.Request) error {
	return http.ErrUseLastResponse
}

type DomainResolver struct {
	items  map[string]map[string]struct{}
	client *http.Client
}

func NewDomainResolver(d map[string][]string) *DomainResolver {
	c := &DomainResolver{
		items: map[string]map[string]struct{}{},
		client: &http.Client{
			CheckRedirect: disableRedirectsFunc,
			Timeout:       2 * time.Second,
		},
	}

	for domain, keys := range d {
		c.items[domain] = map[string]struct{}{}
		for _, key := range keys {
			c.items[domain][key] = struct{}{}
		}
	}
	return c
}

func (d *DomainResolver) Set(domain string, key string) (string, bool) {
	keys, ok := d.items[domain]
	if !ok {
		return "", false
	}
	keys[key] = struct{}{}
	return key, true
}

func (d *DomainResolver) HasKey(domain, key string) bool {
	if keys, ok := d.items[domain]; ok {
		if _, ok = keys[key]; ok {
			return true
		}
	}
	return false
}

func (d *DomainResolver) Resolve(raw string) (string, error) {
	url, err := d.parseURL(raw, true)
	if err != nil {
		return raw, err
	}

	url, err = d.normalizeURL(url)
	if err != nil {
		return raw, err
	}

	url, err = d.resolveURL(url)
	if err != nil {
		return raw, err
	}

	return url.String(), nil
}

func (d *DomainResolver) parseURL(raw string, resolve bool) (*url.URL, error) {
	url, err := url.Parse(raw)
	if err != nil {
		return nil, err
	}

	if !url.IsAbs() {
		return nil, fmt.Errorf("Cannot Resolve a non-absolute URL")
	}

	url, err = d.normalizeURL(url)
	if err != nil {
		return nil, err
	}

	if resolve {
		url, err = d.resolveURL(url)
		if err != nil {
			return nil, err
		}
	}
	return url, nil
}

func (d *DomainResolver) normalizeURL(url *url.URL) (*url.URL, error) {
	hostname := url.Hostname()
	h := strings.Split(url.Hostname(), ".")

	if len(h) >= 2 {
		hostname = strings.Join(h[len(h)-2:], ".")
	}

	query := url.Query()

	for key := range query {
		if !d.HasKey(hostname, key) {
			query.Del(key)
		}
	}

	url.RawQuery = query.Encode()
	url.Fragment = ""

	return url, nil
}

func (d *DomainResolver) resolveURL(url *url.URL) (*url.URL, error) {
	resp, err := d.client.Head(url.String())
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		return url, nil
	}

	if resp.StatusCode == 301 ||
		resp.StatusCode == 302 ||
		resp.StatusCode == 303 ||
		resp.StatusCode == 304 ||
		resp.StatusCode == 305 ||
		resp.StatusCode == 306 ||
		resp.StatusCode == 307 ||
		resp.StatusCode == 308 {
		u := resp.Header.Get("Location")
		url, err = d.parseURL(u, false)
		if err != nil {
			return nil, err
		}
	}
	return url, nil
}

func MD5Encode(URL string) string {
	algorithm := md5.New()
	algorithm.Write([]byte(URL))
	return hex.EncodeToString(algorithm.Sum(nil))
}
