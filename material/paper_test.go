package material_test

import (
	"html/template"
	"testing"
	"os"

	
	"github.com/RossMerr/Breeze/material"
)

func Test_Paper(t *testing.T) {
	paper := material.NewPaper(material.PaperParams{})

	
    tmpl, _ := template.New("test").Parse(paper.String())
	tmpl.Execute(os.Stdout, paper)
}