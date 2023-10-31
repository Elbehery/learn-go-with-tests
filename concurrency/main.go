package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(websiteChecker WebsiteChecker, urls []string) map[string]bool {
	res := map[string]bool{}
	results := make(chan result)

	for _, url := range urls {
		go func(u string) {
			results <- result{
				string: u,
				bool:   websiteChecker(u),
			}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-results
		res[r.string] = r.bool
	}

	return res
}
