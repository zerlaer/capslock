package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/go-vgo/robotgo"
)

// App struct
type App struct {
	ctx      context.Context
	isRunning bool
	ticker   *time.Ticker
	interval int // 时间间隔，单位：秒
	
	// 定时关机相关
	isShuttingDown bool
	shutdownTimer  *time.Timer
	shutdownTime   int // 关机时间，单位：秒
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		interval: 60, // 默认60秒
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// StartPreventLock starts the lock prevention
func (a *App) StartPreventLock() string {
	if a.isRunning {
		return "已经在运行中"
	}

	// 启动定时任务
	a.ticker = time.NewTicker(time.Duration(a.interval) * time.Second)
	a.isRunning = true

	go func() {
		for range a.ticker.C {
			// 模拟按下和释放Cap Lock键
			robotgo.KeyTap("capslock")
			log.Printf("已按下Caps Lock键 (间隔: %d秒)", a.interval)
			// 再次按下恢复原来的状态
			time.Sleep(100 * time.Millisecond)
			robotgo.KeyTap("capslock")
			log.Printf("已释放Caps Lock键 (间隔: %d秒)", a.interval)
		}
	}()

	return "防锁屏功能已启动"
}

// StopPreventLock stops the lock prevention
func (a *App) StopPreventLock() string {
	if !a.isRunning {
		return "已经停止"
	}

	a.ticker.Stop()
	a.isRunning = false

	return "防锁屏功能已停止"
}

// GetStatus returns the current status
func (a *App) GetStatus() map[string]interface{} {
	return map[string]interface{}{
		"running":       a.isRunning,
		"interval":      a.interval,
		"isShuttingDown": a.isShuttingDown,
		"shutdownTime":   a.shutdownTime,
	}
}

// SetInterval sets the interval for lock prevention
func (a *App) SetInterval(seconds int) string {
	if seconds <= 0 {
		return "时间间隔必须大于0"
	}

	// 如果正在运行，先停止再重启
	if a.isRunning {
		a.StopPreventLock()
		a.interval = seconds
		a.StartPreventLock()
		return fmt.Sprintf("时间间隔已更新为 %d 秒", seconds)
	}

	a.interval = seconds
	return fmt.Sprintf("时间间隔已设置为 %d 秒", seconds)
}

// SetShutdownTime sets the shutdown time
func (a *App) SetShutdownTime(seconds int) string {
	if seconds <= 0 {
		return "关机时间必须大于0"
	}

	a.shutdownTime = seconds
	return fmt.Sprintf("关机时间已设置为 %d 秒", seconds)
}

// StartShutdownTimer starts the shutdown timer
func (a *App) StartShutdownTimer() string {
	if a.isShuttingDown {
		return "定时关机已经在运行中"
	}

	if a.shutdownTime <= 0 {
		return "请先设置关机时间"
	}

	// 创建关机定时器
	a.shutdownTimer = time.NewTimer(time.Duration(a.shutdownTime) * time.Second)
	a.isShuttingDown = true

	go func() {
		<-a.shutdownTimer.C
		log.Printf("定时关机时间到，执行关机操作")
		
		// 执行关机命令（Windows）
		cmd := exec.Command("shutdown", "/s", "/t", "0")
		if err := cmd.Run(); err != nil {
			log.Printf("执行关机命令失败: %v", err)
		}
		log.Printf("定时关机时间到，执行关机操作")
		
		a.isShuttingDown = false
	}()

	return fmt.Sprintf("定时关机已启动，将在 %d 秒后关机", a.shutdownTime)
}

// CancelShutdownTimer cancels the shutdown timer
func (a *App) CancelShutdownTimer() string {
	if !a.isShuttingDown {
		return "没有正在执行的定时关机"
	}

	a.shutdownTimer.Stop()
	a.isShuttingDown = false

	return "定时关机已取消"
}
