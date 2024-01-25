package link_parser

import (
	"golang.org/x/net/html"
	"io"
	"strings"
)

type HTMLLink struct {
	Link, Text string
}

func ExtractLinks(r io.Reader) ([]HTMLLink, error) {
	node, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	links := make([]HTMLLink, 0)
	f := func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			link := parseNodeLink(node.Attr)
			if len(link) == 0 {
				return
			}
			hLink := HTMLLink{Link: link, Text: ""}
			forAllNodesRecursively(node, func(childNode *html.Node) {
				if childNode.Type == html.TextNode {
					hLink.Text = hLink.Text + " " + childNode.Data
				}
			})
			hLink.Text = strings.TrimSpace(hLink.Text)
			links = append(links, hLink)
		}
	}
	forAllNodesRecursively(node, f)
	return links, nil
}

func forAllNodesRecursively(node *html.Node, process func(nodeInput *html.Node)) {
	process(node)
	for n := node.FirstChild; n != nil; n = n.NextSibling {
		forAllNodesRecursively(n, process)
	}
}

func parseNodeLink(attr []html.Attribute) string {
	for i := range attr {
		if attr[i].Key == "href" {
			return attr[i].Val
		}
	}
	return ""
}

//if node.Parent != nil {
//	fmt.Printf("Processing node with data %s and type %v and parent %s \n", node.Data, node.Type, node.Parent.Data)
//}
