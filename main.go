/*
 * Copyright 2016 ThoughtWorks, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"github.com/gocd-contrib/gocd-golang-agent/agent"
	"time"
	"flag"
	"fmt"
	"os"
)

var (
	Version = "0.0"
	Githash = "No Version Provided"
)

func main() {

	versonPtr := flag.Bool("version", false, "Show GoCD Golang Agent Verson")
	flag.Parse()

	if *versonPtr {
		fmt.Println("GoCD Golang Agent Verson : ", Version, " (", Githash, ")")
		os.Exit(0)
	}

	agent.Initialize()
	for {
		err := agent.Start()
		if err != nil {
			agent.LogInfo("something wrong: %v", err.Error())
		}
		agent.LogInfo("sleep 10 seconds and restart")
		time.Sleep(10 * time.Second)
	}
}
