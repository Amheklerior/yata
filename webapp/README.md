# The Yata web application

This is the **web application** for a smooth interaction with the server.

It's a React app built with [Vite](https://vitejs.dev/) and [TypeScript](https://www.typescriptlang.org/).

It uses TanStack's [React Query](https://tanstack.com/query/v4) to handle the data fetching and caching, and a couple of [Radix UI](https://www.radix-ui.com/) components to implement the form and the toast notifications in an accessible way.

## Prerequisites

To be able to run and build the web application, you need to have [Node.js](https://nodejs.org/en/) to be installed on your system.

If you don't have it installed, you can install it with the following command (on MacOS):

```sh
brew install node
```

or via [nvm](https://github.com/nvm-sh/nvm) and similar tools.

NOTE. refer to the [Node.js documentation](https://nodejs.org/en/docs/) for more details.

## Run the project locally

```sh
$ pnpm install # install the dependencies
$ pnpm dev # run the project in development mode
```

## Build the project

```sh
$ pnpm build # build the project
```

Checkout the [`package.json`](./package.json) file for all the scripts available.

## Project structure

The project is structured as follows:

```sh
├── index.html      # the entry point
└── src             # the source code
    ├── assets      # the assets dir
    ├── layouts     # where the baase Layout compoent is defined 
    ├── components  # the components dir, where all other components are defined
    ├── contexts    # here is were the global notification ctx is defined
    ├── lib         # the lib dir (here you can find the api integration layo)
    └── index.css   # the global styles (Tailwind)
```

**Worth having a closer look:**

- [`index.css`](./src/index.css) • the global styles
- [`AddForm.tsx`](./src/components/AddForm.tsx) • the input form for adding new tasks
- [`TaskItem.tsx`](./src/components/TaskItem.tsx) • the task item component, from where the user can mark it as complete or delete it
- [`NotificationCenter.tsx`](./src/components/NotificationCenter.tsx) • the notification center, where the toast notifications are rendered
- [`lib`](./src/lib/) • the lib dir, where the api integration layer is defined, with types, validation schemas, and queries

## Areas of improvement

- define a more robust api integration layer
- add some test coverage
