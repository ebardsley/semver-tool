package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/coreos/go-semver/semver"
)

var (
	bumpMajor = flag.Bool("major", false, "Bump the major version.")
	bumpMinor = flag.Bool("minor", false, "Bump the major version.")
	bumpPatch = flag.Bool("patch", false, "Bump the major version.")
	errUsage  = fmt.Errorf("exactly one of -major, -minor, or -patch can be specified")
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s <-major|-minor|-patch> <version>\n\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}
}

func bumpSemver(v string, major bool, minor bool, patch bool) (string, error) {
	if b2i(major)+b2i(minor)+b2i(patch) != 1 {
		return "", errUsage
	}
	s, err := semver.NewVersion(v)
	if err != nil {
		return "", err
	}

	if major {
		s.BumpMajor()
	} else if minor {
		s.BumpMinor()
	} else if patch {
		s.BumpPatch()
	}
	return s.String(), nil
}

func b2i(b bool) int8 {
	if b {
		return 1
	}
	return 0
}

func main() {
	flag.Parse()
	v, err := bumpSemver(flag.Arg(0), *bumpMajor, *bumpMinor, *bumpPatch)
	if err != nil {

		if err == errUsage {
			flag.Usage()

		}
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	fmt.Println(v)
	os.Exit(0)
}
