# GoLang Starter For Web REST API Project

## Modules

- **Fiber** for Router
- **ent** for ORM
- **ozzo-validation** for Request validator
- **go-ini** for configuration ini files
- etc.

## Features:

- Register
- Login JWT Token
- `me` endpoint (protected routes)

## Structures

- `config` directory : All configuration setup
- `config.ini` files to define configuration parameters
- `ent` for entity ORM database layer
- `handlers` directory to define all request reponse handler
- `middleware` directory to define all app middleware
- `routes` directory to define all app routes

## Commands

- `make schema {--name}` : to create new ORM entity schema
- `make run` : to run the main app
- `make generate` : to generate all entity file

&copy; Catur Andi Pamungkas
