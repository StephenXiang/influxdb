package profile

import (
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

func New(cpuprofile, memprofile string) func() {
	// prof stores the file locations of active profiles.
	var prof struct {
		cpu *os.File
		mem *os.File
	}

	if cpuprofile != "" {
		f, err := os.Create(cpuprofile)
		if err != nil {
			log.Fatalf("cpuprofile: %v", err)
		}
		prof.cpu = f
		_ = pprof.StartCPUProfile(prof.cpu)
	}

	if memprofile != "" {
		f, err := os.Create(memprofile)
		if err != nil {
			log.Fatalf("memprofile: %v", err)
		}
		prof.mem = f
		runtime.MemProfileRate = 4096
	}

	return func() {
		if prof.cpu != nil {
			pprof.StopCPUProfile()
			_ = prof.cpu.Close()
		}
		if prof.mem != nil {
			_ = pprof.Lookup("heap").WriteTo(prof.mem, 0)
			_ = prof.mem.Close()
		}
	}
}
