 first choose a module path (we'll use example/user/hello) and create a go.mod file that declares it:
$ go mod init github.com/odedviner/cobra
go: creating new go.mod: module example/user/hello

The first statement in a Go source file must be package name. Executable commands must always use package main.

Next, create a file named hello.go inside that directory containing the following Go code:


$ go install github.com/spf13/cobra-cli@latest
$ ~/go/bin/cobra-cli init
$ pwd
cobra_test/oded_cobra/cmd
$ go build -o oded



This will compile your Go program and save the executable as myprogram in the current directory. 
go build -o myprogram


```
$ ./main 
Error: requires at least 1 arg(s), only received 0
Usage:
  cobra [flags]
  cobra [command]

Available Commands:
  completion     Generate the autocompletion script for the specified shell
  help           Help about any command
  mons           Output mon endpoints
  restore-quorum When quorum is lost, restore quorum to the remaining healthy mon

Flags:
  -h, --help   help for cobra

Use "cobra [command] --help" for more information about a command.
```