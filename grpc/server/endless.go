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
func (g *GrpcHeart) endlessTcpRegisterAndListen() error {
	flag.Parse()
	add, err := net.ResolveTCPAddr("tcp4", g.Port)
	if err != nil {
		return err
	}
	if *argReload {
		log.Println("start reload server")
		// why newfile 3
		// here -》 https://github.com/shitingbao/endless#readme in bottom
		f := os.NewFile(3, "")
		l, err := net.FileListener(f)
		if err != nil {
			return err
		}
		g.listen = l
	} else {
		l, err := net.ListenTCP("tcp", add)
		if err != nil {
			return err
		}
		g.listen = l
	}

	go g.serverLoad()
	g.signalHandler()
	return nil
}

// signalHandler
// When a signal is received, perform different actions
// When syscall.SIGUSR2 come in，start reload
func (g *GrpcHeart) signalHandler() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2)
	for {
		sig := <-ch
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			signal.Stop(ch)
			return
		case syscall.SIGUSR2:
			if err := g.reload(); err != nil {
				log.Fatalf("restart error: %v", err)
			}
			g.wg.Wait()
			return
		}
	}
}

// reload
// Save current net and standard information（in，out and err）
// then start it
// be careful not use cmd.Run(),it will block
func (g *GrpcHeart) reload() error {
	defer g.listen.Close()
	// 待定 可能需要 close 父进程的 listen 不然父进程和子进程一起接受连接
	// 但是 close 时，子进程如果还没开始监听，就会丢失连接
	f, err := g.listen.(*net.TCPListener).File()
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
