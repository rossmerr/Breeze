package widgets

import breeze "github.com/RossMerr/Breeze"

type Center struct {
	breeze.Widget
	Params CenterParams
}

func NewCenter(params CenterParams) breeze.Widget {
	return &Center{Params: params}
}

type CenterParams struct {
	Child breeze.Widget
}
