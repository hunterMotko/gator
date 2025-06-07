package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/huntermotko/gator/internal/config"
	"github.com/huntermotko/gator/internal/database"
	_ "github.com/lib/pq"
)

type State struct {
	db  *database.Queries
	cfg *config.Config
}

func ExecuteCli() {
	if len(os.Args) > 1 && os.Args[1] == "init" {
		if err := handlerInit(); err != nil {
			log.Fatalf("INIT ERR: %v\n", err)
		}
		log.Println("success config: $HOME/.gatorconfig.json")
		os.Exit(0)
	}

	conf, err := config.Read()
	if err != nil {
		log.Println("Did you gator init?")
		log.Fatalf("cannot read config error: %v\n", err)
	}

	db, err := sql.Open("postgres", conf.DB_Url)
	if err != nil {
		log.Fatalf("cannot read config error: %v\n", err)
	}

	dbQueries := database.New(db)
	state := State{
		db:  dbQueries,
		cfg: &conf,
	}
	cmds := Commands{
		Cmds: make(map[string]func(*State, Command) error),
	}

	cmds.Register("login", HandlerLogin)
	cmds.Register("register", HandlerRegister)
	cmds.Register("users", handlerCheckUsers)
	cmds.Register("agg", handlerAgg)
	cmds.Register("feeds", handlerFeeds)
	cmds.Register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.Register("follow", middlewareLoggedIn(handlerFollow))
	cmds.Register("following", middlewareLoggedIn(handlerFollowing))
	cmds.Register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cmds.Register("browse", middlewareLoggedIn(handlerBrowse))
	cmds.Register("reset", handlerReset)

	if len(os.Args) < 2 {
		fmt.Printf("GATOR CLI\n\n")
		fmt.Printf("Available Commands: \n")
		for k := range cmds.Cmds {
			fmt.Printf("- %s\n", k)
		}
		os.Exit(0)
	}

	command := Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	if err := cmds.Run(&state, command); err != nil {
		log.Fatalf("run error: %v", err)
	}
}
