package linux

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArch(t *testing.T) {
	for from, to := range map[string]string{
		"linuxamd64":             "amd64",
		"linux386":               "i386",
		"linuxarm64":             "arm64",
		"linuxarm5":              "armel",
		"linuxarm6":              "armhf",
		"linuxarm7":              "armv7l",
		"linuxppc64":             "ppc64",
		"linuxppc64le":           "ppc64le",
		"linuxwhat":              "what",
		"linuxmips64lesoftfloat": "mips64el",
		"linuxmipslehardfloat":   "mipsel",
		"linuxmipssoftfloat":     "mips",
		"linuxmips64hardfloat":   "mips64",
	} {
		t.Run(fmt.Sprintf("%s to %s", from, to), func(t *testing.T) {
			require.Equal(t, to, Arch(from))
		})
	}
}
