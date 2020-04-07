# The Friend Management RESTFul Api

<p align="center">
  <img height="115px" src="https://www.docker.com/sites/default/files/social/docker_facebook_share.png"/>
  <img height="115px" src="https://logos-download.com/wp-content/uploads/2019/01/Golang_Logo.png"/>
  <img height="115px" src="https://www.mysql.com/common/logos/logo-mysql-170x115.png"/>
</p>

## Table of Contents
- [Introduction](#introduction)        
    - [Features](#features)
    - [Quick Installation](#quick-installation)
    - [Quick Start](#quick-start)      
- [Unit Testing](#unit-testing)
    - [Prerequisite](#prerequisite)
    - [How to run](#how-to-run)

## Introduction

A simple api built using Golang 1.4.x

### Features

- Golang 1.4.x
- Docker
- MySQL
- Unit Testing

### Quick Installation
```
# Docker Desktop for Mac
https://download.docker.com/mac/stable/Docker.dmg

# Docker Desktop for Windows
https://download.docker.com/win/stable/Docker%20Desktop%20Installer.exe

# Verify it
docker version
```

### Quick Start
```
# Build the project
docker-compose build

# Launch the project
docker-compose up

# RESTFul Api
http://localhost:8080/api

# Verify it
http://localhost:8080/ping

# Swagger
http://localhost:8080/swagger/index.html
```

## Unit Testing

A simple MSTest project to help demonstrate how to do unit using .NET Core

### Prerequisite

```
# Launch MySQL
docker-compose -f docker-compose.testing.yml up
```

### How to run

Run them directly from Visual Studio.

Or from the terminal, in the solution root, simply run:

```
# Run all tests
dotnet test

# Run all tests  with coverage
dotnet test /p:CollectCoverage=true /p:CoverletOutput=TestResults/ /p:CoverletOutputFormat=lcov

# Run method. Ex: Register_Ok
dotnet test --filter Register_Ok
```