package ingest

type CustomDomainQueryString struct {
	data map[string]map[string]struct{}
}

func NewCustomDomainQueryString(d map[string][]string) *CustomDomainQueryString {
	c := &CustomDomainQueryString{
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

func (d *CustomDomainQueryString) Set(domain string, key string) (string, bool) {
	keys, ok := d.data[domain]
	if !ok {
		return "", false
	}
	keys[key] = struct{}{}
	return key, true
}

func (d *CustomDomainQueryString) Get(domain string) map[string]struct{} {
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
