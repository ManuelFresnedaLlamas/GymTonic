package common

import (
	"fmt"
	"time"
)

func ToPostgresDate(t time.Time) string {
	return fmt.Sprintf("%d-%d-%d", t.Year(), t.Month(), t.Day())
}
