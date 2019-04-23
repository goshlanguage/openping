package ping

// Poll polls, is a helper for our cmd client
func Poll(s Store, url string) (err error) {
	request, err := GetRequest(url)
	if err != nil {
		return err
	}
	doc, err := GetDocument(request)
	if err != nil {
		return err
	}
	s.Update(url, doc)
	return nil
}
