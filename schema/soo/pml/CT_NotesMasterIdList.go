// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package pml

import (
	"encoding/xml"

	"github.com/unidoc/unioffice"
)

type CT_NotesMasterIdList struct {
	// Notes Master ID
	NotesMasterId *CT_NotesMasterIdListEntry
}

func NewCT_NotesMasterIdList() *CT_NotesMasterIdList {
	ret := &CT_NotesMasterIdList{}
	return ret
}

func (m *CT_NotesMasterIdList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.NotesMasterId != nil {
		senotesMasterId := xml.StartElement{Name: xml.Name{Local: "p:notesMasterId"}}
		e.EncodeElement(m.NotesMasterId, senotesMasterId)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_NotesMasterIdList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_NotesMasterIdList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "notesMasterId"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "notesMasterId"}:
				m.NotesMasterId = NewCT_NotesMasterIdListEntry()
				if err := d.DecodeElement(m.NotesMasterId, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_NotesMasterIdList %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_NotesMasterIdList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_NotesMasterIdList and its children
func (m *CT_NotesMasterIdList) Validate() error {
	return m.ValidateWithPath("CT_NotesMasterIdList")
}

// ValidateWithPath validates the CT_NotesMasterIdList and its children, prefixing error messages with path
func (m *CT_NotesMasterIdList) ValidateWithPath(path string) error {
	if m.NotesMasterId != nil {
		if err := m.NotesMasterId.ValidateWithPath(path + "/NotesMasterId"); err != nil {
			return err
		}
	}
	return nil
}
