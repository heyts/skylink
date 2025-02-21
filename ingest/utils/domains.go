package utils

// FIXME: Replace with a DomainResolver struct
// that does the normalization and resolution in
// one go:
// dr := DomainResolver(domainsMap)
// url := dr.ResolveURL("https://google.com?f=3432")
//
// - createCanonicalURL
// - normalizeCanonicalURL
// - resolveCanonicalURL
type DomainResolver struct {
	data map[string]map[string]struct{}
}

func NewDomainResolver(d map[string][]string) *DomainResolver {
	c := &DomainResolver{
		data: map[string]map[string]struct{}{},
	}

	for domain, keys := range d {
		c.data[domain] = map[string]struct{}{}
		for _, key := range keys {
			c.data[domain][key] = struct{}{}
		}
	}
	return c
}

func (d *DomainResolver) Set(domain string, key string) (string, bool) {
	keys, ok := d.data[domain]
	if !ok {
		return "", false
	}
	keys[key] = struct{}{}
	return key, true
}

func (d *DomainResolver) Get(domain string) map[string]struct{} {
	ret := map[string]struct{}{}
	keys, ok := d.data[domain]
	if !ok {
		return ret
	}

	for k, _ := range keys {
		ret[k] = struct{}{}
	}

	return ret
}

func (d *DomainResolver) Resolve(url string) string {
	return ""
}
