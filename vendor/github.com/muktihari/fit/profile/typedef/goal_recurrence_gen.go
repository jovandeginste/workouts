// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type GoalRecurrence byte

const (
	GoalRecurrenceOff     GoalRecurrence = 0
	GoalRecurrenceDaily   GoalRecurrence = 1
	GoalRecurrenceWeekly  GoalRecurrence = 2
	GoalRecurrenceMonthly GoalRecurrence = 3
	GoalRecurrenceYearly  GoalRecurrence = 4
	GoalRecurrenceCustom  GoalRecurrence = 5
	GoalRecurrenceInvalid GoalRecurrence = 0xFF
)

func (g GoalRecurrence) Byte() byte { return byte(g) }

func (g GoalRecurrence) String() string {
	switch g {
	case GoalRecurrenceOff:
		return "off"
	case GoalRecurrenceDaily:
		return "daily"
	case GoalRecurrenceWeekly:
		return "weekly"
	case GoalRecurrenceMonthly:
		return "monthly"
	case GoalRecurrenceYearly:
		return "yearly"
	case GoalRecurrenceCustom:
		return "custom"
	default:
		return "GoalRecurrenceInvalid(" + strconv.Itoa(int(g)) + ")"
	}
}

// FromString parse string into GoalRecurrence constant it's represent, return GoalRecurrenceInvalid if not found.
func GoalRecurrenceFromString(s string) GoalRecurrence {
	switch s {
	case "off":
		return GoalRecurrenceOff
	case "daily":
		return GoalRecurrenceDaily
	case "weekly":
		return GoalRecurrenceWeekly
	case "monthly":
		return GoalRecurrenceMonthly
	case "yearly":
		return GoalRecurrenceYearly
	case "custom":
		return GoalRecurrenceCustom
	default:
		return GoalRecurrenceInvalid
	}
}

// List returns all constants.
func ListGoalRecurrence() []GoalRecurrence {
	return []GoalRecurrence{
		GoalRecurrenceOff,
		GoalRecurrenceDaily,
		GoalRecurrenceWeekly,
		GoalRecurrenceMonthly,
		GoalRecurrenceYearly,
		GoalRecurrenceCustom,
	}
}