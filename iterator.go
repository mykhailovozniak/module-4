package module_4

import "fmt"

type collection interface {
	createIterator() iterator
}

type iterator interface {
	hasNext() bool
	getNext() *headline
}

type headlineIterator struct {
	index int
	headlines [] *headline
}

type headline struct {
	title, body, author string
}

func (h *headlineIterator) hasNext() bool {
	return h.index < len(h.headlines)
}

func (h *headlineIterator) getNext() *headline {
	if h.hasNext() {
		headline := h.headlines[h.index]
		h.index++

		return headline
	}
	return nil
}

type headlineCollection struct {
	headlines [] *headline
}

func (h *headlineCollection) createIterator() iterator {
	return &headlineIterator{
		headlines: h.headlines,
	}
}

func iteratorClientCode() {
	topHeadline := &headline{
		title: "Top Headline",
		body: "Some body that top describes headline",
		author: "Original auth of topHeadline",
	}

	financeHeadline := &headline{
		title: "Finance Headline",
		body: "Some body that describes finance headline",
		author: "Original auth of Finance headline",
	}

	headlineCollection := &headlineCollection{
		headlines: [] *headline{topHeadline, financeHeadline},
	}

	iterator := headlineCollection.createIterator()
	for iterator.hasNext() {
		headline := iterator.getNext()
		fmt.Printf("Headline is %+v\n", headline)
	}
}
