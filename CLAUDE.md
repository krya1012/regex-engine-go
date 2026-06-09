# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a [Hyperskill](https://hyperskill.org/projects/254) course project: building a regex engine from scratch in Go, without using the standard `regexp` package.

## Commands

**Run the program:**
```bash
cd "Regex Engine (Go)/task" && go run main.go
```

**Run tests (Python-based Hyperskill test runner):**
```bash
cd "Regex Engine (Go)/task" && python tests.py
```
Tests require the `hstest` Python package (installed via `pip install -r requirements.txt`).

**Build:**
```bash
cd "Regex Engine (Go)/task" && go build main.go
```

## Architecture

All implementation lives in `Regex Engine (Go)/task/main.go`. The engine reads input as `regex|text` from stdin and prints `true` or `false`.

The project is structured as incremental stages, each adding new capabilities:

1. **Single character strings** — match one char regex vs one char input; `.` is a wildcard; empty regex → `true`; empty input (non-empty regex) → `false`
2. **Equal length strings** — recursive character-by-character matching using stage 1 as a helper
3. **Different length strings** — slide the regex across all positions in the input string
4. **`^` and `$` anchors** — `^` forces match at start only; `$` forces match at end only (inserted as termination condition after empty-regex check but before empty-string check)
5. **Repetition operators `?`, `*`, `+`** — added as conditions inside the equal-length recursive function; each operator looks at the preceding character and the next part of the regex
6. **Escaping** — `\` before any metacharacter treats it as a literal; `\\` matches a literal backslash

The recommended design uses three layers:
- A single-character matcher (handles `.` wildcard and `\` escape prefix)
- A recursive equal-length matcher (handles `?`/`*`/`+` and `^`/`$`)
- A sliding-window entry point (handles different-length matching and `^` anchoring)
