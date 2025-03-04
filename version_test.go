package shwild_test

import (
	"github.com/stretchr/testify/require"
	"github.com/synesissoftware/shwild.Go"

	"testing"
)

const (
	Expected_VersionMajor uint16 = 0
	Expected_VersionMinor uint16 = 2
	Expected_VersionPatch uint16 = 3
	Expected_VersionAB    uint16 = 0x4001
)

func Test_Version_Elements(t *testing.T) {
	require.Equal(t, Expected_VersionMajor, shwild.VersionMajor)
	require.Equal(t, Expected_VersionMinor, shwild.VersionMinor)
	require.Equal(t, Expected_VersionPatch, shwild.VersionPatch)
	require.Equal(t, Expected_VersionAB, shwild.VersionAB)
}

func Test_Version(t *testing.T) {
	require.Equal(t, uint64(0x0000_0002_0003_4001), shwild.Version)
}

func Test_Version_String(t *testing.T) {
	require.Equal(t, "0.2.3-alpha1", shwild.VersionString())
}
