## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/petere/autopex/blob/886690f4cdf4977eb1917cde6696b7dac3fcc43c/README.md)
- [Version 0 SQL objects](https://github.com/petere/autopex/blob/886690f4cdf4977eb1917cde6696b7dac3fcc43c/autopex--0.sql)
- [Event-trigger implementation](https://github.com/petere/autopex/blob/886690f4cdf4977eb1917cde6696b7dac3fcc43c/autopex.c)

`autopex` is an archived PostgreSQL 9.3-era experiment that intercepts `CREATE EXTENSION` with an event trigger. Together with the external Pex tool, it attempts to download, build, and install an extension automatically when PostgreSQL cannot find it locally.

```sql
CREATE EXTENSION autopex;
SELECT evtname, evtevent, evtenabled
FROM pg_event_trigger
WHERE evtname = 'autopex';
```

The operating-system PostgreSQL account must have a Pex installation and build environment. This design turns a database DDL request into network access, arbitrary upstream source retrieval, compilation, and writes to the server installation. It therefore expands database superuser access into operating-system supply-chain execution.

Do not use `autopex` on production or current PostgreSQL. Prefer reviewed, reproducible packages installed by configuration management. If preserving a historical test environment, deny untrusted `CREATE` privileges, isolate outbound network access, pin every artifact and checksum, and run only inside a disposable host. Dropping the extension should be verified to remove its event trigger before further DDL.
