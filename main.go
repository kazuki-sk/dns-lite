package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	arg := flag.String("configdir", "./config.toml", "Specify the location of the configuration file")
	flag.Parse()

	conf, err := ReadConf(*arg)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	log.Println("Init SQL")
	db, err := OpenDB(conf.DB.Type, conf.DB.DNS)

	log.Println("Create Table")
	if err := db.AutoMigrate(&Zone{}, &Record{}, &SOA{}); err != nil {
		log.Fatal(err)
	}
}
