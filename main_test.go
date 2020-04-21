package main_test

import (
	c "context"
	"testing"

	"github.com/go-bdd/gobdd"
	"github.com/go-bdd/gobdd/context"
	sapp "github.com/go-bdd/sample-app"
)

type appKey struct{}
type createAccountErr struct{}
type loginErr struct{}

func TestFeatures(t *testing.T) {
	app := sapp.NewApp()
	s := gobdd.NewSuite(t, gobdd.WithBeforeScenario(func(ctx context.Context) {
		ctx.Set(appKey{}, app)
	}))
	s.AddStep(`I create a new account with username ([\da-zA-Z0-9]+)`, iCreateNewAccountWithUsername)
	s.AddStep(`the creation of the account (succeeded|failed)`, theCreationOfAccount)
	s.AddStep(`I log in to the system using username ([\da-zA-Z0-9]+)`, iLogInToTheSystemUsingUsername)
	s.AddStep(`the logging in ([\da-zA-Z0-9]+)`, loggingIn)
	go app.Run(c.Background())
	s.Run()
}

func iCreateNewAccountWithUsername(t gobdd.TestingT, ctx context.Context, u string) context.Context {
	app := getApp(ctx)
	err := app.CreateNewAccount(u, "pass")
	ctx.Set(createAccountErr{}, err)
	return ctx
}

func iLogInToTheSystemUsingUsername(t gobdd.TestingT, ctx context.Context, username string) context.Context {
	app := getApp(ctx)
	err := app.Login(username, "pass")
	ctx.Set(loginErr{}, err)
	return ctx
}

func loggingIn(t gobdd.TestingT, ctx context.Context, result string) context.Context {
	logErrRaw, err := ctx.Get(loginErr{})
	if err != nil {
		t.Error("missing login result")
		return ctx
	}

	if result == "succeeded" && logErrRaw != nil {
		t.Error("the login failed")
		return ctx
	}

	if result == "failed" && logErrRaw == nil {
		t.Error("the login succeeded")
	}

	return ctx
}

func theCreationOfAccount(t gobdd.TestingT, ctx context.Context, result string) context.Context {
	givenErr, err := ctx.Get(createAccountErr{})
	if err != nil {
		t.Error("could not find result of creating the account")
	}
	if result == "succeeded" && givenErr != nil {
		t.Error("expected to success but failed")
	}
	if result == "fails" && givenErr == nil {
		t.Error("expected to fail but succeeded")
	}
	return ctx
}

func getApp(ctx context.Context) *sapp.App {
	app, err := ctx.Get(appKey{})
	if err != nil {
		panic(err)
	}

	return app.(*sapp.App)
}
