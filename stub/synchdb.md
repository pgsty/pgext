## Usage

Sources:

- [Official documentation](https://docs.synchdb.com)
- [Pinned upstream README](https://github.com/Hornetlabs/synchdb/blob/48d5d11e95d1c9711827ed9bade8255123434c87/README.md)
- [Version 1.0 SQL definitions](https://github.com/Hornetlabs/synchdb/blob/48d5d11e95d1c9711827ed9bade8255123434c87/synchdb--1.0.sql)
- [Extension control file](https://github.com/Hornetlabs/synchdb/blob/48d5d11e95d1c9711827ed9bade8255123434c87/synchdb.control)

`synchdb` is a mixed C and Java PostgreSQL extension for initial snapshot plus change-data-capture replication from MySQL, SQL Server, Oracle, and OpenLog Replicator sources into PostgreSQL. It embeds Debezium connectors through the JVM and applies events inside PostgreSQL; it is not a bidirectional or PostgreSQL-to-PostgreSQL replication system.

### Core Workflow

Install the extension, register a named connector, start its background worker, and inspect state. Upstream uses the literal text `'null'` for unspecified table and snapshot fields because the registration function is strict.

```sql
CREATE EXTENSION synchdb CASCADE;

SELECT synchdb_add_conninfo(
    'mysqlconn',
    '127.0.0.1',
    3306,
    'cdc_user',
    'source-secret',
    'inventory',
    'postgres',
    'inventory.orders,inventory.customers',
    'null',
    'mysql'
);

SELECT synchdb_start_engine_bgw('mysqlconn');
SELECT * FROM synchdb_state_view;
```

Stop a connector before changing its definition or performing maintenance:

```sql
SELECT synchdb_stop_engine_bgw('mysqlconn');
```

### Management and Observability

- `synchdb_add_conninfo(...)` stores a connector in `synchdb_conninfo`; `synchdb_del_conninfo(name)` removes it.
- `synchdb_start_engine_bgw(name)` starts snapshot plus streaming, while its two-argument overload selects a snapshot mode. `synchdb_pause_engine(name)`, `synchdb_resume_engine(name)`, and `synchdb_stop_engine_bgw(name)` control execution.
- `synchdb_state_view` exposes stage, state, process ID, last error, and source offset.
- `synchdb_genstats`, `synchdb_snapstats`, and `synchdb_cdcstats` expose general, snapshot, and change-event counters; `synchdb_reset_stats(name)` resets a connector's counters.
- Object-map, datatype-translation, JMX, OpenLog Replicator, and offset functions support advanced connector configuration; follow the version-matched official guide rather than guessing their argument semantics.

### Prerequisites and Operations

Version `1.0` requires `pgcrypto`, a JVM/JNI runtime, and the Java components installed alongside the C library. The pinned README lists Java 17 or later, PostgreSQL 16 through 18, and a Unix-like operating system; Oracle FDW snapshots and OpenLog Replicator builds add their own native dependencies.

Stored connector passwords are encrypted, but the initial secret still crosses the SQL statement and may reach client history, statement logs, audit logs, or monitoring. Use a tightly controlled administration session and least-privileged source accounts. Initial snapshots can be large, and DDL/type conversion, offsets, source retention, restarts, duplicate delivery, schema drift, and apply errors determine correctness. Test a representative source, monitor `synchdb_state_view` and the stats views, verify row counts and keys, and rehearse pause, resume, rebuild, and recovery before production use.
