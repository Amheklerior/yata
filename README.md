# YATA

This is **YATA** (_Yet Anotehr Todo App_).

## Run the project

To run the project, open the terminal window run the following command:

```sh
docker compose up
```

Then you can open the webapp at [http://localhost:80](http://localhost:80).

NOTE. you need to have docker installed on your machine.

## Assumptions

- freedom in ui/ux desing
- freedom in choosing the tech stack (outside the FE)
- no multitenant
- no auth
- no ci/cd
- no infra
- no internationalization

## Project structure

The project is structured in a BE server and a FE webapp.

You can find the server code in the [`./server`](./server) directory, and the webapp code in the [`./webapp`](./webapp) directory.

Refer to the [server doc](./server/README.md) and the [webapp doc](./webapp/README.md) for more details.

### Why a Go server?

I simply thought it was the perfect opportunity for me to deepen my knowledge of Go and learn more about how to build a REST API with it.

## Possible areas of improvement

general:

- setup a CI/CD pipeline with GitHub Actions
- setup a git hooks to run the linter and tests on commit/push
- setup an automatic changelog generator

on the server side:

- parametrize the **HOST** and **PORT** the server listen to in an .env file
- integrating [Air](https://github.com/air-verse/air) (live server reloading during development) to save a bit of time
- replace the [`getTaskIdFromURLParam`](./internal/api/tasks_handler.go) function with a custom middleware to attach to the `/:id` subroutes.
- implement safe cuncurrent access to the database (using a mutex) -- skipped because the idea was to do the following instead...
- implement a persistent [SQLite](https://www.sqlite.org/index.html) database instead of the in-memory one

on the webapp side:

- implement a more robust http client (or swap it with a third party lib)
- improve API error handling
- improve data validation
- add motion to the UI to provide a better UX (eg. entry and exit transit animations on updating the list)
- integrate a tab focus-trap to prevent the user from tabbing outside the viewport when using the keyboard
- add E2E tests for the webapp

## AI usage

AI usage has been kept to a minimum, with some worth-mentioning exceptions:

- resolve issues with the CLI tool
- resolve the cors issue
- implement the sound fx
- help in writing the README docs
- scaffold the Dockerfile and the docker-compose.yaml files
- removing the default browser ring on interactive elements
- solve other minor issues along the way
