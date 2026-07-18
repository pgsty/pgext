## Usage

Sources:

- [Alibaba Cloud rds_ccl SQL-throttling documentation](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/use-the-rds-ccl-extension-to-perform-sql-throttling-on-an-apsaradb-rds-for-postgresql-instance)
- [Alibaba Cloud RDS product overview](https://www.alibabacloud.com/help/en/rds/product-overview/what-is-apsaradb-rds-what-is-apsaradb-rds)

`rds_ccl` is an Alibaba Cloud ApsaraDB RDS for PostgreSQL extension that throttles statements by normalized `query_id`. Rules cap concurrency and queue excess executions to protect an instance from bursts or expensive recurring queries.

```sql
CREATE EXTENSION rds_ccl;
SELECT * FROM rds_show_current_db_ccl_rule();
```

Enable the instance parameters `rds_enable_ccl=on` and `compute_query_id=auto` or `on` before use, and create rules with a privileged RDS account. Rules must be reloaded after restart or primary/secondary failover. An overly low concurrency or queue limit can block legitimate traffic, so derive rules from measured query IDs, stage them gradually, monitor queues, and keep an immediate disable path.
