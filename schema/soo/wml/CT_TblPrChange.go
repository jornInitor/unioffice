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

	"github.com/unidoc/unioffice"
)

type CT_TblPrChange struct {
	AuthorAttr string
	DateAttr   *time.Time
	// Annotation Identifier
	IdAttr int64
	TblPr  *CT_TblPrBase
}

func NewCT_TblPrChange() *CT_TblPrChange {
	ret := &CT_TblPrChange{}
	ret.TblPr = NewCT_TblPrBase()
	return ret
}

func (m *CT_TblPrChange) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:author"},
		Value: fmt.Sprintf("%v", m.AuthorAttr)})
	if m.DateAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:date"},
			Value: fmt.Sprintf("%v", *m.DateAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	e.EncodeToken(start)
	setblPr := xml.StartElement{Name: xml.Name{Local: "w:tblPr"}}
	e.EncodeElement(m.TblPr, setblPr)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TblPrChange) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.TblPr = NewCT_TblPrBase()
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
lCT_TblPrChange:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "tblPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "tblPr"}:
				if err := d.DecodeElement(m.TblPr, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_TblPrChange %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TblPrChange
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TblPrChange and its children
func (m *CT_TblPrChange) Validate() error {
	return m.ValidateWithPath("CT_TblPrChange")
}

// ValidateWithPath validates the CT_TblPrChange and its children, prefixing error messages with path
func (m *CT_TblPrChange) ValidateWithPath(path string) error {
	if err := m.TblPr.ValidateWithPath(path + "/TblPr"); err != nil {
		return err
	}
	return nil
}