package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/jquery"
	"github.com/yuin/gopher-lua"
)

var jq = jquery.NewJQuery

func PlayLua(s string) {
	L := lua.NewState()
	defer L.Close()
	if err := L.DoString(s); err != nil {
		panic(err)
	}
}

func main() {
	editor := js.Global.Get("ace").Call("edit", "editor")
	session := editor.Call("getSession")
	session.Call("setMode", "ace/mode/lua")
	cons := jq("#cons")

	jq("#run").On("click", func() {
		PlayLua(editor.Call("getValue").String())
		cons.SetScrollTop(int(cons.Prop("scrollHeight").(float64)))
	})

	jq("#clean").On("click", func() {
		cons.SetHtml("")
	})
}
