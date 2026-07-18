## Usage

Sources:

- [Project README](https://github.com/vkrinitsyn/rppd/blob/51e263d221d0f1dbafa8667eda5eaf176e389788/README.md)
- [Extension package metadata](https://github.com/vkrinitsyn/rppd/blob/51e263d221d0f1dbafa8667eda5eaf176e389788/rppd-pg/Cargo.toml)
- [Trigger and status implementation](https://github.com/vkrinitsyn/rppd/blob/51e263d221d0f1dbafa8667eda5eaf176e389788/rppd-pg/src/lib.rs)
- [Extension-managed tables](https://github.com/vkrinitsyn/rppd/blob/51e263d221d0f1dbafa8667eda5eaf176e389788/rppd-pg/pg_setup.sql)

`rppd` 0.2.2 is the PostgreSQL trigger component of Rust Python Postgres Discovery. Table events are sent to an external RPPD service, which selects Python source stored in extension-managed tables and runs it against PostgreSQL or an optional etcd queue. It also models node registration, execution logs, and cron schedules.

### Inspect extension state

After installing the native library and creating the extension, inspect configuration before attaching application triggers:

```sql
CREATE EXTENSION rppd;

SELECT id, host, host_name, active_since, master
FROM rppd_config;

SELECT id, schema_table, topic, queue, fn_logging
FROM rppd_function;

SELECT rppd_info();
```

An application table is connected to the service through the trigger function:

```sql
CREATE TRIGGER orders_rppd
AFTER INSERT OR UPDATE OR DELETE ON orders
FOR EACH ROW EXECUTE FUNCTION rppd_event();
```

A matching, administrator-controlled row in `rppd_function` determines the Python code, table or queue topic, serialization, logging, and cleanup behavior. Start and validate the external RPPD service before enabling the trigger.

### Remote-code and transaction boundary

The extension is marked superuser-required, targets Linux, and deliberately turns database rows into executable Python. Anyone able to modify `rppd_function`, its trigger path, service configuration, or etcd messages can potentially influence code execution with the daemon's database and operating-system access. Revoke direct table writes, use a separately sandboxed service identity, authenticate and encrypt every network hop, sign reviewed code outside the database, and audit every configuration change.

Trigger delivery and external execution are not the same transaction as the source-row change. The upstream README notes that an inserted row may not yet be visible to the service and suggests delaying access, which is not a correctness guarantee. Design idempotency, retries, ordering, crash recovery, and observability explicitly; do not assume exactly-once execution. Stored output and errors can also expose application data, so bound retention and access to `rppd_function_log`.
