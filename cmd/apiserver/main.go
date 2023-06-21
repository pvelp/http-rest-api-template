package main

import (
	"flag"
	"log"

	"github.com/pvelp/http-rest-api-template/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	// _, err := toml.Decode(configPath, config)
	// fmt.Println(config)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
