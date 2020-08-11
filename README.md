# Kevin's template for gotests

This is the template set that I use to generate tests for [gotests](https://github.com/cweill/gotests) (forked [here](https://github.com/kpaulisse/gotests)). It is based on [testify](https://github.com/stretchr/testify) and my opinions and preferences.

If you disagree with my opinions and preferences, nothing is forcing you to use this. :smile_cat:

## Features

- Creates a separate struct for the test cases and an array based on that struct. I find this to be more readable than the default templates which use an anonymous struct.

- Does not use `Args` but instead promotes each parameter as a capitalized, first-class member of the `testcase` struct.

- Has excellent support for the following types of functions/methods:

  - Return only an error -- place the expected error string in the `ExpectedError` field and use `assert.EqualError`.
  - Return only one value -- place the expected result in the `ExpectedResult` field and use `assert.Equal`.
  - Return one value and an error -- the `ExpectedError` and `ExpectedResult` fields are used, and the result is checked only if an error has not occurred.

- Generates code that compiles, albeit some of it commented-out, for cases it doesn't handle well.

## Maturity

This project is maintained only as I need capabilities for personal or professional tasks.

This is whatever comes before "alpha" (and frankly, it's not likely ever to progress beyond that). It's not being operated or supported as a project intended for others to use and contribute to at this point. That's why there's no CI and all development is via YOLO pushes to master for now.

## Installation

You need to use my fork of `gotests` because I've made some changes:

- Added [sprig](https://github.com/Masterminds/sprig) template functions
- Incorporated [PR#80 - gotests generate test names with MixedCaps](https://github.com/cweill/gotests/pull/80) (updated branch to resolve merge conflicts as of August 2020 and get tests to pass)
- Latest unreleased gotests features, specifically the `-parallel` flag

Here's how to install my fork of `gotests` (warning - may replace your stock copy, so only do this if you know what you're doing):

1. `git clone https://github.com/kpaulisse/gotests.git`
1. `cd gotests`
1. `go install ./...`

This should create `gotests` in your `go/bin` directory. Make sure that `which gotests` finds this (adjust your PATH if not).

You'll then clone this repository somewhere that you want to store templates.

1. `cd /path/to/templates`
1. `git clone git@github.com:kpaulisse/gotests-template.git kpaulisse-gotests-template`

Now, configure your editor (here's an example for vscode):

Go: Generate Tests Flags

- Add Item: `-template_dir=/path/to/templates/kpaulisse-gotests-template/templates`
- Add Item: `-parallel` (assuming you want this)

## Limitations

- Requires my fork of `gotests` (see [Installation](#installation))
- Lacks a test suite
- Doesn't create examples for non-string types (instead comments these out)
- Doesn't strip build flag comments from test files (it just copies all header comments as-is)
