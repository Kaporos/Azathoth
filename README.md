# Azathoth

This project is still an in-development proof of concept.

## Usage

To run the server, simply run

    go run .

To connect, simply run

    ssh 127.0.0.1 -p 4022

## Dev

During dev, i made some hot-reload scripts. (those scripts require you to have `just` installed, as well as `inotify-tools`).
If you don't have just, you can look inside justfile to see how it works.

To build the server with hot-reload run

    just run

To connect and reconnect automatically run

    just connect


