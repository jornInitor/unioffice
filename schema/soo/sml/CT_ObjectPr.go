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
	"strconv"

	"github.com/jornInitor/unioffice"
)

type CT_ObjectPr struct {
	// Locked Flag
	LockedAttr *bool
	// Default Size Flag
	DefaultSizeAttr *bool
	// Print Flag
	PrintAttr *bool
	// Disabled Flag
	DisabledAttr *bool
	// UI Object Flag
	UiObjectAttr *bool
	// Automatic Fill Flag
	AutoFillAttr *bool
	// Automatic Line Flag
	AutoLineAttr *bool
	// Automatic Size Flag
	AutoPictAttr *bool
	// Custom Function
	MacroAttr *string
	// Alternative Text
	AltTextAttr *string
	// Dynamic Data Exchange Flag
	DdeAttr *bool
	IdAttr  *string
	Anchor  *CT_ObjectAnchor
}

func NewCT_ObjectPr() *CT_ObjectPr {
	ret := &CT_ObjectPr{}
	ret.Anchor = NewCT_ObjectAnchor()
	return ret
}

func (m *CT_ObjectPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.LockedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "locked"},
			Value: fmt.Sprintf("%d", b2i(*m.LockedAttr))})
	}
	if m.DefaultSizeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "defaultSize"},
			Value: fmt.Sprintf("%d", b2i(*m.DefaultSizeAttr))})
	}
	if m.PrintAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "print"},
			Value: fmt.Sprintf("%d", b2i(*m.PrintAttr))})
	}
	if m.DisabledAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "disabled"},
			Value: fmt.Sprintf("%d", b2i(*m.DisabledAttr))})
	}
	if m.UiObjectAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "uiObject"},
			Value: fmt.Sprintf("%d", b2i(*m.UiObjectAttr))})
	}
	if m.AutoFillAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "autoFill"},
			Value: fmt.Sprintf("%d", b2i(*m.AutoFillAttr))})
	}
	if m.AutoLineAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "autoLine"},
			Value: fmt.Sprintf("%d", b2i(*m.AutoLineAttr))})
	}
	if m.AutoPictAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "autoPict"},
			Value: fmt.Sprintf("%d", b2i(*m.AutoPictAttr))})
	}
	if m.MacroAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "macro"},
			Value: fmt.Sprintf("%v", *m.MacroAttr)})
	}
	if m.AltTextAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "altText"},
			Value: fmt.Sprintf("%v", *m.AltTextAttr)})
	}
	if m.DdeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dde"},
			Value: fmt.Sprintf("%d", b2i(*m.DdeAttr))})
	}
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	e.EncodeToken(start)
	seanchor := xml.StartElement{Name: xml.Name{Local: "ma:anchor"}}
	e.EncodeElement(m.Anchor, seanchor)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ObjectPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Anchor = NewCT_ObjectAnchor()
	for _, attr := range start.Attr {
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "id" ||
			attr.Name.Space == "http://purl.oclc.org/ooxml/officeDocument/relationships" && attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
			continue
		}
		if attr.Name.Local == "autoLine" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AutoLineAttr = &parsed
			continue
		}
		if attr.Name.Local == "print" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PrintAttr = &parsed
			continue
		}
		if attr.Name.Local == "disabled" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DisabledAttr = &parsed
			continue
		}
		if attr.Name.Local == "uiObject" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UiObjectAttr = &parsed
			continue
		}
		if attr.Name.Local == "autoFill" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AutoFillAttr = &parsed
			continue
		}
		if attr.Name.Local == "locked" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.LockedAttr = &parsed
			continue
		}
		if attr.Name.Local == "autoPict" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AutoPictAttr = &parsed
			continue
		}
		if attr.Name.Local == "macro" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.MacroAttr = &parsed
			continue
		}
		if attr.Name.Local == "altText" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AltTextAttr = &parsed
			continue
		}
		if attr.Name.Local == "dde" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DdeAttr = &parsed
			continue
		}
		if attr.Name.Local == "defaultSize" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DefaultSizeAttr = &parsed
			continue
		}
	}
lCT_ObjectPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "anchor"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "anchor"}:
				if err := d.DecodeElement(m.Anchor, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_ObjectPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ObjectPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ObjectPr and its children
func (m *CT_ObjectPr) Validate() error {
	return m.ValidateWithPath("CT_ObjectPr")
}

// ValidateWithPath validates the CT_ObjectPr and its children, prefixing error messages with path
func (m *CT_ObjectPr) ValidateWithPath(path string) error {
	if err := m.Anchor.ValidateWithPath(path + "/Anchor"); err != nil {
		return err
	}
	return nil
}
