package material

import (
	breeze "github.com/RossMerr/Breeze"
	"github.com/RossMerr/Breeze/theme"
	"github.com/RossMerr/Breeze/ui"
	"github.com/RossMerr/Breeze/widgets"
)

// AppBar a material design app bar
type AppBar struct {
	Params AppBarParams
}

// NewAppBar creates a material design app bar
func NewAppBar(params AppBarParams) AppBar {
	return AppBar{Params: params}
}

func (s AppBar) String() string {
	return ""
}

type AppBarParams struct {
	Action           breeze.Widget
	ActionsIconTheme theme.IconThemeData
	BackgroundColor  ui.Color
	Bottom           PreferredSizeWidget
	ButtomOpacity    float32
	Brightness       ui.Brightness
	CenterTitle      bool
	Elevation        float32
	FlexibleSpace    breeze.Widget
	IconTheme        theme.IconThemeData
	Leading          breeze.Widget
	PreferredSize    ui.Size
	Primary          bool
	TextTheme        theme.TextTheme
	Title            widgets.Text
	TitleSpacing     float32
	ToolbarOpacity   float32
}
