## Usage

Sources:

- [pgmonitor-extension v2.2.0 README](https://github.com/CrunchyData/pgmonitor-extension/blob/v2.2.0/README.md)
- [pgmonitor v2.2.0 control file](https://github.com/CrunchyData/pgmonitor-extension/blob/v2.2.0/pgmonitor.control)
- [pgmonitor-extension v2.2.0 release notes](https://github.com/CrunchyData/pgmonitor-extension/releases/tag/v2.2.0)

pgmonitor exposes PostgreSQL monitoring metrics through a curated set of views, materialized views, and tables for external collectors. Its SQL metrics work without a background worker; the optional pgmonitor_bgw worker periodically refreshes materialized data.

### Create the Extension

Create a dedicated schema and install pgmonitor there:

    CREATE SCHEMA pgmonitor_ext;
    CREATE EXTENSION pgmonitor SCHEMA pgmonitor_ext;

Grant collectors only the access they need to the metric objects. Some underlying PostgreSQL statistics remain subject to built-in role and row-visibility rules.

### Collect Metrics

External agents can select the active objects described by the extension's configuration tables:

    SELECT *
    FROM pgmonitor_ext.metric_views
    WHERE active;

    SELECT *
    FROM pgmonitor_ext.metric_matviews
    WHERE active;

    SELECT *
    FROM pgmonitor_ext.metric_tables
    WHERE active;

These tables describe metric name, activation, scope, and refresh interval. Query the installed definitions rather than assuming every metric is enabled on every PostgreSQL version.

The metric surface includes activity, database and table statistics, locks, replication, WAL and archive status, vacuum progress, settings, checkpoints, and extension-specific views when their dependencies are available.

### Refresh Materialized Metrics Manually

Without the background worker, invoke the refresh procedure for the configured schema and metric:

    CALL pgmonitor_ext.refresh_metrics(
      'pgmonitor_ext',
      'pg_stat_statements'
    );

Use names returned by metric_matviews; do not assume the example metric is installed or active. The extension retains a legacy refresh function for compatibility, but new integrations should use the documented procedure.

### Optional Background Worker

To schedule refreshes inside PostgreSQL:

    shared_preload_libraries = 'pgmonitor_bgw'
    pgmonitor_bgw.dbname = 'postgres,app'
    pgmonitor_bgw.role = 'postgres'
    pgmonitor_bgw.interval = 30

Restart PostgreSQL after changing shared_preload_libraries. pgmonitor_bgw.dbname is required and lists the databases to maintain. Upstream v2.2 currently requires the worker role to be a superuser; use the narrowest controlled role and protect its credentials and settings.

### Object Index

- metric_views: directly queried metric views and their collection metadata.
- metric_matviews: materialized metrics and refresh intervals.
- metric_tables: table-backed metrics and maintenance metadata.
- refresh_metrics(schema, name): refresh procedure for one configured metric.
- pgmonitor_bgw.dbname: databases processed by the optional worker.
- pgmonitor_bgw.role: role used for refresh work.
- pgmonitor_bgw.interval: worker loop interval in seconds.

### Version 2.2 and Caveats

Version 2.2 removes the settings-checksum metric, fixes the legacy refresh path on PostgreSQL 13, and reduces routine log noise.

- Metric queries add load to shared statistics, catalogs, and extension objects. Set collection intervals from measured cost.
- A healthy collector connection does not prove materialized views are fresh; monitor their timestamps and worker logs.
- The extension supplies database metrics, not host, filesystem, or process metrics.
- Installing pgmonitor does not automatically configure Prometheus, exporters, dashboards, or alert rules.
