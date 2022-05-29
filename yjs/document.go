package yjs

import (
	_ "embed"
	"log"
	v8 "rogchap.com/v8go"
)

type Document struct {
	context *v8.Context
	script *v8.UnboundScript
}

//go:embed dist/bundle.js
var source string

var isolate *v8.Isolate
var compiledScript *v8.UnboundScript
var globalTemplate *v8.ObjectTemplate

func NewDocument(initialText *string, initialObjectJson *string) *Document {
	var result Document
	var err error

	if isolate == nil {
		isolate = v8.NewIsolate()

		compiledScript, err = isolate.CompileUnboundScript(source, "app.js", v8.CompileOptions{}) // compile script in new isolate with cached data
		if err != nil {
			panic(err)
		}

		globalTemplate = v8.NewObjectTemplate(isolate) // a template that represents a JS Object
	}

	result.context = v8.NewContext(isolate, globalTemplate)

	_, err = compiledScript.Run(result.context)
	if err != nil {
		panic(err)
	}

	if initialText != nil {
		err = result.set("documentText", *initialText)
	} else {
		err = result.set("documentText", v8.Undefined(isolate))
	}
	if err != nil {
		panic(err)
	}

	if initialObjectJson != nil {
		err = result.set("documentObject", *initialObjectJson)
	} else {
		err = result.set("documentObject", v8.Undefined(isolate))
	}
	if err != nil {
		panic(err)
	}

	value, err := result.context.RunScript("entry.initialize()", "app.js")
	if err != nil {
		panic(err)
	}

	log.Println("returned: ", value.String())

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

func (d *Document) ToJSON() (string, error) {
	value, err := d.context.RunScript("entry.toJSON()", "app.js")

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

func (d *Document) Close() {
	d.context.Close()
}