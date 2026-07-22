## Usage

Sources:

- [Aiven Extras README at the reviewed commit](https://github.com/aiven/aiven-extras/blob/ed2a4a676034b1ca3edb767d5805c482262d3f9b/README.md)
- [Control template at the reviewed commit](https://github.com/aiven/aiven-extras/blob/ed2a4a676034b1ca3edb767d5805c482262d3f9b/aiven_extras.control.in)
- [Build version at the reviewed commit](https://github.com/aiven/aiven-extras/blob/ed2a4a676034b1ca3edb767d5805c482262d3f9b/Makefile)

`aiven_extras` is an Aiven-maintained collection of narrowly scoped wrappers for operations that normally require elevated PostgreSQL privileges. After a privileged operator installs it, eligible non-superusers can administer selected logical-replication objects, configure session-local `auto_explain`, adjust selected pgaudit settings, and run controlled dblink operations. It installs into the fixed `aiven_extras` schema and does not grant general superuser access.

### Logical Replication Workflow

The calling role still needs `REPLICATION`. Use schema-qualified table names and restrict execution privileges on the wrappers to the operational roles that need them.

```sql
CREATE EXTENSION aiven_extras;

SELECT aiven_extras.pg_create_publication(
  'pub1', 'INSERT,UPDATE', 'public.foo', 'public.bar'
);

SELECT * FROM aiven_extras.pg_list_all_subscriptions();
```

The replication surface includes `pg_create_publication`, `pg_create_publication_for_all_tables`, `pg_create_subscription`, `pg_drop_subscription`, `pg_alter_subscription_enable`, `pg_alter_subscription_disable`, `pg_alter_subscription_refresh_publication`, `pg_list_all_subscriptions`, and `pg_stat_replication_list`. Subscription creation still performs network and replication-slot work; it is not made transactional by the wrapper.

### Session Diagnostics

`auto_explain_load` loads `auto_explain` for the current session. The `set_auto_explain_*` family then changes that session's settings only.

```sql
SELECT aiven_extras.auto_explain_load();
SELECT aiven_extras.set_auto_explain_log_min_duration('2000');
SELECT aiven_extras.set_auto_explain_log_analyze('on');
```

Other exposed helpers include `claim_public_schema_ownership`, `dblink_record_execute`, `dblink_slot_create_or_drop`, `explain_statement`, `session_replication_role`, `set_pgaudit_parameter`, and `set_pgaudit_role_parameter`.

### Operational Notes

- The current reviewed Makefile identifies version `1.1.20`, matching the catalog. Pin the provider package or source commit whose behavior you tested.
- Publication and subscription wrappers are privileged interfaces, not a privilege bypass. Audit their grants and revoke `PUBLIC` execution where the deployment's role model requires it.
- Subscription connection strings can contain credentials. Keep them out of logs, query history, and error-reporting pipelines, and prefer provider-supported secret handling.
- Enabling plan analysis can add execution overhead and expose query text or values in logs. Use `auto_explain` selectively, especially with analysis enabled.
