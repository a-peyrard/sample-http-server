set shell := ["bash", "-uc"]

# Defaults ========================================================================================
project := 'sample-http-server'

# Commands ========================================================================================
# show this help
help:
	@just --list

# run the server
server:
  @go run ./main.go

inject:
  @k6 run --vus 10 --duration 30s k6/test.js