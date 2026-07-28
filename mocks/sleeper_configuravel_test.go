package main

import (
	"testing"
	"time"
)

func TestSleeperConfiguravel(t *testing.T) {
	tempoPausa := 5 * time.Second

	tempoSpy := &TempoSpy{}

	sleeper := SleeperConfiguravel{tempoPausa, tempoSpy.Pausa}
	sleeper.Sleep()

	if tempoSpy.duracaoPausa != tempoPausa {
		t.Errorf("deveria ter pausado por %v, mas pausou por %v", tempoPausa, tempoSpy.duracaoPausa)
	}
}
