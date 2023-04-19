# Boilerplate REST service using Gin web framework (Golang)

[![Go](https://github.com/lakhinsu/gin-boilerplate/actions/workflows/go.yml/badge.svg)](https://github.com/lakhinsu/gin-boilerplate/actions/workflows/go.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/lakhinsu/gin-boilerplate)](https://goreportcard.com/report/github.com/lakhinsu/gin-boilerplate)

## Introduction

This repository contains a boilerplate REST API service that connects to mifos api using [Gin](https://gin-gonic.com/) web framework. 

This boilerplate uses most common and highly adpated libraries like [zerolog](https://github.com/rs/zerolog) for logging and [viper](https://github.com/spf13/viper) for env configuration.

This boilerplate was sourced from  
[blog](https://medium.com/@hinsulak/rest-api-service-boilerplate-using-gin-web-framework-golang-c4eeb48b14f) 
which explains the process behind creation.

## Features
- HTTP and HTTPS support
- .env file and OS environment variables support
- Models
- Controllers
- Routers
- Request ID middleware
- Request logging middleware
- CORS middleware


## Run locally
Follow these steps to run this service locally
- Get dependencies
`go get .`
- Set environment variables - use the sample .env file provided in the repository

Execute `go run main.go` command at the repository root to start the service.
## Running in Docker
This repo includes the sample Dockerfile and docker-compose.yaml files that can be used as a reference.
