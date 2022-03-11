package main

import (
	"bufio"
	flag2 "carson.io/pkg/flag"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func initBuildCmd() *flag2.Command {
	cmd := flag2.NewCommand(flag.NewFlagSet("build", flag.ContinueOnError),
		map[string][]flag2.CmdField{
			"bool": {
				{"all", true, "create gh-pages: src -> docs"},
			},
			"string": {
				{"o", "..\\docs\\", "The output directory"},
			},
		})
	cmd.MainFunc = func(args []string) error {
		if err := cmd.Parse(args, true); err != nil {
			return err
		}
		outputDir := (cmd.Lookup("o")).Value.(flag.Getter).Get().(string)
		return build(outputDir)
	}
	return cmd
}

type msgFunc func(args []string) error

func startCMD(quitChan *chan error) {
	cmdBuild := initBuildCmd()
	var flagSetAll []*flag.FlagSet
	for _, curCmd := range []*flag2.Command{cmdBuild} {
		flagSetAll = append(flagSetAll, curCmd.FlagSet)
	}

	menuHelp := func(args []string) error {
		for _, curFlagSet := range flagSetAll {
			curFlagSet.Usage()
		}
		return nil
	}
	msgMap := map[string]msgFunc{
		"help":  menuHelp,
		"-help": menuHelp,
		"-h":    menuHelp,
		"quit": func(args []string) error {
			*quitChan <- errors.New("terminal close")
			return nil
		},
		"cls": func(args []string) error {
			var clearMap map[string]func() error
			clearMap = make(map[string]func() error)
			clearMap["linux"] = func() error {
				cmd := exec.Command("clear")
				cmd.Stdout = os.Stdout
				return cmd.Run()
			}
			clearMap["windows"] = func() error {
				cmd := exec.Command("cmd", "/c", "cls") // /c: Close
				cmd.Stdout = os.Stdout
				return cmd.Run()
			}
			clearFunc, ok := clearMap[runtime.GOOS]
			if !ok {
				return errors.New("your platform is unsupported! i can't clear terminal screen :(")

			}
			return clearFunc()
		},
		"build": cmdBuild.MainFunc,
		"run": func(args []string) error {
			go func() {
				_ = run()
			}()
			return nil
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Enter CMD: ")
		scanner.Scan()
		text := scanner.Text()
		args := strings.Split(text, " ")
		if handleFunc, exists := msgMap[args[0]]; exists {
			if err := handleFunc(args[1:]); err != nil {
				_, _ = fmt.Fprintln(os.Stderr, err.Error())
			}
		}
	}
}
