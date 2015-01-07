package controllers

import (
	"time"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Run() revel.Result {
	return c.Render()
}

func (c App) Results() revel.Result {
	return c.Render()
}

func (c App) Running() revel.Result {
	for {
		time.Sleep(6)
		t := time.Now().String()
		return c.Render(t)
	}
}
