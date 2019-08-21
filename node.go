package breeze

type Node interface {
	//Render() string
}

type Nodes struct {
	children	[]Node
}


// func (s Node)String() string {
// 	return ""
// //	return `<{{.Params.Component}} class"MuiPaper-root MuiPaper-elevation1  MuiPaper-rounded">{{.}}</{{.Params.Component}}>`
// }