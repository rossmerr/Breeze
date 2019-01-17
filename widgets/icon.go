package widgets

import (
	breeze "github.com/RossMerr/Breeze"
	"github.com/RossMerr/Breeze/textDirection"
)

// Icon a graphical icon widget drawn with a glyph from a font described in an IconData such as material's predefined IconDatas in Icons
type Icon struct {
	breeze.Widget
	Params IconParams
}

// NewIcon creates an icon
func NewIcon(params IconParams) breeze.Widget {
	return &Icon{Params: params}
}

type IconParams struct {
	Icon          IconData
	Color         Color
	Size          int
	TextDirection textDirection.TextDirection
}

// IconData a description of an icon fulfilled by a font glyph.
type IconData struct {
	CodePoint          int
	FontFamily         string
	FontPackage        string
	MatchTextDirection bool
}
