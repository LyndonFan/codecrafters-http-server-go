#!/bin/sh
localserver_file=localserver
touch $localserver_file
go build -o "$localserver_file" app/*.go
chmod +x $localserver_file