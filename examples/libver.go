package main

import (
	shwild "github.com/synesissoftware/shwild.Go"
	"github.com/synesissoftware/ver2go"

	"fmt"
)

func main() {
	fmt.Printf("shwild v%s\n", ver2go.CalcVersionString(shwild.VersionMajor, shwild.VersionMinor, shwild.VersionPatch, shwild.VersionAB))
	fmt.Printf("ver2go v%s\n", ver2go.CalcVersionString(ver2go.VersionMajor, ver2go.VersionMinor, ver2go.VersionPatch, ver2go.VersionAB))
}
