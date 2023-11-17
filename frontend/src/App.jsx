import { useState } from "react";
import { useEffect } from "react";
import logo from "./assets/images/logo-universal.png";
import "./App.css";
import { Greet } from "../wailsjs/go/main/App";
import { MakeNewSpell } from "../wailsjs/go/main/App";
import { Temp } from "../wailsjs/go/main/App";
function Spell({ name, level, type, description }) {
  return (
    <div className="Spell">
      <h1>{name}</h1>
      <ul>
        <li>Level: {level}</li>
        <li>Type: {type}</li>
        <li>Description: {description}</li>
      </ul>
    </div>
  );
}

function App() {
  const [temp, setTemp] = useState([]);
  useEffect(() => {
    Spells();
  }, []); // Call Spells() when the component mounts
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

  function Spells() {
    // Assume Temp is a function that fetches the list of spells from your backend
    // You may need to adjust this based on how your backend is set up
    Temp().then((result) => setTemp(result));
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
        <label for="chooseOption">Element of magic spell</label>
        <select id="chooseOption" onChange={updateSpellType}>
          <option value="Fire">Fire</option>
          <option value="Water">Water</option>
          <option value="Earth">earth</option>
          <option value="Air">Air</option>
          <option value="Light">Light</option>
          <option value="Dark">Dark</option>
        </select>
        <input
          id="name"
          className="input"
          onChange={updateDes}
          autoComplete="off"
          name="input"
          type="text"
          placeholder="Des"
        />
        <button className="btn" onClick={SaveFile}>
          Save
        </button>
        <div>
          <div>
            <h1>Spells List</h1>
            {temp.map((spell) => (
              <Spell
                key={spell.name}
                name={spell.name}
                level={spell.levelOfSpell}
                type={spell.nameOfType}
                description={spell.description}
              />
            ))}
          </div>
        </div>
        //
      </div>
    </div>
  );
}

export default App;
