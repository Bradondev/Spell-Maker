package main

import (
	"context"
	"fmt"
	"os"
)

type Spell struct{
	Name string `json:"name"`
	Description string `json:"des"`
	NameOfType string 
	LevelOfSpell string
	Stats SpellStats
}



type SpellStats struct{
	damage int32
	Debuff string
	ManaCost int32
	CastTime int32
	Range string
}





 var Spells = []Spell{}
 var NewSpell Spell
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

func(a *App) MakeNewSpell(Name string ,Des string,Level string, SpellType string) {
	NewSpell.Name = Name
	NewSpell.Description = Des
	NewSpell.LevelOfSpell = Level
	NewSpell.NameOfType= SpellType
	Spells = append(Spells, NewSpell)
	fmt.Println(Spells)
}

func (a *App) SaveNewSpell(newSpell Spell) {
	err := os.WriteFile(newSpell.Name + ".txt",[]byte(newSpell.Description), 0644) //create a new file
	if err != nil {
	   fmt.Println(err)
	   return
	}
	fmt.Println("File is created successfully.")
}