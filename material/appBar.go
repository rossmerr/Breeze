package material

import "github.com/RossMerr/Breeze/widgets"

// AppBar a material design app bar
type AppBar struct {
	Params AppBarParams
}

// NewAppBar creates a material design app bar
func NewAppBar(params AppBarParams) AppBar {
	return AppBar{Params: params}
}

func (s AppBar)String() string {
	return ""
}

type AppBarParams struct {
	Title widgets.Text
}
