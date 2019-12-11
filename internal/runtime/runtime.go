// MIT License
//
// Copyright 2019 Canonical Ledgers, LLC
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.

package runtime

import (
	"fmt"

	"crawshaw.io/sqlite/sqlitex"
	"github.com/wasmerio/go-ext-wasm/wasmer"
)

type VM struct {
	wasmer.Instance
}

func NewVM(mod *wasmer.Module) (*VM, error) {
	imp, err := imports()
	if err != nil {
		return nil, err
	}

	inst, err := mod.InstantiateWithImports(imp)
	if err != nil {
		return nil, fmt.Errorf("wasmer.Module.InstantiateWithImports(): %w", err)
	}
	return &VM{inst}, nil
}

func (vm *VM) Call(ctx *Context,
	fname string, args ...interface{}) (v wasmer.Value, txErr, err error) {

	f, ok := vm.Exports[fname]
	if !ok {
		txErr = fmt.Errorf("unknown function")
		return
	}

	defer func(release func(*error)) {
		if err != nil {
			release(&err)
			return
		}
		release(&txErr)
	}(sqlitex.Save(ctx.Chain.Conn))

	vm.SetContextData(ctx)

	v, err = f(args...)
	if err != nil {
		if err.Error() == fmt.Sprintf(
			"Failed to call the `%s` exported function.", fname) {
			var errStr string
			errStr, err = wasmer.GetLastError()
			if err == nil {
				if errStr != ErrorExecLimitExceededString {
					err = fmt.Errorf(errStr)
					return
				}
				txErr = ErrorExecLimitExceeded{}
			}
		}
	}
	if ctx.Err != nil {
		switch ctx.Err.(type) {
		case ErrorRevert, ErrorExecLimitExceeded:
			txErr = ctx.Err
		case ErrorSelfDestruct:
			txErr = nil
		default:
			err = ctx.Err
		}
	}
	return
}
