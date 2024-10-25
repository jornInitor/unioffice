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

	"github.com/jornInitor/unioffice/schema/soo/ofc/sharedTypes"
)

type CT_DocPartName struct {
	// Name Value
	ValAttr string
	// Built-In Entry
	DecoratedAttr *sharedTypes.ST_OnOff
}

func NewCT_DocPartName() *CT_DocPartName {
	ret := &CT_DocPartName{}
	return ret
}

func (m *CT_DocPartName) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	if m.DecoratedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:decorated"},
			Value: fmt.Sprintf("%v", *m.DecoratedAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DocPartName) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ValAttr = parsed
			continue
		}
		if attr.Name.Local == "decorated" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.DecoratedAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_DocPartName: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_DocPartName and its children
func (m *CT_DocPartName) Validate() error {
	return m.ValidateWithPath("CT_DocPartName")
}

// ValidateWithPath validates the CT_DocPartName and its children, prefixing error messages with path
func (m *CT_DocPartName) ValidateWithPath(path string) error {
	if m.DecoratedAttr != nil {
		if err := m.DecoratedAttr.ValidateWithPath(path + "/DecoratedAttr"); err != nil {
			return err
		}
	}
	return nil
}
