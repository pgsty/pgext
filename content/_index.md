---
title: "PostgreSQL Extension Cloud"
cascade:
  type: docs
breadcrumbs: false
excludeSearch: true
width: full
toc: false
comment: false
---


Harness the synergistic superpower of the PostgreSQL extensions ecosystem with three things:

- [**Catalog**](/list) : Find the extension you need, with unparalleled [**431**](/e/) extensions included
- [**Repository**](/repo) : Get pre-built RPM/DEB packages on **14** mainstream [**Linux Distributions**](/os)
- [**PIG Package Manager**](/pig/) : The Missing Package Manager for PostgreSQL & Extension Ecosystem

{{< cards cols=4 >}}  
  {{< card link="/list" title="Extension Catalog" icon="clipboard-list" subtitle="The complete list of 430+ available PostgreSQL extensions" >}}
  {{< card link="/repo" title="Software Repository" icon="cube" subtitle="The APT / DNF repo that deliver PostgreSQL extensions in native Linux format" >}}
  {{< card link="/pig" title="Package Manager" icon="cash" subtitle="The missing package manager for PostgreSQL & Extensions Ecosystem" >}}
{{< /cards >}}


```bash
curl -fsSL https://repo.pigsty.io/pig | bash  # install pig cli
pig repo set                  # setup upstream repository on your linux
pig install pg18              # install PostgreSQL 18 kernel pkg from PGDG
pig install pg_duckdb -v 18   # e.g. install pg_duckdb extension for PG 18
```


