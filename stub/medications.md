## Usage

Sources:

- [Official upstream README](https://github.com/constructive-io/pgpm-demo/blob/7793d0dca37c96d92e0cd4c16e25c8140142039e/packages/medications/README.md)
- [Official extension control file (medications.control)](https://github.com/constructive-io/pgpm-demo/blob/7793d0dca37c96d92e0cd4c16e25c8140142039e/packages/medications/medications.control)

`medications` — **🛠 Built by the Constructive team — creators of modular Postgres tooling for secure, composable backends. If you like our work, contribute on GitHub.**. Use it when an application needs this specific database capability. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION medications;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.0`.
- Install the confirmed extension dependencies first: `plpgsql`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
