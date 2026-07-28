## Usage

Sources:

- [Official project page](https://github.com/timescale/promscale_extension)

`promscale` — Deprecated tables, types, and functions supporting Promscale. Use it when collecting or interpreting the corresponding PostgreSQL statistics. The reviewed upstream material marks this capability deprecated.

### Core Workflow

This component has no confirmed standalone `CREATE EXTENSION` workflow in the reviewed source. Build, load, or enable it only through the exact upstream mechanism, then verify the resulting server behavior in an isolated database.

### Requirements and Caveats

- The catalog records version `0.8.0`.
- Install the confirmed extension dependencies first: `timescaledb`.
- Upstream material contains an explicit deprecation boundary.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
