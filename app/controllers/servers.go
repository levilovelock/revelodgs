package controllers

import (
	"rps/app/models"
	"rps/app/routes"

	"github.com/revel/revel"
)

// Responsible for managing interactions with servers in a REST like way
type Servers struct {
	Application
}

func (c *Servers) checkUser() revel.Result {
	if user := c.connected(); user == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(routes.Application.Index())
	}
	return nil
}

func (c *Servers) List() revel.Result {
	results, err := c.Txn.Select(models.Server{},
		`select * from Server where userid = ?`, c.connected().UserId)
	if err != nil {
		panic(err)
	}

	var servers []*models.Server
	for _, r := range results {
		s := r.(*models.Server)
		servers = append(servers, s)
	}

	return c.RenderJson(servers)
}

func (c *Servers) ListGames() revel.Result {
	results, err := c.Txn.Select(models.Game{}, `select * from Game`)
	if err != nil {
		panic(err)
	}

	var Games []*models.Game
	for _, r := range results {
		g := r.(*models.Game)
		Games = append(Games, g)
	}

	return c.RenderJson(Games)
}

func (c *Servers) Show(id int) revel.Result {
	return c.Render()
}

func (c *Servers) Delete(id int) revel.Result {
	return c.Render()
}

func (c *Servers) New() revel.Result {
	return c.Render()
}

func (c *Servers) Create() revel.Result {
	return c.Render()
}
