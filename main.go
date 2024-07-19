//package main
//
//import (
//	"context"
//	"encoding/json"
//	"fmt"
//	"log"
//	"os"
//	"strings"
//
//	"github.com/chromedp/cdproto/cdp"
//	"github.com/chromedp/chromedp"
//)
//
//type ASData struct {
//	ASHandle         string   `json:"as_handle"`
//	ASNName          string   `json:"asn_name"`
//	OrganizationName string   `json:"organization_name"`
//	OrganizationID   string   `json:"organization_id"`
//	Country          string   `json:"country"`
//	RegionalRegistry string   `json:"regional_registry"`
//	IPv4CIDRs        []string `json:"ipv4_cidrs"`
//	IPv6CIDRs        []string `json:"ipv6_cidrs"`
//}
//
//func main() {
//	// Initialize a chrome instance
//	ctx, cancel := chromedp.NewContext(
//		context.Background(),
//		chromedp.WithLogf(log.Printf),
//	)
//	defer cancel()
//
//	// Navigate to the target web page and select the HTML elements of interest
//	var nodes []*cdp.Node
//	if err := chromedp.Run(ctx,
//		chromedp.Navigate("https://asnlookup.com/asn/AS15169"),
//		chromedp.Nodes(".container .table tbody tr", &nodes, chromedp.ByQueryAll),
//	); err != nil {
//		log.Fatalf("Failed to fetch nodes: %v", err)
//	}
//
//	// Extract and format the data
//	asData := ASData{
//		IPv4CIDRs: []string{},
//		IPv6CIDRs: []string{},
//	}
//
//	for _, node := range nodes {
//		var title, content string
//		for _, child := range node.Children {
//			if child.NodeName == "th" {
//				title = child.NodeValue
//			}
//			if child.NodeName == "td" {
//				content = child.NodeValue
//				if strings.Contains(title, "IPv4 CIDRs") {
//					// Handle IPv4 CIDRs
//					links := strings.Split(content, "\n")
//					for _, link := range links {
//						link = strings.TrimSpace(link)
//						if link != "" {
//							asData.IPv4CIDRs = append(asData.IPv4CIDRs, link)
//						}
//					}
//				} else if strings.Contains(title, "IPv6 CIDRs") {
//					// Handle IPv6 CIDRs
//					links := strings.Split(content, "\n")
//					for _, link := range links {
//						link = strings.TrimSpace(link)
//						if link != "" {
//							asData.IPv6CIDRs = append(asData.IPv6CIDRs, link)
//						}
//					}
//				} else {
//					// Handle other fields
//					switch title {
//					case "AS Handle":
//						asData.ASHandle = content
//					case "ASN Name":
//						asData.ASNName = content
//					case "Organization Name":
//						asData.OrganizationName = content
//					case "Organization ID":
//						asData.OrganizationID = content
//					case "Country":
//						asData.Country = content
//					case "Regional Registry":
//						asData.RegionalRegistry = content
//					}
//				}
//			}
//		}
//	}
//
//	// Convert the data to JSON
//	file, err := os.Create("as_data.json")
//	if err != nil {
//		log.Fatalf("Failed to create file: %v", err)
//	}
//	defer file.Close()
//
//	encoder := json.NewEncoder(file)
//	encoder.SetIndent("", "  ")
//	if err := encoder.Encode(asData); err != nil {
//		log.Fatalf("Failed to write JSON: %v", err)
//	}
//
//	fmt.Println("Data has been written to as_data.json")
//}



