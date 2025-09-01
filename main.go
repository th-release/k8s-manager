package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	utils "cth.release/common/utils"
	web "cth.release/web"
)

func main() {
	port := flag.String("PORT", "8080", "PORT Integer Default: 8080")
	kubeConfig := flag.String("KUBE", "~/.kube/config", "kubernetes config path Default: ~/.kube/config")

	flag.Parse()

	Port, err := strconv.Atoi(*port)

	if err != nil {
		fmt.Println("Error: --PORT is Not Integer")
		flag.Usage()
		os.Exit(1)
	}

	if Port > 65535 || Port < 1 {
		fmt.Println("Error: --PORT is Invalid ex(Port > 65535 or Port < 1)")
		flag.Usage()
		os.Exit(1)
	}

	err = utils.SetConfig(&utils.Config{
		Port:       Port,
		KubeConfig: *kubeConfig,
	})

	if err != nil {
		log.Fatalf("Error saving config : %s", err.Error())
	}

	log.Println("Kubernetes Manager")

	fmt.Println(" _   __        _                               _             ")
	fmt.Println("| | / /       | |                             | |             ")
	fmt.Println("| |/ /  _   _ | |__    ___  _ __  _ __    ___ | |_   ___  ___ ")
	fmt.Println("|    \\ | | | || '_ \\  / _ \\| '__|| '_ \\  / _ \\| __| / _ \\/ __|")
	fmt.Println("| |\\  \\| |_| || |_) ||  __/| |   | | | ||  __/| |_ |  __/\\__ \\")
	fmt.Println("\\_| \\_/ \\__,_||_.__/  \\___||_|   |_| |_| \\___| \\__| \\___||___/	")

	fmt.Println()

	fmt.Println("Kubernetes Manager On Port: " + *port)

	config := utils.GetConfig()

	app := web.InitServer(config)

	if app == nil {
		log.Fatalln("Init Server Error")
		return
	}

	app.App.Listen(fmt.Sprintf(":%d", config.Port))
}
