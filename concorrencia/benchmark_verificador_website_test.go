package concorrencia

import (
	"testing"
	"time"
)

func slowStubVerificadorWebsite(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkVerificadorWebsites(b *testing.B) {
	urls := make([]string, 100)

	for i := range urls {
		urls[i] = "uma url"
	}

	for b.Loop() {
		VerificaWebSites(slowStubVerificadorWebsite, urls)
	}
}
