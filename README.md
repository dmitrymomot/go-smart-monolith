# go-smart-monolith

[![License](https://img.shields.io/github/license/dmitrymomot/go-smart-monolith)](https://github.com/dmitrymomot/go-smart-monolith/blob/main/LICENSE)
[![Tests](https://github.com/dmitrymomot/go-smart-monolith/actions/workflows/tests.yml/badge.svg)](https://github.com/dmitrymomot/go-smart-monolith/actions/workflows/tests.yml)
[![CodeQL Analysis](https://github.com/dmitrymomot/go-smart-monolith/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/dmitrymomot/go-smart-monolith/actions/workflows/codeql-analysis.yml)
[![GolangCI Lint](https://github.com/dmitrymomot/go-smart-monolith/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/dmitrymomot/go-smart-monolith/actions/workflows/golangci-lint.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/dmitrymomot/go-smart-monolith)](https://goreportcard.com/report/github.com/dmitrymomot/go-smart-monolith)

This is an example of a monolithic application built ready for microservice architecture. 
Each service is built as a separate package and can be easily extracted into a separate repository.

## Structure

The application is structured in a way that allows for easy transition to microservices. The application is split into the following parts:

![System Design of the App](monolith.png "System Design of the App")

In case you decide to split the application into microservices, you can easily do it:

![System Design of the App](microservice.png "System Design of the App")

## Usefull links

- [The Twelve-Factor App](https://12factor.net/)
- [The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Microservices architecture](https://microservices.io/)
- [Microservices test architecture](https://threedots.tech/post/microservices-test-architecture/)
- [Increasing Cohesion in Go with Generic Decorators](https://threedots.tech/post/increasing-cohesion-in-go-with-generic-decorators/)
- [Basic CQRS](https://threedots.tech/post/basic-cqrs-in-go/)
- [Go Project Layout](https://github.com/golang-standards/project-layout)