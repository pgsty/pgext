## Usage

Sources:

- [Official upstream README](https://github.com/jlockerman/table_builder_poc/blob/b24fb265a70d8dcabc0b9d8331f9d256b212f740/Readme.md)
- [Official extension control file (table_builder.control)](https://github.com/jlockerman/table_builder_poc/blob/b24fb265a70d8dcabc0b9d8331f9d256b212f740/extension/table_builder.control)
- [Official implementation source](https://github.com/jlockerman/table_builder_poc/blob/b24fb265a70d8dcabc0b9d8331f9d256b212f740/extension/src/lib.rs)

`table_builder` — Proof-of-concept for a safe rust table builder and querier. Use it when an application needs this specific database capability. Upstream describes it as a proof of concept.

### Core Workflow

```sql
CREATE EXTENSION table_builder;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Upstream describes the project as a proof of concept.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
