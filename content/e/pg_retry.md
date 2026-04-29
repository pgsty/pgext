---
title: "pg_retry"
linkTitle: "pg_retry"
description: "Retry SQL statements on transient errors with exponential backoff"
weight: 4100
categories: ["UTIL"]
width: full
---

[**pg_retry**](https://github.com/Agent-Hellboy/pg_retry) : Retry SQL statements on transient errors with exponential backoff


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4100** | {{< badge content="pg_retry" link="https://github.com/Agent-Hellboy/pg_retry" >}} | {{< ext "pg_retry" >}} | `1.0.0` | {{< category "UTIL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_retry` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "pg_retry_18" "green" >}} {{< bg "17" "pg_retry_17" "green" >}} {{< bg "16" "pg_retry_16" "red" >}} {{< bg "15" "pg_retry_15" "red" >}} {{< bg "14" "pg_retry_14" "red" >}} | `pg_retry_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "postgresql-18-retry" "green" >}} {{< bg "17" "postgresql-17-retry" "green" >}} {{< bg "16" "postgresql-16-retry" "red" >}} {{< bg "15" "postgresql-15-retry" "red" >}} {{< bg "14" "postgresql-14-retry" "red" >}} | `postgresql-$v-retry` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_retry_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_retry_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_retry_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_retry_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_retry_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_retry_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_retry_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_retry_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_retry_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_retry_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_retry_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_retry_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_retry_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_retry_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_retry_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_retry_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_retry_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_retry_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_retry_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_retry_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_retry_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_retry_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_retry_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_retry_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_retry_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_retry_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_retry_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_retry_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_retry_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_retry_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-retry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-retry : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-retry : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-retry : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-retry : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-retry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-retry : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-retry : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-retry : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-retry : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-retry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-retry : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-retry : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-retry : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-retry : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-retry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-retry : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-retry : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-retry : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-retry : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-retry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-retry : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-retry : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-retry : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-retry : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-retry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-retry : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-retry : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-retry : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-retry : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-retry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-retry : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-retry : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-retry : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-retry : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-retry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-retry : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-retry : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-retry : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-retry : MISS 0" "red" >}}      |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-retry : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-retry : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-retry : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-retry : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-retry : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-retry : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-retry : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-retry : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-retry : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-retry : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_retry_18` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 17.4 KiB | [pg_retry_18-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_retry_18-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_retry_18` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.4 KiB | [pg_retry_18-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_retry_18-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_retry_18` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.6 KiB | [pg_retry_18-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_retry_18-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_retry_18` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.4 KiB | [pg_retry_18-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_retry_18-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_retry_18` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.6 KiB | [pg_retry_18-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_retry_18-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_retry_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.6 KiB | [pg_retry_18-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_retry_18-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-retry` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 19.6 KiB | [postgresql-18-retry_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-retry/postgresql-18-retry_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-retry` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 19.6 KiB | [postgresql-18-retry_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-retry/postgresql-18-retry_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-retry` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 19.6 KiB | [postgresql-18-retry_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-retry/postgresql-18-retry_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-retry` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.6 KiB | [postgresql-18-retry_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-retry/postgresql-18-retry_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-retry` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 20.5 KiB | [postgresql-18-retry_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-retry/postgresql-18-retry_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-retry` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 20.5 KiB | [postgresql-18-retry_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-retry/postgresql-18-retry_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-retry` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 20.5 KiB | [postgresql-18-retry_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-retry/postgresql-18-retry_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-retry` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 20.2 KiB | [postgresql-18-retry_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-retry/postgresql-18-retry_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_retry_17` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 17.4 KiB | [pg_retry_17-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_retry_17-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_retry_17` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.4 KiB | [pg_retry_17-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_retry_17-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_retry_17` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.6 KiB | [pg_retry_17-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_retry_17-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_retry_17` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.5 KiB | [pg_retry_17-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_retry_17-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_retry_17` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.6 KiB | [pg_retry_17-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_retry_17-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_retry_17` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.6 KiB | [pg_retry_17-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_retry_17-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-retry` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 19.6 KiB | [postgresql-17-retry_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-retry/postgresql-17-retry_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-retry` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 19.5 KiB | [postgresql-17-retry_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-retry/postgresql-17-retry_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-retry` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 19.5 KiB | [postgresql-17-retry_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-retry/postgresql-17-retry_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-retry` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.5 KiB | [postgresql-17-retry_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-retry/postgresql-17-retry_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-retry` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 21.3 KiB | [postgresql-17-retry_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-retry/postgresql-17-retry_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-retry` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 21.2 KiB | [postgresql-17-retry_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-retry/postgresql-17-retry_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-retry` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 20.5 KiB | [postgresql-17-retry_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-retry/postgresql-17-retry_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-retry` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 20.3 KiB | [postgresql-17-retry_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-retry/postgresql-17-retry_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Agent-Hellboy/pg_retry" title="Repository" icon="github" subtitle="github.com/Agent-Hellboy/pg_retry" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_retry-1.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_retry;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_retry;		# install via package name, for the active PG version

pig install pg_retry -v 18;   # install for PG 18
pig install pg_retry -v 17;   # install for PG 17

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_retry;
```




## Usage

> [pg_retry: Retry SQL statements on transient errors with exponential backoff](https://github.com/Agent-Hellboy/pg_retry)

### Function Signature

```sql
retry.retry(
  sql TEXT,                          -- SQL statement to run (exactly one)
  max_tries INT DEFAULT 3,           -- total attempts (1 + retries), >= 1
  base_delay_ms INT DEFAULT 50,      -- initial backoff delay in ms
  max_delay_ms INT DEFAULT 1000,     -- cap for exponential backoff
  retry_sqlstates TEXT[] DEFAULT ARRAY['40001','40P01','55P03','57014']
) RETURNS INT                       -- number of rows processed
```

Default retryable SQLSTATEs: `40001` (serialization_failure), `40P01` (deadlock_detected), `55P03` (lock_not_available), `57014` (query_canceled).

### Examples

Basic retry with defaults:

```sql
SELECT retry.retry('UPDATE accounts SET balance = balance - 100 WHERE id = 1');
```

Custom retry parameters:

```sql
SELECT retry.retry(
    'INSERT INTO audit_log (event) VALUES (''test'')',
    5,      -- max_tries
    100,    -- base_delay_ms
    5000    -- max_delay_ms
);
```

### GUC Configuration

```sql
ALTER SYSTEM SET pg_retry.default_max_tries = 5;
ALTER SYSTEM SET pg_retry.default_base_delay_ms = 100;
ALTER SYSTEM SET pg_retry.default_max_delay_ms = 5000;
ALTER SYSTEM SET pg_retry.default_sqlstates = '40001,40P01,55P03,57014';
SELECT pg_reload_conf();
```

### Safety Rules

- Only one SQL statement per call (multi-statement fails)
- Transaction control statements (BEGIN, COMMIT, ROLLBACK) are prohibited
- Parameters are validated (max_tries >= 1, non-negative delays, base <= max delay)
