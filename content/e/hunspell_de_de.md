---
title: "hunspell_de_de"
linkTitle: "hunspell_de_de"
description: "German Hunspell Dictionary"
weight: 2271
categories: ["FTS"]
width: full
---

[**hunspell**](https://github.com/postgrespro/hunspell_dicts) : German Hunspell Dictionary


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2271** | {{< badge content="hunspell_de_de" link="https://github.com/postgrespro/hunspell_dicts" >}} | {{< ext "hunspell_de_de" "hunspell" >}} | `1.0` | {{< category "FTS" >}} | {{< license "PostgreSQL" >}} | {{< language "Data" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "hunspell_cs_cz" >}} {{< ext "hunspell_en_us" >}} {{< ext "hunspell_fr" >}} {{< ext "hunspell_nl_nl" >}} {{< ext "hunspell_ne_np" >}} {{< ext "hunspell_nn_no" >}} {{< ext "hunspell_pt_pt" >}} {{< ext "hunspell_ru_ru" >}} |
|    **Siblings**   | {{< ext "hunspell_cs_cz" >}} {{< ext "hunspell_en_us" >}} {{< ext "hunspell_fr" >}} {{< ext "hunspell_ne_np" >}} {{< ext "hunspell_nl_nl" >}} {{< ext "hunspell_nn_no" >}} {{< ext "hunspell_pt_pt" >}} {{< ext "hunspell_ru_ru" >}} {{< ext "hunspell_ru_ru_aot" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `hunspell` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "hunspell_18" "green" >}} {{< bg "17" "hunspell_17" "green" >}} {{< bg "16" "hunspell_16" "green" >}} {{< bg "15" "hunspell_15" "green" >}} {{< bg "14" "hunspell_14" "green" >}} | `hunspell_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "postgresql-18-hunspell" "green" >}} {{< bg "17" "postgresql-17-hunspell" "green" >}} {{< bg "16" "postgresql-16-hunspell" "green" >}} {{< bg "15" "postgresql-15-hunspell" "green" >}} {{< bg "14" "postgresql-14-hunspell" "green" >}} | `postgresql-$v-hunspell` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0" "hunspell_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "hunspell_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "hunspell_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "hunspell_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "hunspell_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0" "hunspell_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "hunspell_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "hunspell_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "hunspell_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "hunspell_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0" "hunspell_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "hunspell_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "hunspell_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "hunspell_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "hunspell_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0" "hunspell_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "hunspell_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "hunspell_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "hunspell_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "hunspell_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0" "hunspell_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "hunspell_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "hunspell_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "hunspell_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "hunspell_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0" "hunspell_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "hunspell_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "hunspell_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "hunspell_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "hunspell_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-hunspell : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-hunspell : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-hunspell : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-hunspell : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-hunspell : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-hunspell : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-hunspell : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-hunspell : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-hunspell : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-hunspell : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-hunspell : AVAIL 1" "green" >}} |


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/postgrespro/hunspell_dicts" title="Repository" icon="github" subtitle="github.com/postgrespro/hunspell_dicts" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="hunspell-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg hunspell;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install hunspell;		# install via package name, for the active PG version
pig install hunspell_de_de;		# install by extension name, for the current active PG version

pig install hunspell_de_de -v 18;   # install for PG 18
pig install hunspell_de_de -v 17;   # install for PG 17
pig install hunspell_de_de -v 16;   # install for PG 16
pig install hunspell_de_de -v 15;   # install for PG 15
pig install hunspell_de_de -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION hunspell_de_de;
```




## Usage

> [hunspell_de_de: German Hunspell dictionary for PostgreSQL](https://github.com/postgrespro/hunspell_dicts)

German Hunspell dictionary and text search configuration for PostgreSQL full-text search.

```sql
CREATE EXTENSION hunspell_de_de;

SELECT ts_lexize('german_hunspell', 'Geschichten');

SELECT to_tsvector('german_hunspell', 'Geschichten');
```

This extension provides the `german_hunspell` dictionary and text search configuration.
