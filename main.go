package main

import (
	"fmt"
	"log"

	utils "cth.release/common/utils"
	web "cth.release/web"
)

func main() {
	log.Println("Kubernetes Manager")

	config := utils.GetConfig()

	fmt.Println(" _   __        _                               _             ")
	fmt.Println("| | / /       | |                             | |             ")
	fmt.Println("| |/ /  _   _ | |__    ___  _ __  _ __    ___ | |_   ___  ___ ")
	fmt.Println("|    \\ | | | || '_ \\  / _ \\| '__|| '_ \\  / _ \\| __| / _ \\/ __|")
	fmt.Println("| |\\  \\| |_| || |_) ||  __/| |   | | | ||  __/| |_ |  __/\\__ \\")
	fmt.Println("\\_| \\_/ \\__,_||_.__/  \\___||_|   |_| |_| \\___| \\__| \\___||___/	")

	fmt.Println()

	fmt.Println("Kubernetes Manager On Port: " + config.Port)

	app := web.InitServer(config)

	if app == nil {
		log.Fatalln("Init Server Error")
		return
	}

	app.App.Listen(fmt.Sprintf(":%s", config.Port))
}
