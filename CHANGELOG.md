# Changelog

All notable changes to this project will be documented in this file.
Format follows [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

## [Unreleased]

## [0.1.0] — Stage 1: Single character strings

### Added
- `matchOne`: compares a single-character regex against a single-character input
- Wildcard `.` matches any input character
- Empty regex always returns `true`; non-empty regex against empty input returns `false`
- Reads `regex|text` from stdin, prints `true`/`false` to stdout
