// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package pml_test

import (
	"encoding/xml"
	"testing"

	"github.com/jornInitor/unioffice/schema/soo/pml"
)

func TestEG_ShowTypeConstructor(t *testing.T) {
	v := pml.NewEG_ShowType()
	if v == nil {
		t.Errorf("pml.NewEG_ShowType must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.EG_ShowType should validate: %s", err)
	}
}

func TestEG_ShowTypeMarshalUnmarshal(t *testing.T) {
	v := pml.NewEG_ShowType()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewEG_ShowType()
	xml.Unmarshal(buf, v2)
}
