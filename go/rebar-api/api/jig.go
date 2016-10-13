package api

import (
	"path"

	"github.com/digitalrebar/digitalrebar/go/rebar-api/datatypes"
)

// Jig wraps datatypes.Jig to provide client API functionality
type Jig struct {
	datatypes.Jig
	Timestamps
	apiHelper
}

// Jigger is anything that a Jig can be bound to.
type Jigger interface {
	Crudder
	jigs()
}

// Jigs returns all of the Jigs.
func (c *Client) Jigs(scope ...Jigger) (res []*Jig, err error) {
	res = make([]*Jig, 0)
	paths := make([]string, len(scope))
	for i := range scope {
		paths[i] = fragTo(scope[i])
	}
	paths = append(paths, "jigs")
	return res, c.List(path.Join(datatypes.API_PATH, path.Join(paths...)), &res)
}
