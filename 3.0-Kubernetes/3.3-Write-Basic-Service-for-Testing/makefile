SHELL := /bin/bash

run:
	go run main.go

# We set the build variable here to "local".
# When we do "make build", it builds an executable which leaves a binary file behind.
# 3.3-Write-Basic-Service-for-Testing.exe and when we run it, 
# ./3.3-Write-Basic-Service-for-Testing.exe
# 2023/01/03 23:06:55 Starting Service local
build:
	go build -ldflags "-X main.build=local" 