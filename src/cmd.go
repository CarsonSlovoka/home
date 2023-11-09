package main

import (
	"bufio"
	. "carson.io/pkg/utils"
	"context"
	"flag"
	"fmt"
	. "github.com/CarsonSlovoka/go-pkg/v2/op"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
)

type ParaBuild struct {
	outputDir string
	isForce   bool
}

func initFlagSetBuild(p *ParaBuild) *flag.FlagSet {
	f := flag.NewFlagSet("build", flag.ContinueOnError)
	f.BoolVar(&p.isForce, "f", false, "Forced overwrite of output folders")
	f.StringVar(&p.outputDir, "o", "..\\docs\\", "The output directory")
	return f
}

type ParaRun struct {
	isLocalMode bool
}

func initFlagRun(p *ParaRun) *flag.FlagSet {
	f := flag.NewFlagSet("run", flag.ContinueOnError)
	f.BoolVar(&p.isLocalMode, "f", true, "run on 127.0.0.1 by default")
	return f
}

func shutdownServer(server *http.Server) error {
	if server == nil {
		PInfo.Println("The server has shutdown already.")
		return nil
	}
	return server.Shutdown(context.Background())
}

func startCMD(wg *sync.WaitGroup) {
	paraBuild := new(ParaBuild)
	paraRun := new(ParaRun)
	flagSetBuild := initFlagSetBuild(paraBuild)
	flagSetRun := initFlagRun(paraRun)
	flagSetShutdownServer := flag.NewFlagSet("shutdownServer", flag.ContinueOnError)

	var (
		server   *http.Server
		listener net.Listener
	)
	handleMsg := func(args []string) (string, error) {
		var err error
		switch args[0] {
		case "h":
			fallthrough
		case "help":
			flagSetBuild.Usage()
			flagSetRun.Usage()
			flagSetShutdownServer.Usage()
		case "quit":
			if server != nil {
				err = shutdownServer(server)
			}
			return "quit", err
		case "build":
			{
				if err = flagSetBuild.Parse(args[1:]); err != nil {
					// 會自動觸發預設的錯誤
					return "", PErr.Errorf("parse build error")
				}
				p := paraBuild
				if !p.isForce {
					if _, err := os.Stat(p.outputDir); !os.IsNotExist(err) {
						return "", PErr.Errorf("the directory (%s) exists already", p.outputDir)
					}
				}
				POk.Println("Start build...")
				err = build(p.outputDir)
				POk.Println("End build")
			}

		case "cls":
			{
				var cmd *exec.Cmd
				switch runtime.GOOS {
				case "linux":
					cmd = exec.Command("clear")
					cmd.Stdout = os.Stdout
				case "windows":
					cmd = exec.Command("cmd", "/c", "cls") // /c: Close
					cmd.Stdout = os.Stdout
				default:
					return "", PErr.Errorf("your platform is unsupported! i can't clear terminal screen :(")
				}
				return "", cmd.Run()
			}
		case "shutdownServer":
			return "", shutdownServer(server)
		case "run":
			{
				if server != nil {
					return "", PErr.Errorf("run already.")
				}

				if err = flagSetRun.Parse(args[1:]); err != nil {
					return "", PErr.Errorf("parse run error")
				}
				p := paraRun
				server, listener = BuildServer(p.isLocalMode)

				go func() {
					if err2 := server.Serve(listener); err2 != nil {
						server = nil // 表示關閉，重新指派到nil去，方便再啟動
						log.Println(POk.Sprintf("close the server.\n"))
					}
				}()

				port := fmt.Sprintf("%d", listener.Addr().(*net.TCPAddr).Port)
				url := If(p.isLocalMode,
					"localhost:"+port+"/",
					listener.Addr().String(), // 127.0.0.1
				)
				if runtime.GOOS == "windows" {
					if err := exec.Command("rundll32", "url.dll,FileProtocolHandler", "http://"+url).Start(); err != nil {
						panic(err)
					}
				}
				POk.Printf("http://%s\n", url)
			}
		default:
			return "", fmt.Errorf("unknown command %q", strings.Join(args, " "))
		}
		return "", err
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Enter CMD (enter 'quit' to leave): ")
		scanner.Scan()
		input := scanner.Text()
		args := strings.Split(input, " ")
		response, err := handleMsg(args)
		if err != nil {
			_, _ = PErr.Fprintln(os.Stderr, err)
			continue
		}
		if response == "quit" {
			break
		}
	}
	wg.Done()
}
