## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/log_functions/log_functions-1.0.0/README)

`log_functions` — log_functions is a PostgreSQL module that logs each function executed. Use it when collecting or interpreting the corresponding PostgreSQL statistics. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

The reviewed distribution uses a legacy SQL or non-control installation layout, so it does not establish a modern standalone `CREATE EXTENSION` and upgrade workflow. Follow the pinned upstream installation mechanism and verify the installed objects in an isolated database.

### Requirements and Caveats

- The catalog records version `1.0.0`.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
