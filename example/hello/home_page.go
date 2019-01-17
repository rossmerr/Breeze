package hello

import (
	"fmt"

	. "github.com/RossMerr/Breeze"
	"github.com/RossMerr/Breeze/mainAxisAlignment"
	. "github.com/RossMerr/Breeze/material"
	"github.com/RossMerr/Breeze/material/icons"
	. "github.com/RossMerr/Breeze/widgets"
)

// MyHomePage is a home page widget.
type MyHomePage struct {
	Core
	counter int
}

// Build renders the MyHomePage widget. Implements Widget interface.
func (m *MyHomePage) Build(ctx BuildContext) Widget {
	return NewScaffold(ScaffoldParams{
		AppBar: NewAppBar(AppBarParams{
			Title: NewText(TextParams{
				Text: "My Home Page",
			}),
		}),
		Body: NewCenter(CenterParams{
			Child: NewColumn(ColumnParams{
				MainAxisAlignment: mainAxisAlignment.Center,
				Children: []Widget{
					NewText(TextParams{
						Text: "You have pushed the button this many times:",
					}),
					NewText(TextParams{
						Text:  fmt.Sprintf("%d", m.counter),
						Style: ctx.Theme.TextTheme.Display1,
					}),
				},
			}),
		}),
		FloatingActionButton: NewFloatingActionButton(
			FloatingActionButtonParams{
				OnPressed: m.incrementCounter,
				Tooltip:   "Increment",
				Child: NewIcon(IconParams{
					Icon: icons.Add(),
				}),
			},
		),
	})
}

// incrementCounter increments app's counter by one.
func (m *MyHomePage) incrementCounter() {
	m.counter++
	Rerender(m)
}
