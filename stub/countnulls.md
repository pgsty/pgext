## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/countnulls/countnulls-1.0.0/README)
- [Official extension SQL (countnulls.sql)](https://api.pgxn.org/src/countnulls/countnulls-1.0.0/countnulls.sql)

`countnulls` — Simple function to count the number of NULL arguments. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

The reviewed distribution uses a legacy SQL or non-control installation layout, so it does not establish a modern standalone `CREATE EXTENSION` and upgrade workflow. Follow the pinned upstream installation mechanism and verify the installed objects in an isolated database.

### Important Objects

- `countnulls("any")` is an extension function and returns `int`.
- `countnulls("any", "any")` is an extension function and returns `int`.
- `countnulls("any", "any", "any")` is an extension function and returns `int`.
- `countnulls("any", "any", "any", "any")` is an extension function and returns `int`.
- `countnulls("any", "any", "any", "any", "any")` is an extension function and returns `int`.
- `countnulls("any", "any", "any", "any", "any", "any")` is an extension function and returns `int`.
- `countnulls("any", "any", "any", "any", "any", "any", "any")` is an extension function and returns `int`.
- `countnulls("any", "any", "any", "any", "any", "any", "any", "any")` is an extension function and returns `int`.
- `countnulls(VARIADIC "any")` is an extension function and returns `int`.

### Requirements and Caveats

- The catalog records version `1.0.0`.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
