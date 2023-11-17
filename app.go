package main

import (
	"context"
	"fmt"
	"io"
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
 var NameOfSpellsInFolder []string 
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

		f, err := os.Open("Spells")
		if err != nil {
			fmt.Println(err)
			return
		}
		files, err := f.Readdir(0)
		if err != nil {
			fmt.Println(err)
			return
		}
	//
		for _, v := range files {
			NameOfSpellsInFolder = append(NameOfSpellsInFolder,v.Name())
			fmt.Println(v.Name())
			if v.IsDir(){
				fmt.Println("open dir")
				n, err := os.Open("Spells/"+v.Name())
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("Open new text file")
				SpellFiles, err := n.Readdir(0)
				if err != nil {
				fmt.Println(err)
				return}
				for _, P := range SpellFiles {
					file, err := os.Open("Spells/"+v.Name()+"/"+P.Name())
					if err != nil {
						fmt.Println(err)
					}
					b, err := io.ReadAll(file)
					fmt.Print(string(b))
					
				 }

	////
				fmt.Println(v.Name() +" is a dir")
			}
		}
		fmt.Println(NameOfSpellsInFolder)
	
}
////
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
	f, err := os.Create("Spells/" +NewSpell.NameOfType +"/" + NewSpell.Name)
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

func MakeNewDirectory(NewSpellType string){
	_, err := os.Open(NewSpellType)
	if err != nil {
	if os.IsNotExist(err) {
		directoryname := "Spells/"+NewSpellType
		err2 := os.Mkdir(directoryname, 0755) //create a directory and give it required permissions
		if err2 != nil {
		SaveSpellToFolder()
		fmt.Println(err2)
		 //print the error on the console
		return}}}
	SaveSpellToFolder()
	fmt.Println("This folder has already been made")
	}
//



func (a *App) SaveNewSpell(newSpell Spell) {
	//
}