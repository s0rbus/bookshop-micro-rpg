package main

import (
	"fmt"
	"os"

	"github.com/dop251/goja"
	"github.com/dop251/goja/parser"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
	"github.com/s0rbus/bookshop-micro-rpg/api"
)

var vm *goja.Runtime

func GetVM() *goja.Runtime {
	return vm
}

func LoadExpansion(dir string, name string) (api.ExpansionStruct, error) {
	var expansion api.ExpansionStruct
	vm = goja.New()
	new(require.Registry).Enable(vm)
	console.Enable(vm)
	vm.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))
	file, err := os.Open(fmt.Sprintf("%s/%s", dir, name))
	if err != nil {
		return expansion, fmt.Errorf("error opening the expansion file %v ", err)
	}
	p, err := parser.ParseFile(nil, name, file, 0)
	if err != nil {
		return expansion, fmt.Errorf("error parsing file %v ", err)
	}
	prog, err := goja.CompileAST(p, true)
	if err != nil {
		return expansion, fmt.Errorf("error compiling the script %v ", err)
	}
	_, err = vm.RunProgram(prog)
	if err != nil {
		return expansion, fmt.Errorf("error running the script %v ", err)
	}

	var Name func() string
	err = vm.ExportTo(vm.Get("getName"), &Name)
	if err != nil {
		return expansion, fmt.Errorf("error getting function %v", err)
	}
	expansion.Name = Name
	var GetRequiredThrows func() int
	err = vm.ExportTo(vm.Get("getRequiredThrows"), &GetRequiredThrows)
	if err != nil {
		return expansion, fmt.Errorf("error getting function %v", err)
	}
	expansion.GetRequiredThrows = GetRequiredThrows
	var Run func(int, []int) ([]string, error)
	err = vm.ExportTo(vm.Get("run"), &Run)
	if err != nil {
		return expansion, fmt.Errorf("error getting function %v", err)
	}
	expansion.Run = Run
	var SetVerbose func(v bool)
	err = vm.ExportTo(vm.Get("setVerbose"), &SetVerbose)
	if err != nil {
		return expansion, fmt.Errorf("error getting function %v", err)
	}
	expansion.SetVerbose = SetVerbose
	return expansion, nil

}
