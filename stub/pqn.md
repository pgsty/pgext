## Usage

Sources:

- [docs/getting-started/pqn-installation.md](https://github.com/pgquery-narrative/pgquerynarrative/blob/b0d91876f50ea2f3519ff12ee27c01c5e253e7a3/docs/getting-started/pqn-installation.md)
- [docs/integrations/postgres-extension.md](https://github.com/pgquery-narrative/pgquerynarrative/blob/b0d91876f50ea2f3519ff12ee27c01c5e253e7a3/docs/integrations/postgres-extension.md)
- [infra/pqn-extension/pqn.control](https://github.com/pgquery-narrative/pgquerynarrative/blob/b0d91876f50ea2f3519ff12ee27c01c5e253e7a3/infra/pqn-extension/pqn.control)
- [infra/pqn-extension/pqn--1.0.sql](https://github.com/pgquery-narrative/pgquerynarrative/blob/b0d91876f50ea2f3519ff12ee27c01c5e253e7a3/infra/pqn-extension/pqn--1.0.sql)
- [infra/pqn-extension/pqn--1.0--1.1.sql](https://github.com/pgquery-narrative/pgquerynarrative/blob/b0d91876f50ea2f3519ff12ee27c01c5e253e7a3/infra/pqn-extension/pqn--1.0--1.1.sql)

`pqn` 1.1 performs restricted query investigation inside PostgreSQL, using selected exposure views and a separate evidence ledger. It does not require the external PgQueryNarrative service.

### Enablement

The supported workflow is PostgreSQL 16–18. As a superuser, create the extension and initialize its ledger. A non-superuser installer first needs a DBA to run the upstream role script and grant the documented memberships.

```sql
CREATE EXTENSION pqn;
SELECT pqn_api.init();
SELECT pqn_api.verify_setup();
```

### Investigation Workflow

`pqn_api.expose_sql` prints proposed exposure SQL; review it before calling `pqn_api.expose`. `pqn_api.enroll` assigns an existing login to a viewer, analyst or administrator group and records limits. `pqn_api.plan`, `pqn_api.run`, `pqn_api.investigate` and `pqn_api.prove` support plans, bounded query execution, findings and comparison evidence. The separate command-line tool also proposes rewrites.

The default exposure scope limits visible columns. Full scope permits planning and measurement over hidden columns and can reveal information through counts or plans. Review grants and function-owner roles before enrolling users.

### Operations

`pg_stat_statements` and its preload/restart are needed only for workload ranking, not for the base extension. Login timeout settings can be changed by users; strict enforcement needs the external scheduler to call `pqn_api.enforce_limits()` with suitable monitoring and cancellation privileges. `DROP EXTENSION pqn` removes extension code but retains separately initialized evidence in `pqn_ledger`; review the upstream uninstall workflow for views, roles and retained data.
