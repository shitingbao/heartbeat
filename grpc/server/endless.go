package server

import (
	"flag"
	"log"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

var (
	// 设置一个重启的参数，用于区分正常开启还是升级
	argReload = flag.Bool("reload", false, "listen on fd open 3 (internal use only)")
)

// EndlessTcpRegisterAndListen
func (e *GrpcHeart) endlessTcpRegisterAndListen() error {
	flag.Parse()
	add, err := net.ResolveTCPAddr("tcp4", e.Port)
	if err != nil {
		return err
	}
	if *argReload {
		// why newfile 3
		// here -》 https://github.com/shitingbao/endless#readme in bottom
		f := os.NewFile(3, "")
		l, err := net.FileListener(f)
		if err != nil {
			return err
		}
		e.listen = l
	} else {
		l, err := net.ListenTCP("tcp", add)
		if err != nil {
			return err
		}
		e.listen = l
	}

	go e.serverLoad()
	e.signalHandler()
	return nil
}

// signalHandler
// When a signal is received, perform different actions
// When syscall.SIGUSR2 come in，start reload
func (e *GrpcHeart) signalHandler() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2)
	for {
		sig := <-ch
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			signal.Stop(ch)
			return
		case syscall.SIGUSR2:
			if err := e.reload(); err != nil {
				log.Fatalf("restart error: %v", err)
			}
			e.wg.Wait()
			return
		}
	}
}

// reload
// Save current net and standard information（in，out and err）
// then start it
// be careful not use cmd.Run(),it will block
func (e *GrpcHeart) reload() error {
	f, err := e.listen.(*net.TCPListener).File()
	if err != nil {
		return err
	}
	cmd := exec.Command(os.Args[0], "-reload")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.ExtraFiles = append(cmd.ExtraFiles, f)
	return cmd.Start()
}
