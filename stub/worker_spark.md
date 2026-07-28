## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/worker_spark/worker_spark-0.0.1/README.md)

`worker_spark` — A background worker for PostgreSQL 9.3 which executes a procedure at regular intervals. Use it for the corresponding scheduling, temporal, or time-series workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

The reviewed distribution uses a legacy SQL or non-control installation layout, so it does not establish a modern standalone `CREATE EXTENSION` and upgrade workflow. Follow the pinned upstream installation mechanism and verify the installed objects in an isolated database.

### Requirements and Caveats

- The catalog records version `0.0.1`.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
