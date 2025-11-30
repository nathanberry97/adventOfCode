# Advent of Code

> This is a TypeScript repository using Nx to store my
> [Advent of Code](https://adventofcode.com/) solutions.

## About

Advent of Code is an annual set of Christmas-themed programming challenges that
can be solved in any programming language.
This repository contains my solutions implemented in TypeScript, organized
using the Nx monorepo tool.

## Prerequisites

- Node.js (v18 or higher recommended)
- npm

## Getting Started

1. Clone the repository
2. Install dependencies: `make init`

## Project Structure

The repository is organized as an Nx workspace with the following structure:

- `AoC/` - Contains yearly projects with solutions for each day
- `libs/` - Shared utilities and helper functions
- `tools/` - Custom Nx generators and tooling
- Project naming convention: `aoc-<year>` (e.g., `aoc-2023`)

## Available `make` commands:

```txt
Advent of Code

Usage:
  init                  Initialise the project: install node dependencies
  test                  Test a specific project: make test PROJECT=<project-name> DAY=<day>
  build                 Build a specific project: make build PROJECT=<project-name>
  run                   Run a specific project: make run PROJECT=<project-name> DAY=<day>
  generate-new-year     Generate a new year: make generate-new-year YEAR=<year>
  list-aoc              List AoC projects in this repo
  clean                 Remove build output and Nx cache
```
