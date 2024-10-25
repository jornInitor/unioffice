// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package math_test

import (
	"encoding/xml"
	"testing"

	"github.com/jornInitor/unioffice/schema/soo/ofc/math"
)

func TestCT_TextConstructor(t *testing.T) {
	v := math.NewCT_Text()
	if v == nil {
		t.Errorf("math.NewCT_Text must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_Text should validate: %s", err)
	}
}

func TestCT_TextMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_Text()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_Text()
	xml.Unmarshal(buf, v2)
}
