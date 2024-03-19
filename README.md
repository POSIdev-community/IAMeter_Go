# IAMeter_Go
The repository contains a collection of tests to evaluate the capability of various source code analyzers to search for vulnerabilities. Some examples have deliberately false vulnerabilities that may prompt false triggers from analyzers. Other examples are hard to detect for various reasons.

In real applications, both cases may lead to unclear scan results, so you will have to handpick the ones that matter.

The test collection of examples is a set of Go files.

The examples are in the files whose name reflects the aim of the code tested. Some files have several examples.

The code examples contain comments about whether vulnerabilities are expected to be found. The explanation indicates whether the code is false positive or false negative.

To use examples, you need to run `go run main.go`
