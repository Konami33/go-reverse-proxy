#!/bin/bash

cd origin-server
go run main.go
cd ..

cd origin-server-2
go run main.go
cd ..

cd proxy-server
go run main.go
cd ..
