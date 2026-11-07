---
title: "pg_proctab"
linkTitle: "pg_proctab"
description: "PostgreSQL extension to access the OS process table"
weight: 6460
categories: ["STAT"]
languages: ["SQL"]
licenses: ["Apache-2.0"]
repos: ["PIGSTY"]
page_width: full
---

[**pgnodemx**](https://github.com/CrunchyData/pgnodemx) : PostgreSQL extension to access the OS process table


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6460** | {{< badge content="pg_proctab" link="https://github.com/CrunchyData/pgnodemx" >}} | {{< ext "pg_proctab" "pgnodemx" >}} | `2.0.1` | {{< category "STAT" >}} | {{< license "Apache-2.0" >}} | {{< language "SQL" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "pgnodemx" >}} |
|   **See Also**    | {{< ext "system_stats" >}} {{< ext "pg_stat_kcache" >}} {{< ext "pgmeminfo" >}} |
|    **Siblings**   | {{< ext "pgnodemx" >}} |

> [!Note] Compatibility stub; PGDG RPM incomplete.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgnodemx` | `pgnodemx` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.1` | {{< bg "18" "pgnodemx_18" "green" >}} {{< bg "17" "pgnodemx_17" "green" >}} {{< bg "16" "pgnodemx_16" "green" >}} {{< bg "15" "pgnodemx_15" "green" >}} {{< bg "14" "pgnodemx_14" "green" >}} | `pgnodemx_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.0.1` | {{< bg "18" "postgresql-18-pgnodemx" "green" >}} {{< bg "17" "postgresql-17-pgnodemx" "green" >}} {{< bg "16" "postgresql-16-pgnodemx" "green" >}} {{< bg "15" "postgresql-15-pgnodemx" "green" >}} {{< bg "14" "postgresql-14-pgnodemx" "green" >}} | `postgresql-$v-pgnodemx` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.0.1" "pgnodemx_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.0.1" "pgnodemx_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.0.1" "pgnodemx_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.0.1" "pgnodemx_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.0.1" "pgnodemx_14 : AVAIL 3" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.0.1" "pgnodemx_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.0.1" "pgnodemx_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.0.1" "pgnodemx_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.0.1" "pgnodemx_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.0.1" "pgnodemx_14 : AVAIL 3" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.0.1" "pgnodemx_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.1" "pgnodemx_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.1" "pgnodemx_16 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.1" "pgnodemx_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.1" "pgnodemx_14 : AVAIL 4" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.0.1" "pgnodemx_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.1" "pgnodemx_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.1" "pgnodemx_16 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.1" "pgnodemx_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.1" "pgnodemx_14 : AVAIL 4" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.0.1" "pgnodemx_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.1" "pgnodemx_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.1" "pgnodemx_16 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.1" "pgnodemx_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.1" "pgnodemx_14 : AVAIL 4" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.0.1" "pgnodemx_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.1" "pgnodemx_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.1" "pgnodemx_16 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.1" "pgnodemx_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.1" "pgnodemx_14 : AVAIL 4" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 2.0.1" "postgresql-18-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-17-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-pgnodemx : AVAIL 3" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 2.0.1" "postgresql-18-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-17-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-pgnodemx : AVAIL 3" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 2.0.1" "postgresql-18-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-17-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-pgnodemx : AVAIL 3" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 2.0.1" "postgresql-18-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-17-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-pgnodemx : AVAIL 3" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 2.0.1" "postgresql-18-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-17-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-pgnodemx : AVAIL 3" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 2.0.1" "postgresql-18-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-17-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-pgnodemx : AVAIL 3" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 2.0.1" "postgresql-18-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-17-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-pgnodemx : AVAIL 3" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 2.0.1" "postgresql-18-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-17-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-pgnodemx : AVAIL 3" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 2.0.1" "postgresql-18-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-17-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-pgnodemx : AVAIL 3" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 2.0.1" "postgresql-18-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-17-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-pgnodemx : AVAIL 3" "blue" >}} |
{.matrix .wide .expand}


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/CrunchyData/pgnodemx" title="Repository" icon="github" subtitle="github.com/CrunchyData/pgnodemx" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgnodemx-2.0.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pgnodemx;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pgnodemx;		# install via package name, for the active PG version
pig install pg_proctab;		# install by extension name, for the current active PG version

pig install pg_proctab -v 18;   # install for PG 18
pig install pg_proctab -v 17;   # install for PG 17
pig install pg_proctab -v 16;   # install for PG 16
pig install pg_proctab -v 15;   # install for PG 15
pig install pg_proctab -v 14;   # install for PG 14

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_proctab CASCADE; -- requires pgnodemx
```

## Usage

Sources:

- [pgnodemx 2.0.1 documentation](https://github.com/pgnodemx/pgnodemx/blob/v2.0.1/README.md)
- [Compatibility control file](https://github.com/pgnodemx/pgnodemx/blob/v2.0.1/pg_proctab.control)
- [Compatibility SQL](https://github.com/pgnodemx/pgnodemx/blob/v2.0.1/pg_proctab--0.0.10-compat.sql)

The `pg_proctab` shipped in `pgnodemx` provides five SQL wrappers for the legacy process-monitoring interface. It depends on `pgnodemx`; the package version is 2.0.1, while its SQL installation version is `0.0.10-compat`.

### Enable

Add `pgnodemx` to `shared_preload_libraries`, preserving other entries, then restart PostgreSQL:

```ini
shared_preload_libraries = 'pgnodemx'
```

Install both extensions in the same schema as a superuser:

```sql
CREATE EXTENSION pgnodemx;
CREATE EXTENSION pg_proctab VERSION '0.0.10-compat';
```

The explicit version is required because the main control file has no default version. The audited PGDG RPM packages omit that main control file; use a package containing it before running this command. Querying the underlying OS metrics requires superuser privileges or membership in `pg_monitor`.

### Query

```sql
SELECT a.pid, a.query, p.utime, p.stime, p.vsize, p.rss
FROM pg_stat_activity a
JOIN pg_proctab() p ON a.pid = p.pid;

SELECT "user", nice, system, idle, iowait FROM pg_cputime();
SELECT load1, load5, load15, last_pid FROM pg_loadavg();
SELECT memused, memfree, memshared, membuffers, memcached,
       swapused, swapfree, swapcached
FROM pg_memusage();
SELECT devname, reads_completed, writes_completed,
       sectors_read, sectors_written
FROM pg_diskusage();
```

These functions read Linux process and system statistics visible to the PostgreSQL server. Container namespaces and OS permissions affect what is visible. In this compatibility layer, `rss` is reported in KiB, memory totals in MiB, and `memshared` comes from `Shmem`.
