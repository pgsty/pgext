## Usage

Sources:

- [Console guide](https://github.com/mkleczek/pgwrh/blob/v1.0.0-alpha1/pgwrh_ui/README.md)
- [Control file](https://github.com/mkleczek/pgwrh/blob/v1.0.0-alpha1/pgwrh_ui/pgwrh_ui.control)
- [Viewer role setup](https://github.com/mkleczek/pgwrh/blob/v1.0.0-alpha1/pgwrh_ui/readonly.sql)
- [Operator role setup](https://github.com/mkleczek/pgwrh/blob/v1.0.0-alpha1/pgwrh_ui/operator.sql)
- [PostgREST configuration](https://github.com/mkleczek/pgwrh/blob/v1.0.0-alpha1/pgwrh_ui/postgrest.conf)

`pgwrh_ui` 1.0.0-alpha1 provides a browser console for a PostgreSQL 18 pgwrh controller. It displays placement, replica state, and rollout blockers, and offers optional management actions. Install it only in the controller database. A separate PostgREST process serves its SQL endpoints and embedded browser assets.

### Enable the Console

The database needs `pgwrh` 1.0.0-alpha1 and PL/pgSQL. The console does not require `pgwrh_wait` or its preload. PostgREST must support custom media handlers; the upstream Compose example uses 14.16. As a trusted administrator on the controller:

```sql
CREATE EXTENSION pgwrh_ui;
```

Create the deployment roles using the supplied `readonly.sql` and, if management is needed, `operator.sql`. Roles are not created or removed by the extension. From the source checkout, with `CONTROLLER_ADMIN_URI` set to an administrator connection:

```sh
psql -X -v ON_ERROR_STOP=1 "$CONTROLLER_ADMIN_URI" -f pgwrh_ui/readonly.sql
```

```sql
CREATE ROLE pgwrh_ui_authenticator LOGIN NOINHERIT;
GRANT pgwrh_ui_viewer TO pgwrh_ui_authenticator;
```

Configure authentication for that login using the deployment’s normal method, then start the external service with the supplied configuration:

```sh
export PGRST_DB_URI='postgresql://pgwrh_ui_authenticator@localhost/controller_database'
postgrest pgwrh_ui/postgrest.conf
```

The supplied configuration exposes the console at the local `/rpc/index` route with viewer access. Expose only the `pgwrh_ui` schema. Reload PostgREST’s schema cache after extension or function changes:

```sql
NOTIFY pgrst, 'reload schema';
```

### Access and Management

`pgwrh_ui_viewer` can view every group on the controller. `pgwrh_ui_operator` can additionally change replica weights and routing status, register replicas, and start, commit, or roll back rollouts. Grant operator membership only after running the operator setup script and configuring authentication. The console has no built-in login page or token manager; remote access needs authenticated PostgREST roles or an authenticated reverse proxy.

Replica provisioning, replication credentials, source grants, group creation, and table policies remain SQL administration tasks. A group has one shared draft; stale submissions return HTTP 409 and must be reviewed again. Maintenance routing changes do not delete data or stop assigned replication.

### Monitoring Limits

Readiness reports have no per-replica heartbeat timestamp or durable history. Zero blockers does not prove live reachability. Displayed WAL lag is a byte distance, not elapsed time or a read-consistency promise. Controller failure also makes the console unavailable. This alpha supports fresh installation only; package installation does not start PostgREST.
