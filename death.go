package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type ffunc func(*chan struct{})

// Death manages the death of your application.
type Death struct {
	wg       *sync.WaitGroup
	sigChan  chan os.Signal
	fTimeout time.Duration
	fFunc    ffunc
	done     chan struct{}
}

func NewDeath(f ffunc) (death *Death) {
	death = &Death{
		sigChan:  make(chan os.Signal, 1),
		wg:       &sync.WaitGroup{},
		fTimeout: 3 * time.Second,
		fFunc:    f,
		done:     make(chan struct{}),
	}

	signal.Notify(death.sigChan, syscall.SIGINT, syscall.SIGTERM)
	death.wg.Add(1)
	go death.waitForSignal()

	return death
}

func (d *Death) waitForSignal() {
	defer d.wg.Done()
	for {
		select {
		case <-d.sigChan:
			if d.fFunc != nil {
				go d.fFunc(&d.done)
			}
			return
		}
	}

}

func (d *Death) SetFinalize(f ffunc) {
	d.fFunc = f
}

func (d *Death) Wait() {
	d.wg.Wait()

	i := 1
	timer := time.NewTicker(d.fTimeout)
	for {
		select {
		case <-timer.C:
			fmt.Printf("Finalize ing : %v\n", time.Duration(i)*d.fTimeout)
			i++
		case <-d.done:
			fmt.Println("Finalize Finished")
			return
		}
	}
}
