// +build codegen

package api

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestUniqueInputAndOutputs(t *testing.T) {
	cases := [][]struct {
		expectedInput  string
		expectedOutput string
		operation      string
		input          string
		output         string
	}{
		{
			{
				expectedInput:  "FooOperationInput",
				expectedOutput: "FooOperationOutput",
				operation:      "FooOperation",
				input:          "FooInputShape",
				output:         "FooOutputShape",
			},
			{
				expectedInput:  "BarOperationInput",
				expectedOutput: "BarOperationOutput",
				operation:      "BarOperation",
				input:          "FooInputShape",
				output:         "FooOutputShape",
			},
		},
		{
			{
				expectedInput:  "FooOperationInput",
				expectedOutput: "FooOperationOutput",
				operation:      "FooOperation",
				input:          "FooInputShape",
				output:         "FooOutputShape",
			},
		},
		{
			{
				expectedInput:  "FooOperationInput",
				expectedOutput: "FooOperationOutput",
				operation:      "FooOperation",
				input:          "FooInputShape",
				output:         "FooOutputShape",
			},
		},
		{
			{
				expectedInput:  "FooOperationInput",
				expectedOutput: "FooOperationOutput",
				operation:      "FooOperation",
				input:          "FooInputShape",
				output:         "FooOutputShape",
			},
		},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			a := &API{
				Operations: map[string]*Operation{},
				Shapes:     map[string]*Shape{},
			}

			expected := map[string][]string{}
			for _, op := range c {
				o := &Operation{
					Name:         op.operation,
					ExportedName: op.operation,
					InputRef: ShapeRef{
						API:       a,
						ShapeName: op.input,
						Shape: &Shape{
							API:       a,
							ShapeName: op.input,
						},
					},
					OutputRef: ShapeRef{
						API:       a,
						ShapeName: op.input,
						Shape: &Shape{
							API:       a,
							ShapeName: op.input,
						},
					},
				}
				o.InputRef.Shape.refs = append(o.InputRef.Shape.refs, &o.InputRef)
				o.OutputRef.Shape.refs = append(o.OutputRef.Shape.refs, &o.OutputRef)

				a.Operations[o.Name] = o

				a.Shapes[op.input] = o.InputRef.Shape
				a.Shapes[op.output] = o.OutputRef.Shape

				expected[op.operation] = append(expected[op.operation],
					op.expectedInput,
					op.expectedOutput,
				)
			}

			a.fixStutterNames()
			a.applyShapeNameAliases()
			a.createInputOutputShapes()
			for k, v := range expected {
				if a.Operations[k].InputRef.Shape.ShapeName != v[0] {
					t.Errorf("Error %s case: Expected %q, but received %q", k, v[0], a.Operations[k].InputRef.Shape.ShapeName)
				}
				if a.Operations[k].OutputRef.Shape.ShapeName != v[1] {
					t.Errorf("Error %s case: Expected %q, but received %q", k, v[1], a.Operations[k].OutputRef.Shape.ShapeName)
				}
			}
		})

	}
}

func TestCollidingFields(t *testing.T) {
	cases := map[string]struct {
		MemberRefs  map[string]*ShapeRef
		Expect      []string
		IsException bool
	}{
		"SimpleMembers": {
			MemberRefs: map[string]*ShapeRef{
				"Code":     {},
				"Foo":      {},
				"Message":  {},
				"OrigErr":  {},
				"String":   {},
				"Validate": {},
			},
			Expect: []string{
				"Code",
				"Foo",
				"Message",
				"OrigErr",
				"String_",
				"Validate_",
			},
		},
	}

	for k, c := range cases {
		t.Run(k, func(t *testing.T) {
			a := &API{
				Shapes: map[string]*Shape{
					"shapename": {
						ShapeName:  k,
						MemberRefs: c.MemberRefs,
						Exception:  c.IsException,
					},
				},
			}

			a.renameCollidingFields()

			for i, name := range a.Shapes["shapename"].MemberNames() {
				if e, a := c.Expect[i], name; e != a {
					t.Errorf("expect %v, got %v", e, a)
				}
			}
		})
	}
}

