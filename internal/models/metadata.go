package models

import (
	"time"
)

type MetaData struct {
	ID           string
	InsertedTime time.Time
	UpdatedTime  time.Time
}
