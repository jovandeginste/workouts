// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/internal/sliceutil"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"time"
)

// StressLevel is a StressLevel message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type StressLevel struct {
	StressLevelTime  time.Time // Units: s; Time stress score was calculated
	StressLevelValue int16

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewStressLevel creates new StressLevel struct based on given mesg.
// If mesg is nil, it will return StressLevel with all fields being set to its corresponding invalid value.
func NewStressLevel(mesg *proto.Message) *StressLevel {
	vals := [2]proto.Value{}

	var unknownFields []proto.Field
	var developerFields []proto.DeveloperField
	if mesg != nil {
		arr := pool.Get().(*[poolsize]proto.Field)
		unknownFields = arr[:0]
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 1 || mesg.Fields[i].Name == factory.NameUnknown {
				unknownFields = append(unknownFields, mesg.Fields[i])
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		unknownFields = sliceutil.Clone(unknownFields)
		*arr = [poolsize]proto.Field{}
		pool.Put(arr)
		developerFields = mesg.DeveloperFields
	}

	return &StressLevel{
		StressLevelValue: vals[0].Int16(),
		StressLevelTime:  datetime.ToTime(vals[1].Uint32()),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts StressLevel into proto.Message. If options is nil, default options will be used.
func (m *StressLevel) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumStressLevel}

	if m.StressLevelValue != basetype.Sint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Int16(m.StressLevelValue)
		fields = append(fields, field)
	}
	if !m.StressLevelTime.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint32(uint32(m.StressLevelTime.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}

	for i := range m.UnknownFields {
		fields = append(fields, m.UnknownFields[i])
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)
	*arr = [poolsize]proto.Field{}
	pool.Put(arr)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// StressLevelTimeUint32 returns StressLevelTime in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *StressLevel) StressLevelTimeUint32() uint32 { return datetime.ToUint32(m.StressLevelTime) }

// SetStressLevelValue sets StressLevelValue value.
func (m *StressLevel) SetStressLevelValue(v int16) *StressLevel {
	m.StressLevelValue = v
	return m
}

// SetStressLevelTime sets StressLevelTime value.
//
// Units: s; Time stress score was calculated
func (m *StressLevel) SetStressLevelTime(v time.Time) *StressLevel {
	m.StressLevelTime = v
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *StressLevel) SetUnknownFields(unknownFields ...proto.Field) *StressLevel {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *StressLevel) SetDeveloperFields(developerFields ...proto.DeveloperField) *StressLevel {
	m.DeveloperFields = developerFields
	return m
}