# Teltech Backend Coding Challenge

Simple web application in Go which accepts math problems via the URL and returns the response in JSON.

## Reflection

### Project setup

Create github repo and use standard golang project layout.

### Framework vs No framework

Was considering Gin + cache lib, but decided to go with default net/http package impl and own map cache implementation.

### Implementation

Create math logic, connect it to http handlers and then test until it works. :)

## Updates

**13.10.2020**

- Package separation and simple math logic implementation.
- REST handlers implementation, first draft version.

**14.10.2020**

- REST handlers implementation second draft version.
- Table tests for REST handlers.
- JSON marshalling, refactoring


## Questions