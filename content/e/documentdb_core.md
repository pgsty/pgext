---
title: "documentdb_core"
linkTitle: "documentdb_core"
description: "Core API surface for DocumentDB on PostgreSQL"
weight: 9010
categories: ["SIM"]
width: full
---

[**documentdb**](https://github.com/documentdb/documentdb) : Core API surface for DocumentDB on PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9010** | {{< badge content="documentdb_core" link="https://github.com/documentdb/documentdb" >}} | {{< ext "documentdb_core" "documentdb" >}} | `0.110` | {{< category "SIM" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "documentdb" >}} {{< ext "documentdb_distributed" >}} |
|   **See Also**    | {{< ext "mongo_fdw" >}} {{< ext "rum" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} {{< ext "pg_cron" >}} {{< ext "postgis" >}} {{< ext "vector" >}} |
|    **Siblings**   | {{< ext "documentdb" >}} {{< ext "documentdb_distributed" >}} {{< ext "documentdb_extended_rum" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.110` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "red" >}} | `documentdb` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.110` | {{< bg "18" "documentdb_18" "green" >}} {{< bg "17" "documentdb_17" "green" >}} {{< bg "16" "documentdb_16" "green" >}} {{< bg "15" "documentdb_15" "green" >}} {{< bg "14" "documentdb_14" "red" >}} | `documentdb_$v` | `postgresql$v-contrib`, `pg_cron_$v`, `pgvector_$v`, `rum_$v`, `postgis36_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.110` | {{< bg "18" "postgresql-18-documentdb" "green" >}} {{< bg "17" "postgresql-17-documentdb" "green" >}} {{< bg "16" "postgresql-16-documentdb" "green" >}} {{< bg "15" "postgresql-15-documentdb" "green" >}} {{< bg "14" "postgresql-14-documentdb" "red" >}} | `postgresql-$v-documentdb` | `postgresql-$v-cron`, `postgresql-$v-pgvector`, `postgresql-$v-rum`, `postgresql-$v-postgis-3` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.110" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.110" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.110" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.110" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.110" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.110" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.110" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.110" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.110" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.110" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.110" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.110" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.110" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.110" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.110" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.110" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.110" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.110" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.110" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.110" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.110" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.110" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.110" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.110" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.110" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.110" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.110" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.110" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.110" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.110" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.110" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.110" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.110" "postgresql-18-documentdb : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.110" "postgresql-17-documentdb : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.110" "postgresql-16-documentdb : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.110" "postgresql-15-documentdb : AVAIL 2" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.110" "postgresql-18-documentdb : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.110" "postgresql-17-documentdb : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.110" "postgresql-16-documentdb : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.110" "postgresql-15-documentdb : AVAIL 2" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.110" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.110" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.110" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.110" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.110" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.110" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.110" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.110" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.110" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.110" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.110" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.110" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.110" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.110" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.110" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.110" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.110" "postgresql-18-documentdb : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.110" "postgresql-17-documentdb : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.110" "postgresql-16-documentdb : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.110" "postgresql-15-documentdb : AVAIL 2" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.110" "postgresql-18-documentdb : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.110" "postgresql-17-documentdb : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.110" "postgresql-16-documentdb : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.110" "postgresql-15-documentdb : AVAIL 2" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/documentdb/documentdb" title="Repository" icon="github" subtitle="github.com/documentdb/documentdb" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="documentdb-0.110-0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg documentdb;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install documentdb;		# install via package name, for the active PG version
pig install documentdb_core;		# install by extension name, for the current active PG version

pig install documentdb_core -v 18;   # install for PG 18
pig install documentdb_core -v 17;   # install for PG 17
pig install documentdb_core -v 16;   # install for PG 16
pig install documentdb_core -v 15;   # install for PG 15

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_documentdb, pg_documentdb_core';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION documentdb_core;
```



## Usage

> [documentdb_core: Core API surface for DocumentDB for PostgreSQL](https://github.com/documentdb/documentdb)

DocumentDB provides MongoDB-compatible document database functionality built on PostgreSQL. The `documentdb_core` extension introduces BSON datatype support and operations for native Postgres.

### BSON Data Type

The extension adds a native BSON (Binary JSON) data type to PostgreSQL, enabling storage and manipulation of MongoDB-style documents.

### Basic Document Operations

Documents are managed through MongoDB-compatible CRUD operations via the DocumentDB API layer:

```python
import pymongo

client = pymongo.MongoClient(
    'mongodb://user:pass@localhost:10260/?tls=true&tlsAllowInvalidCertificates=true'
)

db = client["myDatabase"]
collection = db.create_collection("myCollection")

# Insert documents
collection.insert_one({
    'name': 'John Doe',
    'email': 'john@email.com',
    'address': '123 Main St'
})

collection.insert_many([
    {'name': 'Jane Smith', 'email': 'jane@email.com'},
    {'name': 'Alice Johnson', 'email': 'alice@email.com'}
])

# Query documents
for doc in collection.find():
    print(doc)

single = collection.find_one({'name': 'John Doe'})
```

### Aggregation Pipelines

```python
pipeline = [
    {'$match': {'name': 'Alice Johnson'}},
    {'$project': {'_id': 0, 'name': 1, 'email': 1}}
]

results = collection.aggregate(pipeline)
for doc in results:
    print(doc)
```

### Components

- **documentdb_core**: BSON datatype support and operations for native Postgres
- **documentdb (pg_documentdb)**: Public API surface providing CRUD functionality
- **pg_documentdb_gw**: Gateway protocol translation layer (MongoDB wire protocol to PostgreSQL)

The extension supports full-text search, geospatial queries, and vector search on BSON documents.
