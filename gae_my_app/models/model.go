package models

import (
	"time"
)

type Topic struct {
	Title   string
	Content string
	Created time.Time
}
