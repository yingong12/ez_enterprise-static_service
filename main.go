package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"static_service/http"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {
	//加载环境变量
	envPath := ".env"
	flag.StringVar(&envPath, "c", ".env", "配置文件")
	flag.Parse()
	if err := godotenv.Load(envPath); err != nil {
		panic(err)
	}
	log.Println("env loadded from file ", envPath)
	err, _ := http.Start()
	if err != nil {
		panic(err)
	}
	log.Println("Httpserver started ")

	//wait for sys signals
	exitChan := make(chan os.Signal)
	signal.Notify(exitChan, os.Interrupt, os.Kill, syscall.SIGTERM)
	select {
	case sig := <-exitChan:
		log.Println("Doing cleaning works before shutdown...")
		log.Println("You abandoned me, bye bye", sig)
	}
}
