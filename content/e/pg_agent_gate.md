---
title: "pg_agent_gate"
linkTitle: "pg_agent_gate"
description: "SQL proposal, verification and execution gate with human approval for agent roles."
weight: 7470
categories: ["SEC"]
languages: ["Rust"]
licenses: ["Apache-2.0"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_agent_gate**](https://github.com/Manuelreyesbravo/pg_agent_gate) : SQL proposal, verification and execution gate with human approval for agent roles.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7470** | {{< badge content="pg_agent_gate" link="https://github.com/Manuelreyesbravo/pg_agent_gate" >}} | {{< ext "pg_agent_gate" >}} | `0.3.0` | {{< category "SEC" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `agent_gate` `agent_gate_internal` |
|   **See Also**    | {{< ext "pg_living_assertions" >}} {{< ext "pg_policy" >}} |

> [!Note] Upstream PG18-19; packages PG18. Preload before the first agent statement; READ COMMITTED required. Human approval preserves commit guards; pg_living_assertions optional. Build pgrx 0.19.3, upstream 0.19.2. Trusted code and external effects remain outside the sandbox.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "red" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_agent_gate` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "pg_agent_gate_18" "green" >}} {{< bg "17" "pg_agent_gate_17" "red" >}} {{< bg "16" "pg_agent_gate_16" "red" >}} {{< bg "15" "pg_agent_gate_15" "red" >}} {{< bg "14" "pg_agent_gate_14" "red" >}} | `pg_agent_gate_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "postgresql-18-pg-agent-gate" "green" >}} {{< bg "17" "postgresql-17-pg-agent-gate" "red" >}} {{< bg "16" "postgresql-16-pg-agent-gate" "red" >}} {{< bg "15" "postgresql-15-pg-agent-gate" "red" >}} {{< bg "14" "postgresql-14-pg-agent-gate" "red" >}} | `postgresql-$v-pg-agent-gate` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "pg_agent_gate_18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_agent_gate_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_agent_gate_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_agent_gate_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_agent_gate_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "pg_agent_gate_18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_agent_gate_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_agent_gate_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_agent_gate_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_agent_gate_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "pg_agent_gate_18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_agent_gate_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_agent_gate_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_agent_gate_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_agent_gate_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "pg_agent_gate_18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_agent_gate_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_agent_gate_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_agent_gate_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_agent_gate_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "pg_agent_gate_18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_agent_gate_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_agent_gate_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_agent_gate_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_agent_gate_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "pg_agent_gate_18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_agent_gate_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_agent_gate_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_agent_gate_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_agent_gate_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-agent-gate : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-agent-gate : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-agent-gate : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-agent-gate : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-agent-gate : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-agent-gate : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-agent-gate : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-agent-gate : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-agent-gate : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-agent-gate : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-agent-gate : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-agent-gate : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-agent-gate : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-agent-gate : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-agent-gate : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-agent-gate : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-agent-gate : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-agent-gate : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-agent-gate : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-agent-gate : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-agent-gate : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-agent-gate : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-agent-gate : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-agent-gate : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-agent-gate : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-agent-gate : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-agent-gate : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-agent-gate : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-agent-gate : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-agent-gate : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-agent-gate : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-agent-gate : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-agent-gate : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-agent-gate : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-agent-gate : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-agent-gate : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-agent-gate : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-agent-gate : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-agent-gate : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-agent-gate : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-agent-gate : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-agent-gate : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-agent-gate : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-agent-gate : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-agent-gate : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-agent-gate : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-agent-gate : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-agent-gate : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-agent-gate : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-agent-gate : N/A 0" "gray" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_agent_gate_18` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 524.4 KiB | [pg_agent_gate_18-0.3.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_agent_gate_18-0.3.0-1PGSTY.el8.x86_64.rpm) |
| `pg_agent_gate_18` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 457.4 KiB | [pg_agent_gate_18-0.3.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_agent_gate_18-0.3.0-1PGSTY.el8.aarch64.rpm) |
| `pg_agent_gate_18` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 526.7 KiB | [pg_agent_gate_18-0.3.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_agent_gate_18-0.3.0-1PGSTY.el9.x86_64.rpm) |
| `pg_agent_gate_18` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 485.6 KiB | [pg_agent_gate_18-0.3.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_agent_gate_18-0.3.0-1PGSTY.el9.aarch64.rpm) |
| `pg_agent_gate_18` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 526.6 KiB | [pg_agent_gate_18-0.3.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_agent_gate_18-0.3.0-1PGSTY.el10.x86_64.rpm) |
| `pg_agent_gate_18` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 462.1 KiB | [pg_agent_gate_18-0.3.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_agent_gate_18-0.3.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-agent-gate` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 461.1 KiB | [postgresql-18-pg-agent-gate_0.3.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-agent-gate/postgresql-18-pg-agent-gate_0.3.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-agent-gate` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 397.0 KiB | [postgresql-18-pg-agent-gate_0.3.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-agent-gate/postgresql-18-pg-agent-gate_0.3.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-agent-gate` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 460.9 KiB | [postgresql-18-pg-agent-gate_0.3.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-agent-gate/postgresql-18-pg-agent-gate_0.3.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-agent-gate` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 397.3 KiB | [postgresql-18-pg-agent-gate_0.3.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-agent-gate/postgresql-18-pg-agent-gate_0.3.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-agent-gate` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 500.3 KiB | [postgresql-18-pg-agent-gate_0.3.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-agent-gate/postgresql-18-pg-agent-gate_0.3.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-agent-gate` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 459.3 KiB | [postgresql-18-pg-agent-gate_0.3.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-agent-gate/postgresql-18-pg-agent-gate_0.3.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-agent-gate` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 494.7 KiB | [postgresql-18-pg-agent-gate_0.3.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-agent-gate/postgresql-18-pg-agent-gate_0.3.0-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-agent-gate` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 451.4 KiB | [postgresql-18-pg-agent-gate_0.3.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-agent-gate/postgresql-18-pg-agent-gate_0.3.0-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-agent-gate` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 493.2 KiB | [postgresql-18-pg-agent-gate_0.3.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-agent-gate/postgresql-18-pg-agent-gate_0.3.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-agent-gate` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 451.0 KiB | [postgresql-18-pg-agent-gate_0.3.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-agent-gate/postgresql-18-pg-agent-gate_0.3.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Manuelreyesbravo/pg_agent_gate" title="Repository" icon="github" subtitle="github.com/Manuelreyesbravo/pg_agent_gate" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_agent_gate-0.3.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_agent_gate;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_agent_gate;		# install via package name, for the active PG version

pig install pg_agent_gate -v 18;   # install for PG 18

```


[**Config**](https://pigsty.io/docs/pgsql/ext/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_agent_gate';
```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_agent_gate;
```

## Usage

Sources:

- [README](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/v0.3.0/README.md)
- [Control](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/v0.3.0/pg_agent_gate.control)
- [CHANGELOG](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/v0.3.0/CHANGELOG.md)
- [SQL upgrade](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/v0.3.0/sql/pg_agent_gate--0.2.15--0.3.0.sql)

`pg_agent_gate` 0.3.0 routes registered agent roles through proposals, checks, dry runs and guarded execution on PostgreSQL 18/19, with optional human approval before commit. The library must be loaded before the first agent statement, through shared preload or a role-level session preload followed by reconnection.

### Core Workflow

```sql
CREATE EXTENSION pg_agent_gate;
```

```sql
SELECT agent_gate.whoami();
SELECT agent_gate.discover('customers');
SELECT agent_gate.propose('SELECT 1', 'Check connectivity');
```

### Operational Boundaries

Install as a superuser and register ordinary roles with narrow table privileges, row limits and a DDL allow-list. The API includes `agent_gate.whoami`, `agent_gate.discover`, `agent_gate.propose`, `agent_gate.dry_run`, `agent_gate.commit` and `agent_gate.acts`. Verbs require READ COMMITTED. Proposal-time session settings are checked again before dry-run/commit, and reads that invoke known state-changing built-ins are rejected. Re-registering a name does not inherit old proposals. `agent_gate.max_proposal_bytes` and `agent_gate.max_intent_bytes` bound input sizes. Existing duplicate committed proposals can block the upgrade and require inspection. Optional bound checks use `pg_living_assertions`. Rollback cannot undo every sequence or external effect; existing privileged code must remain trusted. The server log can contain literal SQL. Attempt records participate in the outer transaction, and the default fast durability mode can lose recent records on a crash.

### Human Approval

Set `p_requires_approval` when calling `agent_gate.register_agent` to require consent for that agent's proposals. A member of `agent_gate_approver` records consent with `agent_gate.approve`; the agent then commits its own proposal. Approval executes no SQL and does not bypass row limits, assertions or commit-time checks. The agent cannot approve itself, and `propose_and_commit` cannot bypass the approval requirement.

The fixed `agent_gate_internal` schema stores proposals, executions and approvals; agents do not access those tables directly. The execution record identifies the approver in `approved_by`.

### Library and SQL Upgrades

The library retains the 0.2.15 protections against function-running argument casts and lone-semicolon crashes. Replace the library and restart or reconnect every process that preloads it, then update each database to 0.3.0 using the supplied SQL upgrade. The 0.3.0 migration adds approval objects and is not a no-op. Treat DDL permission as schema power that row-count limits do not bound.
