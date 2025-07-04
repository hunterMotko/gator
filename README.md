# Gator CLI

A command-line interface (CLI) program built with Go, leveraging PostgreSQL for
data storage, `sqlc` for type-safe SQL, and `goose` for database migrations.

## Prerequisites

Before you can run the `gator` CLI, you'll need to have the following installed
on your system:

* **PostgreSQL**: This program uses a PostgreSQL database to store its data.
    You can download and install PostgreSQL from the official website:
    [https://www.postgresql.org/download/](https://www.postgresql.org/download/)
* **Go**: The program is written in Go. You can download and install Go from
    the official website: [https://go.dev/doc/install](https://go.dev/doc/install)

## Installation

To install the `gator` CLI, open your terminal and run the following command:

```bash
go install [github.com/hunterMotko/gator@latest](https://github.com/hunterMotko/gator@latest) 
```

Run `gator init` to initalize you .gatorconfig.json file

## Usage

Commands: `gator [commmand] [params ...]`

* reset
* register: Register a user
* login: Login as a user
* users: List users
* addfeed: Add a RSS feed
* feeds: List all feeds for that user
* browse: Browse users posts 
* agg: Srape current feeds
* follow: Follow a feed
* following: List all feeds the user is following
* unfollow: Unfollow a feed

## TODO

* Add sorting and filtering options to the browse command
* Add pagination to the browse command
* Add concurrency to the agg command so that it can fetch more frequently
* Add a search command that allows for fuzzy searching of posts
* Add bookmarking or liking posts
* Add a TUI that allows you to select a post in the terminal and view it in a more readable format (either in the terminal or open in a browser)
* Add an HTTP API (and authentication/authorization) that allows other users to interact with the service remotely
* Write a service manager that keeps the agg command running in the background and restarts it if it crashes
