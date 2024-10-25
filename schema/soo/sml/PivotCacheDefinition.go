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
	"strconv"

	"github.com/jornInitor/unioffice"
)

type PivotCacheDefinition struct {
	CT_PivotCacheDefinition
}

func NewPivotCacheDefinition() *PivotCacheDefinition {
	ret := &PivotCacheDefinition{}
	ret.CT_PivotCacheDefinition = *NewCT_PivotCacheDefinition()
	return ret
}

func (m *PivotCacheDefinition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:ma"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xdr"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "ma:pivotCacheDefinition"
	return m.CT_PivotCacheDefinition.MarshalXML(e, start)
}

func (m *PivotCacheDefinition) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_PivotCacheDefinition = *NewCT_PivotCacheDefinition()
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
		if attr.Name.Local == "upgradeOnRefresh" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UpgradeOnRefreshAttr = &parsed
			continue
		}
		if attr.Name.Local == "tupleCache" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.TupleCacheAttr = &parsed
			continue
		}
		if attr.Name.Local == "saveData" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SaveDataAttr = &parsed
			continue
		}
		if attr.Name.Local == "supportSubquery" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SupportSubqueryAttr = &parsed
			continue
		}
		if attr.Name.Local == "optimizeMemory" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.OptimizeMemoryAttr = &parsed
			continue
		}
		if attr.Name.Local == "supportAdvancedDrill" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SupportAdvancedDrillAttr = &parsed
			continue
		}
		if attr.Name.Local == "refreshedBy" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RefreshedByAttr = &parsed
			continue
		}
		if attr.Name.Local == "refreshedDateIso" {
			parsed, err := ParseStdlibTime(attr.Value)
			if err != nil {
				return err
			}
			m.RefreshedDateIsoAttr = &parsed
			continue
		}
		if attr.Name.Local == "invalid" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.InvalidAttr = &parsed
			continue
		}
		if attr.Name.Local == "backgroundQuery" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.BackgroundQueryAttr = &parsed
			continue
		}
		if attr.Name.Local == "missingItemsLimit" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.MissingItemsLimitAttr = &pt
			continue
		}
		if attr.Name.Local == "refreshedVersion" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 8)
			if err != nil {
				return err
			}
			pt := uint8(parsed)
			m.RefreshedVersionAttr = &pt
			continue
		}
		if attr.Name.Local == "refreshOnLoad" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RefreshOnLoadAttr = &parsed
			continue
		}
		if attr.Name.Local == "refreshedDate" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.RefreshedDateAttr = &parsed
			continue
		}
		if attr.Name.Local == "recordCount" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.RecordCountAttr = &pt
			continue
		}
		if attr.Name.Local == "createdVersion" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 8)
			if err != nil {
				return err
			}
			pt := uint8(parsed)
			m.CreatedVersionAttr = &pt
			continue
		}
		if attr.Name.Local == "minRefreshableVersion" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 8)
			if err != nil {
				return err
			}
			pt := uint8(parsed)
			m.MinRefreshableVersionAttr = &pt
			continue
		}
		if attr.Name.Local == "enableRefresh" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.EnableRefreshAttr = &parsed
			continue
		}
	}
lPivotCacheDefinition:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "cacheSource"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "cacheSource"}:
				if err := d.DecodeElement(m.CacheSource, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "cacheFields"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "cacheFields"}:
				if err := d.DecodeElement(m.CacheFields, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "cacheHierarchies"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "cacheHierarchies"}:
				m.CacheHierarchies = NewCT_CacheHierarchies()
				if err := d.DecodeElement(m.CacheHierarchies, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "kpis"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "kpis"}:
				m.Kpis = NewCT_PCDKPIs()
				if err := d.DecodeElement(m.Kpis, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "tupleCache"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "tupleCache"}:
				m.TupleCache = NewCT_TupleCache()
				if err := d.DecodeElement(m.TupleCache, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "calculatedItems"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "calculatedItems"}:
				m.CalculatedItems = NewCT_CalculatedItems()
				if err := d.DecodeElement(m.CalculatedItems, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "calculatedMembers"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "calculatedMembers"}:
				m.CalculatedMembers = NewCT_CalculatedMembers()
				if err := d.DecodeElement(m.CalculatedMembers, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "dimensions"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "dimensions"}:
				m.Dimensions = NewCT_Dimensions()
				if err := d.DecodeElement(m.Dimensions, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "measureGroups"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "measureGroups"}:
				m.MeasureGroups = NewCT_MeasureGroups()
				if err := d.DecodeElement(m.MeasureGroups, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "maps"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "maps"}:
				m.Maps = NewCT_MeasureDimensionMaps()
				if err := d.DecodeElement(m.Maps, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on PivotCacheDefinition %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lPivotCacheDefinition
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the PivotCacheDefinition and its children
func (m *PivotCacheDefinition) Validate() error {
	return m.ValidateWithPath("PivotCacheDefinition")
}

// ValidateWithPath validates the PivotCacheDefinition and its children, prefixing error messages with path
func (m *PivotCacheDefinition) ValidateWithPath(path string) error {
	if err := m.CT_PivotCacheDefinition.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
