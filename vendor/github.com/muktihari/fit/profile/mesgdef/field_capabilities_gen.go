// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/internal/sliceutil"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// FieldCapabilities is a FieldCapabilities message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type FieldCapabilities struct {
	MessageIndex typedef.MessageIndex
	MesgNum      typedef.MesgNum
	Count        uint16
	File         typedef.File
	FieldNum     uint8

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewFieldCapabilities creates new FieldCapabilities struct based on given mesg.
// If mesg is nil, it will return FieldCapabilities with all fields being set to its corresponding invalid value.
func NewFieldCapabilities(mesg *proto.Message) *FieldCapabilities {
	vals := [255]proto.Value{}

	var unknownFields []proto.Field
	var developerFields []proto.DeveloperField
	if mesg != nil {
		arr := pool.Get().(*[poolsize]proto.Field)
		unknownFields = arr[:0]
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 254 || mesg.Fields[i].Name == factory.NameUnknown {
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

	return &FieldCapabilities{
		MessageIndex: typedef.MessageIndex(vals[254].Uint16()),
		File:         typedef.File(vals[0].Uint8()),
		MesgNum:      typedef.MesgNum(vals[1].Uint16()),
		FieldNum:     vals[2].Uint8(),
		Count:        vals[3].Uint16(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts FieldCapabilities into proto.Message. If options is nil, default options will be used.
func (m *FieldCapabilities) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumFieldCapabilities}

	if m.MessageIndex != typedef.MessageIndexInvalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = proto.Uint16(uint16(m.MessageIndex))
		fields = append(fields, field)
	}
	if m.File != typedef.FileInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint8(byte(m.File))
		fields = append(fields, field)
	}
	if m.MesgNum != typedef.MesgNumInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint16(uint16(m.MesgNum))
		fields = append(fields, field)
	}
	if m.FieldNum != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint8(m.FieldNum)
		fields = append(fields, field)
	}
	if m.Count != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint16(m.Count)
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

// SetMessageIndex sets MessageIndex value.
func (m *FieldCapabilities) SetMessageIndex(v typedef.MessageIndex) *FieldCapabilities {
	m.MessageIndex = v
	return m
}

// SetFile sets File value.
func (m *FieldCapabilities) SetFile(v typedef.File) *FieldCapabilities {
	m.File = v
	return m
}

// SetMesgNum sets MesgNum value.
func (m *FieldCapabilities) SetMesgNum(v typedef.MesgNum) *FieldCapabilities {
	m.MesgNum = v
	return m
}

// SetFieldNum sets FieldNum value.
func (m *FieldCapabilities) SetFieldNum(v uint8) *FieldCapabilities {
	m.FieldNum = v
	return m
}

// SetCount sets Count value.
func (m *FieldCapabilities) SetCount(v uint16) *FieldCapabilities {
	m.Count = v
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *FieldCapabilities) SetUnknownFields(unknownFields ...proto.Field) *FieldCapabilities {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *FieldCapabilities) SetDeveloperFields(developerFields ...proto.DeveloperField) *FieldCapabilities {
	m.DeveloperFields = developerFields
	return m
}