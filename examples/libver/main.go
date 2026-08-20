package main

import (
	shwild "github.com/synesissoftware/shwild.Go"
	ver2go "github.com/synesissoftware/ver2go"

	"fmt"
)

func main() {
	fmt.Printf("shwild v%s\n", shwild.VersionString())
	fmt.Printf("ver2go v%s\n", ver2go.VersionString())
}
