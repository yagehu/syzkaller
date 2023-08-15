package main

import "github.com/google/syzkaller/pkg/compiler"

type wasi struct {}

func (*wasi) prepare(sourceDir string, build bool, arches []*Arch) error {
	return nil
}

func (*wasi) prepareArch(*Arch) error {
	return nil
}

func (*wasi) processFile(
	arch *Arch,
	info *compiler.ConstInfo,
) (map[string]uint64, map[string]bool, error) {
	var (
		res        = make(map[string]uint64)
		undeclared = make(map[string]bool)
	)

	return res, undeclared, nil
}
