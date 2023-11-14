import { useState } from "react";
import logo from "./assets/images/logo-universal.png";
import "./App.css";
import { Greet } from "../wailsjs/go/main/App";
import { SaveNewSpell } from "../wailsjs/go/main/App";
import { MakeNewSpell } from "../wailsjs/go/main/App";

function App() {
  const [resultText, setResultText] = useState(
    "Please enter your name below 👇"
  );
  const [name, setName] = useState("");
  const updateName = (e) => setName(e.target.value);
  const updateResultText = (result) => setResultText(result);

  const [selectedFile, setSelectedFile] = useState(null);

  const [Des, SetDes] = useState("");
  const updateDes = (e) => SetDes(e.target.value);

  const [Level, SetLevel] = useState("");
  const updateLevel = (e) => SetLevel(e.target.value);

  const [SpellType, SetSpellType] = useState("");
  const updateSpellType = (e) => SetSpellType(e.target.value);

  function makeFile() {
    SaveNewSpell();
  }
  function SaveFile() {
    MakeNewSpell(name, Des, Level, SpellType);
  }

  function greet() {
    Greet(name).then(updateResultText);
  }

  return (
    <div id="App">
      <div id="result" className="result">
        {resultText}
      </div>
      <div id="input" className="input-box">
        <input
          id="name"
          className="input"
          onChange={updateName}
          autoComplete="off"
          name="input"
          type="text"
          placeholder="Name of Spell"
        />
        <input
          id="name"
          className="input"
          onChange={updateLevel}
          autoComplete="off"
          name="input"
          type="text"
          placeholder="Level"
        />
        <input
          id="name"
          className="input"
          onChange={updateSpellType}
          autoComplete="off"
          name="input"
          type="text"
          placeholder="SpellType"
        />
        <input
          id="name"
          className="input"
          onChange={updateDes}
          autoComplete="off"
          name="input"
          type="text"
          placeholder="Des"
        />

        <input
          type="file"
          value={selectedFile}
          onChange={(e) => setSelectedFile(e.target.files[0])}
        />
        <button className="btn" onClick={SaveFile}>
          Save
        </button>
      </div>
    </div>
  );
}

export default App;
