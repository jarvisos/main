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

   File: processstack.go
   Author: W. Max Lees <max.lees@gmail.com>
   Date: 06.14.2015
*/

package processstack

import (
	"fmt"
	"os/exec"
)

func StartProcesses() error {
	return processStack.startProcesses()
}

func WaitProcesses() {
	processStack.waitProcesses()
}

var processStack processes

type processes struct {
	nlp       *exec.Cmd
	appserver *exec.Cmd
}

// Generate the processes
func (proc *processes) startProcesses() error {
	fmt.Printf("Starting NLP...\n")
	proc.nlp = exec.Command("nlp")
	err := proc.nlp.Start()
	if err != nil {
		fmt.Printf("Couldn't initialize NLP process: %v\n", err)
		return err
	}

	fmt.Printf("Starting App Server...\n")
	proc.appserver = exec.Command("appserver")
	err = proc.appserver.Start()
	if err != nil {
		fmt.Printf("Couldn't initialize App Server process: %v\n", err)
		return err
	}

	return nil
}

// Wait for processes to terminate
func (proc *processes) waitProcesses() {
	proc.appserver.Wait()
	proc.nlp.Wait()
}
