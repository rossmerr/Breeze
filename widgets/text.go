package widgets

import (
	breeze "github.com/RossMerr/Breeze"
	"github.com/RossMerr/Breeze/theme"
)

type Text struct {
	breeze.Widget

	Params TextParams
}

func NewText(params TextParams) Text {
	return Text{Params: params}
}

type TextParams struct {
	Text  string
	Style theme.TextStyle
}
