package main

import (
	"fmt"
	"github.com/pastushenkoy/adventOfCode/utils"
	"regexp"
	"sort"
	"strconv"
	"time"
)

type EventType int

const (
	BeginsShift EventType = 1
	FallsAsleep EventType = 2
	WakesUp     EventType = 3
)

type Event struct {
	timestamp time.Time
	guardId   int
	eventType EventType
}

type Guard struct {
	id             int
	sleepIntervals []SleepInterval
	sleepDuration float64
}

type SleepInterval struct {
	begin time.Time
	end   time.Time
}

func (guard *Guard) calculateSleepDuration() {
	durationMinutes := 0.
	for _, interval := range guard.sleepIntervals{
		sub := interval.end.Sub(interval.begin)
		durationMinutes += sub.Minutes()
	}
	guard.sleepDuration = durationMinutes
}

func main() {
	data := utils.ReadFileOfStrings("input.txt")

	guards := getGuardsData(data)

	guard1, minute1 := calculateStrategy1(guards)

	fmt.Println(fmt.Sprintf("Strategy 1. Guard Id is %v. Minute is %v. The multiplication is %v", guard1, minute1, minute1*guard1));


	guardId2, minuteId2 := calculateStrategy2(guards)

	fmt.Println(fmt.Sprintf("Strategy 2. Guard Id is %v. Minute is %v. The multiplication is %v", guardId2, minuteId2, minuteId2*guardId2));
}

func calculateStrategy2(guards []*Guard) (guardId int, minuteId int) {
	mostSleepingMinute := -1
	mostSleptDaysCount := 0
	mostSleepingGuardId := -1
	for i := 0; i < 60; i++ {
		for _, guard := range guards {
			guardSleptDays := 0
			for _, interval := range guard.sleepIntervals {
				if interval.begin.Minute() <= i && i <= interval.end.Minute() {
					guardSleptDays++
				}
			}
			if mostSleptDaysCount < guardSleptDays {
				mostSleepingMinute = i
				mostSleptDaysCount = guardSleptDays
				mostSleepingGuardId = guard.id
			}
		}
	}
	return mostSleepingGuardId, mostSleepingMinute
}

func calculateStrategy1(guards []*Guard) (id int, minute int) {
	for _, guard := range guards {
		guard.calculateSleepDuration()
	}
	lazyGuard := guards[0]
	for _, guard := range guards {
		guard.calculateSleepDuration()
		if guard.sleepDuration > lazyGuard.sleepDuration {
			lazyGuard = guard
		}
	}
	mostSleepingMinute := -1
	sleepingIntervalsCount := 0
	for i := 1; i < 60; i++ {
		intervalCount := 0
		for _, interval := range lazyGuard.sleepIntervals {
			if interval.begin.Minute() <= i && i <= interval.end.Minute() {
				intervalCount++
			}
		}
		if intervalCount > sleepingIntervalsCount {
			sleepingIntervalsCount = intervalCount
			mostSleepingMinute = i
		}
	}
	return lazyGuard.id, mostSleepingMinute
}

func getGuardsData(data []string) []*Guard {
	events := make([]Event, 0)
	for _, row := range data {
		events = append(events, parseEvent(row))
	}
	sort.Slice(events, func(i, j int) bool {
		return events[i].timestamp.Before(events[j].timestamp)
	})
	guards := make([]*Guard, 0)
	var currentGuard *Guard
	var found bool
	for _, event := range events {
		switch event.eventType {
		case BeginsShift:
			currentGuard, found = getGuardById(guards, event.guardId)
			if !found {
				currentGuard = &Guard{
					id:             event.guardId,
					sleepIntervals: make([]SleepInterval, 0),
				}
				guards = append(guards, currentGuard)
			}

		case FallsAsleep:
			currentGuard.sleepIntervals = append(currentGuard.sleepIntervals, SleepInterval{begin: event.timestamp})
		case WakesUp:
			currentInterval := len(currentGuard.sleepIntervals) - 1
			currentGuard.sleepIntervals[currentInterval].end = event.timestamp.Add(-time.Minute)
		}
	}
	return guards
}

func getGuardById(guards []*Guard, id int) (*Guard, bool) {
	for _, guard := range guards{
		if guard.id == id{
			return guard, true
		}
	}
	return &Guard{}, false
}


func parseEvent(row string) Event {
	re := regexp.MustCompile("\\[(\\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2})\\] (falls asleep|wakes up|Guard #(\\d+) begins shift)")
	submatch := re.FindStringSubmatch(row)
	timestamp, _ := time.Parse("2006-01-02 15:04", submatch[1])
	var event EventType
	var id int
	switch submatch[2] {
	case "falls asleep":
		event = FallsAsleep
		id = 0
	case "wakes up":
		event = WakesUp
		id = 0
	default:
		event = BeginsShift
		idStr, _ := strconv.ParseInt(submatch[3], 10, 32);
		id = int(idStr)
	}

	return Event{
		timestamp: timestamp,
		eventType: event,
		guardId:   id,
	}

}
