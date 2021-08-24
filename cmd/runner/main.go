package main

import (
	"encoding/json"
	"fmt"
	"github.com/hustoj/runner/runner"
	"os"
)

func main() {

	args := os.Args
	if (len(args) < 2) {
		fmt.Printf("Usage:\n\trunner <config.json>\n\trunner init\n")
		return
	}

	if (args[1] == "init") {
		runner.InitConfig()
		return
	}

	setting := runner.LoadConfig(args[1])
	runner.InitLogger(setting.LogPath, setting.Verbose)

	task := runner.RunningTask{}
	task.Init(setting)
	task.Run()

	result := task.GetResult()
	content, _ := json.Marshal(result)
	fmt.Println(string(content))
}
