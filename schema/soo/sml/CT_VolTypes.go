// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package sml

import (
	"encoding/xml"
	"fmt"

	"github.com/jornInitor/unioffice"
)

type CT_VolTypes struct {
	// Volatile Dependency Type
	VolType []*CT_VolType
	ExtLst  *CT_ExtensionList
}

func NewCT_VolTypes() *CT_VolTypes {
	ret := &CT_VolTypes{}
	return ret
}

func (m *CT_VolTypes) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	sevolType := xml.StartElement{Name: xml.Name{Local: "ma:volType"}}
	for _, c := range m.VolType {
		e.EncodeElement(c, sevolType)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_VolTypes) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_VolTypes:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "volType"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "volType"}:
				tmp := NewCT_VolType()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.VolType = append(m.VolType, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_VolTypes %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_VolTypes
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_VolTypes and its children
func (m *CT_VolTypes) Validate() error {
	return m.ValidateWithPath("CT_VolTypes")
}

// ValidateWithPath validates the CT_VolTypes and its children, prefixing error messages with path
func (m *CT_VolTypes) ValidateWithPath(path string) error {
	for i, v := range m.VolType {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/VolType[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
