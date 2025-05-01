@echo off
cd /d "%~dp0.."

go build -ldflags "-s -w" -o ".\bin\moneta.exe" ".\cmd\cli"
