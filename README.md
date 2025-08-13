# Pokémon CLI (pokecli)

A simple and interactive **Command Line Interface** tool written in Go that fetches Pokémon information from the [PokeAPI](https://pokeapi.co/).  
Built using the **Cobra** library for structured CLI commands.

---

## ✨ Features

- Fetch Pokémon details by name (type, abilities, height, weight, stats, etc.).
- Search Pokémon by type (e.g., fire, water, electric).
- Well-structured CLI with subcommands using Cobra.
- Clean, maintainable project structure with separate files for commands.

---

## 📦 Installation

### Prerequisites
- [Go](https://go.dev/doc/install) installed (version 1.18+ recommended)
- Internet connection (to fetch Pokémon data)

### Clone the repository
```bash
git clone https://github.com/yourusername/pokecli.git
cd pokecli
Install dependencies
bash
Copy
Edit
go mod tidy
🚀 Usage
Build the CLI
bash
Copy
Edit
go build -o pokecli
Run the CLI
bash
Copy
Edit
./pokecli
Available Commands
Get Pokémon by Name
bash
Copy
Edit
./pokecli get pikachu
Example output:

vbnet
Copy
Edit
Name: pikachu
Height: 4
Weight: 60
Types: electric
Abilities: static, lightning-rod
List Pokémon by Type
bash
Copy
Edit
./pokecli type fire
Example output:

python-repl
Copy
Edit
Charmander
Vulpix
Growlithe
...
📂 Project Structure
go
Copy
Edit
pokecli/
│── cmd/
│   ├── root.go       # Base CLI command
│   ├── get.go        # 'get' subcommand for Pokémon details
│   ├── type.go       # 'type' subcommand for listing Pokémon by type
│── main.go           # Entry point
│── go.mod
│── README.md
🛠 Built With
Go — Programming language

Cobra — CLI framework

PokeAPI — Pokémon REST API

📜 License
This project is licensed under the MIT License — see the LICENSE file for details.

