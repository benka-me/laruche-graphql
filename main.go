package main

 import (
 	"github.com/benka-me/laruche/go-pkg/discover"
 	"github.com/urfave/cli"
 	"github.com/benka-me/laruche-graphql/go-pkg/http/rpc"
	"google.golang.org/grpc"
 	"os"
 	"fmt"
 	"log"
 )

type App struct {
    rpc.Clients // This struct contains connection tou your services
}

func Run(engine discover.Engine) error {
	app := &App{
		Clients : rpc.InitClients(engine, grpc.WithInsecure()), // Init clients of dependencies services, please change options.
	}
	fmt.Println(app)
	return nil
}

 func main() {
 	app := cli.NewApp()
 	app.Commands = cli.Commands{
 		{
 			Name: "dev",
 			Action: func(context *cli.Context) error {
 				if len(os.Args) < 3 {
 			        fmt.Println("usage: exec author/service-name")
 			        os.Exit(0)
 				}
 				engine, err := discover.ParseEngine(os.Args[2], true)
 		        if err != nil {
 			        log.Fatal(err)
 		        }
 				return Run(*engine)
 			},
 		},
 	}

 	app.Action = func(c *cli.Context) error {
 		if len(os.Args) < 2 {
 			fmt.Println("usage: exec author/service-name")
 			os.Exit(0)
 		}
 		engine, err := discover.ParseEngine(os.Args[1], false)
 		if err != nil {
 			log.Fatal(err)
 		}
 	    return Run(*engine)
 	}

 	err := app.Run(os.Args)
 	if err != nil {
 		log.Fatal(err)
 	}
 }


