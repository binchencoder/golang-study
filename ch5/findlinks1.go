// Findlinks prints the links in an HTML document read from standard input.
package main

import (
	"golang.org/x/net/html"
	"os"
	"fmt"
)

//
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	//fmt.Printf("HTML Doc Tree: %v\n", doc)
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// Visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	// 循环条件初始化/条件判断/循环后条件变化
	//for c := n.FirstChild; c != nil; c = c.NextSibling {
	//	//fmt.Printf("node c:%+v\n", c)
	//	links = visit(links, c)
	//}

	return visitSibling(links, n.FirstChild)
}

func visitSibling(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}

	links = visit(links, n)

	sibling := n.NextSibling;
	if sibling == nil {
		return links
	}

	return visitSibling(links, sibling)
}


