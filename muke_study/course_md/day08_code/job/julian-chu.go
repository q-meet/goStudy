package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Page struct {
	FileName string
	Body     string
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	urls := map[string]string{
		"https://baidu.com":  "baidu.html",
		"https://google.com": "google.html",
		"https://bing.com":   "bing.html",
	}

	fetcher := NewFetcher()

	for url := range urls {
		filename := urls[url]
		go fetcher.DoHttpGet(ctx, url, filename)
	}
	page := fetcher.FirstResult()
	cancel()
	fmt.Println(page)
	filename := fmt.Sprintf("/tmp/%s", page.FileName)
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.WriteString(page.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

}

type Fetcher struct {
	client *http.Client
	ch     chan Page
}

func NewFetcher() *Fetcher {
	ch := make(chan Page, 0)
	client := &http.Client{}
	return &Fetcher{client, ch}
}

func (f Fetcher) FirstResult() Page {
	return <-f.ch
}

func (f Fetcher) DoHttpGet(ctx context.Context, url, filename string) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	resp, err := f.client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("got http status code: %d from %s\n", resp.StatusCode, req.URL)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	p := Page{
		FileName: filename,
		Body:     string(body),
	}
	select {
	case f.ch <- p:
	case <-ctx.Done():
	}
}