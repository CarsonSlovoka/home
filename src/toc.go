package main

import (
	. "carson.io/pkg/utils"
	"fmt"
	"html/template"
	"regexp"
	"strconv"
)

type TOCNode struct {
	Depth    int
	Content  string
	parent   *TOCNode
	Children []*TOCNode
}

func renderToc(nodes []*TOCNode, ulClassName string) template.HTML {
	var result string
	if ulClassName != "" {
		result = fmt.Sprintf(`<ul class="%s">`, ulClassName)
	} else {
		result = "<ul>"
	}

	for _, node := range nodes {
		result += "<li>" + node.Content
		if len(node.Children) > 0 {
			result += string(renderToc(node.Children, ""))
		}
		result += "</li>"
	}
	result += "</ul>"
	return template.HTML(result)
}

func ParseHTMLAsTOC(content string) []*TOCNode {
	var rootNode []*TOCNode

	reToc := regexp.MustCompile(`(?m)^<h(\d)(.*)>(.*)<\/h\d>`)
	matchList := reToc.FindAllStringSubmatch(content, -1)
	var preNode *TOCNode
	for _, match := range matchList {
		depthStr, _, heading := match[1], match[2], match[3] // match[0]是所有匹配的項目，0之後才是每一個group的內容
		depth, err := strconv.Atoi(depthStr)
		if err != nil {
			PErr.Printf("error strconv.Atoi %s\n", err)
			return nil
		}
		curNode := &TOCNode{depth, heading, preNode, nil}
		if rootNode == nil {
			rootNode = make([]*TOCNode, 0)
			rootNode = append(rootNode, curNode)
			preNode = curNode
			continue
		}

		if preNode != nil && depth > preNode.Depth {
			if preNode.Children == nil {
				preNode.Children = make([]*TOCNode, 0)
			}
			preNode.Children = append(preNode.Children, curNode)
			preNode = curNode
			continue
		}

		// 往回找，直到前一個深度與它相等
		for {
			preNode = preNode.parent
			if preNode == nil {
				rootNode = append(rootNode, curNode)
				preNode = curNode
				break
			}
			if preNode.Depth < curNode.Depth {
				preNode.Children = append(preNode.Children, curNode)
				preNode = curNode
				break
			}
		}
	}
	return rootNode
}
