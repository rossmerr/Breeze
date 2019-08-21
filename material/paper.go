package material

import (	
	"github.com/RossMerr/Breeze/widgets"
	"github.com/RossMerr/Breeze"

)


// Paper The background of an application resembles the flat, opaque texture of a sheet of paper, and an application’s behavior mimics paper’s ability to be re-sized, shuffled, and bound together in multiple sheets.
type Paper struct {
	Params PaperParams
	children	breeze.Nodes
}

// NewPaper creates a material design paper
func NewPaper(params PaperParams) Paper {
	if params.Component == "" {
		params.Component = widgets.ElementType("div")
	}

	return Paper{Params: params, children: breeze.Nodes{}}
}

func (s Paper) String() string {
	return `<{{.Params.Component}} class="MuiPaper-root MuiPaper-elevation{{.Params.Elevation}}  {{if .Params.Square }} MuiPaper-rounded {{end}}">{{ .children }}></{{.Params.Component}}>`
}

type PaperParams struct {
	// The component used for the root node. Either a string to use a DOM element or a component.
	Component widgets.ElementType
	// Shadow depth, corresponds to dp in the spec. It accepts values between 0 and 24 inclusive.
	Elevation int
	// If true, rounded corners are disabled.
	Square bool
}