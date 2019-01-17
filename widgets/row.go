package widgets

import (
	breeze "github.com/RossMerr/Breeze"
	"github.com/RossMerr/Breeze/widgets/mainAxisAlignment"
)

type Row struct {
	breeze.Widget
	Params RowParams
}

func NewRow(params RowParams) breeze.Widget {
	return &Row{Params: params}
}

type RowParams struct {
	MainAxisAlignment mainAxisAlignment.MainAxisAlignment
	Children          []breeze.Widget
}
