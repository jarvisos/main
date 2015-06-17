/*
  Copyright 2015 W. Max Lees

  This file is part of jarvisos.

  Jarvisos is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  Jarvisos is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.

  You should have received a copy of the GNU General Public License
  along with jarvisos.  If not, see <http://www.gnu.org/licenses/>.

  File: jarvis.go
  Author: W. Max Lees <max.lees@gmail.com>
  Date: 06.12.2015
*/

package main

import (
	"bufio"
	"fmt"
	"github.com/jarvisos/appserver/nlpclient"
	"os"
	"strings"
	"time"
)

// The entry point for Jarvis OS. This creates all other necessary Jarvis
// processes.
func main() {
	fmt.Printf("Jarvis OS\n\n")

	// Start processes
	/* err := processstack.StartProcesses()
	if err != nil {
		return
	}
	fmt.Printf("\n") */

	appServer, err := appserverclient.NewClient("localhost:7491", time.Second)
	if err != nil {
		fmt.Printf("Error connecting to app server: %v\n", err)
		return
	}

	// Main loop
	for in := ""; in != "exit\n"; in = getInput() {
		if strings.HasPrefix(in, "exec") {
			call := strings.TrimPrefix(in, "exec")
			_, err := appServer.DirectCall(call)
			if err != nil {
				fmt.Printf("Error making direct call: %v\n", err)
			}
		}
		//processstack.PrintStdout()
	}

	// Shutdown processes
	//processstack.WaitProcesses()
}

// Get user input
func getInput() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("[Jarvis OS]$ ")
	in, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	return in
}
