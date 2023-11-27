import { useState } from "react";
import { useEffect } from "react";
import logo from "./assets/images/logo-universal.png";
import "./App.css";
import { Greet } from "../wailsjs/go/main/App";
import { MakeNewSpell } from "../wailsjs/go/main/App";
import { Temp } from "../wailsjs/go/main/App";

function Spell({ name, level, type, description }) {
  console.log(type);
  const getColumnTypeClass = () => {
    switch (type) {
      case "Water":
        return "water-column";
      case "Earth":
        return "earth-column";
      case "Fire":
        return "fire-column";
      case "Air":
        return "air-column";
      case "Light":
        return "light-column";
      case "Dark":
        return "dark-column";
      default:
        return "default-column";
    }
  };

  const columnClass = getColumnTypeClass();

  return (
    <div className={`grid-item ${columnClass}`}>
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
    Spells();
  }

  function Spells() {
    // Assume Temp is a function that fetches the list of spells from your backend
    // You may need to adjust this based on how your backend is set up
    Temp().then((result) => setTemp(result));
  }

  return (
    <div id="App">
      <div id="input" className="input-area">
        <input
          id="name"
          className="input-group"
          onChange={updateName}
          autoComplete="off"
          name="input"
          type="text"
          placeholder="Name of Spell"
        />
        <input
          id="name"
          className="input-group"
          onChange={updateLevel}
          autoComplete="off"
          name="input"
          type="text"
          placeholder="Level"
        />

        <label className="input-group" for="chooseOption">
          Element of magic spell
        </label>
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
          className="input-group"
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
          <div class="grid-container">
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
      </div>
    </div>
  );
}

export default App;
