package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"github.com/shirou/gopsutil/v3/process"
)

const (
	processName = "League of Legends.exe"
)

var (
	cmds = []string{"alt", "`"}
)

func main() {
	defer hook.End()
	// 注册按键事件
	hook.Register(hook.KeyDown, cmds, func(e hook.Event) {
		// 关闭进程
		pids, err := robotgo.FindIds(processName)
		if err != nil || len(pids) == 0 {
			fmt.Printf("找不到进程 %s\n", processName)
			return
		}

		for _, pid := range pids {
			go kill(pid)
		}
	})

	// 保持进程运行，等待按键事件
	s := hook.Start()
	<-hook.Process(s)
}

func kill(pid int) {
	newProcess, err := process.NewProcess(int32(pid))
	if err != nil {
		fmt.Printf("进程 %s 不存在：%s\n", processName, err)
		return
	}
	err = newProcess.Kill()
	if err != nil {
		fmt.Printf("关闭进程 %s 失败：%s\n", processName, err)
		return
	}

	fmt.Printf("关闭进程 %s 成功\n", processName)
}
