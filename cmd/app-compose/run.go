/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"

	"github.com/edermanoel94/app-compose/internal/manager"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run multiple services",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		config := args[0]

		data, err := ioutil.ReadFile(config)

		if err != nil {
			return err
		}

		services := make([]manager.Service, 0)

		if err := json.Unmarshal(data, &services); err != nil {
			return err
		}

		chErr := make(chan error)

		go func() {
			err := <-chErr
			log.Fatal(err)
		}()

		wg := sync.WaitGroup{}

		for _, service := range services {

			wg.Add(1)
			go func(s manager.Service) {

				defer wg.Done()

				err := s.Execute()

				if err != nil {
					chErr <- err
					return
				}

			}(service)
		}

		wg.Wait()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
