package main

import (
	"log"

	"github.com/SARVESHVARADKAR123/task-manager-cli/internal/config"
	"github.com/SARVESHVARADKAR123/task-manager-cli/internal/app/service/task"

	"github.com/SARVESHVARADKAR123/task-manager-cli/internal/infra/clock"
	"github.com/SARVESHVARADKAR123/task-manager-cli/internal/infra/repo/json"
	"github.com/SARVESHVARADKAR123/task-manager-cli/internal/ui/cli"


)


func main(){
	//1. load config
	cfg:= config.Load()
	
	//2. Infrastructure
	repo:= json.NewTaskRepo(cfg.DataPath)
	clk:=clock.NewSystemClock()

	//3. Application
	taskService:=task.New(repo,clk)

	// 4. CLI
	root := cli.NewRoot(taskService)

	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
	
}