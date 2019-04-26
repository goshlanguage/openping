package ping

import "time"

// Poll polls, is a helper for our cmd client
func Poll(s Store, url string) (rc int, latency time.Duration, err error) {
	request, err := GetRequest(url)
	if err != nil {
		return 0, 0, err
	}
	doc, rc, latency, err := GetDocument(request)
	if err != nil {
		return 0, 0, err
	}
	s.Update(url, rc, latency, doc)
	return rc, latency, nil
}
