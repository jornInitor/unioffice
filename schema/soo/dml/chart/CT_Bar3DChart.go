// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package chart

import (
	"encoding/xml"
	"fmt"

	"github.com/jornInitor/unioffice"
)

type CT_Bar3DChart struct {
	BarDir     *CT_BarDir
	Grouping   *CT_BarGrouping
	VaryColors *CT_Boolean
	Ser        []*CT_BarSer
	DLbls      *CT_DLbls
	GapWidth   *CT_GapAmount
	GapDepth   *CT_GapAmount
	Shape      *CT_Shape
	AxId       []*CT_UnsignedInt
	ExtLst     *CT_ExtensionList
}

func NewCT_Bar3DChart() *CT_Bar3DChart {
	ret := &CT_Bar3DChart{}
	ret.BarDir = NewCT_BarDir()
	return ret
}

func (m *CT_Bar3DChart) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	sebarDir := xml.StartElement{Name: xml.Name{Local: "c:barDir"}}
	e.EncodeElement(m.BarDir, sebarDir)
	if m.Grouping != nil {
		segrouping := xml.StartElement{Name: xml.Name{Local: "c:grouping"}}
		e.EncodeElement(m.Grouping, segrouping)
	}
	if m.VaryColors != nil {
		sevaryColors := xml.StartElement{Name: xml.Name{Local: "c:varyColors"}}
		e.EncodeElement(m.VaryColors, sevaryColors)
	}
	if m.Ser != nil {
		seser := xml.StartElement{Name: xml.Name{Local: "c:ser"}}
		for _, c := range m.Ser {
			e.EncodeElement(c, seser)
		}
	}
	if m.DLbls != nil {
		sedLbls := xml.StartElement{Name: xml.Name{Local: "c:dLbls"}}
		e.EncodeElement(m.DLbls, sedLbls)
	}
	if m.GapWidth != nil {
		segapWidth := xml.StartElement{Name: xml.Name{Local: "c:gapWidth"}}
		e.EncodeElement(m.GapWidth, segapWidth)
	}
	if m.GapDepth != nil {
		segapDepth := xml.StartElement{Name: xml.Name{Local: "c:gapDepth"}}
		e.EncodeElement(m.GapDepth, segapDepth)
	}
	if m.Shape != nil {
		seshape := xml.StartElement{Name: xml.Name{Local: "c:shape"}}
		e.EncodeElement(m.Shape, seshape)
	}
	seaxId := xml.StartElement{Name: xml.Name{Local: "c:axId"}}
	for _, c := range m.AxId {
		e.EncodeElement(c, seaxId)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "c:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Bar3DChart) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.BarDir = NewCT_BarDir()
lCT_Bar3DChart:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "barDir"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "barDir"}:
				if err := d.DecodeElement(m.BarDir, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "grouping"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "grouping"}:
				m.Grouping = NewCT_BarGrouping()
				if err := d.DecodeElement(m.Grouping, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "varyColors"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "varyColors"}:
				m.VaryColors = NewCT_Boolean()
				if err := d.DecodeElement(m.VaryColors, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "ser"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "ser"}:
				tmp := NewCT_BarSer()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Ser = append(m.Ser, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "dLbls"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "dLbls"}:
				m.DLbls = NewCT_DLbls()
				if err := d.DecodeElement(m.DLbls, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "gapWidth"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "gapWidth"}:
				m.GapWidth = NewCT_GapAmount()
				if err := d.DecodeElement(m.GapWidth, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "gapDepth"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "gapDepth"}:
				m.GapDepth = NewCT_GapAmount()
				if err := d.DecodeElement(m.GapDepth, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "shape"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "shape"}:
				m.Shape = NewCT_Shape()
				if err := d.DecodeElement(m.Shape, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "axId"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "axId"}:
				tmp := NewCT_UnsignedInt()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AxId = append(m.AxId, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_Bar3DChart %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Bar3DChart
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Bar3DChart and its children
func (m *CT_Bar3DChart) Validate() error {
	return m.ValidateWithPath("CT_Bar3DChart")
}

// ValidateWithPath validates the CT_Bar3DChart and its children, prefixing error messages with path
func (m *CT_Bar3DChart) ValidateWithPath(path string) error {
	if err := m.BarDir.ValidateWithPath(path + "/BarDir"); err != nil {
		return err
	}
	if m.Grouping != nil {
		if err := m.Grouping.ValidateWithPath(path + "/Grouping"); err != nil {
			return err
		}
	}
	if m.VaryColors != nil {
		if err := m.VaryColors.ValidateWithPath(path + "/VaryColors"); err != nil {
			return err
		}
	}
	for i, v := range m.Ser {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Ser[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.DLbls != nil {
		if err := m.DLbls.ValidateWithPath(path + "/DLbls"); err != nil {
			return err
		}
	}
	if m.GapWidth != nil {
		if err := m.GapWidth.ValidateWithPath(path + "/GapWidth"); err != nil {
			return err
		}
	}
	if m.GapDepth != nil {
		if err := m.GapDepth.ValidateWithPath(path + "/GapDepth"); err != nil {
			return err
		}
	}
	if m.Shape != nil {
		if err := m.Shape.ValidateWithPath(path + "/Shape"); err != nil {
			return err
		}
	}
	for i, v := range m.AxId {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AxId[%d]", path, i)); err != nil {
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
