# Regex Engine (Go)

A regex engine built from scratch in Go as part of the [Hyperskill project #254](https://hyperskill.org/projects/254). The standard `regexp` package is intentionally not used — the engine is implemented step by step across six stages, supporting literals, the `.` wildcard, `^`/`$` anchors, repetition operators (`?`, `*`, `+`), and escape sequences.

## Usage

Input is read from stdin as `regex|text`. The engine prints `true` or `false`.

```bash
echo "col.ur|colour" | go run "Regex Engine (Go)/task/main.go"
# true
```

## Running tests

```bash
cd "Regex Engine (Go)/task" && python tests.py
```

Requires the `hstest` Python package:

```bash
pip install -r requirements.txt
```

## Stages

| Stage | Feature | Status |
|-------|---------|--------|
| 1 | Single character matching, `.` wildcard, empty-string rules | ✅ |
| 2 | Equal-length string matching (recursive) | ✅ |
| 3 | Different-length strings (sliding window) | ✅ |
| 4 | `^` and `$` anchors | ✅ |
| 5 | Repetition operators: `?`, `*`, `+` | ✅ |
| 6 | Escape sequences (`\`) | ✅ |
