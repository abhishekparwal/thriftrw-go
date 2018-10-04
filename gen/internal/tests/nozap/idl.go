// Code generated by thriftrw v1.14.0. DO NOT EDIT.
// @generated

package nozap

import "go.uber.org/thriftrw/thriftreflect"

// ThriftModule represents the IDL file used to generate this package.
var ThriftModule = &thriftreflect.ThriftModule{
	Name:     "nozap",
	Package:  "go.uber.org/thriftrw/gen/internal/tests/nozap",
	FilePath: "nozap.thrift",
	SHA1:     "05f7228060eeb97fbd181a7a2660a6483799ac78",
	Raw:      rawIDL,
}

const rawIDL = "enum EnumDefault {\n    Foo, Bar, Baz\n}\n\nstruct PrimitiveRequiredStruct {\n    1: required bool boolField\n    2: required byte byteField\n    3: required i16 int16Field\n    4: required i32 int32Field\n    5: required i64 int64Field\n    6: required double doubleField\n    7: required string stringField\n    8: required binary binaryField\n    9: required list<string> listOfStrings\n    10: required set<i32> setOfInts\n    11: required map<i64, double> mapOfIntsToDoubles\n}\n\ntypedef map<string, string> StringMap\ntypedef PrimitiveRequiredStruct Primitives\ntypedef list<string> StringList\n"
