// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/unidoc/unioffice"
)

type WdCT_WrapTopBottom struct {
	DistTAttr    *uint32
	DistBAttr    *uint32
	EffectExtent *WdCT_EffectExtent
}

func NewWdCT_WrapTopBottom() *WdCT_WrapTopBottom {
	ret := &WdCT_WrapTopBottom{}
	return ret
}

func (m *WdCT_WrapTopBottom) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.DistTAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distT"},
			Value: fmt.Sprintf("%v", *m.DistTAttr)})
	}
	if m.DistBAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distB"},
			Value: fmt.Sprintf("%v", *m.DistBAttr)})
	}
	e.EncodeToken(start)
	if m.EffectExtent != nil {
		seeffectExtent := xml.StartElement{Name: xml.Name{Local: "wp:effectExtent"}}
		e.EncodeElement(m.EffectExtent, seeffectExtent)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *WdCT_WrapTopBottom) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "distT" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DistTAttr = &pt
			continue
		}
		if attr.Name.Local == "distB" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DistBAttr = &pt
			continue
		}
	}
lWdCT_WrapTopBottom:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "effectExtent"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "effectExtent"}:
				m.EffectExtent = NewWdCT_EffectExtent()
				if err := d.DecodeElement(m.EffectExtent, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on WdCT_WrapTopBottom %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lWdCT_WrapTopBottom
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the WdCT_WrapTopBottom and its children
func (m *WdCT_WrapTopBottom) Validate() error {
	return m.ValidateWithPath("WdCT_WrapTopBottom")
}

// ValidateWithPath validates the WdCT_WrapTopBottom and its children, prefixing error messages with path
func (m *WdCT_WrapTopBottom) ValidateWithPath(path string) error {
	if m.EffectExtent != nil {
		if err := m.EffectExtent.ValidateWithPath(path + "/EffectExtent"); err != nil {
			return err
		}
	}
	return nil
}
