

## Usage

> [pgl_ddl_deploy: automated ddl deployment using pglogical](https://github.com/enova/pgl_ddl_deploy)

Transparent DDL replication for PostgreSQL 9.5+ supporting both pglogical and native logical replication. Automatically propagates DDL changes (CREATE TABLE, ALTER TABLE, etc.) to subscribers.

### Enabling

```sql
CREATE EXTENSION pgl_ddl_deploy;
```

### Configuration

Insert configuration into the `pgl_ddl_deploy.set_configs` table:

```sql
-- Replicate DDL for all user schemas and auto-add new tables
INSERT INTO pgl_ddl_deploy.set_configs (set_name, include_schema_regex, driver)
VALUES ('default', '.*', 'native'::pgl_ddl_deploy.driver);

-- Or with pglogical driver
INSERT INTO pgl_ddl_deploy.set_configs (set_name, include_schema_regex, driver)
VALUES ('default', '.*', 'pglogical'::pgl_ddl_deploy.driver);

-- Maintain only specific tables already in replication (ALTER TABLE only)
INSERT INTO pgl_ddl_deploy.set_configs (set_name, include_only_repset_tables, driver)
VALUES ('my_tables', TRUE, 'native'::pgl_ddl_deploy.driver);
```

### Deploy Event Triggers

After configuring, deploy the event triggers:

```sql
SELECT pgl_ddl_deploy.deploy(set_config_id) FROM pgl_ddl_deploy.set_configs;
```

### Key Configuration Options

- `driver`: `native` or `pglogical`
- `set_name`: publication name or pglogical replication set name
- `include_schema_regex`: regex to match schemas for DDL replication
- `include_only_repset_tables`: if true, only ALTER TABLE for tables already in replication
- `lock_safe_deployment`: if true, DDL executes in a low lock_timeout loop on subscriber
- `allow_multi_statements`: if true, multi-statement DDL can be propagated
- `queue_subscriber_failures`: if true, failed DDL on subscriber is queued for retry
- `ddl_only_replication`: replicate schema only without auto-adding tables to data replication

### Monitoring

```sql
-- View unhandled DDL events
SELECT * FROM pgl_ddl_deploy.unhandled;

-- View failed subscriber DDL
SELECT * FROM pgl_ddl_deploy.subscriber_logs WHERE NOT succeeded;

-- Retry failed DDL on subscriber
SELECT pgl_ddl_deploy.retry_all_subscriber_logs();
```

### Checking Resolved Schemas

```sql
SELECT pgl_ddl_deploy.resolved_regex_include_schemas(set_config_id)
FROM pgl_ddl_deploy.set_configs;
```
