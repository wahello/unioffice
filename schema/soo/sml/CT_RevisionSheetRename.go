// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/unidoc/unioffice"
)

type CT_RevisionSheetRename struct {
	// Sheet Id
	SheetIdAttr uint32
	// Old Sheet Name
	OldNameAttr string
	// New Sheet Name
	NewNameAttr string
	ExtLst      *CT_ExtensionList
	RIdAttr     *uint32
	UaAttr      *bool
	RaAttr      *bool
}

func NewCT_RevisionSheetRename() *CT_RevisionSheetRename {
	ret := &CT_RevisionSheetRename{}
	return ret
}

func (m *CT_RevisionSheetRename) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sheetId"},
		Value: fmt.Sprintf("%v", m.SheetIdAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "oldName"},
		Value: fmt.Sprintf("%v", m.OldNameAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "newName"},
		Value: fmt.Sprintf("%v", m.NewNameAttr)})
	if m.RIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rId"},
			Value: fmt.Sprintf("%v", *m.RIdAttr)})
	}
	if m.UaAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ua"},
			Value: fmt.Sprintf("%d", b2i(*m.UaAttr))})
	}
	if m.RaAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ra"},
			Value: fmt.Sprintf("%d", b2i(*m.RaAttr))})
	}
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_RevisionSheetRename) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "sheetId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.SheetIdAttr = uint32(parsed)
			continue
		}
		if attr.Name.Local == "oldName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.OldNameAttr = parsed
			continue
		}
		if attr.Name.Local == "newName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NewNameAttr = parsed
			continue
		}
		if attr.Name.Local == "rId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.RIdAttr = &pt
			continue
		}
		if attr.Name.Local == "ua" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UaAttr = &parsed
			continue
		}
		if attr.Name.Local == "ra" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RaAttr = &parsed
			continue
		}
	}
lCT_RevisionSheetRename:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_RevisionSheetRename %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_RevisionSheetRename
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_RevisionSheetRename and its children
func (m *CT_RevisionSheetRename) Validate() error {
	return m.ValidateWithPath("CT_RevisionSheetRename")
}

// ValidateWithPath validates the CT_RevisionSheetRename and its children, prefixing error messages with path
func (m *CT_RevisionSheetRename) ValidateWithPath(path string) error {
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
