// Copyright (c) 2026 Arslaan Pathan
// This software is licensed under the ARPL. See LICENSE for details.

package build

import (
	"github.com/yuin/gopher-lua"
)

type Kiirofile struct {
	Path string
}

func (k *Kiirofile) Load(L *lua.LState) error {
	if err := L.DoFile(k.Path); err != nil {
		// just return it, they can do the err detection
		return err
	}
	return nil
}

func (k *Kiirofile) RunBuild(L *lua.LState) error {
	// just return it because if its an error notmyfault they detect it
	// notmyfault you forgot to Kiirofile.Load()
	// or that the Kiirofile maintainer typed it wrong
	return L.CallByParam(lua.P{
		Fn: L.GetGlobal("build"),
		NRet: 0,
		Protect: true,
	})
}

func (k *Kiirofile) RunPackage(L *lua.LState) error {
	// just return it because if its an error notmyfault they detect it
	// notmyfault you forgot to Kiirofile.Load()
	// or that the Kiirofile maintainer typed it wrong
	return L.CallByParam(lua.P{
		Fn: L.GetGlobal("package"),
		NRet: 0,
		Protect: true,
	})
}

// TODO: Add auto-detection and methods for MakeBuild, CargoBuild, CMakeBuild, etc for the projects that don't have a Kiirofile
// That's probably for later because right now we need basic package manager with Kiirofile parsing working
