package template

import (
	"github.com/vela-security/vela-public/assert"
	"github.com/vela-security/vela-public/lua"
)

var xEnv assert.Environment

func NewL(L *lua.LState) *Template {
	return &Template{
		co: xEnv.Clone(L),
	}
}

func WithEnv(env assert.Environment) {
	xEnv = env
}
