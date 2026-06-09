# Changelog

All notable changes to this project will be documented in this file.
Format follows [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

## [Unreleased]

## [0.5.0] — Stage 5: Controlling repetition

### Added
- `?` operator: matches the preceding character zero or one time
- `*` operator: matches the preceding character zero or more times
- `+` operator: matches the preceding character one or more times
- All three operators compose with `.` wildcard and `^`/`$` anchors

## [0.4.0] — Stage 4: Implementing the operators ^ and $

### Added
- `^` anchor: regex starting with `^` matches only at the beginning of the input
- `$` anchor: regex ending with `$` matches only at the end of the input
- Combined `^...$` forces an exact full-string match

## [0.3.0] — Stage 3: Working with strings of different length

### Added
- `matchAnywhere`: sliding-window entry point that checks `matchEqual` at every position in the input string

## [0.2.0] — Stage 2: Matching two equal length strings

### Added
- `matchEqual`: recursive equal-length regex matching, reusing `matchOne` per character

## [0.1.0] — Stage 1: Single character strings

### Added
- `matchOne`: compares a single-character regex against a single-character input
- Wildcard `.` matches any input character
- Empty regex always returns `true`; non-empty regex against empty input returns `false`
- Reads `regex|text` from stdin, prints `true`/`false` to stdout
