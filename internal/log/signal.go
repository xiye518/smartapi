// +build !windows

package log

import (
	"os"
	"os/signal"
	"syscall"
)

/*HandleSignalChangeLogLevel 实现了根据信号量修改日志等级的方法
使用方法 go HandleSignalChangeLogLevel()
监听3个信号:
	kill -10 pid 修改日志等级为 all
	kill -12 pid 修改日志等级为 error
	kill -14 pid 修改日志等级为 info
	kill -12 $(cat /tmp/xalarmgo.pid)
*/
func HandleSignalChangeLogLevel() {
	//mission 改为同一个输入，滚动修改日志等级，计数（信号输入次数），一个升序，一个降序
	chSignalUser1 := make(chan os.Signal)
	chSignalUser2 := make(chan os.Signal)
	chSignalUser3 := make(chan os.Signal)
	signal.Notify(chSignalUser1, syscall.SIGUSR1) //kill -10 pid 修改日志等级为 all
	signal.Notify(chSignalUser2, syscall.SIGUSR2) //kill -12 pid 修改日志等级为 error
	signal.Notify(chSignalUser3, syscall.SIGALRM) //kill -14 pid 修改日志等级为 info
	for {
		select {
		case <-chSignalUser1:
			SetLogLevelAll()
		case <-chSignalUser2:
			SetLogLevelError()
		case <-chSignalUser3:
			SetLogLevelInfo()
		}
	}
}
