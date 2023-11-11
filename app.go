package main

import (
	"context"
	"fmt"
	"os"
)

type Spell struct{
	Name string
}


// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}


func (a *App) MakeNewSpell(newSpell Spell) {
	err := os.WriteFile("myfile.txt", []byte("Hello, alexa!"), 0644) //create a new file
	if err != nil {
	   fmt.Println(err)
	   return
	}
	fmt.Println("File is created successfully.")
}