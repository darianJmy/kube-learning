/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"github.com/spf13/cobra"
	"k8s.io/component-base/logs"
	"kube-learning/practise/cobra-practise/alice/app"
	"os"
)

func main() {

	command := app.NewAliceCommand()

	code := run(command)

	os.Exit(code)
}

func run(command *cobra.Command) int {
	defer logs.FlushLogs()
	//rand.Seed(time.Now().UnixNano())

	if err := command.Execute(); err != nil {
		return 1
	}
	return 0
}
