package utils

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/net/html"
)

const MAX_REDIRECTS = 10

type LinkMeta map[string]string

func (m LinkMeta) Get(k string) string {
	return m[k]
}

func (m LinkMeta) GetOrDefault(k, def string) string {
	val, ok := m[k]
	if !ok {
		return def
	}
	return val
}

func (m LinkMeta) Set(k, v string) string {
	m[k] = v
	return v
}

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

func (d *DomainResolver) Resolve(raw string) (string, *LinkMeta, error) {
	url, err := url.Parse(raw)
	if err != nil {
		return "", nil, err
	}

	if !url.IsAbs() {
		return "", nil, fmt.Errorf("cannot resolve relative URL: %s", url)
	}

	url, meta, err := d.resolveURL(url)
	if err != nil {
		return raw, nil, err
	}

	return url.String(), meta, nil
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

func (d *DomainResolver) resolveURL(url *url.URL) (*url.URL, *LinkMeta, error) {
	redirectCount := 0
	var meta *LinkMeta

	for {
		if redirectCount > MAX_REDIRECTS {
			return nil, nil, fmt.Errorf("too many redirects")
		}
		resp, err := d.client.Get(url.String())
		if err != nil {
			return nil, nil, err
		}
		defer resp.Body.Close()

		s := resp.StatusCode

		switch {
		case s == http.StatusOK:
			meta, err = d.serializeMeta(resp.Body)
			if err != nil {
				return nil, nil, err
			}

			url, err = d.normalizeURL(url)
			if err != nil {
				return nil, nil, err
			}

			return url, meta, nil

		case s >= http.StatusMovedPermanently && s <= http.StatusPermanentRedirect:
			curr := resp.Header.Get("Location")
			if curr == "" {
				return nil, nil, fmt.Errorf("malformed \"location\" header: %q", curr)
			}
			url, err = url.Parse(curr)
			if err != nil {
				return nil, nil, err
			}
			redirectCount++

		default:
			return nil, nil, fmt.Errorf("unexpected http status code: %q, skipping", resp.Status)
		}
	}
}

func (d *DomainResolver) serializeMeta(r io.Reader) (*LinkMeta, error) {
	document, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	head, err := d.ParseHeadSection(document)
	if err != nil {
		return nil, err
	}

	title, err := d.ParseTitleTag(head)
	if err != nil {
		return nil, err
	}

	meta, err := d.ParseMeta(head.FirstChild)
	if err != nil {
		return nil, err
	}

	meta.Set("title", title)
	return meta, nil
}

func (d *DomainResolver) ParseHeadSection(doc *html.Node) (*html.Node, error) {
	var head *html.Node
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node == nil {
			return
		}
		if node.Type == html.ElementNode && node.Data == "head" {
			head = node
			return
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}

	crawler(doc)
	if head != nil {
		return head, nil
	}
	return nil, errors.New("no <head> element found in the document")
}

func (d *DomainResolver) ParseTitleTag(doc *html.Node) (string, error) {
	var title *html.Node
	var crawler func(*html.Node)

	crawler = func(node *html.Node) {
		if node == nil {
			return
		}

		if node.Type == html.ElementNode && node.Data == "title" {
			title = node
			return
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}

	crawler(doc)

	if title != nil && title.FirstChild != nil {
		return title.FirstChild.Data, nil
	}
	return "", errors.New("no <title> element found in the document")
}

func (d *DomainResolver) ParseMeta(doc *html.Node) (*LinkMeta, error) {
	attrs := &LinkMeta{}

	for doc != nil {
		attr := doc.Attr

		if len(attr) == 2 && attr[0].Key == "property" && strings.HasPrefix(attr[0].Val, "og:") {
			attrs.Set(attr[0].Val, attr[1].Val)
		}
		doc = doc.NextSibling
	}

	return attrs, nil
}

func MD5Encode(URL string) string {
	algorithm := md5.New()
	algorithm.Write([]byte(URL))
	return hex.EncodeToString(algorithm.Sum(nil))
}
