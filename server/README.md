# The YATA server

The YATA server is a simple REST API built in Go, with [Chi](https://github.com/go-chi/chi).

## Prerequisites

To be able to run and build the server, you need to have [Go](https://go.dev/) to be installed on your system.

If you don't have it installed, you can install it with the following command (on MacOS):

```sh
brew install go
```

NOTE. refer to the [Go documentation](https://go.dev/doc/install) for more details.

## The Makefile

There's a [`Makefile`](./makefile) to help you run the server in development mode.

```sh
$ cd ./server
$ make dev # run the server in development mode
$ make run # build and run the server from the /dist dir
$ make build # build the server into the /dist dir
$ make fmt # format the code
$ make test # run the tests
$ make clean # clean the build output dir
```

## Project structure

The project is structured as follows:

```sh
├── internal # this is where the goodies live
│   ├── api # the http handlers
│   ├── app # the application definition
│   ├── routes # the chi router
│   └── store # the store interface and its in-memory implementation
│
└─── main.go # the entry point
```

## The tests

Tests are in the [/api](./internal/api/) directory, colocated with the code they test, in the `_test.go` files.

I decided to implement top level tests in order to test the whole stack from handling the HTTP requests to the database implementation.

## Other things I wanted to implement

- integrating [Air](https://github.com/air-verse/air) (live server reloading during development) to save a bit of time
- replace the [`getTaskIdFromURLParam`](./internal/api/tasks_handler.go) function with a custom middleware to attach to the `/:id` subroutes.
- implement safe cuncurrent access to the database (using a mutex) -- skipped because the idea was to do the following instead...
- implement a persistent [SQLite](https://www.sqlite.org/index.html) database instead of the in-memory one
