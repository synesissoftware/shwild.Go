# shwild.Go <!-- omit in toc -->

[![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)
[![GitHub release](https://img.shields.io/github/v/release/synesissoftware/shwild.Go.svg)](https://github.com/synesissoftware/shwild.Go/releases/latest)
[![Last Commit](https://img.shields.io/github/last-commit/synesissoftware/shwild.Go)](https://github.com/synesissoftware/shwild.Go/commits/master)
[![Go](https://github.com/synesissoftware/shwild.Go/actions/workflows/go.yml/badge.svg)](https://github.com/synesissoftware/shwild.Go/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/synesissoftware/shwild.Go)](https://goreportcard.com/report/github.com/synesissoftware/shwild.Go)
[![Go Reference](https://pkg.go.dev/badge/github.com/synesissoftware/shwild.Go.svg)](https://pkg.go.dev/github.com/synesissoftware/shwild.Go)

**S**hell-C**o**mpatible W**ILDc**ards for Go


## Introduction

**shwild** is a small, simple library that provides shell-compatible wildcard matching. It implemented in several languages: **shwild.Go** is the **Go** implementation.


## Table of Contents <!-- omit in toc -->

- [Introduction](#introduction)
- [Installation](#installation)
- [Components](#components)
	- [Standalone match function](#standalone-match-function)
	- [Compiled pattern](#compiled-pattern)
- [Examples](#examples)
- [Project Information](#project-information)
	- [Where to get help](#where-to-get-help)
	- [Contribution guidelines](#contribution-guidelines)
	- [Dependencies](#dependencies)
		- [Development/Example/Testing Dependencies](#developmentexampletesting-dependencies)
	- [Related projects](#related-projects)
	- [License](#license)

## Installation

```Go

import shwild "github.com/synesissoftware/shwild.Go"
```

## Components

Two means of pattern matching are provided:
* standalone match function; and
* compiled pattern;


### Standalone match function

```Go
func Match(pattern string, s string, args ...any) (bool, error)
```

`shwild.Match` evaluates string `s` against `pattern`, subject to additional arguments that moderate behaviour, and returns a `bool` that indicates match if the function succeeds; if if fails the `error` contains information about why.


### Compiled pattern

```Go
func Compile(pattern string, args ...any) (CompiledPattern, error)

func (cp CompiledPattern) Match(s string) (bool, error)
```

`shwild.Compile` compiles `pattern` into a `CompiledPattern` instance, which may then be used to evaluate string `s` against `pattern`, subject to additional arguments that moderate behaviour, and returns a `bool` that indicates match if the function succeeds; if if fails the `error` contains information about why.


## Examples

Examples are provided in the ```examples``` directory, along with a markdown description for each. A detailed list TOC of them is provided in [EXAMPLES.md](./EXAMPLES.md).


## Project Information


### Where to get help

[GitHub Page](https://github.com/synesissoftware/shwild.Go "GitHub Page")

### Contribution guidelines

Defect reports, feature requests, and pull requests are welcome on https://github.com/synesissoftware/shwild.Go.


### Dependencies

* [**ver2go**](https://github.com/synesissoftware/ver2go/);


#### Development/Example/Testing Dependencies

* [**CLASP.Go**](https://github.com/synesissoftware/CLASP.Go/);
* [**STEGoL**](https://github.com/synesissoftware/STEGoL/);
* [**testify**](https://github.com/stretchr/testify);



### Related projects

* [**shwild**](https://github.com/synesissoftware/shwild/);
* [**shwild.Rust**](https://github.com/synesissoftware/shwild.Rust/);


### License

**shwild.Go** is released under the 3-clause BSD license. See [LICENSE](./LICENSE) for details.
