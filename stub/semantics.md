## Usage

Sources:

- [Official upstream README](https://github.com/aquametalabs/aquameta/blob/ec5dbe6b84ea61b5e8b4b10263a0911b848b0025/README.md)
- [Official extension control file (semantics.control)](https://github.com/aquametalabs/aquameta/blob/ec5dbe6b84ea61b5e8b4b10263a0911b848b0025/extensions/semantics/semantics.control)

`semantics` — Aquameta Semantics Extension. Use it when an application needs this specific database capability. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION semantics;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.5.0`.
- Install the confirmed extension dependencies first: `meta`, `widget`.
- The control file marks the extension as non-relocatable.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
