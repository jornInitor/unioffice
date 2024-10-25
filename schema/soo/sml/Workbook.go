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

	"github.com/jornInitor/unioffice"
)

type Workbook struct {
	CT_Workbook
}

func NewWorkbook() *Workbook {
	ret := &Workbook{}
	ret.CT_Workbook = *NewCT_Workbook()
	return ret
}

func (m *Workbook) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:ma"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xdr"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "ma:workbook"
	return m.CT_Workbook.MarshalXML(e, start)
}

func (m *Workbook) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Workbook = *NewCT_Workbook()
	for _, attr := range start.Attr {
		if attr.Name.Local == "conformance" {
			m.ConformanceAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lWorkbook:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "fileVersion"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "fileVersion"}:
				m.FileVersion = NewCT_FileVersion()
				if err := d.DecodeElement(m.FileVersion, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "fileSharing"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "fileSharing"}:
				m.FileSharing = NewCT_FileSharing()
				if err := d.DecodeElement(m.FileSharing, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "workbookPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "workbookPr"}:
				m.WorkbookPr = NewCT_WorkbookPr()
				if err := d.DecodeElement(m.WorkbookPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "workbookProtection"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "workbookProtection"}:
				m.WorkbookProtection = NewCT_WorkbookProtection()
				if err := d.DecodeElement(m.WorkbookProtection, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "bookViews"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "bookViews"}:
				m.BookViews = NewCT_BookViews()
				if err := d.DecodeElement(m.BookViews, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "sheets"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "sheets"}:
				if err := d.DecodeElement(m.Sheets, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "functionGroups"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "functionGroups"}:
				m.FunctionGroups = NewCT_FunctionGroups()
				if err := d.DecodeElement(m.FunctionGroups, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "externalReferences"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "externalReferences"}:
				m.ExternalReferences = NewCT_ExternalReferences()
				if err := d.DecodeElement(m.ExternalReferences, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "definedNames"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "definedNames"}:
				m.DefinedNames = NewCT_DefinedNames()
				if err := d.DecodeElement(m.DefinedNames, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "calcPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "calcPr"}:
				m.CalcPr = NewCT_CalcPr()
				if err := d.DecodeElement(m.CalcPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "oleSize"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "oleSize"}:
				m.OleSize = NewCT_OleSize()
				if err := d.DecodeElement(m.OleSize, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "customWorkbookViews"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "customWorkbookViews"}:
				m.CustomWorkbookViews = NewCT_CustomWorkbookViews()
				if err := d.DecodeElement(m.CustomWorkbookViews, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "pivotCaches"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "pivotCaches"}:
				m.PivotCaches = NewCT_PivotCaches()
				if err := d.DecodeElement(m.PivotCaches, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "smartTagPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "smartTagPr"}:
				m.SmartTagPr = NewCT_SmartTagPr()
				if err := d.DecodeElement(m.SmartTagPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "smartTagTypes"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "smartTagTypes"}:
				m.SmartTagTypes = NewCT_SmartTagTypes()
				if err := d.DecodeElement(m.SmartTagTypes, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "webPublishing"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "webPublishing"}:
				m.WebPublishing = NewCT_WebPublishing()
				if err := d.DecodeElement(m.WebPublishing, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "fileRecoveryPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "fileRecoveryPr"}:
				tmp := NewCT_FileRecoveryPr()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.FileRecoveryPr = append(m.FileRecoveryPr, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "webPublishObjects"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "webPublishObjects"}:
				m.WebPublishObjects = NewCT_WebPublishObjects()
				if err := d.DecodeElement(m.WebPublishObjects, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on Workbook %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lWorkbook
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Workbook and its children
func (m *Workbook) Validate() error {
	return m.ValidateWithPath("Workbook")
}

// ValidateWithPath validates the Workbook and its children, prefixing error messages with path
func (m *Workbook) ValidateWithPath(path string) error {
	if err := m.CT_Workbook.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
