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

	type Metadata struct {
		LastChanged time.Time
	}

	type Vertex struct {
		X        int
		Y        int
		Metadata Metadata
	}

	vert := Vertex{X: 1, Y: 2, Metadata: Metadata{LastChanged: time.Now()}}

	fmt.Println(vert.X, vert.Y, vert.Metadata)
}
