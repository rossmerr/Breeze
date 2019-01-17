package material

import breeze "github.com/RossMerr/Breeze"

// Scaffold implements the basic material design visual layout structure.
type Scaffold struct {
	breeze.Widget
	Params ScaffoldParams
}

// NewScaffold creates a visual scaffold for material design widgets.
func NewScaffold(params ScaffoldParams) breeze.Widget {
	return &Scaffold{Params: params}
}

type ScaffoldParams struct {
	AppBar               AppBar
	Body                 breeze.Widget
	FloatingActionButton FloatingActionButton
}
