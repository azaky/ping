package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func tick(interval time.Duration, do func(time.Time)) *time.Ticker {
	ticker := time.NewTicker(interval)
	go func() {
		for t := range ticker.C {
			do(t)
		}
	}()
	return ticker
}

func hit(url string) func(time.Time) {
	return func(t time.Time) {
		res, err := http.DefaultClient.Get(url)
		if err != nil {
			log.Printf("Error hitting [%s]: %s", url, err.Error())
		} else if res.StatusCode != http.StatusOK {
			log.Printf("Error hitting [%s]: %s", url, res.Status)
		} else {
			log.Printf("OK hitting [%s]", url)
		}
	}
}

func main() {
	interval, err := strconv.ParseInt(os.Getenv("I"), 10, 64)
	if err != nil {
		log.Fatalf("envvar I (interval) is required: %s", err.Error())
	}

	urls := strings.Split(os.Getenv("X"), ",")
	urls = append(urls, "http://127.0.0.1:"+os.Getenv("PORT"))

	var tickers []*time.Ticker
	for _, url := range urls {
		if len(url) == 0 {
			continue
		}

		ticker := tick(time.Duration(interval)*time.Millisecond, hit(url))
		tickers = append(tickers, ticker)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Print("/ was hit")
	})
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
