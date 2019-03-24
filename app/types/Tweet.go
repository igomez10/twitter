package main

import "time"

// Tweet holds the data for a single tweet, included its content and metadata
type Tweet struct {
	ID       int
	Content  string
	postedAt time.Time
	device   string
	Author   int
}
