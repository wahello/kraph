package store

import (
	"github.com/milosgajdos/kraph/api"
	"github.com/milosgajdos/kraph/query"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding"
)

// DOTAttrs are attributes for Graphiz DOT graph
type DOTAttrs interface {
	Attrs
	// DOTAttributes aare required by gonum.org/v1/gonum/graph/encoding/dot
	DOTAttributes() []encoding.Attribute
}

// GraphAttributes are graph attributes
type GraphAttributes interface {
	// Attributes returns attributes as a slice of encoding.Attribute
	Attributes() []encoding.Attribute
}

// Attrs provide a simple key-value store
// for storing arbitrary entity attributes
type Attrs interface {
	GraphAttributes
	// Keys returns all attribute keys
	Keys() []string
	// Get returns the attribute value for the given key
	Get(string) string
	// Set sets the value of the attribute for the given key
	Set(string, string)
}

// Metadata provides a simple key-valule store
// for arbitrary entity data of arbitrary type
type Metadata interface {
	// Keys returns all metadata keys
	Keys() []string
	// Get returns the attribute value for the given key
	Get(string) interface{}
	// Set sets the value of the attribute for the given key
	Set(string, interface{})
}

// Entity is an arbitrary store entity
type Entity interface {
	// Attrs returns attributes
	Attrs() Attrs
	// Metadata returns metadata
	Metadata() Metadata
	// Attributes returns attributes as a slice of encoding.Attribute
	Attributes() []encoding.Attribute
}

// DOTNode is a GraphViz DOT Node
type DOTNode interface {
	Node
	// DOTID returns Graphiz DOT ID
	DOTID() string
	// SetDOTID sets Graphiz DOT ID
	SetDOTID(string)
}

// Node is a graph node
type Node interface {
	Entity
	graph.Node
	// Name returns node name
	Name() string
}

// Edge is an edge between two nodes
type Edge interface {
	Entity
	graph.WeightedEdge
}

// Graph is a graph of API objects
type Graph interface {
	graph.Graph
	// Subgraph returns a subgraph of the graph starting at Node
	// up to the given depth or it returns an error
	SubGraph(Node, int) (graph.Graph, error)
}

// DOTGraph returns Graphiz DOT store
type DOTGraph interface {
	Graph
	// DOTID returns DOT graph ID
	DOTID() string
	// DOTAttributers returns global graph DOT attributes
	DOTAttributers() (graph, node, edge encoding.Attributer)
	// DOT returns Graphviz graph
	DOT() (string, error)
}

// Store allows to store and query the graph of API objects
type Store interface {
	Graph
	// Add adds an api.Object to the store and returns a Node
	Add(api.Object, ...Option) (Node, error)
	// Link links two nodes and returns the new edge between them
	Link(Node, Node, ...Option) (Edge, error)
	// Delete deletes an entity from the store
	Delete(Entity, ...Option) error
	// Query queries the store and returns the results
	Query(...query.Option) ([]Entity, error)
}