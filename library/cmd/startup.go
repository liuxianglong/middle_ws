package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

type StartUpRunnable interface {
	Run() error
}

type StartUp struct {
	wait              sync.WaitGroup
	terminal          chan interface{}
	heartbeatInterval uint32
	loops             []LoopFunc
	ctx               context.Context
}

type LoopFunc func(<-chan interface{}, *sync.WaitGroup)

type IntervalLoopFunc func(args ...interface{})

func (startup *StartUp) Init(ctx context.Context) error {
	startup.terminal = make(chan interface{})
	startup.heartbeatInterval = 30
	startup.ctx = ctx
	return nil
}

func (startup *StartUp) IntervalLoop(interval int64, f IntervalLoopFunc, args ...interface{}) {
	startup.loops = append(startup.loops, func(close <-chan interface{}, wait *sync.WaitGroup) {
		defer wait.Done()
		ticker := time.NewTicker(time.Duration(interval) * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				f(args)
			case <-close:
				return
			}
		}
	})
}

func (startup *StartUp) Start(loops ...LoopFunc) error {
	startup.loops = append(startup.loops, loops...)
	return startup.startWithHeartbeat()
}

func (startup *StartUp) startWithHeartbeat() error {
	startup.wait.Add(len(startup.loops) + 2)
	go startup.onHeartbeat()
	go startup.onSystemKill()
	for _, runner := range startup.loops {
		go runner(startup.terminal, &startup.wait)
	}
	startup.wait.Wait()
	return nil
}

func (startup *StartUp) onSystemKill() {
	defer startup.wait.Done()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(sigs)
	sig := <-sigs
	g.Log().Infof(startup.ctx, "pid %d receive sig %d", os.Getpid(), sig)
	startup.Stop()
}

func (startup *StartUp) Stop() {
	close(startup.terminal)
}

func (startup *StartUp) onHeartbeat() {
	defer startup.wait.Done()
	hbTicker := time.NewTicker(time.Duration(startup.heartbeatInterval) * time.Second)
	for {
		select {
		case <-hbTicker.C:
			m := &runtime.MemStats{}
			runtime.ReadMemStats(m)
			g.Log().Printf(startup.ctx, "### Current memory usage: %dKb ###", m.Alloc/1024)
		case <-startup.terminal:
			g.Log().Infof(startup.ctx, "pid[%d]|terminate | received terminate signal", os.Getpid())
			return
		}
	}
}
