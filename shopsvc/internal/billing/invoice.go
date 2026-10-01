package billing

import "time"

type Invoice struct {
	Number   string
	IssuedAt time.Time
	TotalCts int64
}

func (i Invoice) Overdue(now time.Time, terms time.Duration) bool {
	return now.After(i.IssuedAt.Add(terms))
}
