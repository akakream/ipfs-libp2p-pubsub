package dagpb

// Code generated by go-ipld-prime gengo.  DO NOT EDIT.

import (
	ipld "github.com/ipld/go-ipld-prime"
)

var _ ipld.Node = nil // suppress errors when this dependency is not referenced
// Type is a struct embeding a NodePrototype/Type for every Node implementation in this package.
// One of its major uses is to start the construction of a value.
// You can use it like this:
//
//	dagpb.Type.YourTypeName.NewBuilder().BeginMap() //...
//
// and:
//
//	dagpb.Type.OtherTypeName.NewBuilder().AssignString("x") // ...
var Type typeSlab

type typeSlab struct {
	Bytes         _Bytes__Prototype
	Bytes__Repr   _Bytes__ReprPrototype
	Int           _Int__Prototype
	Int__Repr     _Int__ReprPrototype
	Link          _Link__Prototype
	Link__Repr    _Link__ReprPrototype
	PBLink        _PBLink__Prototype
	PBLink__Repr  _PBLink__ReprPrototype
	PBLinks       _PBLinks__Prototype
	PBLinks__Repr _PBLinks__ReprPrototype
	PBNode        _PBNode__Prototype
	PBNode__Repr  _PBNode__ReprPrototype
	String        _String__Prototype
	String__Repr  _String__ReprPrototype
}

// --- type definitions follow ---

// Bytes matches the IPLD Schema type "Bytes".  It has bytes kind.
type Bytes = *_Bytes
type _Bytes struct{ x []byte }

// Int matches the IPLD Schema type "Int".  It has int kind.
type Int = *_Int
type _Int struct{ x int64 }

// Link matches the IPLD Schema type "Link".  It has link kind.
type Link = *_Link
type _Link struct{ x ipld.Link }

// PBLink matches the IPLD Schema type "PBLink".  It has Struct type-kind, and may be interrogated like map kind.
type PBLink = *_PBLink
type _PBLink struct {
	Hash  _Link
	Name  _String__Maybe
	Tsize _Int__Maybe
}

// PBLinks matches the IPLD Schema type "PBLinks".  It has list kind.
type PBLinks = *_PBLinks
type _PBLinks struct {
	x []_PBLink
}

// PBNode matches the IPLD Schema type "PBNode".  It has Struct type-kind, and may be interrogated like map kind.
type PBNode = *_PBNode
type _PBNode struct {
	Links _PBLinks
	Data  _Bytes__Maybe
}

// String matches the IPLD Schema type "String".  It has string kind.
type String = *_String
type _String struct{ x string }
