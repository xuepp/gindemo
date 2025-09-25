package main

import "log"

func main() {
	engine, cfg, err := InitApp()
	if err != nil {
		log.Fatalf("failed to init app: %v", err)
	}

	log.Printf("server running at %s", cfg.Addr)
	if err := engine.Run(cfg.Addr); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
