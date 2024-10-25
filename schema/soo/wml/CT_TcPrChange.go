// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package wml

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"

	"github.com/jornInitor/unioffice"
)

type CT_TcPrChange struct {
	AuthorAttr string
	DateAttr   *time.Time
	// Annotation Identifier
	IdAttr int64
	TcPr   *CT_TcPrInner
}

func NewCT_TcPrChange() *CT_TcPrChange {
	ret := &CT_TcPrChange{}
	ret.TcPr = NewCT_TcPrInner()
	return ret
}

func (m *CT_TcPrChange) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:author"},
		Value: fmt.Sprintf("%v", m.AuthorAttr)})
	if m.DateAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:date"},
			Value: fmt.Sprintf("%v", *m.DateAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	e.EncodeToken(start)
	setcPr := xml.StartElement{Name: xml.Name{Local: "w:tcPr"}}
	e.EncodeElement(m.TcPr, setcPr)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TcPrChange) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.TcPr = NewCT_TcPrInner()
	for _, attr := range start.Attr {
		if attr.Name.Local == "author" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AuthorAttr = parsed
			continue
		}
		if attr.Name.Local == "date" {
			parsed, err := ParseStdlibTime(attr.Value)
			if err != nil {
				return err
			}
			m.DateAttr = &parsed
			continue
		}
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.IdAttr = parsed
			continue
		}
	}
lCT_TcPrChange:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "tcPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "tcPr"}:
				if err := d.DecodeElement(m.TcPr, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_TcPrChange %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TcPrChange
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TcPrChange and its children
func (m *CT_TcPrChange) Validate() error {
	return m.ValidateWithPath("CT_TcPrChange")
}

// ValidateWithPath validates the CT_TcPrChange and its children, prefixing error messages with path
func (m *CT_TcPrChange) ValidateWithPath(path string) error {
	if err := m.TcPr.ValidateWithPath(path + "/TcPr"); err != nil {
		return err
	}
	return nil
}
