package controllers

import "github.com/robfig/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	greeting := "Aloha World"
	return c.Render(greeting)
}


func (c App) Hello(myName string) revel.Result {
	return c.Render(myName)
}