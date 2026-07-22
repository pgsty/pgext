---
title: "pgfr_analyze"
linkTitle: "pgfr_analyze"
description: "Reporting and analysis functions for pgfr_record"
weight: 6061
categories: ["STAT"]
width: full
---

[**pg_flight_recorder**](https://github.com/dventimisupabase/pg_flight_recorder) : Reporting and analysis functions for pgfr_record


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6061** | {{< badge content="pgfr_analyze" link="https://github.com/dventimisupabase/pg_flight_recorder" >}} | {{< ext "pgfr_analyze" "pg_flight_recorder" >}} | `2.29.2` | {{< category "STAT" >}} | {{< license "Apache-2.0" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pgfr_analyze` |
|   **Requires**    | {{< ext "pgfr_record" >}} |
|   **See Also**    | {{< ext "pgfr_record" >}} {{< ext "pg_profile" >}} {{< ext "pg_stat_statements" >}} |
|    **Siblings**   | {{< ext "pgfr_record" >}} |

> [!Note] Secondary extension shipped by pg_flight_recorder; requires pgfr_record.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.29.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "red" >}} | `pg_flight_recorder` | `pgfr_record` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.29.2` | {{< bg "18" "pg_flight_recorder_18" "green" >}} {{< bg "17" "pg_flight_recorder_17" "green" >}} {{< bg "16" "pg_flight_recorder_16" "green" >}} {{< bg "15" "pg_flight_recorder_15" "green" >}} {{< bg "14" "pg_flight_recorder_14" "red" >}} | `pg_flight_recorder_$v` | `pg_cron_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.29.2` | {{< bg "18" "postgresql-18-pg-flight-recorder" "green" >}} {{< bg "17" "postgresql-17-pg-flight-recorder" "green" >}} {{< bg "16" "postgresql-16-pg-flight-recorder" "green" >}} {{< bg "15" "postgresql-15-pg-flight-recorder" "green" >}} {{< bg "14" "postgresql-14-pg-flight-recorder" "red" >}} | `postgresql-$v-pg-flight-recorder` | `postgresql-$v-cron` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_flight_recorder_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_flight_recorder_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_flight_recorder_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_flight_recorder_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_flight_recorder_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "pg_flight_recorder_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_flight_recorder_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-18-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-17-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-16-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-15-pg-flight-recorder : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-flight-recorder : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-18-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-17-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-16-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-15-pg-flight-recorder : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-flight-recorder : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-18-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-17-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-16-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-15-pg-flight-recorder : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-flight-recorder : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-18-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-17-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-16-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-15-pg-flight-recorder : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-flight-recorder : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-18-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-17-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-16-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-15-pg-flight-recorder : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-flight-recorder : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-18-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-17-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-16-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-15-pg-flight-recorder : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-flight-recorder : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-18-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-17-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-16-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-15-pg-flight-recorder : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-flight-recorder : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-18-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-17-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-16-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-15-pg-flight-recorder : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-flight-recorder : MISS 0" "red" >}}      |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-18-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-17-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-16-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-15-pg-flight-recorder : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-flight-recorder : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-18-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-17-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-16-pg-flight-recorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.29.2" "postgresql-15-pg-flight-recorder : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-flight-recorder : MISS 0" "red" >}}      |


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/dventimisupabase/pg_flight_recorder" title="Repository" icon="github" subtitle="github.com/dventimisupabase/pg_flight_recorder" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_flight_recorder-2.29.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_flight_recorder;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_flight_recorder;		# install via package name, for the active PG version
pig install pgfr_analyze;		# install by extension name, for the current active PG version

pig install pgfr_analyze -v 18;   # install for PG 18
pig install pgfr_analyze -v 17;   # install for PG 17
pig install pgfr_analyze -v 16;   # install for PG 16
pig install pgfr_analyze -v 15;   # install for PG 15

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgfr_analyze CASCADE; -- requires pgfr_record
```

## Usage

Sources:

- [pgfr_analyze v2.29.2 README](https://github.com/dventimisupabase/pg_flight_recorder/blob/v2.29.2/pgfr_analyze/README.md)
- [pgfr_analyze control file](https://github.com/dventimisupabase/pg_flight_recorder/blob/v2.29.2/pgfr_analyze/extension.control)
- [pg_flight_recorder v2.29.2 reference](https://github.com/dventimisupabase/pg_flight_recorder/blob/v2.29.2/REFERENCE.md)
- [v2.29.2 release notes](https://github.com/dventimisupabase/pg_flight_recorder/releases/tag/v2.29.2)

pgfr_analyze is the read-oriented analysis half of PostgreSQL Flight Recorder. It interprets history captured by pgfr_record to compare periods, summarize waits, assemble incident timelines, detect regressions, and estimate blast radius. Install it after pgfr_record and use it for diagnosis rather than collection.

### Install the Analysis Layer

    CREATE EXTENSION pg_cron;
    CREATE EXTENSION pgfr_record;
    CREATE EXTENSION pgfr_analyze;
    SELECT pgfr_record.enable();

pgfr_analyze depends on recorder objects and has no useful history until pgfr_record has collected samples. It does not require its own background worker.

### Start an Incident Investigation

Summarize what happened around a known timestamp:

    SELECT *
    FROM pgfr_analyze.what_happened_at(
      now() - interval '15 minutes'
    );

Build an ordered incident view over a period:

    SELECT *
    FROM pgfr_analyze.incident_timeline(
      now() - interval '30 minutes',
      now()
    );

Compare a suspect period with a baseline using compare, then narrow the result with wait_summary or anomaly_report. Always inspect the installed function signatures with psql or pg_get_function_arguments because optional parameters can change between releases.

### Analysis Index

- compare: contrast metrics across baseline and incident windows.
- wait_summary: aggregate sampled wait events.
- report and anomaly_report: produce broad or anomaly-focused summaries.
- what_happened_at: inspect observations around one timestamp.
- incident_timeline: order notable events over a range.
- detect_regressions and detect_query_storms: flag worsening behavior or query bursts.
- blast_radius: identify affected sessions or workloads.
- capacity_summary and capacity_report: summarize capacity-related signals.
- configuration analysis: relate relevant setting changes to the period.

### Interpretation Workflow

1. Confirm pgfr_record.health_check() and the available sample interval.
2. Choose explicit baseline and incident windows with comparable workload.
3. Use compare and wait_summary to locate the dominant change.
4. Correlate activity, locks, replication, vacuum, and configuration evidence.
5. Validate hypotheses against PostgreSQL logs, query plans, application telemetry, and host metrics.

### Caveats

- Results are observations and heuristics, not proof of causality. Sparse sampling can miss short events.
- Counter resets, extension upgrades, restarts, retention gaps, and workload seasonality can distort comparisons.
- Some findings are richer when pg_stat_statements is enabled and sufficiently sized.
- Access to analysis output can expose query text and operational metadata; restrict privileges accordingly.
- v2.29.2 mainly improves managed-service scheduling behavior in pgfr_record. It does not replace the need to verify that collection jobs actually ran.
