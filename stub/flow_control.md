## Usage

Sources:

- [Amazon RDS Multi-AZ DB cluster concepts](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html)
- [Amazon RDS for PostgreSQL extension versions](https://docs.aws.amazon.com/AmazonRDS/latest/PostgreSQLReleaseNotes/postgresql-extensions.html)

flow_control is an AWS-managed extension for RDS for PostgreSQL Multi-AZ DB clusters. Reader background workers report apply lag to the writer; when any reader exceeds the configured threshold, the writer worker adds delay at transaction end to reduce write pressure.

### Operation and observation

AWS documents version 1.0. On a supported engine version and region, add the extension to the DB parameter group's preload setting, reboot the instances as required by RDS, create it in the database, and inspect its statistics:

```sql
CREATE EXTENSION flow_control;

SHOW flow_control.target_standby_apply_lag;

SELECT * FROM get_flow_control_stats();

SELECT pid, wait_event_type, wait_event, query
FROM pg_stat_activity
WHERE wait_event_type = 'Extension';
```

The documented default lag threshold is two minutes. Throttled sessions surface an Extension wait event in the activity view and Performance Insights. Parameter availability and change procedures are controlled by the RDS engine version and DB parameter group.

### Scope and tradeoffs

This is not a portable extension for self-managed PostgreSQL, Aurora PostgreSQL, or a generic primary/standby installation. It is part of the RDS Multi-AZ DB cluster service and depends on AWS-managed communication among its instances.

Flow control is most useful for short, highly concurrent OLTP transactions. It is less effective when long batch transactions cause the lag, reduces writer throughput when active, and does not guarantee that lag stays below the target. AWS documents disabling it by removing flow_control from the preload list and rebooting the DB instance; test the write-latency and failover consequences before changing production settings.
