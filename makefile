SHELL := /bin/bash


run:
	go run ./app/main.go -config $(PWD)/config.yml