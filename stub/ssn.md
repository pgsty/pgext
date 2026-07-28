## Usage

Sources:

- [Official extension SQL (ssn.sql)](https://api.pgxn.org/src/ssn/ssn-1.0.0/sql/ssn.sql)

`ssn` — SSN type implementation for PostgreSQL. Use it when application data needs this type, domain, or its operators. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

The reviewed distribution uses a legacy SQL or non-control installation layout, so it does not establish a modern standalone `CREATE EXTENSION` and upgrade workflow. Follow the pinned upstream installation mechanism and verify the installed objects in an isolated database.

### Important Objects

- `ssns` is a table installed or managed by the extension.

### Requirements and Caveats

- The catalog records version `1.0.0`.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
