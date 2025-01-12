#!/bin/bash
cd ./mpod
go build -o /usr/bin/mpod main.go

if alias mpod &>/dev/null; then
	exit 0
else
	alias mpod="/usr/bin/mpod"
fi
