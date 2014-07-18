package main
import (
	"log"
	"flag"
	"io/ioutil"
	"launchpad.net/goyaml"
	"fmt"
	//"time"
	//"os"
	//db "zltg/db"
)

const (
	FATAL = iota
	ERROR
	WARN
	INFO
	DEBUG
)

var log_level int
var logger log.Logger

var config map[string]string

func set_log_level(ll string) {
	switch ll {
	case "DEBUG":
		log_level = DEBUG
	case "INFO":
		log_level = INFO
	case "WARN":
		log_level = WARN
	case "ERROR":
		log_level = ERROR
	case "FATAL":
		log_level = FATAL
	}
}

func zlog(ll int, message string) {
	names := [5]string{"FATAL","ERROR","WARN","INFO","DEBUG"}
	if ll <= log_level {
		log.Printf("%5s %s\n", names[ll], message)
	}
}

func initLogging(){
	
}

func main() {
	/*****************
   * set up logging
   *****************/
	log_level = INFO
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	log.Println("testing the logger")

	/************************
	 * read in configuration
   ***********************/
	var config_file string
	flag.StringVar(&config_file, "c", "./zltg.yaml", "Config file in YAML format")
	flag.Parse()
	log.Println("Config file=", config_file)
	config_string, _ := ioutil.ReadFile(config_file)
	//log.Printf("config contents=%s\n", config_string)
	err := goyaml.Unmarshal(config_string, &config)
	log.Println(err)
	for key, value := range config {
		log.Println(key, ": ", value)
	}
	lls, ok := config["log_level"]
	if ok {
		set_log_level(lls)
		log.Printf("Log Level is %s = %d \n", config["log_level"], log_level)
		zlog(DEBUG, fmt.Sprintf("Log Level is %s = %d \n", config["log_level"], log_level))
	}
	InitDb(config["db"])
}



