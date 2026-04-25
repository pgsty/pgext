---
title: "omni_vfs"
linkTitle: "omni_vfs"
description: "Virtual File System"
weight: 2974
categories: ["FEAT"]
width: full
---

[**omnigres**](https://docs.omnigres.org/omni_vfs/reference/) : Virtual File System


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2974** | {{< badge content="omni_vfs" link="https://docs.omnigres.org/omni_vfs/reference/" >}} | {{< ext "omni_vfs" "omnigres" >}} | `0.2.2` | {{< category "FEAT" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `omni_vfs` |
|   **Requires**    | {{< ext "omni_vfs_types_v1" >}} {{< ext "dblink" >}} |
|    **Need By**    | {{< ext "omni_schema" >}} |
|    **Siblings**   | {{< ext "omni" >}} {{< ext "omni_auth" >}} {{< ext "omni_aws" >}} {{< ext "omni_cloudevents" >}} {{< ext "omni_containers" >}} {{< ext "omni_credentials" >}} {{< ext "omni_email" >}} {{< ext "omni_http" >}} {{< ext "omni_httpc" >}} {{< ext "omni_httpd" >}} {{< ext "omni_id" >}} {{< ext "omni_json" >}} {{< ext "omni_kube" >}} {{< ext "omni_ledger" >}} {{< ext "omni_manifest" >}} {{< ext "omni_mimetypes" >}} {{< ext "omni_os" >}} {{< ext "omni_polyfill" >}} {{< ext "omni_python" >}} {{< ext "omni_regex" >}} {{< ext "omni_rest" >}} {{< ext "omni_schema" >}} {{< ext "omni_seq" >}} {{< ext "omni_service" >}} {{< ext "omni_session" >}} {{< ext "omni_sql" >}} {{< ext "omni_sqlite" >}} {{< ext "omni_test" >}} {{< ext "omni_txn" >}} {{< ext "omni_types" >}} {{< ext "omni_var" >}} {{< ext "omni_vfs_types_v1" >}} {{< ext "omni_web" >}} {{< ext "omni_worker" >}} {{< ext "omni_xml" >}} {{< ext "omni_yaml" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `omnigres` | `omni_vfs_types_v1`, `dblink` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.2` | {{< bg "18" "omnigres_18" "green" >}} {{< bg "17" "omnigres_17" "green" >}} {{< bg "16" "omnigres_16" "green" >}} {{< bg "15" "omnigres_15" "green" >}} {{< bg "14" "omnigres_14" "green" >}} | `omnigres_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.2` | {{< bg "18" "postgresql-18-omnigres" "green" >}} {{< bg "17" "postgresql-17-omnigres" "green" >}} {{< bg "16" "postgresql-16-omnigres" "green" >}} {{< bg "15" "postgresql-15-omnigres" "green" >}} {{< bg "14" "postgresql-14-omnigres" "green" >}} | `postgresql-$v-omnigres` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "omnigres_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 20250120" "omnigres_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "omnigres_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "omnigres_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "omnigres_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "omnigres_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 20250120" "omnigres_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "omnigres_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "omnigres_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "omnigres_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "omnigres_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 20250507" "omnigres_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250507" "omnigres_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250507" "omnigres_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250507" "omnigres_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "omnigres_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 20250507" "omnigres_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250507" "omnigres_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250507" "omnigres_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250507" "omnigres_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "omnigres_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "omnigres_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "omnigres_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "omnigres_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "omnigres_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "omnigres_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "omnigres_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "omnigres_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "omnigres_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "omnigres_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "postgresql-18-omnigres : MISS 0" "red" >}}      | {{< bg "PIGSTY 20250120" "postgresql-17-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-16-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-15-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-14-omnigres : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "postgresql-18-omnigres : MISS 0" "red" >}}      | {{< bg "PIGSTY 20250120" "postgresql-17-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-16-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-15-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-14-omnigres : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 20251108" "postgresql-18-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-17-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-16-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-15-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-14-omnigres : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 20251108" "postgresql-18-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-17-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-16-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-15-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-14-omnigres : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "postgresql-18-omnigres : MISS 0" "red" >}}      | {{< bg "PIGSTY 20250120" "postgresql-17-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-16-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-15-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-14-omnigres : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "postgresql-18-omnigres : MISS 0" "red" >}}      | {{< bg "PIGSTY 20250120" "postgresql-17-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-16-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-15-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-14-omnigres : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 20251108" "postgresql-18-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-17-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-16-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-15-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-14-omnigres : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 20251108" "postgresql-18-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-17-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-16-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-15-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-14-omnigres : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-omnigres : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-omnigres : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-omnigres : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-omnigres : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-omnigres : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-omnigres : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-omnigres : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-omnigres : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-omnigres : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-omnigres : MISS 0" "red" >}}      |


## Source

{{< cards cols=3 >}}
{{< card link="https://docs.omnigres.org/omni_vfs/reference/" title="Repository" icon="link" subtitle="docs.omnigres.org/omni_vfs/reference/" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="omnigres-20251108.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg omnigres;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install omnigres;		# install via package name, for the active PG version
pig install omni_vfs;		# install by extension name, for the current active PG version

pig install omni_vfs -v 18;   # install for PG 18
pig install omni_vfs -v 17;   # install for PG 17
pig install omni_vfs -v 16;   # install for PG 16
pig install omni_vfs -v 15;   # install for PG 15
pig install omni_vfs -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION omni_vfs CASCADE; -- requires omni_vfs_types_v1, dblink
```




## Usage

> [omni_vfs: Virtual File System](https://docs.omnigres.org/omni_vfs/reference/)

The `omni_vfs` extension provides a unified Virtual File System API for PostgreSQL with multiple backend support.

### Core Functions

**`omni_vfs.list(fs, path)`** -- Lists directory contents, returns set of `(name, kind)`.

**`omni_vfs.list_recursively(fs, path, max)`** -- Recursive listing. Use `max` to limit results.

**`omni_vfs.file_info(fs, path)`** -- Returns file metadata: `size`, `created_at`, `accessed_at`, `modified_at`, `kind`.

**`omni_vfs.read(fs, path, file_offset, chunk_size)`** -- Reads file content as `bytea`.

**`omni_vfs.write(fs, path, content, create_file, append)`** -- Writes content, returns bytes written.

### Backends

**Local filesystem:**

```sql
SELECT omni_vfs.local_fs('/mount/path');
-- No access outside this directory is permitted
```

**Database-backed (table_fs):**

```sql
SELECT omni_vfs.table_fs('filesystem_name');
```

**Remote filesystem:**

```sql
SELECT omni_vfs.remote_fs('dbname=otherdb host=127.0.0.1', $$omni_vfs.local_fs('/path')$$);
```

### Example

```sql
WITH fs AS (SELECT omni_vfs.local_fs('/data') AS fs),
     entries AS (SELECT fs, (omni_vfs.list(fs, '')).* FROM fs)
SELECT entries.name, entries.kind,
       (omni_vfs.file_info(fs, '/' || entries.name)).*
FROM entries;
```
