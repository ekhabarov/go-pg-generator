package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	flags "github.com/jessevdk/go-flags"
	_ "github.com/lib/pq"
)

type Options struct {
	Host         string   `short:"s" long:"server" description:"Server name or IP address" default:"127.0.0.1"`
	Port         int      `short:"p" long:"port" description:"Port" default:"5432"`
	User         string   `short:"u" long:"user" description:"Database user."`
	Password     string   `short:"w" long:"password" description:"Database password."`
	Database     string   `short:"d" long:"database" description:"Database name." required:"true"`
	Tables       []string `short:"t" long:"tables" description:"Tables to export." required:"true"`
	SSLMode      string   `long:"ssl" description:"SSL mode [require|verify-full|verify-ca|disable]" default:"disable"`
	FilePerTable bool     `short:"f" long:"file-per-table" description:"Save each structure to its own .go file."`
	PackageName  string   `long:"package" description:"Package name for generated files."`
}

func main() {
	var options Options

	parser := flags.NewParser(&options, flags.Default)

	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		}
		os.Exit(1)
	}

	if options.User == "" {
		//obtain current logged user
		u, err := user.Current()
		if err != nil {
			log.Fatalln("unable to get current os user: ", err)
		}
		options.User = u.Username
	}

	db, err := dbConnect(options)
	if err != nil {
		log.Fatalf("unable to connect to database: %s", err)
	}
	defer db.Close()

	p := "models"
	if options.PackageName != "" {
		p = options.PackageName
	}
	d := fmt.Sprintf("package %s \n", p)

	if !options.FilePerTable {
		fmt.Println(d)
	}

	for _, t := range options.Tables {
		cols, err := columnList(db, t)
		if err != nil {
			log.Fatalln(err)
		}

		data := getStruct(t, cols)

		if options.FilePerTable {
			if err := saveToFile(t, []byte(d+data)); err != nil {
				log.Fatalln(err)
			}
			continue
		}
		fmt.Println(data)
	}
}
