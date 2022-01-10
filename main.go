package main

import "fmt"
"
"

func main() {
	cfg := config.localconfig
	err = config.LoadConfig(&cfg, "./config/")
	if err != nil {
		panic(fmt.Sprintf("Could not Load Config %v", err))
	}
	httpServer := api.RestAPIStartServer("localhost", "9092", cfg.ContextPath)
}
