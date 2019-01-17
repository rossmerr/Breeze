package widgets

import (
	breeze "github.com/RossMerr/Breeze"
	"github.com/RossMerr/Breeze/mainAxisAlignment"
)

type Column struct {
	breeze.Widget
	Params ColumnParams
}

func NewColumn(params ColumnParams) breeze.Widget {
	return &Column{Params: params}
}

type ColumnParams struct {
	MainAxisAlignment mainAxisAlignment.MainAxisAlignment
	Children          []breeze.Widget
}
