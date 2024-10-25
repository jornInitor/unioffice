// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package chartDrawing_test

import (
	"encoding/xml"
	"testing"

	"github.com/jornInitor/unioffice/schema/soo/dml/chartDrawing"
)

func TestCT_GroupShapeChoiceConstructor(t *testing.T) {
	v := chartDrawing.NewCT_GroupShapeChoice()
	if v == nil {
		t.Errorf("chartDrawing.NewCT_GroupShapeChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chartDrawing.CT_GroupShapeChoice should validate: %s", err)
	}
}

func TestCT_GroupShapeChoiceMarshalUnmarshal(t *testing.T) {
	v := chartDrawing.NewCT_GroupShapeChoice()
	buf, _ := xml.Marshal(v)
	v2 := chartDrawing.NewCT_GroupShapeChoice()
	xml.Unmarshal(buf, v2)
}
