package material

import breeze "github.com/RossMerr/Breeze"

// FloatingActionButton a material design floating action button
type FloatingActionButton struct {
	Params FloatingActionButtonParams
}

// NewFloatingActionButton creates a circular floating action button
func NewFloatingActionButton(params FloatingActionButtonParams) FloatingActionButton {
	return FloatingActionButton{Params: params}
}

type FloatingActionButtonParams struct {
	OnPressed func()
	Tooltip   string
	Child     breeze.Widget
}
