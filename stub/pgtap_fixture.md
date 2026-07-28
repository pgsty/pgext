## Usage

Sources:

- [Official upstream README](https://github.com/tyrusj/pgtap_fixture/blob/5696f7dabdcce205efb78879df9b8de7123e0071/README.md)
- [Official extension control file (pgtap_fixture.control)](https://github.com/tyrusj/pgtap_fixture/blob/5696f7dabdcce205efb78879df9b8de7123e0071/install/pgtap_fixture.control)

`pgtap_fixture` — A PostgreSQL extension that allows more complex fixtures to be created for pgTAP tests. Use it when an application needs this specific database capability. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION pgtap_fixture;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The catalog records version `0.0.1`.
- Install the confirmed extension dependencies first: `plpgsql`, `pgtap`.
- The control file marks the extension as relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
