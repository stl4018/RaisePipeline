package main

import (
	"os"
	"testing"
)

func TestParser(t *testing.T) {

	// Test no commands
	os.Args = []string{"./raise"}
	parseNoCommands, err := RaiseParser()
	if err != nil {
		t.Errorf("Error was not nil: %v", err)
	}
	if len(parseNoCommands) != 0 {
		t.Errorf("Expected no commands, got %v instead", parseNoCommands)
	}

	// Test no-alt commands
	os.Args = []string{"./raise_project", "-pull=somewhere.com", "-build", "-launch"}
	parseNoAltCommands, err := RaiseParser()
	if err != nil {
		t.Errorf("Error was not nil: %v", err)
	}
	if len(parseNoAltCommands) != 3 || parseNoAltCommands[0] != "pull=somewhere.com" || parseNoAltCommands[1] != "build" || parseNoAltCommands[2] != "launch" {
		t.Errorf("Expected ordered no-alt commands, got %v instead", parseNoAltCommands)
	}

	// Test out of order no-alt commands
	os.Args = []string{"./raise", "-pull=somewhere.com", "-launch", "-build"}
	parseNoAltCommands, err = RaiseParser()
	if err != nil {
		t.Errorf("Error was not nil: %v", err)
	}
	if len(parseNoAltCommands) != 3 || parseNoAltCommands[0] != "pull=somewhere.com" || parseNoAltCommands[1] != "launch" || parseNoAltCommands[2] != "build" {
		t.Errorf("Expected un-ordered no-alt commands, got %v instead", parseNoAltCommands)
	}

	// Test alt commands
	os.Args = []string{"./raise", "-AltPull=somewhere.com", "-AltBuildAndLaunch"}
	parseAltCommands, err := RaiseParser()
	if err != nil {
		t.Errorf("Error was not nil: %v", err)
	}
	if len(parseAltCommands) != 3 || parseAltCommands[0] != "pull=somewhere.com" || parseAltCommands[1] != "build" || parseAltCommands[2] != "launch" {
		t.Errorf("Expected ordered alt commands, got %v instead", parseAltCommands)
	}

	// Test multi-language commands
	os.Args = []string{"./raise", "-AltPull=somewhere.com", "-build"}
	_, err = RaiseParser()
	if err == nil {
		t.Errorf("Expected error but got nil")
	}

	// Test illegible commands
	os.Args = []string{"./raise", "-AltPull=somewhere.com", "-oogabooga"}
	_, err = RaiseParser()
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
}

func TestRouter(t *testing.T) {
	// Test no commands
	commands := []string{}
	err := RaiseRouter(commands)
	if err == nil {
		t.Errorf("Expected error but got nil")
	}

	// Test ordered commands
	commands = []string{"-build", "-launch"}
	err = RaiseRouter(commands)
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	// Test unordered commands
	commands = []string{"-launch", "-build"}
	err = RaiseRouter(commands)
	if err == nil {
		t.Errorf("Expected error but got nil")
	}

}
