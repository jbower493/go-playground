package main

import (
	"fmt"
	"time"
)

func formatToNemTime(dateTime time.Time) (string, error) {
	loc, err := time.LoadLocation("Australia/Brisbane")
	formatted := dateTime.In(loc).Format(time.DateTime)
	return formatted, err
}

func main() {
	now := time.Now()
	formatted, _ := formatToNemTime(now)

	tenThirtyPmUtc, _ := time.Parse(time.RFC3339, "2026-05-23T22:30:00+01:00")
	tenThirtyPmUtcFormatted, _ := formatToNemTime(tenThirtyPmUtc)

	fmt.Println(formatted)
	fmt.Println(tenThirtyPmUtcFormatted)
}
