---
title: Package
description: Extension Packages and Alias
icon: Package
---


Mange extensions and packages are not that simple, here are two common extension examples:

![](/img/pigsty/ext-mapping.png)

| Entity        | Example `pgvector`                       | Example `postgis`...                                       |
|---------------|------------------------------------------|------------------------------------------------------------|
| **Extension** | `vector`                                 | `postgis`, `postgis_topology`, `postgis_raster`,...        |
| **Package**   | `pgvector`                               | `postgis`                                                  |
| **OS PKG**    | pgvector_17                              | postgresql-16-postgis-3                                    |
| **RPM/DEB**   | pgvector_17_0.8.0-1PGDG.rhel8.x86_64.rpm | postgresql-17-postgis-3_3.5.2+dfsg-1.pgdg22.04+1_amd64.deb |

To install the right RPM / DEB with minimal effort, we need to use the abstract layer: **package alias**.
So you can install these extensions by specifying the "Normalized" names, like `pgvector` or `postgis`.
Without knowing any details about PG & OS version, Arch, Extension versions, and any other details.

Package alias `pkg` are used for extension download & install, but you'll have to use the extension name `ext` when `CREATE EXTENSION` in the database (like the `vector` in `meta` database).
And beware some extensions require explicit preloading, like the `timescaledb` in the above example.

Besides, all the extensions are categorized into 16 major categories, we also have alias for the entire extension category
so that you can download and install them in batch, such as:

```yaml title="replace 17 with 16,15,14,13,..."
repo_extra_packages: [ pg17-main ,pg17-core ,pg17-time ,pg17-gis ,pg17-rag ,pg17-fts ,pg17-olap ,pg17-feat ,pg17-lang ,pg17-type ,pg17-util ,pg17-func ,pg17-admin ,pg17-stat ,pg17-sec ,pg17-fdw ,pg17-sim ,pg17-etl]
pg_extensions: [pg17-time ,pg17-gis ,pg17-rag ,pg17-fts ,pg17-feat ,pg17-lang ,pg17-type ,pg17-util ,pg17-func ,pg17-admin ,pg17-stat ,pg17-sec ,pg17-fdw ,pg17-sim ,pg17-etl ] #,pg17-olap]
```

All extensions **CAN** be installed simultaneously, except the `olap` category, where `citus` conflict with `hydra`, and `pg_duckdb` conflict with `pg_mooncake`.
So you can download them all, but install one at a time.



