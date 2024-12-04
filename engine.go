package main

import "errors"

type Engine struct {
	data map[string]string
}

func NewEngien() *Engine {
	return &Engine{
		data: make(map[string]string),
	}
}

func (e *Engine) Set(key, value string) error {
	e.data[key] = value
	return nil
}

func (e *Engine) Get(key string) (string, error) {
	val, ok := e.data[key]
	if !ok {
		return "", errors.New("Key not found!")
	}
	return val, nil
}

