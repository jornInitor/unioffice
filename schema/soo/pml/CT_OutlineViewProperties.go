// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package pml

import (
	"encoding/xml"

	"github.com/unidoc/unioffice"
)

type CT_OutlineViewProperties struct {
	// Common View Properties
	CViewPr *CT_CommonViewProperties
	// List of Presentation Slides
	SldLst *CT_OutlineViewSlideList
	ExtLst *CT_ExtensionList
}

func NewCT_OutlineViewProperties() *CT_OutlineViewProperties {
	ret := &CT_OutlineViewProperties{}
	ret.CViewPr = NewCT_CommonViewProperties()
	return ret
}

func (m *CT_OutlineViewProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	secViewPr := xml.StartElement{Name: xml.Name{Local: "p:cViewPr"}}
	e.EncodeElement(m.CViewPr, secViewPr)
	if m.SldLst != nil {
		sesldLst := xml.StartElement{Name: xml.Name{Local: "p:sldLst"}}
		e.EncodeElement(m.SldLst, sesldLst)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_OutlineViewProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CViewPr = NewCT_CommonViewProperties()
lCT_OutlineViewProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "cViewPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "cViewPr"}:
				if err := d.DecodeElement(m.CViewPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "sldLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "sldLst"}:
				m.SldLst = NewCT_OutlineViewSlideList()
				if err := d.DecodeElement(m.SldLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_OutlineViewProperties %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_OutlineViewProperties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_OutlineViewProperties and its children
func (m *CT_OutlineViewProperties) Validate() error {
	return m.ValidateWithPath("CT_OutlineViewProperties")
}

// ValidateWithPath validates the CT_OutlineViewProperties and its children, prefixing error messages with path
func (m *CT_OutlineViewProperties) ValidateWithPath(path string) error {
	if err := m.CViewPr.ValidateWithPath(path + "/CViewPr"); err != nil {
		return err
	}
	if m.SldLst != nil {
		if err := m.SldLst.ValidateWithPath(path + "/SldLst"); err != nil {
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