func TestCreateInputOutputShapes(t *testing.T) {
	meta := Metadata{
		APIVersion:          "0000-00-00",
		EndpointPrefix:      "rpcservice",
		JSONVersion:         "1.1",
		Protocol:            "json",
		ServiceAbbreviation: "RPCService",
		ServiceFullName:     "RPC Service",
		ServiceID:           "RPCService",
		SignatureVersion:    "v4",
		TargetPrefix:        "RPCService_00000000",
		UID:                 "RPCService-0000-00-00",
	}

	type OpExpect struct {
		Input  string
		Output string
	}

	cases := map[string]struct {
		API          *API
		ExpectOps    map[string]OpExpect
		ExpectShapes []string
	}{
		"allRename": {
			API: &API{Metadata: meta,
				Operations: map[string]*Operation{
					"FirstOp": {Name: "FirstOp",
						InputRef:  ShapeRef{ShapeName: "FirstOpRequest"},
						OutputRef: ShapeRef{ShapeName: "FirstOpResponse"},
					},
					"SecondOp": {Name: "SecondOp",
						InputRef:  ShapeRef{ShapeName: "SecondOpRequest"},
						OutputRef: ShapeRef{ShapeName: "SecondOpResponse"},
					},
				},
				Shapes: map[string]*Shape{
					"FirstOpRequest":   {ShapeName: "FirstOpRequest", Type: "structure"},
					"FirstOpResponse":  {ShapeName: "FirstOpResponse", Type: "structure"},
					"SecondOpRequest":  {ShapeName: "SecondOpRequest", Type: "structure"},
					"SecondOpResponse": {ShapeName: "SecondOpResponse", Type: "structure"},
				},
			},
			ExpectOps: map[string]OpExpect{
				"FirstOp": {
					Input:  "FirstOpInput",
					Output: "FirstOpOutput",
				},
				"SecondOp": {
					Input:  "SecondOpInput",
					Output: "SecondOpOutput",
				},
			},
			ExpectShapes: []string{
				"FirstOpInput", "FirstOpOutput",
				"SecondOpInput", "SecondOpOutput",
			},
		},
		"noRename": {
			API: &API{Metadata: meta,
				Operations: map[string]*Operation{
					"FirstOp": {Name: "FirstOp",
						InputRef:  ShapeRef{ShapeName: "FirstOpInput"},
						OutputRef: ShapeRef{ShapeName: "FirstOpOutput"},
					},
					"SecondOp": {Name: "SecondOp",
						InputRef:  ShapeRef{ShapeName: "SecondOpInput"},
						OutputRef: ShapeRef{ShapeName: "SecondOpOutput"},
					},
				},
				Shapes: map[string]*Shape{
					"FirstOpInput":   {ShapeName: "FirstOpInput", Type: "structure"},
					"FirstOpOutput":  {ShapeName: "FirstOpOutput", Type: "structure"},
					"SecondOpInput":  {ShapeName: "SecondOpInput", Type: "structure"},
					"SecondOpOutput": {ShapeName: "SecondOpOutput", Type: "structure"},
				},
			},
			ExpectOps: map[string]OpExpect{
				"FirstOp": {
					Input:  "FirstOpInput",
					Output: "FirstOpOutput",
				},
				"SecondOp": {
					Input:  "SecondOpInput",
					Output: "SecondOpOutput",
				},
			},
			ExpectShapes: []string{
				"FirstOpInput", "FirstOpOutput",
				"SecondOpInput", "SecondOpOutput",
			},
		},
		"renameWithNested": {
			API: &API{Metadata: meta,
				Operations: map[string]*Operation{
					"FirstOp": {Name: "FirstOp",
						InputRef:  ShapeRef{ShapeName: "FirstOpWriteMe"},
						OutputRef: ShapeRef{ShapeName: "FirstOpReadMe"},
					},
					"SecondOp": {Name: "SecondOp",
						InputRef:  ShapeRef{ShapeName: "SecondOpWriteMe"},
						OutputRef: ShapeRef{ShapeName: "SecondOpReadMe"},
					},
				},
				Shapes: map[string]*Shape{
					"FirstOpWriteMe": {ShapeName: "FirstOpWriteMe", Type: "structure",
						MemberRefs: map[string]*ShapeRef{
							"Foo": {ShapeName: "String"},
						},
					},
					"FirstOpReadMe": {ShapeName: "FirstOpReadMe", Type: "structure",
						MemberRefs: map[string]*ShapeRef{
							"Bar":  {ShapeName: "Struct"},
							"Once": {ShapeName: "Once"},
						},
					},
					"SecondOpWriteMe": {ShapeName: "SecondOpWriteMe", Type: "structure"},
					"SecondOpReadMe":  {ShapeName: "SecondOpReadMe", Type: "structure"},
					"Once":            {ShapeName: "Once", Type: "string"},
					"String":          {ShapeName: "String", Type: "string"},
					"Struct": {ShapeName: "Struct", Type: "structure",
						MemberRefs: map[string]*ShapeRef{
							"Foo": {ShapeName: "String"},
							"Bar": {ShapeName: "Struct"},
						},
					},
				},
			},
			ExpectOps: map[string]OpExpect{
				"FirstOp": {
					Input:  "FirstOpInput",
					Output: "FirstOpOutput",
				},
				"SecondOp": {
					Input:  "SecondOpInput",
					Output: "SecondOpOutput",
				},
			},
			ExpectShapes: []string{
				"FirstOpInput", "FirstOpOutput",
				"Once",
				"SecondOpInput", "SecondOpOutput",
				"String", "Struct",
			},
		},
		"aliasedInput": {
			API: &API{Metadata: meta,
				Operations: map[string]*Operation{
					"FirstOp": {Name: "FirstOp",
						InputRef:  ShapeRef{ShapeName: "FirstOpRequest"},
						OutputRef: ShapeRef{ShapeName: "FirstOpResponse"},
					},
				},
				Shapes: map[string]*Shape{
					"FirstOpRequest": {ShapeName: "FirstOpRequest", Type: "structure",
						AliasedShapeName: true,
					},
					"FirstOpResponse": {ShapeName: "FirstOpResponse", Type: "structure"},
				},
			},
			ExpectOps: map[string]OpExpect{
				"FirstOp": {
					Input:  "FirstOpRequest",
					Output: "FirstOpOutput",
				},
			},
			ExpectShapes: []string{
				"FirstOpOutput", "FirstOpRequest",
			},
		},
		"aliasedOutput": {
			API: &API{Metadata: meta,
				Operations: map[string]*Operation{
					"FirstOp": {Name: "FirstOp",
						InputRef:  ShapeRef{ShapeName: "FirstOpRequest"},
						OutputRef: ShapeRef{ShapeName: "FirstOpResponse"},
					},
				},
				Shapes: map[string]*Shape{
					"FirstOpRequest": {ShapeName: "FirstOpRequest", Type: "structure"},
					"FirstOpResponse": {ShapeName: "FirstOpResponse", Type: "structure",
						AliasedShapeName: true,
					},
				},
			},
			ExpectOps: map[string]OpExpect{
				"FirstOp": {
					Input:  "FirstOpInput",
					Output: "FirstOpResponse",
				},
			},
			ExpectShapes: []string{
				"FirstOpInput", "FirstOpResponse",
			},
		},
		"resusedShape": {
			API: &API{Metadata: meta,
				Operations: map[string]*Operation{
					"FirstOp": {Name: "FirstOp",
						InputRef:  ShapeRef{ShapeName: "FirstOpRequest"},
						OutputRef: ShapeRef{ShapeName: "ReusedShape"},
					},
				},
				Shapes: map[string]*Shape{
					"FirstOpRequest": {ShapeName: "FirstOpRequest", Type: "structure",
						MemberRefs: map[string]*ShapeRef{
							"Foo": {ShapeName: "ReusedShape"},
							"ooF": {ShapeName: "ReusedShapeList"},
						},
					},
					"ReusedShape": {ShapeName: "ReusedShape", Type: "structure"},
					"ReusedShapeList": {ShapeName: "ReusedShapeList", Type: "list",
						MemberRef: ShapeRef{ShapeName: "ReusedShape"},
					},
				},
			},
			ExpectOps: map[string]OpExpect{
				"FirstOp": {
					Input:  "FirstOpInput",
					Output: "FirstOpOutput",
				},
			},
			ExpectShapes: []string{
				"FirstOpInput", "FirstOpOutput",
				"ReusedShape", "ReusedShapeList",
			},
		},
		"aliasedResusedShape": {
			API: &API{Metadata: meta,
				Operations: map[string]*Operation{
					"FirstOp": {Name: "FirstOp",
						InputRef:  ShapeRef{ShapeName: "FirstOpRequest"},
						OutputRef: ShapeRef{ShapeName: "ReusedShape"},
					},
				},
				Shapes: map[string]*Shape{
					"FirstOpRequest": {ShapeName: "FirstOpRequest", Type: "structure",
						MemberRefs: map[string]*ShapeRef{
							"Foo": {ShapeName: "ReusedShape"},
							"ooF": {ShapeName: "ReusedShapeList"},
						},
					},
					"ReusedShape": {ShapeName: "ReusedShape", Type: "structure",
						AliasedShapeName: true,
					},
					"ReusedShapeList": {ShapeName: "ReusedShapeList", Type: "list",
						MemberRef: ShapeRef{ShapeName: "ReusedShape"},
					},
				},
			},
			ExpectOps: map[string]OpExpect{
				"FirstOp": {
					Input:  "FirstOpInput",
					Output: "ReusedShape",
				},
			},
			ExpectShapes: []string{
				"FirstOpInput",
				"ReusedShape", "ReusedShapeList",
			},
		},
		"unsetInput": {
			API: &API{Metadata: meta,
				Operations: map[string]*Operation{
					"FirstOp": {Name: "FirstOp",
						OutputRef: ShapeRef{ShapeName: "FirstOpResponse"},
					},
				},
				Shapes: map[string]*Shape{
					"FirstOpResponse": {ShapeName: "FirstOpResponse", Type: "structure"},
				},
			},
			ExpectOps: map[string]OpExpect{
				"FirstOp": {
					Input:  "FirstOpInput",
					Output: "FirstOpOutput",
				},
			},
			ExpectShapes: []string{
				"FirstOpInput", "FirstOpOutput",
			},
		},
		"unsetOutput": {
			API: &API{Metadata: meta,
				Operations: map[string]*Operation{
					"FirstOp": {Name: "FirstOp",
						InputRef: ShapeRef{ShapeName: "FirstOpRequest"},
					},
				},
				Shapes: map[string]*Shape{
					"FirstOpRequest": {ShapeName: "FirstOpRequest", Type: "structure"},
				},
			},
			ExpectOps: map[string]OpExpect{
				"FirstOp": {
					Input:  "FirstOpInput",
					Output: "FirstOpOutput",
				},
			},
			ExpectShapes: []string{
				"FirstOpInput", "FirstOpOutput",
			},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			a := c.API
			err := a.Setup()
			if err != nil {
				t.Error(err)
			}

			for opName, op := range a.Operations {
				if e, a := op.InputRef.ShapeName, op.InputRef.Shape.ShapeName; e != a {
					t.Errorf("expect input ref and shape names to match, %s, %s", e, a)
				}
				if e, a := c.ExpectOps[opName].Input, op.InputRef.ShapeName; e != a {
					t.Errorf("expect %v input shape, got %v", e, a)
				}

				if e, a := op.OutputRef.ShapeName, op.OutputRef.Shape.ShapeName; e != a {
					t.Errorf("expect output ref and shape names to match, %s, %s", e, a)
				}
				if e, a := c.ExpectOps[opName].Output, op.OutputRef.ShapeName; e != a {
					t.Errorf("expect %v output shape, got %v", e, a)
				}
			}

			if e, a := c.ExpectShapes, a.ShapeNames(); !reflect.DeepEqual(e, a) {
				t.Errorf("expect %v shapes, got %v", e, a)
			}
		})
	}
}

