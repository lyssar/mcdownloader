package utils

import (
	"github.com/gookit/color"
)

func PrintInfo(a interface{}) {
	color.Note.Print("\u2773 ")
	color.Note.Println(a)
}
