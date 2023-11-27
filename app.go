package main

import (
	"context"
	"fmt"
	//"io"
	"os"
	"encoding/json"
)

type Spell struct{
	Name string `json:"name"`
	Description string `json:"description"`
	NameOfType string `json:"nameOfType"`
	LevelOfSpell string`json:"levelOfSpell"`
	Stats SpellStats `json:"stats"`
}



type SpellStats struct{
	Damage int32 `json:"damage"`
	Debuff string `json:"Debuff"`
	ManaCost int32	`json:"mana_cost"`
	CastTime int32	`json:"casttime"`
	Range string	`json:"range"`
}

 var Spells []Spell
 var temp Spell
 
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
func (a *App) startup(ctx context.Context)  {
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

	for _, v := range files {
		NameOfSpellsInFolder = append(NameOfSpellsInFolder, v.Name())
		fmt.Println(v.Name())

		if v.IsDir() {
			fmt.Println("open dir")
			n, err := os.Open("Spells/" + v.Name())
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Open new text file")
			SpellFiles, err := n.Readdir(0)
			if err != nil {
				fmt.Println(err)
				return
			}

			for _, P := range SpellFiles {
				filePath := "Spells/" + v.Name() + "/" + P.Name()
				fileData, err := os.ReadFile(filePath)
				if err != nil {
					fmt.Println("Error reading file:", err)
					continue // Move on to the next file
				}

				// Create a variable to store the unmarshaled data
				var spell Spell

				// Unmarshal the JSON data into the Spell struct
				err = json.Unmarshal(fileData, &spell)
				if err != nil {
					fmt.Println("Error unmarshaling JSON:", err)
					continue // Move on to the next file
				}

				// Append the unmarshaled data to the Spells slice
				Spells = append(Spells, spell)
			}
		}
	}
	fmt.Println(Spells)
}


func (a *App)Temp()([]Spell ,error){
	return Spells,nil
}

//////
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
		fmt.Println(f)
		return
	}
	
	jsonData, err := json.MarshalIndent(NewSpell, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling struct to JSON:", err)
		return
	}

	// Write JSON data to a file
	err = os.WriteFile("Spells/" +NewSpell.NameOfType +"/" + NewSpell.Name, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON data to file:", err)
		return
	}

	fmt.Println("Struct saved to person.json")
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



