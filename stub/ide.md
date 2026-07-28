## Usage

Sources:

- [Official upstream README](https://github.com/aquametalabs/aquameta/blob/ec5dbe6b84ea61b5e8b4b10263a0911b848b0025/extensions/ide/README.md)
- [Official extension control file (ide.control)](https://github.com/aquametalabs/aquameta/blob/ec5dbe6b84ea61b5e8b4b10263a0911b848b0025/extensions/ide/ide.control)

`ide` — Aquameta IDE. Use it when an application needs this specific database capability. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION ide;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.5.0`.
- Install the confirmed extension dependencies first: `endpoint`, `widget`, `meta`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
