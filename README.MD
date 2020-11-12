# RaisePipeline
A sample pipeline/automation controller system intended to simulate how a real pipeline/automation controller could work

## Specs:

To run program:
```bash
go build
./raise_project
```

Command Line Arguments (simulating a script):
```bash
-pull=somewhere.com     //Specify a place to pull from using -pull
-build                  //build program
-launch                 //launch program
-AltPull=somewhere.com  //Same as pull
-AltBuildAndLaunch      //build and launch program
```

Note: Alt and non-alt command line arguments are intended to simulate two different languages--so you may only use commands from the same language. Also, commands should be in order of pull -> build -> launch, but not all are necessary

## Issues with code:
- It allows for multiple, identical commands to be executed in a row
- The script language parser require hardcoding of foreign languages

## Thought process:
Step 1: See if I understand the assignment correctly
- We want a pipeline than can execute scripts written in other languages
- It interprets the scripts from other languages and executes tasks based on that
- To summarize: I should build a program that interprets scripts from other languages (outside of Golang) and would, if completed, execute pipeline-based tasks based on the foreign script
- Contacting Katherine

Step 2: Plan out this program
- Should have a program that A) Parses input (the script) and B) Calls other functions depending on the results
- The parsed script could be transmitted as an array of commands
- The main function will call two functions, one to parse the input and one to choose which function to call, both of which should have tests
- This is only a sample program, so script does not have to be real and pipeline functions don't have to do anything real 
- Because the script could be in multiple languages, multiple inputs could result in the same pipeline call

Step 3: Write the program
- First write the tests, then write the three important functions (main, parse, pipelineRouter) as well as some sample output functions
- Messaging Katherine for further clarification
- See code

- Step 4: Results
- Initial version done; revising
- Tests not passed
