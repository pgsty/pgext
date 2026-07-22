## Usage

Sources:

- [Official README](https://github.com/asotolongo/resources/blob/911dae0439eeed46760e2ef294047b079a509fb4/README)
- [Extension SQL](https://github.com/asotolongo/resources/blob/911dae0439eeed46760e2ef294047b079a509fb4/resources--0.1.0.sql)
- [Extension control file](https://github.com/asotolongo/resources/blob/911dae0439eeed46760e2ef294047b079a509fb4/resources.control)

resources is a Linux-specific set of `file_fdw` program-backed foreign tables and views for reading CPU, memory, operating-system, filesystem, and PostgreSQL-process information. Each scan runs a fixed shell pipeline on the database server. Upstream warns that the existing version contains bugs.

### Core Workflow

Install only on a compatible Linux host as a superuser, then query the precreated objects:

```sql
CREATE EXTENSION resources CASCADE;

SELECT * FROM resources.cpu_loadavg;
SELECT * FROM resources.meminfo;
SELECT * FROM resources.partitions;
SELECT * FROM resources.v_ps_postgresv2;
```

No application data is stored in these foreign tables; results are parsed from commands such as `cat`, `awk`, `top`, `lscpu`, `df`, and `ps` at query time.

### Important Objects

- `resources.cpu_loadavg` and `resources.cpu_use` expose load averages and CPU-state output.
- `resources.distribution` and `resources.lscpu` expose kernel and CPU information.
- `resources.meminfo` converts Linux memory information to gigabytes.
- `resources.partitions` parses filesystem utilization.
- `resources.ps_postgres` and `resources.ps_postgresv2` parse processes owned by the hard-coded operating-system user postgres.
- `resources.v_ps_postgres` and `resources.v_ps_postgresv2` join process output with PostgreSQL activity data.

### Security and Portability Notes

Program-backed `file_fdw` tables execute operating-system commands with the PostgreSQL service account and therefore require highly privileged setup and access. Revoke foreign-table and server access from untrusted roles. Output formats vary by Linux distribution, locale, command version, cgroup/container environment, and service-account name, so parsing can fail or silently mislabel columns. Some commands sleep or scan process/filesystem state on every query. Do not expose these views as a portable monitoring API; prefer a maintained host exporter outside PostgreSQL.
