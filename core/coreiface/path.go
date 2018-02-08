package iface

import (
	cid "gx/ipfs/QmYVNvtQkeZ6AKSwDrjQTs432QtL6umrrK41EBq3cu7iSP/go-cid"
)

// Path is a generic wrapper for paths used in the API. A path can be resolved
// to a CID using one of Resolve functions in the API.
// TODO: figure out/explain namespaces
type Path interface {
	// String returns the path as a string.
	String() string

	// Namespace returns the first component of the path
	Namespace() string
}

// ResolvedPath is a resolved Path
type ResolvedPath interface {
	// Cid returns cid referred to by path
	Cid() *cid.Cid

	// Root returns cid of root path
	Root() *cid.Cid

	//TODO: Path remainder

	Path
}