**Just use PostgreSQL for everything!** Also check our blog post: [***PostgreSQL is eating the Database World***](https://medium.com/@fengruohang/postgres-is-eating-the-database-world-157c204dcfc4)

[![ecosystem](/ecosystem.gif)](https://medium.com/@fengruohang/postgres-is-eating-the-database-world-157c204dcfc4)

## Catalog

| {{< category "time" >}} | {{< category "gis" >}}  | {{< category "rag" >}}   | {{< category "fts" >}}  | {{< category "olap" >}} | {{< category "feat" >}} | {{< category "lang" >}} | {{< category "type" >}} |
|-------------------------|-------------------------|--------------------------|-------------------------|-------------------------|-------------------------|-------------------------|-------------------------| 
| {{< category "util" >}} | {{< category "func" >}} | {{< category "admin" >}} | {{< category "stat" >}} | {{< category "sec" >}}  | {{< category "fdw" >}}  | {{< category "sim" >}}  | {{< category "etl" >}}  |

{{< cards cols=5 >}}
  {{< card link="/list/cate#time" title="TIME" icon="clock" subtitle="TimescaleDB, Versioning & Temporal Table, Crontab, Async & Background Job Scheduler, ..." >}}
  {{< card link="/list/cate#gis" title="GIS" icon="globe" subtitle="GeoSpatial Data Types, Operators, and Indexes, Hexagonal Indexing, OGR Data FDW, GeoIP & MobilityDB, etc..." >}}
  {{< card link="/list/cate#rag" title="RAG" icon="light-bulb" subtitle="Vector Database with IVFFLAT, HNSW, DiskANN Indexes, AI & ML in SQL interface, Similarity Funcs, etc..." >}}
  {{< card link="/list/cate#fts" title="FTS" icon="search" subtitle="ElasticSearch Alternative with BM25, 2-gram/3-gram Fuzzy Search, Zhparser & Hunspell Segregation Dicts, etc..." >}}
  {{< card link="/list/cate#olap" title="OLAP" icon="chart-bar" subtitle="DuckDB Integration with FDW & PG Lakehouse, Access Parquet from File/S3, Sharding with Citus/Partman/PlProxy, ..." >}}
  {{< card link="/list/cate#feat" title="FEAT" icon="sparkles" subtitle="OpenCypher with AGE, GraphQL, JsonSchema, Hints & Hypo Index, HLL, Rum, IVM, ChemRDKit, and Message Queues,..." >}}
  {{< card link="/list/cate#lang" title="LANG" icon="book-open" subtitle="Develop, Test, Package, and Deliver Stored Procedures written in various PL/Languages: Java, Js, Lua, R, Sh, PRQL, ..." >}}
  {{< card link="/list/cate#type" title="TYPE" icon="cube-transparent" subtitle="Dedicate New Data Types Like: prefix, sember, uint, SIUnit, RoaringBitmap, Rational, Sphere, Hash, RRule, and more..." >}}
  {{< card link="/list/cate#util" title="UTIL" icon="cog" subtitle="Utilities such as send http request, perform gzip/zstd compress, send mails, Regex, ICU, encoding, docs, Encryption,..." >}}
  {{< card link="/list/cate#func" title="FUNC" icon="variable" subtitle="Function such as id generator, aggregations, sketches, vector functions, mathematical functions and digest functions..." >}}
  {{< card link="/list/cate#admin" title="ADMIN" icon="office-building" subtitle="Utilities for Bloat Control, DirtyRead, BufferInspect, DDL Generate, ChecksumVerify, Permission, Priority, Catalog,..." >}}
  {{< card link="/list/cate#stat" title="STAT" icon="presentation-chart-line" subtitle="Observability Catalogs, Monitoring Metrics & Views, Statistics, Query Plans, WaitSampling, SlowLogs, and etc..." >}}
  {{< card link="/list/cate#sec" title="SEC" icon="shield-check" subtitle="Auditing Logs, Enforce Passwords, Keep Secrets, TDE, SM Algorithm, Login Hooks, Log Error, Extension White List, ..." >}}
  {{< card link="/list/cate#fdw" title="FDW" icon="document-download" subtitle="Wrappers & Multicorn for FDW Development, Access other DBMS: MySQL, Mongo, SQLite, MSSQL, Oracle, HDFS, DB2,..." >}}
  {{< card link="/list/cate#sim" title="SIM" icon="terminal" subtitle="Protocol Simulation & heterogeneous DBMS Compatibility: Oracle, MSSQL, DB2, MySQL, Memcached, and Babelfish!" >}}
  {{< card link="/list/cate#etl" title="ETL" icon="truck" subtitle="Logical Replication, Decoding, CDC in protobuf/JSON/Mongo format, Copy & Load & Compare Postgres Databases,..." >}}
{{< /cards >}}

| {{< license "MIT" >}}      | {{< license "ISC" >}}        | {{< license "PostgreSQL" >}} | {{< license "BSD 0-Clause" >}} | {{< license "BSD 2-Clause" >}} | {{< license "BSD 3-Clause" >}} |
|:---------------------------|:-----------------------------|:-----------------------------|:-------------------------------|:-------------------------------|:-------------------------------|
| {{< license "Artistic" >}} | {{< license "Apache-2.0" >}} | {{< license "MPL-2.0" >}}    |                                |                                |                                |
| {{< license "GPL-2.0" >}}  | {{< license "GPL-3.0" >}}    | {{< license "LGPL-2.1" >}}   | {{< license "LGPL-3.0" >}}     | {{< license "AGPL-3.0" >}}     | {{< license "Timescale" >}}    |


| {{< language "c" >}}       | {{< language "c++" >}}       | {{< language "rust" >}}      | {{< language "java" >}}        | {{< language "python" >}}      | {{< language "sql" >}}         | {{< language "data" >}} |
|----------------------------|------------------------------|------------------------------|--------------------------------|--------------------------------|--------------------------------|-------------------------|


## Platform

Pigsty provides complete extension support on these [linux distributions](/os) major versions::

| OS                                                                                     | Vendor | Major |  Minor  | Fullname          | PG Major Version                                                                                                                                                                                                                                            |                     Comment                     |
|:---------------------------------------------------------------------------------------|:-------|:-----:|:-------:|:------------------|:------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:-----------------------------------------------:|
| {{< badge content="el8.x86_64"   link="/os/el8.x86_64"  color="green" >}}              | EL     |   8   |  8.10   | RockyLinux 8 x86  | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       | {{< badge content="Near EOL" color="orange" >}} |
| {{< badge content="el8.aarch64"  link="/os/el8.aarch64" color="green" border=false >}} | EL     |   8   |  8.10   | RockyLinux 8 ARM  | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       | {{< badge content="Near EOL" color="orange" >}} |
| {{< badge content="el9.x86_64"   link="/os/el9.x86_64"  color="green" >}}              | EL     |   9   |   9.6   | RockyLinux 9 x86  | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       |    {{< badge content="OK" color="green" >}}     |
| {{< badge content="el9.aarch64"  link="/os/el9.aarch64" color="green" border=false >}} | EL     |   9   |   9.6   | RockyLinux 9 ARM  | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       |    {{< badge content="OK" color="green" >}}     |
| {{< badge content="el10.x86_64"  link="/os/el10.x86_64" color="green" >}}              | EL     |  10   |  10.0   | RockyLinux 10 x86 | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       |    {{< badge content="OK" color="green" >}}     |
| {{< badge content="el10.aarch64" link="/os/el10.aarch64"color="green" border=false >}} | EL     |  10   |  10.0   | RockyLinux 10 ARM | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       |    {{< badge content="OK" color="green" >}}     |
| {{< badge content="d12.x86_64"   link="/os/d12.x86_64"  color="green" >}}              | Debian |  12   |  12.12  | Debian 12 x86     | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       |    {{< badge content="OK" color="green" >}}     |
| {{< badge content="d12.aarch64"  link="/os/d12.aarch64" color="green" border=false >}} | Debian |  12   |  12.12  | Debian 12 ARM     | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       |    {{< badge content="OK" color="green" >}}     |
| {{< badge content="d13.x86_64"   link="/os/d13.x86_64"  color="green" >}}              | Debian |  13   |  13.1   | Debian 13 x86     | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       |    {{< badge content="OK" color="green" >}}     |
| {{< badge content="d13.aarch64"  link="/os/d13.aarch64" color="green" border=false >}} | Debian |  13   |  13.1   | Debian 13 ARM     | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       |    {{< badge content="OK" color="green" >}}     |
| {{< badge content="u22.x86_64"   link="/os/u22.x86_64"  color="green" >}}              | Ubuntu |  22   | 22.04.5 | Ubuntu 22.04 x86  | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       |    {{< badge content="OK" color="green" >}}     |
| {{< badge content="u22.aarch64"  link="/os/u22.aarch64" color="green" border=false >}} | Ubuntu |  22   | 22.04.5 | Ubuntu 22.04 ARM  | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       |    {{< badge content="OK" color="green" >}}     |
| {{< badge content="u24.x86_64"   link="/os/u24.x86_64"  color="green" >}}              | Ubuntu |  24   | 24.04.3 | Ubuntu 24.04 x86  | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       |    {{< badge content="OK" color="green" >}}     |
| {{< badge content="u24.aarch64"  link="/os/u24.aarch64" color="green" border=false >}} | Ubuntu |  24   | 24.04.3 | Ubuntu 24.04 ARM  | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       |    {{< badge content="OK" color="green" >}}     |


## Users

The **PGSTY.CLOUD** is used by some PostgreSQL Distribution Maker and Vendors:

{{< cards cols=2 >}}
{{< card link="https://github.com/github.com/pgsty/pigsty"     title="Pigsty"     icon="github" subtitle="The author and maintainer, Battery-Included PG RDS" >}}
{{< card link="https://docs.omnigres.org/quick_start/"         title="Omnigres"   icon="github" subtitle="The All-in-One PostgreSQL as a Platform" >}}
{{< card link="https://autobase.tech/docs/extensions/install"  title="AutoBase"   icon="github" subtitle="Self-hosted DBaaS / Postgres automation" >}}
{{< /cards >}}

The repository is hosted on Cloudflare CDN, It is a **free** service to the community.


## About

The PGEXT.CLOUD is maintained by [**PGSTY**](https://github.com/pgsty) / [**VONNG**](https://blog.vonng.com/en/) (rh@vonng.com), and open-sourced under the [**Apache License 2.0**](https://github.com/pgsty/pig/?tab=Apache-2.0-1-ov-file#readme).

| GitHub Repo                                                      | Description                                |
|------------------------------------------------------------------|--------------------------------------------|
| [github.com/pgsty](https://github.com/pgsty)                     | The **PGSTY** Organization                 |
| [github.com/pgsty/pgext](https://github.com/pgsty/pgext)         | This website and extension Data            |
| [github.com/pgsty/pigsty](https://github.com/pgsty/pigsty)       | PostgreSQL Database Distribution           |
| [github.com/pgsty/pig](https://github.com/pgsty/pig)             | PostgreSQL Package Manager                 |
| [github.com/pgsty/ext](https://github.com/pgsty/ext)             | This documentation site, Extension Catalog |
| [github.com/pgsty/rpm](https://github.com/pgsty/rpm)             | RPM Building Specs                         |
| [github.com/pgsty/deb](https://github.com/pgsty/deb)             | DEB Building Specs                         |
| [github.com/pgsty/infra-pkg](https://github.com/pgsty/infra-pkg) | Infra Package Building Specs               |

