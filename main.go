package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// String constants for commands
const (
	BuildCommandString  = "build"
	LaunchCommandString = "launch"
	PullCommandString   = "pull="
)

func main() {

	// Get commands and error from parser
	commands, err := RaiseParser()
	// Handle err
	if err != nil {
		fmt.Println("Error parsing script: " + err.Error())
		return
	}

	// Send commands to router
	err = RaiseRouter(commands)
	// Handle err
	if err != nil {
		fmt.Println("Error executing commands: " + err.Error())
		return
	}

	fmt.Println("Finished successfully")
}

/**********************************************
// RaiseParser: Takes command line arguments and parses the pipeline commands
// Input: None. (Uses os.Args)
// Output: String array of commands, error thrown if invalid command line argument entered
***********************************************/
func RaiseParser() (commands []string, err error) {
	isNonAlt := false
	isAlt := false

	//Iterate through os.Args
	for argIndex, arg := range os.Args {
		// Skip the command for running the program, "./raise"
		if argIndex == 0 {
			continue
		}

		// Parse pull commands
		if len(arg) >= 7 && arg[0:6] == "-pull=" {
			commands = append(commands, PullCommandString+arg[6:len(arg)])
			isNonAlt = true
			continue
		} else if len(arg) >= 10 && arg[0:9] == "-AltPull=" {
			commands = append(commands, PullCommandString+arg[9:len(arg)])
			isAlt = true
			continue
		}

		//Parse remaining commands
		switch arg {
		case "-build":
			commands = append(commands, BuildCommandString)
			isNonAlt = true
		case "-launch":
			commands = append(commands, LaunchCommandString)
			isNonAlt = true
		case "-AltBuildAndLaunch":
			commands = append(commands, BuildCommandString)
			commands = append(commands, LaunchCommandString)
			isAlt = true
		default:
			err = errors.New("Error: Invalid Command")
			return
		}

		// Check for script uniformity
		if isNonAlt && isAlt {
			err = errors.New("Error: multiple languages detected")
		}
	}

	return
}

/**************************************************
// Raise Router: Executes pipeline functionss for a given set of commands
// Input: String array of commands
// Output: Error thrown if a command is invalid or commands are in an improper order
**************************************************/
func RaiseRouter(commands []string) (err error) {
	didBuild := false
	didLaunch := false
	improperOrderErr := errors.New("Error: Commands in an improper order")

	if len(commands) <= 0 {
		err = errors.New("Invalid number of commands parsed from script")
		return
	}

	for _, command := range commands {

		// Check pull commands
		if strings.Contains(command, PullCommandString) {
			if didBuild || didLaunch {
				err = improperOrderErr
				return
			}
			pullLocation := command[len(PullCommandString):len(command)]
			executePull(pullLocation)
			continue
		}

		// Check for remaining commands
		switch command {
		case BuildCommandString:
			if didLaunch {
				err = improperOrderErr
				return
			}
			executeBuild()
			didBuild = true
		case LaunchCommandString:
			executeLaunch()
			didLaunch = true
		default:
			err = errors.New("Invalid command found")
			return
		}
	}

	return
}

/********************************************
// executePull: Simulate a pull
// Input: String of the website location to pull from
// Output: None
********************************************/
func executePull(location string) {
	fmt.Println("Pulling from " + location + " ...")
}

/********************************************
// executeBuild: Simulate a build
// Input: None
// Output: None
********************************************/
func executeBuild() {
	fmt.Println("Building ...")
}

/********************************************
// executeLaunch: Simulate a build
// Input: None
// Output: None
********************************************/
func executeLaunch() {
	fmt.Println("Launching ...")
}
