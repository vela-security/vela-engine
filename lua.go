package engine

import (
	"github.com/vela-security/vela-engine/template"
	"github.com/vela-security/vela-public/assert"
	"github.com/vela-security/vela-public/catch"
	"github.com/vela-security/vela-public/export"
	"github.com/vela-security/vela-public/lua"
	vswitch "github.com/vela-security/vela-switch"
	"path/filepath"
)

var xEnv assert.Environment

/*
local e = vela.engine{
	rule = "rule.d/a.yml",
	tags = {"process" , "github"},
}


local fb = vela.engine.feedback()

local vela.require("process.lua")

local e = vela.engine.load("3rd/process.zip?tags=123&tags=123")
vela.ps().pipe(s.scan)

e.case("feedback = true").pipe(fb.collect)

vela.ps().pipe(e.scan)

engine.with(fb)

*/

func NewEngineLoadL(L *lua.LState) int {
	name := L.CheckString(1)
	info, err := xEnv.Third(name)
	if err != nil {
		L.RaiseError("%s third load fail %v", name, err)
		return 0
	}

	e := &Engine{
		co:    xEnv.Clone(L),
		vsh:   vswitch.NewL(L),
		catch: catch.New(),
	}

	if info.IsZip() {
		e.rules = []string{filepath.Join(info.File(), "*.yaml")}
	} else {
		e.rules = []string{info.File()}
	}

	L.Push(e)
	return 1
}

func NewEngineL(L *lua.LState) int {
	e := NewEngine(L)
	L.Push(e)
	return 1
}

func WithEnv(env assert.Environment) {
	xEnv = env
	template.WithEnv(env)
	kv := lua.NewUserKV()
	kv.Set("feedback", lua.NewFunction(NewFeedbackL))
	kv.Set("load", lua.NewFunction(NewEngineLoadL))
	xEnv.Set("engine", export.New("vela.engine.export",
		export.WithTable(kv),
		export.WithFunc(NewEngineL)),
	)
}
