package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func main() {
	// Define the custom user agent string
	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"

	// Create a new allocator context with the custom user agent
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserAgent(userAgent),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// Create a new Chrome context using the allocator context
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	// Navigate to the target web page and select the HTML elements of interest
	var nodes []*cdp.Node
	if err := chromedp.Run(ctx,
		chromedp.Navigate("https://asnlookup.com/asn/AS15169"),
		chromedp.Nodes(".table tbody tr", &nodes, chromedp.ByQueryAll),
	); err != nil {
		log.Fatalf("Failed to fetch nodes: %v", err)
	}

	// Print the table data
	for _, node := range nodes {
		var rowData []string
		for _, child := range node.Children {
			if child.NodeName == "TH" || child.NodeName == "TD" {
				rowData = append(rowData, child.NodeValue)
			}
		}
		fmt.Println(rowData)
	}
}
