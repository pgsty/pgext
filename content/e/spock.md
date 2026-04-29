---
title: "spock"
linkTitle: "spock"
description: "Multi-master logical replication extension for PostgreSQL"
weight: 9570
categories: ["ETL"]
width: full
---

[**spock**](https://github.com/pgEdge/spock) : Multi-master logical replication extension for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9570** | {{< badge content="spock" link="https://github.com/pgEdge/spock" >}} | {{< ext "spock" >}} | `5.0.5` | {{< category "ETL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `spock` |
|   **See Also**    | {{< ext "lolor" >}} {{< ext "snowflake" >}} |

> [!Note] works on pgedge kernel fork


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `5.0.5` | {{< bg "18" "" "red" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `spock` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `5.0.5` | {{< bg "18" "spock_18" "red" >}} {{< bg "17" "spock_17" "green" >}} {{< bg "16" "spock_16" "red" >}} {{< bg "15" "spock_15" "red" >}} {{< bg "14" "spock_14" "red" >}} | `spock_$v` | `pgedge_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `5.0.5` | {{< bg "18" "pgedge-18-spock" "red" >}} {{< bg "17" "pgedge-17-spock" "green" >}} {{< bg "16" "pgedge-16-spock" "red" >}} {{< bg "15" "pgedge-15-spock" "red" >}} {{< bg "14" "pgedge-14-spock" "red" >}} | `pgedge-$v-spock` | `pgedge-$v` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "spock_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.0.5" "spock_17 : FORK 1" >}}      |      {{< bg "MISS" "spock_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "spock_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "spock_14 : FORK 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "spock_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.0.5" "spock_17 : FORK 1" >}}      |      {{< bg "MISS" "spock_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "spock_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "spock_14 : FORK 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "spock_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.0.5" "spock_17 : FORK 1" >}}      |      {{< bg "MISS" "spock_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "spock_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "spock_14 : FORK 0" "red" >}}      |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "spock_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.0.5" "spock_17 : FORK 1" >}}      |      {{< bg "MISS" "spock_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "spock_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "spock_14 : FORK 0" "red" >}}      |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "spock_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.0.5" "spock_17 : FORK 1" >}}      |      {{< bg "MISS" "spock_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "spock_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "spock_14 : FORK 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "spock_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.0.5" "spock_17 : FORK 1" >}}      |      {{< bg "MISS" "spock_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "spock_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "spock_14 : FORK 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "pgedge-18-spock : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.0.5" "pgedge-17-spock : FORK 1" >}}      |      {{< bg "MISS" "pgedge-16-spock : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-15-spock : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-14-spock : FORK 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "pgedge-18-spock : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.0.5" "pgedge-17-spock : FORK 1" >}}      |      {{< bg "MISS" "pgedge-16-spock : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-15-spock : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-14-spock : FORK 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "pgedge-18-spock : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.0.5" "pgedge-17-spock : FORK 1" >}}      |      {{< bg "MISS" "pgedge-16-spock : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-15-spock : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-14-spock : FORK 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "pgedge-18-spock : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.0.5" "pgedge-17-spock : FORK 1" >}}      |      {{< bg "MISS" "pgedge-16-spock : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-15-spock : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-14-spock : FORK 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "pgedge-18-spock : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.0.5" "pgedge-17-spock : FORK 1" >}}      |      {{< bg "MISS" "pgedge-16-spock : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-15-spock : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-14-spock : FORK 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "pgedge-18-spock : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.0.5" "pgedge-17-spock : FORK 1" >}}      |      {{< bg "MISS" "pgedge-16-spock : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-15-spock : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-14-spock : FORK 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "pgedge-18-spock : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.0.5" "pgedge-17-spock : FORK 1" >}}      |      {{< bg "MISS" "pgedge-16-spock : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-15-spock : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-14-spock : FORK 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "pgedge-18-spock : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.0.5" "pgedge-17-spock : FORK 1" >}}      |      {{< bg "MISS" "pgedge-16-spock : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-15-spock : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-14-spock : FORK 0" "red" >}}      |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "pgedge-18-spock : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-17-spock : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-16-spock : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-15-spock : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-14-spock : FORK 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "pgedge-18-spock : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-17-spock : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-16-spock : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-15-spock : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-14-spock : FORK 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `spock_17` | `5.0.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 195.5 KiB | [spock_17-5.0.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/spock_17-5.0.5-1PIGSTY.el8.x86_64.rpm) |
| `spock_17` | `5.0.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 185.9 KiB | [spock_17-5.0.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/spock_17-5.0.5-1PIGSTY.el8.aarch64.rpm) |
| `spock_17` | `5.0.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 183.2 KiB | [spock_17-5.0.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/spock_17-5.0.5-1PIGSTY.el9.x86_64.rpm) |
| `spock_17` | `5.0.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 179.2 KiB | [spock_17-5.0.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/spock_17-5.0.5-1PIGSTY.el9.aarch64.rpm) |
| `spock_17` | `5.0.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 185.0 KiB | [spock_17-5.0.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/spock_17-5.0.5-1PIGSTY.el10.x86_64.rpm) |
| `spock_17` | `5.0.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 180.9 KiB | [spock_17-5.0.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/spock_17-5.0.5-1PIGSTY.el10.aarch64.rpm) |
| `pgedge-17-spock` | `5.0.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 166.6 KiB | [pgedge-17-spock_5.0.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/spock/pgedge-17-spock_5.0.5-1PIGSTY~bookworm_amd64.deb) |
| `pgedge-17-spock` | `5.0.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 154.0 KiB | [pgedge-17-spock_5.0.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/spock/pgedge-17-spock_5.0.5-1PIGSTY~bookworm_arm64.deb) |
| `pgedge-17-spock` | `5.0.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 166.8 KiB | [pgedge-17-spock_5.0.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/spock/pgedge-17-spock_5.0.5-1PIGSTY~trixie_amd64.deb) |
| `pgedge-17-spock` | `5.0.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 155.0 KiB | [pgedge-17-spock_5.0.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/spock/pgedge-17-spock_5.0.5-1PIGSTY~trixie_arm64.deb) |
| `pgedge-17-spock` | `5.0.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 177.5 KiB | [pgedge-17-spock_5.0.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/spock/pgedge-17-spock_5.0.5-1PIGSTY~jammy_amd64.deb) |
| `pgedge-17-spock` | `5.0.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 172.0 KiB | [pgedge-17-spock_5.0.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/spock/pgedge-17-spock_5.0.5-1PIGSTY~jammy_arm64.deb) |
| `pgedge-17-spock` | `5.0.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 173.6 KiB | [pgedge-17-spock_5.0.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/spock/pgedge-17-spock_5.0.5-1PIGSTY~noble_amd64.deb) |
| `pgedge-17-spock` | `5.0.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 167.8 KiB | [pgedge-17-spock_5.0.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/spock/pgedge-17-spock_5.0.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgEdge/spock" title="Repository" icon="github" subtitle="github.com/pgEdge/spock" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="spock-5.0.5.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg spock;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install spock;		# install via package name, for the active PG version

pig install spock -v 17;   # install for PG 17

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'spock';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION spock;
```



## Usage

> [spock: Multi-master logical replication extension for PostgreSQL](https://github.com/pgEdge/spock)

Multi-master logical replication for PostgreSQL 15+. Each node acts as both publisher and subscriber.

### Configuration

In `postgresql.conf`:

```ini
wal_level = 'logical'
max_worker_processes = 10
max_replication_slots = 10
max_wal_senders = 10
shared_preload_libraries = 'spock'
track_commit_timestamp = on
spock.enable_ddl_replication = on
spock.include_ddl_repset = on
```

### Enabling

```sql
CREATE EXTENSION spock;
```

### Creating Nodes

On each node, create a node identity:

```sql
-- Node 1
SELECT spock.node_create(
    node_name := 'n1',
    dsn := 'host=10.0.0.5 port=5432 dbname=mydb'
);

-- Node 2
SELECT spock.node_create(
    node_name := 'n2',
    dsn := 'host=10.0.0.7 port=5432 dbname=mydb'
);
```

### Creating Subscriptions

For multi-master, each node subscribes to every other node:

```sql
-- On n1: subscribe to n2
SELECT spock.sub_create(
    subscription_name := 'sub_n1n2',
    provider_dsn := 'host=10.0.0.7 port=5432 dbname=mydb'
);

-- On n2: subscribe to n1
SELECT spock.sub_create(
    subscription_name := 'sub_n2n1',
    provider_dsn := 'host=10.0.0.5 port=5432 dbname=mydb'
);
```

### Replication Set Management

```sql
-- Add table to replication
SELECT spock.repset_add_table('default', 'my_table');

-- Remove table from replication
SELECT spock.repset_remove_table('default', 'my_table');

-- Add all tables in a schema
SELECT spock.repset_add_all_tables('default', '{public}');
```

### Key Features

- Multi-master (active-active) replication
- Automatic DDL replication
- Conflict detection and resolution using commit timestamps
- Row and column filtering
- Supports PostgreSQL 15, 16, 17, and 18
- Tables must have primary keys and matching schemas across nodes
