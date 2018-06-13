package nu

import "net/url"

// URL builds a *url.URL using the provided scheme, host, path and queryPairs.
func URL(scheme, host, path string, queryPairs ...string) *url.URL {
	return &url.URL{
		Scheme:   scheme,
		Host:     host,
		Path:     path,
		RawQuery: URLValues(queryPairs...).Encode(),
	}
}
// URLValues populates and returns a url.Values using the provided string key/value pairs.
// Much like strings.NewReplacer, it panics if len(pairs) is not even.
// See https://golang.org/pkg/strings/#NewReplacer
func URLValues(keyValuePairs ...string) url.Values {
	if len(keyValuePairs)%2 != 0 {
		panic("Must provide an even number of key/value strings.")
	}
	values := make(url.Values, len(keyValuePairs)/2)
	for x := 0; x < len(keyValuePairs); x += 2 {
		key := keyValuePairs[x]
		value := keyValuePairs[x+1]
		values.Set(key, value)
	}
	return values
}
