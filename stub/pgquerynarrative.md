## Usage

Sources:

- [docs/integrations/postgres-extension.md](https://github.com/pgquery-narrative/pgquerynarrative/blob/b0d91876f50ea2f3519ff12ee27c01c5e253e7a3/docs/integrations/postgres-extension.md)
- [infra/postgres-extension/pgquerynarrative.control](https://github.com/pgquery-narrative/pgquerynarrative/blob/b0d91876f50ea2f3519ff12ee27c01c5e253e7a3/infra/postgres-extension/pgquerynarrative.control)
- [infra/postgres-extension/pgquerynarrative--1.1.sql](https://github.com/pgquery-narrative/pgquerynarrative/blob/b0d91876f50ea2f3519ff12ee27c01c5e253e7a3/infra/postgres-extension/pgquerynarrative--1.1.sql)

`pgquerynarrative` 1.1 calls a running PgQueryNarrative service from SQL. It is separate from the in-database investigation extension `pqn`.

### Core Workflow

As a superuser, install the HTTP dependency before the wrapper and grant access to an existing role. Configure a reachable, trusted service endpoint.

```sql
CREATE EXTENSION http;
CREATE EXTENSION pgquerynarrative;
SELECT pgquerynarrative_set_api_url('http://localhost:8080');
SELECT pgquerynarrative_grant_access('app_reader');
```

### API and Boundaries

`pgquerynarrative_set_api_key` sets each caller’s session credential. `pgquerynarrative_run_query` runs a read-only query through the service; `pgquerynarrative_generate_report` requests a report, and `pgquerynarrative_list_saved` lists saved queries. `pgquerynarrative_revoke_access` removes access.

Without `http` at install time, the extension creates pending-response functions rather than real HTTP calls. Adding the dependency later does not rewrite those functions automatically. Version 1.1 revokes function execution from `PUBLIC`; the owner controls the stored API URL. Statements containing API-key literals may enter query logs, and SQL is transmitted to the service. The server’s authorization and result limits still apply. PostgreSQL 16–18 is the documented service range; no preload is needed for these SQL wrappers.
