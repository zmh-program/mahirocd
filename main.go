package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"mahirocd/utils"
	"mahirocd/workflow"
)

func main() {
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}

	manager := workflow.NewManager()

	conn := utils.NewConnection(
		fmt.Sprintf("%s/connection", viper.GetString("endpoint")),
		manager.HandleMessage,
	)
	conn.ExecWithBlock()
}
