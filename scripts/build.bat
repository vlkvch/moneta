@echo off
cd /d "%~dp0.."

go build -ldflags "-s -w" -o ".\build\moneta.exe" ".\cmd\cli"
