// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/jornInitor/unioffice/schema/soo/sml"
)

func TestCT_DxfsConstructor(t *testing.T) {
	v := sml.NewCT_Dxfs()
	if v == nil {
		t.Errorf("sml.NewCT_Dxfs must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Dxfs should validate: %s", err)
	}
}

func TestCT_DxfsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Dxfs()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Dxfs()
	xml.Unmarshal(buf, v2)
}
