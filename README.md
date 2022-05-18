# Saving Pokemons

## Overview
The Project I build is to create a Pokémon API from a CSV file using Go with GIN framework

## Setup local development

### Install tools

- [Docker desktop](https://www.docker.com/products/docker-desktop)
- [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
- [Sqlc](https://github.com/kyleconroy/sqlc#installation)

### Setup infrastructure

- Start postgres container:

    ```bash
    make container
    ```

- Create database:

    ```bash
    make createdb
    ```

- Run db migration:

    ```bash
    make migrateup
    ```

### How to run

- Run test:

    ```bash
    make test
    ```

- Run program:

    ```bash
    make server
    ```

- Example for using API:

    ```bash
    http://localhost:8080/pokemon?page_id=1&page_size=100&hp=45&attack=49&defense=49
    ```