# heartbeat

heart is jumping

# Example progress

$ go build //修改代码后，重新构建出 stbserver
$ ps -ef|grep stbserver // 找到正在运行的 stbserver 对应进程号，比如是进程号是： 31088
$ kill -SIGUSR2 31088 // 向 31088 进程发送一个 SIGUSR2 信号，如果有连接，就会显示如下：
501 31021 27220 0 3:09 下午 ttys001 0:00.09 ./stbserver
501 31088 31021 0 3:10 下午 ttys001 0:00.05 ./stbserver -reload

第一个是第一次执行的进程，第二个带有 -reload 参数是升级后的进程，这时候，新的连接就会被新进程接受，
旧进程将不会接受新连接，不过会继续为还没有断开的连接提供服务，直到所有旧连接都断开，然后结束旧连接。如下，只剩一个：

501 31088 1 0 3:10 下午 ?? 0:00.96 ./stbserver -reload
