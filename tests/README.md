# Tests

As the main package cannot be directly tested with `*_test.go` files, this directory contains some testing tools.

## `stresstest.go`
You can run it by using the command `go run stresstest.go`. Some options can be passed to this command:
* `-host` (string, default `http://localhost`): host to test;
* `-port` (int, default `8081`): host port;
* `-count` (int, default `100`): requests count for each URL;
* `-verbose` (bool, default `false`): verbose mode.

Example: `go run stresstest.go -host=localhost -port=8081 -count=1000 -verbose=1`.