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


func MakeNewDirectory(NewSpellType string){
	_, err := os.Open(NewSpellType)
	if err != nil {
	if os.IsNotExist(err) {
		directoryname := NewSpellType
		err2 := os.Mkdir(directoryname, 0755) //create a directory and give it required permissions
		if err2 != nil {
		fmt.Println(err2) //print the error on the console
		return}}}
	SaveSpellToFolder()
	fmt.Println("This folder has already been made")
	}







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
	MakeNewDirectory(NewSpell.NameOfType)
	fmt.Println(Spells)
}


func  SaveSpellToFolder()  {
	f, err := os.Create(NewSpell.NameOfType +"/" + NewSpell.Name)
	if err != nil {
		fmt.Println(err)
		return
	}
	n2, err := f.Write([]byte(NewSpell.Description))
	if err != nil {
		fmt.Println(err)
        f.Close()
		return
	}
	fmt.Println(n2, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}




func (a *App) SaveNewSpell(newSpell Spell) {
	//
}