func TestValidateShapeNameMethod(t *testing.T) {
	cases := map[string]struct {
		inputShapeName    string
		shapeType         string
		expectedShapeName string
		expectedError     string
	}{
		"empty case": {
			inputShapeName:    "",
			shapeType:         "structure",
			expectedShapeName: "",
			expectedError:     "invalid shape name found",
		},
		"No rename": {
			inputShapeName:    "Sample123Shape",
			shapeType:         "structure",
			expectedShapeName: "Sample123Shape",
		},
		"starts with underscores": {
			inputShapeName:    "__Sample123Shape",
			shapeType:         "structure",
			expectedShapeName: "Sample123Shape",
		},
		"Contains underscores": {
			inputShapeName:    "__sample_123_shape__",
			shapeType:         "structure",
			expectedShapeName: "Sample123Shape",
		},
		"Starts with numeric character": {
			inputShapeName:    "123__sampleShape",
			shapeType:         "structure",
			expectedShapeName: "",
			expectedError:     "invalid shape name found",
		},
		"Starts with non alphabetic or non underscore character": {
			inputShapeName:    "&&SampleShape",
			shapeType:         "structure",
			expectedShapeName: "",
			expectedError:     "invalid shape name found",
		},
		"Contains non Alphanumeric or non underscore character": {
			inputShapeName:    "Sample&__Shape",
			shapeType:         "structure",
			expectedShapeName: "",
			expectedError:     "invalid shape name found",
		},
		"Renamed Shape already exists": {
			inputShapeName:    "__sample_shape",
			shapeType:         "structure",
			expectedShapeName: "",
			expectedError:     "rename would result in shape name collision",
		},
		"empty case for enums shape type": {
			inputShapeName:    "",
			shapeType:         "string",
			expectedShapeName: "",
			expectedError:     "invalid shape name found",
		},
		"No rename for enums shape type": {
			inputShapeName:    "Sample123Shape",
			shapeType:         "string",
			expectedShapeName: "Sample123Shape",
		},
		"starts with underscores for enums shape type": {
			inputShapeName:    "__Sample123Shape",
			shapeType:         "string",
			expectedShapeName: "Sample123Shape",
		},
		"Contains underscores for enums shape type": {
			inputShapeName:    "__sample_123_shape__",
			shapeType:         "string",
			expectedShapeName: "Sample123Shape",
		},
		"Starts with numeric character for enums shape type": {
			inputShapeName:    "123__sampleShape",
			shapeType:         "string",
			expectedShapeName: "",
			expectedError:     "invalid shape name found",
		},
		"Starts with non alphabetic or non underscore character for enums shape type": {
			inputShapeName:    "&&SampleShape",
			shapeType:         "string",
			expectedShapeName: "",
			expectedError:     "invalid shape name found",
		},
		"Contains non Alphanumeric or non underscore character for enums shape type": {
			inputShapeName:    "Sample&__Shape",
			shapeType:         "string",
			expectedShapeName: "",
			expectedError:     "invalid shape name found",
		},
		"Renamed Shape already exists for enums shape type": {
			inputShapeName:    "__sample_shape",
			shapeType:         "string",
			expectedShapeName: "",
			expectedError:     "rename would result in shape name collision",
		},
	}

	for name, c := range cases {
		operation := "FooOperation"
		t.Run(name, func(t *testing.T) {
			a := &API{
				Operations: map[string]*Operation{},
				Shapes:     map[string]*Shape{},
			}
			// add another shape with name SampleShape to check for collision
			a.Shapes["SampleShape"] = &Shape{ShapeName: "SampleShape"}
			o := &Operation{
				Name:         operation,
				ExportedName: operation,
				InputRef: ShapeRef{
					API:       a,
					ShapeName: c.inputShapeName,
					Shape: &Shape{
						API:       a,
						ShapeName: c.inputShapeName,
						Type:      c.shapeType,
						Enum:      []string{"x"},
					},
				},
			}
			o.InputRef.Shape.refs = append(o.InputRef.Shape.refs, &o.InputRef)
			a.Operations[o.Name] = o
			a.Shapes[c.inputShapeName] = o.InputRef.Shape

			err := a.validateShapeNames()
			if err != nil || c.expectedError != "" {
				if err == nil {
					t.Fatalf("Received no error, expected error with log: \n \t %v ", c.expectedError)
				}
				if c.expectedError == "" {
					t.Fatalf("Expected no error, got %v", err.Error())
				}
				if e, a := err.Error(), c.expectedError; !strings.Contains(e, a) {
					t.Fatalf("Expected to receive error containing %v, got %v", e, a)
				}
				return
			}

			if e, a := c.expectedShapeName, o.InputRef.Shape.ShapeName; e != a {
				t.Fatalf("Expected shape name to be %v, got %v", e, a)
			}
		})
	}
}
