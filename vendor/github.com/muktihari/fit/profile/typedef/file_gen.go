// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"fmt"
	"strconv"
)

type File byte

const (
	FileDevice           File = 1  // Read only, single file. Must be in root directory.
	FileSettings         File = 2  // Read/write, single file. Directory=Settings
	FileSport            File = 3  // Read/write, multiple files, file number = sport type. Directory=Sports
	FileActivity         File = 4  // Read/erase, multiple files. Directory=Activities
	FileWorkout          File = 5  // Read/write/erase, multiple files. Directory=Workouts
	FileCourse           File = 6  // Read/write/erase, multiple files. Directory=Courses
	FileSchedules        File = 7  // Read/write, single file. Directory=Schedules
	FileWeight           File = 9  // Read only, single file. Circular buffer. All message definitions at start of file. Directory=Weight
	FileTotals           File = 10 // Read only, single file. Directory=Totals
	FileGoals            File = 11 // Read/write, single file. Directory=Goals
	FileBloodPressure    File = 14 // Read only. Directory=Blood Pressure
	FileMonitoringA      File = 15 // Read only. Directory=Monitoring. File number=sub type.
	FileActivitySummary  File = 20 // Read/erase, multiple files. Directory=Activities
	FileMonitoringDaily  File = 28
	FileMonitoringB      File = 32   // Read only. Directory=Monitoring. File number=identifier
	FileSegment          File = 34   // Read/write/erase. Multiple Files. Directory=Segments
	FileSegmentList      File = 35   // Read/write/erase. Single File. Directory=Segments
	FileExdConfiguration File = 40   // Read/write/erase. Single File. Directory=Settings
	FileMfgRangeMin      File = 0xF7 // 0xF7 - 0xFE reserved for manufacturer specific file types
	FileMfgRangeMax      File = 0xFE // 0xF7 - 0xFE reserved for manufacturer specific file types
	FileInvalid          File = 0xFF
)

func (f File) Byte() byte { return byte(f) }

var fileToString = map[File]string{}
var stringToFile = map[string]File{}

func (f File) String() string {
	switch f {
	case FileDevice:
		return "device"
	case FileSettings:
		return "settings"
	case FileSport:
		return "sport"
	case FileActivity:
		return "activity"
	case FileWorkout:
		return "workout"
	case FileCourse:
		return "course"
	case FileSchedules:
		return "schedules"
	case FileWeight:
		return "weight"
	case FileTotals:
		return "totals"
	case FileGoals:
		return "goals"
	case FileBloodPressure:
		return "blood_pressure"
	case FileMonitoringA:
		return "monitoring_a"
	case FileActivitySummary:
		return "activity_summary"
	case FileMonitoringDaily:
		return "monitoring_daily"
	case FileMonitoringB:
		return "monitoring_b"
	case FileSegment:
		return "segment"
	case FileSegmentList:
		return "segment_list"
	case FileExdConfiguration:
		return "exd_configuration"
	case FileMfgRangeMin:
		return "mfg_range_min"
	case FileMfgRangeMax:
		return "mfg_range_max"
	default:
		if val, ok := fileToString[f]; ok {
			return val
		}
		return "FileInvalid(" + strconv.Itoa(int(f)) + ")"
	}
}

// FromString parse string into File constant it's represent, return FileInvalid if not found.
func FileFromString(s string) File {
	switch s {
	case "device":
		return FileDevice
	case "settings":
		return FileSettings
	case "sport":
		return FileSport
	case "activity":
		return FileActivity
	case "workout":
		return FileWorkout
	case "course":
		return FileCourse
	case "schedules":
		return FileSchedules
	case "weight":
		return FileWeight
	case "totals":
		return FileTotals
	case "goals":
		return FileGoals
	case "blood_pressure":
		return FileBloodPressure
	case "monitoring_a":
		return FileMonitoringA
	case "activity_summary":
		return FileActivitySummary
	case "monitoring_daily":
		return FileMonitoringDaily
	case "monitoring_b":
		return FileMonitoringB
	case "segment":
		return FileSegment
	case "segment_list":
		return FileSegmentList
	case "exd_configuration":
		return FileExdConfiguration
	case "mfg_range_min":
		return FileMfgRangeMin
	case "mfg_range_max":
		return FileMfgRangeMax
	default:
		if val, ok := stringToFile[s]; ok {
			return val
		}
		return FileInvalid
	}
}

