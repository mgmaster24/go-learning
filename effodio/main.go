package main

import (
	"fmt"
	"time"
)

func main() {
	timestamps := []string{
		"2022-01-01 00:00:00",
		"2022-01-01 00:00:01",
		"2022-01-01 00:00:02",
		"2022-01-01 00:00:03",
		"2022-01-01 00:00:04",
		"2022-01-01 00:01:04",
		"2022-01-01 00:02:04",
		"2022-01-01 00:03:04",
		"2022-01-01 00:04:04",
		"2022-01-01 00:05:04",
		"2022-01-01 00:06:04",
		"2022-01-01 00:07:04",
		"2022-01-01 00:08:04",
		"2022-01-01 00:09:04",
		"2022-01-01 01:09:04",
		"2022-01-01 02:09:04",
		"2022-01-01 03:09:04",
		"2022-01-01 04:09:04",
		"2022-01-01 05:09:04",
		"2022-01-01 06:09:04",
	}

	layout := "2006-01-02 15:04:05"
	outputs := encodeTimestamps(parseTimeStamps(timestamps, layout))
	fmt.Println(outputs)
}

type Offset struct {
	delta int
	count int
}

type Output struct {
	start   string
	offsets []Offset
}

func parseTimeStamps(timestamps []string, layout string) []time.Time {
	parsedTimestamps := make([]time.Time, len(timestamps))
	for i, ts := range timestamps {
		t, err := time.Parse(layout, ts)
		if err != nil {
			panic(err)
		}

		parsedTimestamps[i] = t
	}

	return parsedTimestamps
}

func encodeTimestamps(timestamps []time.Time) Output {
	var results []Offset
	results = append(results, Offset{
		delta: 0,
		count: 0,
	})

	for i := 1; i < len(timestamps); {
		delta := int(timestamps[i].Sub(timestamps[i-1]).Seconds())
		count := 1
		j := i + 1
		for j < len(timestamps) {
			nextDelta := int(timestamps[j].Sub(timestamps[j-1]).Seconds())
			if nextDelta == delta {
				count++
				j++
			} else {
				break
			}
		}

		results = append(results, Offset{delta: delta, count: count})
		i = j
	}

	return Output{
		start:   fmt.Sprintf("start: %s", timestamps[0]),
		offsets: results,
	}
}
