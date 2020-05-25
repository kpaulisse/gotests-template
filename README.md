# "testify" template for gotests

This is the template set that I use to generate tests for [gotests](https://github.com/cweill/gotests). It is based on [testify](https://github.com/stretchr/testify) and my opinions and preferences.

If you disagree with my opinions and preferences, you can always use the default templates or find others. 😸

## Features

- Creates a separate struct for the test cases and an array based on that struct. I find this to be more readable than the default templates which use an anonymous struct.

- Has excellent support for the following types of functions:

  - Return only an error -- place the expected error string in the `ExpectedError` field.
  - Return only one value -- place the expected result in the `ExpectedResult` field.
  - Return one value and an error -- the `ExpectedError` and `ExpectedResult` fields are used, and the result is checked only if an error has not occurred.

## Limitations

- I like to name my tests as `TestFunctionName` even for a lower-cased function like `functionName`. gotests doesn't expose functionality into the templates to support a function rename, meaning that the test function will be named `Test_functionName` all of the time.
