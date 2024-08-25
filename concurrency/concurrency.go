package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

// commented parts will create a race condition
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			// results[u] = wc(u)
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	// time.Sleep(2 * time.Second)
	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}
	return results
}

func mockWebsiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
}
