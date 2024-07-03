package models

import "time"

type TaskComparator []*Task

func (t TaskComparator) Len() int { return len(t) }

func (t TaskComparator) Less(i, j int) bool {
	iActiveAt, _ := time.Parse("2006-01-02", t[i].ActiveAt)
	jActiveAt, _ := time.Parse("2006-01-02", t[j].ActiveAt)
	if iActiveAt.Before(jActiveAt) {
		return true
	}
	return false
}

func (t TaskComparator) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
