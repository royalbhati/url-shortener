SHELL := /bin/bash


run:
	go run ./app/main.go -config $(PWD)/config.yml

seed:
	go run ./app/admin/main.go seed

migrate:
	go run ./app/admin/main.go migrate