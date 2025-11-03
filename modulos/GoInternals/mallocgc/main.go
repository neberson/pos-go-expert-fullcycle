package main

import (
	"fmt"
	"sync"
)

type span struct {
	size      int
	allocated bool
}

type mheap struct {
	spans []*span
	lock  sync.Mutex
}

type mcentral struct {
	sizeSpans []*span
	lock      sync.Mutex
}

type mcache struct {
	localSpans []*span
}

func NewHeap(size int) *mheap {
	h := &mheap{}

	for i := 0; i < size; i++ {
		h.spans = append(h.spans, &span{size: i + 1})
	}

	return h
}

func (h *mheap) getSpan(size int) *span {
	h.lock.Lock()
	defer h.lock.Unlock()

	for _, span := range h.spans {
		if span.size == size && !span.allocated {
			span.allocated = true
			return span
		}
	}
	return nil
}

func (mc *mcentral) getSpanFromCentral(size int) *span {
	mc.lock.Lock()
	defer mc.lock.Unlock()

	for _, span := range mc.sizeSpans {
		if span.size == size && !span.allocated {
			span.allocated = true
			return span
		}
	}
	return nil
}

func (mc *mcache) getSpanFromCache(size int) *span {
	for _, span := range mc.localSpans {
		if span.size == size && !span.allocated {
			span.allocated = true
			return span
		}
	}
	return nil
}

func main() {
	heap := NewHeap(10)
	mcentral := &mcentral{sizeSpans: heap.spans}
	mcache := &mcache{}

	requestScan := mcache.getSpanFromCache(5)
	if requestScan == nil {
		requestScan = mcentral.getSpanFromCentral(5)
	}
	if requestScan == nil {
		requestScan = heap.getSpan(5)
	}
	if requestScan == nil {
		panic("out of memory")
	}

	fmt.Println("Allocated span")
}
