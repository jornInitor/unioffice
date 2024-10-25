// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package diagram_test

import (
	"encoding/xml"
	"testing"

	"github.com/jornInitor/unioffice/schema/soo/dml/diagram"
)

func TestAG_ConstraintRefAttributesConstructor(t *testing.T) {
	v := diagram.NewAG_ConstraintRefAttributes()
	if v == nil {
		t.Errorf("diagram.NewAG_ConstraintRefAttributes must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.AG_ConstraintRefAttributes should validate: %s", err)
	}
}

func TestAG_ConstraintRefAttributesMarshalUnmarshal(t *testing.T) {
	v := diagram.NewAG_ConstraintRefAttributes()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewAG_ConstraintRefAttributes()
	xml.Unmarshal(buf, v2)
}
