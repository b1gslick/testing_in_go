# what covers

- Testing is a critical part of the development
  process
- well tested code is more maintable, more secure,
  and overall takes less time to write
- go has a rich set of tools for testing build right in
- write unit tests and integrations tests

## Running tests

- running a single
- running a group of tests (test suites)

### Single

go test -run "name"

### Group of tests

add prefix to name, for example _alpha_
and run

go test -run Test_alpha
