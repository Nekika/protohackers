package lib

import (
	"cmp"
	"slices"
	"sync"
)

type Price struct {
	Amount    int32
	Timestamp int32
}

type Repository struct {
	mu     sync.Mutex
	prices []Price
}

func NewRepository() *Repository {
	return &Repository{
		prices: make([]Price, 0),
	}
}

func (r *Repository) Insert(amount, timestamp int32) {
	r.mu.Lock()
	defer r.mu.Unlock()

	price := Price{Amount: amount, Timestamp: timestamp}
	r.prices = append(r.prices, price)
}

func (r *Repository) Average(mintime, maxtime int32) int32 {
	var (
		count, total int32
		prices       = r.Sorted()
	)

	for _, price := range prices {
		if price.Timestamp < mintime {
			continue
		}

		if price.Timestamp > maxtime {
			break
		}

		count += 1
		total += price.Amount
	}

	if count == 0 {
		return 0
	}

	return total / count
}

func (r *Repository) Sorted() []Price {
	var sorted = make([]Price, len(r.prices))

	copy(sorted, r.prices)

	slices.SortFunc(sorted, func(a, b Price) int {
		return cmp.Compare(a.Timestamp, b.Timestamp)
	})

	return sorted
}
