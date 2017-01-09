package main

import (
	"log"
	"os"
	"os/user"

	flags "github.com/jessevdk/go-flags"
	_ "github.com/lib/pq"
)

type Options struct {
	Host     string   `short:"s" long:"server" description:"Server name or IP address" default:"127.0.0.1"`
	Port     int      `short:"p" long:"port" description:"Port" default:"5432"`
	User     string   `short:"u" long:"user" description:"Username"`
	Database string   `short:"d" long:"database" description:"Database name" required:"true"`
	Tables   []string `short:"t" long:"tables" description:"Tables to export" required:"true"`
	SSLMode  string   `long:"ssl" description:"SSL mode (require|verify-full|verify-ca|disable)" default:"disable"`
	Output   string   `short:"o" long:"output" description:"Output filename" default:"out.go"`
}

func main() {
	var options Options

	parser := flags.NewParser(&options, flags.Default)

	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

	if options.User == "" {
		//obtain current logged user
		u, err := user.Current()
		if err != nil {
			log.Println("unable to get current os user: ", err)
		}
		options.User = u.Username
	}

	//fmt.Printf("options: %v\n", options)

	db, err := dbConnect(options)
	if err != nil {
		log.Fatalf("unable to connect to database: %s", err)
	}
	defer db.Close()

	for _, t := range options.Tables {
		cols, err := columnList(db, t)
		check(err)
		saveToFile(t+".go", []byte(GetStruct(t, cols)))
	}

}
