# Caté: The Smart Cat Auto Feeder

Welcome to **Caté**, an all-in-one system designed to help cat owners provide the healthiest possible diet for their pets. Our goal is to reduce the inconvenience of meal preparation and scheduling while ensuring happier and healthier cats.

## Project Overview

- **App (Frontend)**
  - A user-facing interface for cat owners to manage feeding schedules, track health metrics, and configure the system.
- **API (Backend)**
  - A Go-based API layer that provides endpoints for the frontend and other services to communicate with the database and business logic.
- **Worker (Background Jobs)**
  - A Go-based worker that processes heavier or asynchronous tasks (e.g., analytics, notifications, feed scheduling).
- **Migrator (Database Management)**
  - A Go-based service responsible for database schema creation, migration, and versioning.
- **Controller (On-Board Controller in Embedded Rust)**
  - Rust code for the physical hardware that controls feeding mechanisms and other cat feeder operations.

## Prerequisites

- **Docker** (for containerization and distribution)
- **Docker Compose** (for running multiple containers/services together)
- **Node.js & npm** (or Yarn) for Nx CLI usage
- **Nx CLI** (installed or run via npx)
- **GoLang Migrate** (for database migrations) run `brew install golang-migrate`


## Dev Setup

1. Run `yarn install` to install dependencies

## Building the Project

We use Nx to orchestrate builds across all our services. Each service has a docker-build target in its project.json. To build all images at once:

# Build Docker images for all services
npx nx run-many -t docker-build

This will produce Docker images for:
1. **API** (Go-based)
2. **Worker** (Go-based)
3. **Migrator** (Go-based)
4. **Controller** (Rust-based)
5. **App** (the frontend, should it need a Docker build in the future)

## Running the Project

Once you have built all the Docker images, you can run everything in one shot via Docker Compose:

docker compose up

Make sure your docker-compose.yml references the images you built. It could look something like this (simplified example):

version: '3.8'
services:
  api:
    image: api:latest
    ports:
      - "3000:3000"
    # ... additional config

  worker:
    image: worker:latest
    # ... additional config

  migrator:
    image: migrator:latest
    # ... additional config

  controller:
    image: controller:latest
    # ... additional config

  app:
    image: app:latest
    ports:
      - "8080:80"
    # ... additional config

Update this as needed for environment variables, networks, volumes, etc.

## Testing

To run tests for a specific project, you can do:

npx nx test <project-name>

Or to run all tests in the workspace:

npx nx run-many --target=test --all

## Linting

To check for code style, run:

npx nx lint <project-name>

Or to lint all projects:

npx nx run-many --target=lint --all

## Want to Contribute?

We welcome contributions of all kinds, from bug reports, documentation improvements, to feature requests. Please submit a pull request or open an issue to get started!

## Why Cate?

The Cate system is designed to:
- **Automate** feeding schedules to match recommended dietary guidelines.
- **Track** health and nutritional data for each cat, enabling better insight into their diets.
- **Empower** owners with remote access and real-time control no matter where they are.
- **Benefit** cats by providing consistent, portion-controlled meals for improved health.

If you have any questions or suggestions, please open an issue or join the discussion. Happy feeding!

## NX Commands

## Run tasks

To run tasks with Nx use:

```sh
npx nx <target> <project-name>
```

For example:

```sh
npx nx build myproject
```

These targets are either [inferred automatically](https://nx.dev/concepts/inferred-tasks?utm_source=nx_project&utm_medium=readme&utm_campaign=nx_projects) or defined in the `project.json` or `package.json` files.

[More about running tasks in the docs &raquo;](https://nx.dev/features/run-tasks?utm_source=nx_project&utm_medium=readme&utm_campaign=nx_projects)

## Add new projects

While you could add new projects to your workspace manually, you might want to leverage [Nx plugins](https://nx.dev/concepts/nx-plugins?utm_source=nx_project&utm_medium=readme&utm_campaign=nx_projects) and their [code generation](https://nx.dev/features/generate-code?utm_source=nx_project&utm_medium=readme&utm_campaign=nx_projects) feature.

To install a new plugin you can use the `nx add` command. Here's an example of adding the React plugin:
```sh
npx nx add @nx/react
```

Use the plugin's generator to create new projects. For example, to create a new React app or library:

```sh
# Generate an app
npx nx g @nx/react:app demo

# Generate a library
npx nx g @nx/react:lib some-lib
```

You can use `npx nx list` to get a list of installed plugins. Then, run `npx nx list <plugin-name>` to learn about more specific capabilities of a particular plugin. Alternatively, [install Nx Console](https://nx.dev/getting-started/editor-setup?utm_source=nx_project&utm_medium=readme&utm_campaign=nx_projects) to browse plugins and generators in your IDE.

[Learn more about Nx plugins &raquo;](https://nx.dev/concepts/nx-plugins?utm_source=nx_project&utm_medium=readme&utm_campaign=nx_projects) | [Browse the plugin registry &raquo;](https://nx.dev/plugin-registry?utm_source=nx_project&utm_medium=readme&utm_campaign=nx_projects)


[Learn more about Nx on CI](https://nx.dev/ci/intro/ci-with-nx#ready-get-started-with-your-provider?utm_source=nx_project&utm_medium=readme&utm_campaign=nx_projects)

## Install Nx Console

Nx Console is an editor extension that enriches your developer experience. It lets you run tasks, generate code, and improves code autocompletion in your IDE. It is available for VSCode and IntelliJ.

[Install Nx Console &raquo;](https://nx.dev/getting-started/editor-setup?utm_source=nx_project&utm_medium=readme&utm_campaign=nx_projects)

## Useful links

Learn more:

- [Learn more about this workspace setup](https://nx.dev/getting-started/intro#learn-nx?utm_source=nx_project&amp;utm_medium=readme&amp;utm_campaign=nx_projects)
- [Learn about Nx on CI](https://nx.dev/ci/intro/ci-with-nx?utm_source=nx_project&utm_medium=readme&utm_campaign=nx_projects)
- [Releasing Packages with Nx release](https://nx.dev/features/manage-releases?utm_source=nx_project&utm_medium=readme&utm_campaign=nx_projects)
- [What are Nx plugins?](https://nx.dev/concepts/nx-plugins?utm_source=nx_project&utm_medium=readme&utm_campaign=nx_projects)

And join the Nx community:
- [Discord](https://go.nx.dev/community)
- [Follow us on X](https://twitter.com/nxdevtools) or [LinkedIn](https://www.linkedin.com/company/nrwl)
- [Our Youtube channel](https://www.youtube.com/@nxdevtools)
- [Our blog](https://nx.dev/blog?utm_source=nx_project&utm_medium=readme&utm_campaign=nx_projects)
