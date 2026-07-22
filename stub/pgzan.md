## Usage

Sources:

- [Official README and WIP warning](https://github.com/brahmlower/pgzan/blob/b2e45f17059639863dae8fdadff805fe51ca55ac/readme.md)
- [Rust implementation](https://github.com/brahmlower/pgzan/blob/b2e45f17059639863dae8fdadff805fe51ca55ac/src/lib.rs)
- [Extension control file](https://github.com/brahmlower/pgzan/blob/b2e45f17059639863dae8fdadff805fe51ca55ac/pgzan.control)

pgzan is a proof of concept for calling an Oso/Polar authorization policy from PostgreSQL, with a long-term goal of supporting Zanzibar-like row-level authorization. Upstream explicitly warns that it is work in progress and must not approach a production database. The checked code has one hard-coded policy and does not implement the configurable ACL described by the aspirational README.

### Core Workflow

Use only the actual one-argument function in a disposable database. It parses a JSON user with an ID and role, then checks the fixed `update` permission on the fixed Organizations resource:

```sql
CREATE EXTENSION pgzan;

SELECT pgzan_check(
  '{"id":"07b30b3a-8da9-465e-96ef-4054f870cd8a", "role":"manager"}'
);

SELECT pgzan_check(
  '{"id":"07b30b3a-8da9-465e-96ef-4054f870cd8a", "role":"readonly"}'
);
```

The first call returns true and the second false under the compiled policy. This demonstrates evaluation only; it does not connect a session identity, table row, requested operation, or external relationship store.

### Important Objects

- `pgzan_check` parses one JSON text argument and evaluates the hard-coded authorization decision.
- `hello_pgzan` returns a fixed greeting used by the development test.
- The compiled policy defines readonly and manager roles and a fixed set of resource names and permissions.

### Security and Compatibility Notes

The README's illustrative RLS call uses a different function signature from the source and therefore is not an implemented workflow. A caller controls the JSON role value, so using the boolean directly as an RLS decision would let the caller claim manager unless a trusted boundary constructs the input. Parsing errors can abort the statement, and a new Oso evaluator is built on every call. Do not use pgzan to protect data; evaluate the maintained pgauthz extension or an external authorization service instead.
