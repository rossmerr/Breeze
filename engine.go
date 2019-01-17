package breeze

type Engine interface {
	Run(widget Widget)
	Rerender(widget Widget)
}
