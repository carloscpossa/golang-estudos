package main

import "time"

type SleeperConfiguravel struct {
	duracao time.Duration
	pausa   func(time.Duration)
}

func (c *SleeperConfiguravel) Sleep() {
	c.pausa(c.duracao)
}
