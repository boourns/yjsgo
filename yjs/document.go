package yjs

import (
	_ "embed"
	v8 "rogchap.com/v8go"
)

type Document struct {
	isolate *v8.Isolate
	context *v8.Context
	script *v8.UnboundScript
}

//go:embed dist/bundle.js
var source string

func NewDocument(initialState string) *Document {
	var result Document
	var err error

	result.isolate = v8.NewIsolate()

	result.script, err = result.isolate.CompileUnboundScript(source, "app.js", v8.CompileOptions{}) // compile script in new isolate with cached data
	if err != nil {
		panic(err)
	}

	glob := v8.NewObjectTemplate(result.isolate) // a template that represents a JS Object

	result.context = v8.NewContext(result.isolate, glob)

	_, err = result.script.Run(result.context)
	if err != nil {
		panic(err)
	}

	err = result.set("documentText", initialState)
	if err != nil {
		panic(err)
	}

	value, err := result.context.RunScript("entry.initialize()", "app.js")
	if err != nil {
		panic(err)
	}

	if value.String() != "initialized" {
		panic("failed to initialize yjs Document")
	}

	return &result
}

func (d *Document) ToString() (string, error) {
	value, err := d.context.RunScript("entry.toString()", "app.js")

	if err != nil {
		return "", err
	}

	return value.String(), nil
}

func (d *Document) ApplyUpdate(encodedUpdate string) error {
	err := d.set("encodedUpdate", encodedUpdate)
	if err != nil {
		return err
	}

	_, err = d.context.RunScript("entry.applyUpdate()", "app.js")

	return err
}

func (d *Document) EncodeStateAsUpdate(targetStateVector string) (string, error) {
	err := d.set("encodedStateVector", targetStateVector)
	if err != nil {
		return "", err
	}

	result, err := d.context.RunScript("entry.encodeStateAsUpdate()", "app.js")
	if err != nil {
		return "", err
	}
	return result.String(), nil
}

func (d *Document) StateVector() (string, error) {
	result, err := d.context.RunScript("entry.stateVector()", "app.js")
	if err != nil {
		return "", err
	}
	return result.String(), nil
}

func (d *Document) Insert(position uint, content string) error {
	err := d.set("insertPosition", position)
	if err != nil {
		return err
	}
	err = d.set("insertText", content)
	if err != nil {
		return err
	}
	_, err = d.context.RunScript("entry.insert()", "app.js")
	return err
}

func (d *Document) set(name string, value interface{}) error {
	global := d.context.Global()

	return global.Set(name, value)
}
