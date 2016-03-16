package controllers

import "github.com/revel/revel"

type Ng struct {
	*revel.Controller
}

func (c Ng) Index() revel.Result {
	return c.Render()
}

func (c Ng) Note() revel.Result {
	return c.Render()
}
