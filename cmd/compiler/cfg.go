package main

import (
	"encoding/json"
	"fmt"
	"github.com/koding/multiconfig"
	"os"
	"strings"
)

type CompileConfig struct {
	CPU     int    `default:"3"`
	Memory  int    `default:"128"`
	Output  int    `default:"16"`
	Stack   int    `default:"8"`
	Command string `default:"g++"`
	Verbose bool   `default:"true"`
	LogPath string `default:"compiler.log"`
	Args    string `default:"main.cpp -DONLINE_JUDGE -o main -O2 -fmax-errors=10 -Wall --static -lm --std=c++11"`
}

func (config *CompileConfig) GetArgs() []string {
	return strings.Split(config.Args, " ")
}

func loadConfig(path string) *CompileConfig {
	m := multiconfig.NewWithPath(path)
	compileConfig := new(CompileConfig)
	m.MustLoad(compileConfig)
	return compileConfig
}

func initConfig() {
	init_json := multiconfig.New()
	init_data := new(CompileConfig)
	init_json.MustLoad(init_data)

	file, e := os.Create("compiler_config.json")
	if e != nil {
		fmt.Printf("Cannot create file compiler_config.json")
		return
	}

	encoder := json.NewEncoder(file)
	encode_err := encoder.Encode(init_data)
	if (encode_err != nil) {
		fmt.Printf("Cannot create file compiler_config.json")
		file.Close()
		return
	}
	file.Close()
}
