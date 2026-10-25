## Usage

Sources:

- [Official source](https://github.com/jobin2005/postgresql-apt-behaviour-analysis/blob/bc6e7eacefefd2ff2859bf19b28d9d01663016d3/README.md)
- [Official source](https://github.com/jobin2005/postgresql-apt-behaviour-analysis/blob/bc6e7eacefefd2ff2859bf19b28d9d01663016d3/src/apt_guard.control)
- [Official source](https://github.com/jobin2005/postgresql-apt-behaviour-analysis/blob/bc6e7eacefefd2ff2859bf19b28d9d01663016d3/src/apt_guard--1.0.sql)
- [Official source](https://github.com/jobin2005/postgresql-apt-behaviour-analysis/blob/bc6e7eacefefd2ff2859bf19b28d9d01663016d3/Dockerfile.postgres)

`apt_guard` 1.0 records database activity for a research behavioral-security pipeline. The upstream container uses PostgreSQL 15. The native module collects events; external Python and machine-learning services provide the surrounding analysis and mitigation workflow. This is not an independently validated database firewall.

### Collect Events

Configure the shared library and restart PostgreSQL:

```ini
shared_preload_libraries = 'apt_guard'
```

Install as a superuser in the monitored database and inspect collected events:

```sql
CREATE EXTENSION apt_guard;
SELECT * FROM apt_events ORDER BY event_time DESC LIMIT 10;
SELECT * FROM apt_alerts ORDER BY id DESC LIMIT 10;
```

The extension installs tables for users, sessions, events, sequence patterns and alerts. Executor and utility hooks collect command text, timing, outcomes and relation information. The external agent and dashboard require their own deployment; creating the extension does not start them.

### Boundaries and Maintenance

Query text can contain sensitive values. Restrict access to collection tables and to the external services, configure retention, and measure collection overhead before using a busy database. Internal collection-table queries are deliberately skipped, so recorded events are not a complete independent audit trail. Errors and transaction behavior can affect persistence; test the failure paths relevant to the workload.

The container example's credentials are development examples, not production defaults to adopt. The reviewed repository has no license file; no reuse license is assumed. Upgrade the library and SQL objects together, and restart backends that already loaded the old module.
