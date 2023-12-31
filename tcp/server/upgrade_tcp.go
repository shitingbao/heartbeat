// 这里对于长连接
// 总思路，开启新进程，继承老进程的tcp服务
// 老进程等待所有连接关闭后退出
// 新的进程监听新的连接，老进程由于被继承不会继续监听，相当于把端口让出给新进程
package server

import (
	"flag"
	"log"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/shitingbao/heartbeat/core"
)

var (
	// 设置一个重启的参数，用于区分正常开启还是升级
	argReload      = flag.Bool("reload", false, "listen on fd open 3 (internal use only)")
	defaultAddress = ":8080"
)

type EndlessTcp struct {
	address    string
	listen     net.Listener
	wg         *sync.WaitGroup // 该标识标记了父进程的退出逻辑，在进程 listen 的时候 add，并且在信号接收的地方 wait ，在连接全部断开的时候 done，这样连接全部断开的时候，就自动退出了父进程
	readLength int
	conflags   sync.Map
	UserHub    core.Hub

	Duration time.Duration
	callBack func()
}

// default adress is ":8080"
func New(ads ...string) *EndlessTcp {
	e := &EndlessTcp{
		address:    defaultAddress,
		wg:         &sync.WaitGroup{},
		readLength: 256,
		conflags:   sync.Map{},
	}
	if len(ads) > 0 {
		e.address = ads[0]
	}
	return e
}

// SetReadLength 设置每次读取的长度
func (e *EndlessTcp) SetReadLength(readLength int) {
	e.readLength = readLength
}

// EndlessTcpListen监听入口
func (e *EndlessTcp) EndlessTcpRegisterAndListen(u UpgradeRead) error {
	flag.Parse()
	add, err := net.ResolveTCPAddr("tcp4", e.address)
	if err != nil {
		return err
	}
	if *argReload {
		// 获取到cmd中的ExtraFiles内的文件信息，以它为内容启动监听
		// ExtraFiles的内容在reload方法中放入
		log.Println("EndlessTcpRegisterAndListen reload:", *argReload)
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
	go e.listenAccept(u)
	e.signalHandler()
	return nil
}

// 注意不能使用代理的情况连接，可能会出现RemoteAddr相同的情况，导致con连接对象覆盖
func (e *EndlessTcp) listenAccept(u UpgradeRead) {
	log.Println("start listen ", e.address)
	for {
		con, err := e.listen.Accept()
		if err != nil {
			log.Println("Accept:", err)
			return
		}
		e.conflags.Store(con.RemoteAddr().String(), con)
		e.wg.Add(1)
		e.handle(con, u)
	}
}

// read write 方法待定
func (e *EndlessTcp) handle(con net.Conn, u UpgradeRead) {
	go e.read(con, u)
	go e.write(con)
}

func (e *EndlessTcp) read(con net.Conn, u UpgradeRead) {
	for {
		result := make([]byte, e.readLength)
		n, err := con.Read(result)
		if err != nil {
			e.wg.Done()
			adr := con.RemoteAddr().String()
			e.conflags.Delete(adr)

			log.Println("断开 address:", adr)
			return
		}
		u.ReadMessage(&ReadMes{
			N:   n,
			Mes: result,
		})
		e.callBack()
	}
}

func (e *EndlessTcp) write(con net.Conn) error {
	t := time.NewTicker(e.Duration)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			e.conflags.Range(func(_ any, v any) bool {
				cn, ok := v.(net.Conn)
				if ok {
					cn.Write([]byte("1"))
				}
				return true
			})
		}
	}
}

// 信号处理
func (e *EndlessTcp) signalHandler() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2)
	for {
		sig := <-ch
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			signal.Stop(ch)
			log.Printf("stop listen")
			return
		case syscall.SIGUSR2:
			if err := e.reload(); err != nil {
				log.Fatalf("restart error: %v", err)
			}
			// go e.reload()
			log.Println("start wait")
			e.wg.Wait()
			log.Println("stop wait")
			return
		}
	}
}

// 重启方法，这里放入进程中的输入，输出和错误
// 以及最重要的listen对象（放入ExtraFiles中），以文件句柄的形式继承
// 相当于有了所有父进程的属性，然后重新执行该可执行文件
func (e *EndlessTcp) reload() error {
	f, err := e.listen.(*net.TCPListener).File()
	if err != nil {
		log.Println("reload", err)
		return err
	}
	cmd := exec.Command(os.Args[0], "-reload")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.ExtraFiles = append(cmd.ExtraFiles, f)
	return cmd.Start() // 注意这里要用 Start，不能 Run，不然就在新进程中卡住，因为执行可执行程序 Run 中会 wait 等待
}
