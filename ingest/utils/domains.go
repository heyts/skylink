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

const MAX_REDIRECTS = 10

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
	if strings.HasPrefix(key, "!!") {
		return false
	}
	if keys, ok := d.items[domain]; ok {
		if _, ok = keys[key]; ok {
			return true
		}
	}
	return false
}

func (d *DomainResolver) HasDirective(domain, key string) bool {
	if !strings.HasPrefix(key, "!!") {
		return false
	}

	if keys, ok := d.items[domain]; ok {
		if _, ok = keys[key]; ok {
			return true
		}
	}
	return false
}

func (d *DomainResolver) Resolve(raw string) (string, error) {
	url, err := url.Parse(raw)
	if err != nil {
		return "", err
	}

	if !url.IsAbs() {
		return "", fmt.Errorf("cannot resolve relative URL: %s", url)
	}

	url, err = d.resolveURLRecursive(url)
	if err != nil {
		return raw, err
	}

	return url.String(), nil
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

func (d *DomainResolver) resolveURLRecursive(url *url.URL) (*url.URL, error) {
	redirectCount := 0

	for {

		if redirectCount > MAX_REDIRECTS {
			return nil, fmt.Errorf("too many redirects")
		}
		resp, err := d.client.Get(url.String())
		if err != nil {
			return nil, err
		}

		if resp.StatusCode == 200 {
			url, err = d.normalizeURL(url)
			if err != nil {
				return nil, err
			}
			return url, nil
		}

		if resp.StatusCode > 300 || resp.StatusCode < 308 {
			curr := resp.Header.Get("Location")
			if curr == "" {
				return nil, fmt.Errorf("malformed location Header: %s", curr)
			}
			url, err = url.Parse(curr)
			if err != nil {
				return nil, err
			}
			redirectCount++
		}
	}
}

func MD5Encode(URL string) string {
	algorithm := md5.New()
	algorithm.Write([]byte(URL))
	return hex.EncodeToString(algorithm.Sum(nil))
}
