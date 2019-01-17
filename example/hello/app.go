package hello

import . "github.com/RossMerr/Breeze"

// MyApp is our top application widget.
type MyApp struct {
	Core
	homePage *MyHomePage
}

// NewMyApp instantiates a new MyApp widget
func NewMyApp() *MyApp {
	app := &MyApp{}
	app.homePage = &MyHomePage{}
	return app
}

// Build renders the MyApp widget. Implements Widget interface.
func (m *MyApp) Build(ctx BuildContext) Widget {
	return m.homePage
}
