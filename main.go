package main

import (
	"context"
	"errors"
)

func main() {

}

type App struct {
	users map[string]string
}

func NewApp() *App {
	return &App{
		users: map[string]string{},
	}
}

func (app *App) Run(ctx context.Context) {
}

func (app *App) CreateNewAccount(u, p string) error {
	app.users[u] = p
	return nil
}

func (app *App) Login(username, password string) error {
	for u, p := range app.users {
		if u == username && p == password {
			return nil
		}
	}

	return errors.New("invalid login or password")
}
