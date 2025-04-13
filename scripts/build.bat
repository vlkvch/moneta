@echo off

go build -ldflags "-s -w" -o ".\bin\moneta.exe" ".\cmd\cli"
