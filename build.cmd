@echo off
set goos=windows
go build -o build/api.exe cmd/app/main.go
set goos=linux
go build -o build/api cmd/app/main.go