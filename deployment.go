package crowbar

import (
	"log"
	"strconv"
)
// Deployments are the main tool that Crowbar provides to group
// related nodes together.  They carry Deployment-specific
// configuration information (in the form of DeploymentRoles), Nodes
// (a node belongs to exactly one Deployment at any given point in
// time), and NodeRoles (NodeRoles remember which Deployment they were
// first committed in, and always use that when calculating the attrib
// data to be passed in to a run).
//
// Deployments are structured as a tree with the system deployment at
// the root.  All newly-discovered Nodes start in the system
// deployment, where they have their inital set of Roles bound to
// them.
//
// Deployment satisfies the Attriber interface.  GetAttrib and
// SetAttrib calls on a Deployment will be redirected to the
// appropriate DeploymentRole in this Deployment.
type Deployment struct {
	// The unique identifier for a Deployment.
	ID          int64  `json:"id,omitempty"`
	// The state a deployment is in
	State       int    `json:"state,omitempty"`
	// The name of the Deployment.  Must be globally unique.
	Name        string `json:"name,omitempty"`
	// A breif description of what the Deployment is for.
	Description string `json:"description,omitempty"`
	// Whether the deployment is a system deployment.  Right now,
	// there can be only one of these.
	System      bool   `json:"system,omitempty"`
	// The ID of the deployment that is the parent of this one.
	ParentID    int64  `json:"parent_id,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

const (
	DeploymentError = -1
	DeploymentProposed = 0
	DeploymentCommitted = 1
	DeploymentActive = 2
)

func (o *Deployment) Id() string {
	if o.ID != 0 {
		return strconv.FormatInt(o.ID, 10)
	} else if o.Name != "" {
		return o.Name
	} else {
		log.Panic("Deployment has no ID or name")
		return ""
	}
}

func (o *Deployment) ApiName() string {
	return "deployments"
}

// Attribs fetches all the attribs relavent to this Deployment.  Right
// now, that consists of all the Attribs for all the DeploymentRoles
// in this Deployment
func (o *Deployment) Attribs() (res []*Attrib, err error) {
	return Attribs(url(o))
}

// Nodes returns all the Nodes that are in this Deployment.
func (o *Deployment) Nodes() (res []*Node, err error) {
	return Nodes(url(o))
}

// NodeRoles returns all the NodeRoles that are in this Deployment.
func (o *Deployment) NodeRoles() (res []*NodeRole, err error) {
	return NodeRoles(url(o))
}

// Roles returns all the Roles that are in this Deployment.
func (o *Deployment) Roles() (res []*Role, err error) {
	return Roles(url(o))
}

// Deployments returns all of the Deployments.
func Deployments(paths ...string) (res []*Deployment, err error) {
	res = make([]*Deployment, 0)
	return res, session.list(&res,append(paths, "deployments")...)
}
