## Usage

Sources:

- [Official README](https://github.com/fraktalio/fmodel-rust-postgres/blob/3f9c55771582e1fc7fd77858b0aded2fcb766cae/README.md)
- [Extension control file](https://github.com/fraktalio/fmodel-rust-postgres/blob/3f9c55771582e1fc7fd77858b0aded2fcb766cae/fmodel_rust_postgres.control)
- [Rust extension entry point](https://github.com/fraktalio/fmodel-rust-postgres/blob/3f9c55771582e1fc7fd77858b0aded2fcb766cae/src/lib.rs)
- [Event-sourcing schema](https://github.com/fraktalio/fmodel-rust-postgres/blob/3f9c55771582e1fc7fd77858b0aded2fcb766cae/sql/event_sourcing.sql)

`fmodel_rust_postgres` embeds a concrete Rust event-sourcing demonstration in PostgreSQL. The checked source models restaurant and order commands with fmodel-rust; it is a project template and domain example, not a generic SQL-configurable event-sourcing framework.

### Core Workflow

```sql
CREATE EXTENSION fmodel_rust_postgres;

SELECT *
FROM handle(
  '{"CreateRestaurant":{"name":"Ada Cafe","menu":[]}}'::Command
);

SELECT *
FROM handle_all(ARRAY[
  '{"CreateRestaurant":{"name":"North Cafe","menu":[]}}'::Command,
  '{"CreateRestaurant":{"name":"South Cafe","menu":[]}}'::Command
]);
```

`handle(Command)` returns an array of `Event` values for one command. `handle_all(Command[])` processes a command batch in one PostgreSQL transaction, so later commands can observe earlier results and any error rolls the batch back. Installed objects also include `deciders`, `events`, `restaurants`, and `orders`, plus triggers that project restaurant and order events and enforce the previous-event chain.

### Domain and Security Boundaries

The `Command` and `Event` types and their JSON representation are compiled Rust domain types. Adding a domain, changing a payload, or changing projection behavior requires changing the Rust source and rebuilding the extension; arbitrary applications cannot register deciders through SQL. The bootstrap creates generic unqualified object names, so install in a dedicated database or schema plan and inspect collisions before enabling it.

The event tables and functions receive ordinary PostgreSQL default privileges unless administrators revoke them. Direct writes or DDL against the tables can bypass the intended command path, and the extension does not add row-level security, tenant isolation, or an authorization policy. Restrict table writes and function execution, expose an application-specific API, and design backup, replay, retention, and schema-evolution procedures before production use. The pinned build uses pgrx 0.16 with PostgreSQL 13–16 features; match the binary to the exact server major and test upgrades.
