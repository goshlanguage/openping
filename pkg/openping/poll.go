package ping

// Poll polls, is a helper for our cmd client
func Poll(s Store, url string) (rc int, err error) {
	request, err := GetRequest(url)
	if err != nil {
		return 0, err
	}
	doc, rc, err := GetDocument(request)
	if err != nil {
		return 0, err
	}
	s.Update(url, doc)
	return 0, nil
}
