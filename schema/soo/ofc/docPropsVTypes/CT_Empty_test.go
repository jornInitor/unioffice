// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package docPropsVTypes_test

import (
	"encoding/xml"
	"testing"

	"github.com/jornInitor/unioffice/schema/soo/ofc/docPropsVTypes"
)

func TestCT_EmptyConstructor(t *testing.T) {
	v := docPropsVTypes.NewCT_Empty()
	if v == nil {
		t.Errorf("docPropsVTypes.NewCT_Empty must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed docPropsVTypes.CT_Empty should validate: %s", err)
	}
}

func TestCT_EmptyMarshalUnmarshal(t *testing.T) {
	v := docPropsVTypes.NewCT_Empty()
	buf, _ := xml.Marshal(v)
	v2 := docPropsVTypes.NewCT_Empty()
	xml.Unmarshal(buf, v2)
}
