package runner

import (
	"encoding/json"
	"fmt"
	"github.com/koding/multiconfig"
	"os"
	"strings"
)

type TaskConfig struct {
	CPU    int `default:"3"`
	Memory int `default:"256"`
	Output int `default:"16"`
	//Stack  int `default:"8"` // not use yet

	Command  string `default:"./main"`
	//Language int    `default:"2"` // not use yet

	OneTimeCalls []string `default:"execve"`
	AllowedCalls []string `default:"read,write,brk,fstat,uname,mmap,arch_prctl,exit_group,readlink,access,mprotect"`
	AdditionCalls []string `default:""`
	Verbose      bool     `default:"false"`
	Name         string // runName
	//Result       int `default:"4"` // not use yet

	//LogPath  string `default:"/var/log/runner/runner.log"`
	LogPath  string `default:"./runner.log"`
	commands []string
}

func (tc *TaskConfig) GetCommand() string {
	tc.parseCommand()
	return tc.commands[0]
}

func (tc *TaskConfig) parseCommand() {
	if len(tc.commands) == 0 {
		tc.commands = strings.Split(tc.Command, " ")
	}
}

func (tc *TaskConfig) GetArgs() []string {
	tc.parseCommand()
	return tc.commands[1:]
}

func LoadConfig(path string) *TaskConfig {
	m := multiconfig.NewWithPath(path)
	setting = new(TaskConfig)
	m.MustLoad(setting)

	return setting
}

func InitConfig() {
	init_json := multiconfig.New()
	init_data := new(TaskConfig)
	init_json.MustLoad(init_data)

	file, e := os.Create("runner_config.json")
	if e != nil {
		fmt.Printf("Cannot create file runner_config.json")
		return
	}

	encoder := json.NewEncoder(file)
	encode_err := encoder.Encode(init_data)
	if (encode_err != nil) {
		fmt.Printf("Cannot create file runner_config.json")
		file.Close()
		return
	}
	file.Close()

}
