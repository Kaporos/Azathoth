# Collaborating

Hey ! Thanks for wanting to collaborate on this projet.
As this project is still a foetus, the architecture of it **will** change. 

But here is how this works for now:

`main.go` contains the starting point of the app. There is nothing really interesting inside it.

`app/app.go` is the TUI main file. It contains an Update method that is executed at each event, and a View method that generate UI.

`core` package contains all "domain" code. This is where the logic is. (map generation, player code, ...)

`stores` package contains all code to load elements from TOML files.

`components` package contains all code used to render core element to TUI


Commands user can give to the MUD are handled in `core/commands.go`
