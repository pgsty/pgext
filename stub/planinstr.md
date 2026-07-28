## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/planinstr/planinstr-0.0.1/README.md)

`planinstr` — Be sure that you have pg_config installed and in your path. If you used a package management system such as RPM to install PostgreSQL, be sure that the -devel package is also installed. If necessary tell the build process where to find it:. Use it when collecting or interpreting the corresponding PostgreSQL statistics. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

This component has no confirmed standalone `CREATE EXTENSION` workflow in the reviewed source. Build, load, or enable it only through the exact upstream mechanism, then verify the resulting server behavior in an isolated database.

### Requirements and Caveats

- The catalog records version `0.0.1`.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
