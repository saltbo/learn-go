package main

import (
	"context"
	"log"
	"strings"

	"github.com/chromedp/chromedp"
)

func main() {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var res1, res2 string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.etymonline.com/search?q=detection`),
		chromedp.Text(`section a:nth-child(1)`, &res1, chromedp.NodeReady),
		chromedp.Text(`section a:nth-child(2)`, &res2, chromedp.NodeReady),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(strings.TrimSpace(res1), res2)
}