// List returns all constants.
func ListFile() []File {
	list := []File{
		FileDevice,
		FileSettings,
		FileSport,
		FileActivity,
		FileWorkout,
		FileCourse,
		FileSchedules,
		FileWeight,
		FileTotals,
		FileGoals,
		FileBloodPressure,
		FileMonitoringA,
		FileActivitySummary,
		FileMonitoringDaily,
		FileMonitoringB,
		FileSegment,
		FileSegmentList,
		FileExdConfiguration,
		FileMfgRangeMin,
		FileMfgRangeMax,
	}
	for k := range fileToString {
		list = append(list, k)
	}
	return list
}

// FileRegister registers a manufacturer specific File so that the value can be recognized.
// It is recommended to define the constants somewhere else to track your own specifications.
//
// This is intended for those who prefer using this SDK as it is without the need to generate custom SDK using cmd/fitgen.
func FileRegister(v File, s string) error {
	if v >= FileInvalid {
		return fmt.Errorf("could not register outside max range: %d", FileInvalid)
	}

	switch v {
	case FileDevice:
		return fmt.Errorf("duplicate: %d is already exist for FileDevice", v)
	case FileSettings:
		return fmt.Errorf("duplicate: %d is already exist for FileSettings", v)
	case FileSport:
		return fmt.Errorf("duplicate: %d is already exist for FileSport", v)
	case FileActivity:
		return fmt.Errorf("duplicate: %d is already exist for FileActivity", v)
	case FileWorkout:
		return fmt.Errorf("duplicate: %d is already exist for FileWorkout", v)
	case FileCourse:
		return fmt.Errorf("duplicate: %d is already exist for FileCourse", v)
	case FileSchedules:
		return fmt.Errorf("duplicate: %d is already exist for FileSchedules", v)
	case FileWeight:
		return fmt.Errorf("duplicate: %d is already exist for FileWeight", v)
	case FileTotals:
		return fmt.Errorf("duplicate: %d is already exist for FileTotals", v)
	case FileGoals:
		return fmt.Errorf("duplicate: %d is already exist for FileGoals", v)
	case FileBloodPressure:
		return fmt.Errorf("duplicate: %d is already exist for FileBloodPressure", v)
	case FileMonitoringA:
		return fmt.Errorf("duplicate: %d is already exist for FileMonitoringA", v)
	case FileActivitySummary:
		return fmt.Errorf("duplicate: %d is already exist for FileActivitySummary", v)
	case FileMonitoringDaily:
		return fmt.Errorf("duplicate: %d is already exist for FileMonitoringDaily", v)
	case FileMonitoringB:
		return fmt.Errorf("duplicate: %d is already exist for FileMonitoringB", v)
	case FileSegment:
		return fmt.Errorf("duplicate: %d is already exist for FileSegment", v)
	case FileSegmentList:
		return fmt.Errorf("duplicate: %d is already exist for FileSegmentList", v)
	case FileExdConfiguration:
		return fmt.Errorf("duplicate: %d is already exist for FileExdConfiguration", v)
	case FileMfgRangeMin:
		return fmt.Errorf("duplicate: %d is already exist for FileMfgRangeMin", v)
	case FileMfgRangeMax:
		return fmt.Errorf("duplicate: %d is already exist for FileMfgRangeMax", v)
	}

	fileToString[v] = s
	stringToFile[s] = v

	return nil
}