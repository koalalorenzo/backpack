package main

import (
	"fmt"

	"gitlab.com/qm64/backpack/deployer"
)

func main() {
	err := deployer.Deploy(`
job "example" {
	datacenters = ["dc1"]
	type = "service"
	
	group "cache" {
		count = 1
		task "redis" {
		driver = "docker"
		config {
			image = "redis:3.2"
			port_map {
			db = 6379
			}
		}
		}
	}
	}	  
`)
	if err != nil {
		fmt.Printf("e %s", err)
		panic(err)
	}

	fmt.Printf("OK")
}
