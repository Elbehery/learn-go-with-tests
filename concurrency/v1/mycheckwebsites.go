package concurrency

// MyWebsiteChecker checks a url, returns a bool.
type MyWebsiteChecker func(string) bool

// MyCheckWebsites takes a MyWebsiteChecker and a slice of URLs,
// return a map of uri and the result of checking it.
func MyCheckWebsites(checker MyWebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = MyCheckWebsite(url)
	}

	return results
}
