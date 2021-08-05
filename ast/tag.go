package ast

import "fmt"

type TagListNode struct {
	tagListFilterNode *TagListFilterNode
}

func NewTagListNode(tlfn *TagListFilterNode) *TagListNode {
	return &TagListNode{tlfn}
}

func (tln *TagListNode) execute() error { return nil }
func (tln *TagListNode) explain() {
	fmt.Println("list tag ")
	// tln.taskListFilterNode.explain()
}
func (tln *TagListNode) restore() string {
	return "tag list " // + tln.taskListFilterNode.restore()
}

type TagListFilterNode struct {
}

func NewTagListFilterNode(ign *IDGroupNode, cgn *ContentGroupNode) *TagListFilterNode {
	return &TagListFilterNode{}
}
