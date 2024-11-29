package main

import (
	"fmt"
	"slices"
	"time"
)

type Entry struct {
	Price     int32
	Timestamp *time.Time
}

type Historical struct {
	Entries []Entry
}

func (h *Historical) Insert(e Entry) {
	h.Entries = append(h.Entries, e)
	slices.SortFunc(h.Entries, func(a, b Entry) int { return a.Timestamp.Compare(*b.Timestamp) })
}

func (h *Historical) Query(mintime *time.Time, maxtime *time.Time) int32 {
	var (
		n   int32
		sum int32
	)

	fmt.Printf("%#v\n", h)

	for _, entry := range h.Entries {
		if entry.Timestamp.After(*maxtime) {
			break
		}

		if entry.Timestamp.Before(*mintime) {
			continue
		}

		n += 1
		sum += entry.Price
	}

	return sum / n
}
