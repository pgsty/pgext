---
title: "jsonb_plruby"
linkTitle: "jsonb_plruby"
description: "Transform between jsonb and native Ruby data for PL/Ruby"
weight: 3161
categories: ["LANG"]
width: full
---

[**plruby**](https://github.com/commandprompt/plruby) : Transform between jsonb and native Ruby data for PL/Ruby


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3161** | {{< badge content="jsonb_plruby" link="https://github.com/commandprompt/plruby" >}} | {{< ext "jsonb_plruby" "plruby" >}} | `1.0` | {{< category "LANG" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "plruby" >}} |
|   **See Also**    | {{< ext "plruby" >}} {{< ext "hstore_plruby" >}} {{< ext "ltree_plruby" >}} {{< ext "jsonb_plperl" >}} {{< ext "jsonb_plpython3u" >}} |
|    **Siblings**   | {{< ext "plruby" >}} {{< ext "hstore_plruby" >}} {{< ext "ltree_plruby" >}} |

> [!Note] Extension control default_version is 1.0; shipped in the PL/Ruby 2.5.0 package.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `plruby` | `plruby` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.5.0` | {{< bg "18" "plruby_18" "green" >}} {{< bg "17" "plruby_17" "green" >}} {{< bg "16" "plruby_16" "green" >}} {{< bg "15" "plruby_15" "green" >}} {{< bg "14" "plruby_14" "green" >}} | `plruby_$v` | `ruby-libs` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.5.0` | {{< bg "18" "postgresql-18-plruby" "green" >}} {{< bg "17" "postgresql-17-plruby" "green" >}} {{< bg "16" "postgresql-16-plruby" "green" >}} {{< bg "15" "postgresql-15-plruby" "green" >}} {{< bg "14" "postgresql-14-plruby" "green" >}} | `postgresql-$v-plruby` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.5.0" "plruby_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.5.0" "plruby_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.5.0" "plruby_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.5.0" "plruby_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.5.0" "plruby_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.5.0" "plruby_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-plruby : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-plruby : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-plruby : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-plruby : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-plruby : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-plruby : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-plruby : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-plruby : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-plruby : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-plruby : AVAIL 1" "green" >}} |


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/commandprompt/plruby" title="Repository" icon="github" subtitle="github.com/commandprompt/plruby" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="plruby-2.5.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg plruby;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install plruby;		# install via package name, for the active PG version
pig install jsonb_plruby;		# install by extension name, for the current active PG version

pig install jsonb_plruby -v 18;   # install for PG 18
pig install jsonb_plruby -v 17;   # install for PG 17
pig install jsonb_plruby -v 16;   # install for PG 16
pig install jsonb_plruby -v 15;   # install for PG 15
pig install jsonb_plruby -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION jsonb_plruby CASCADE; -- requires plruby
```

## Usage

Sources:

- [PL/Ruby v2.5.0 README](https://github.com/commandprompt/plruby/blob/v2.5.0/README.md)
- [jsonb_plruby v1.0 control file](https://github.com/commandprompt/plruby/blob/v2.5.0/jsonb_plruby/jsonb_plruby.control)
- [jsonb_plruby v1.0 extension SQL](https://github.com/commandprompt/plruby/blob/v2.5.0/jsonb_plruby/jsonb_plruby--1.0.sql)

`jsonb_plruby` installs a PostgreSQL transform between `jsonb` and native Ruby values for the `plruby` language. A transformed `jsonb` argument becomes a Ruby `Hash`, `Array`, `String`, `Integer`, `Float`, `true`, `false`, or `nil`; compatible Ruby values can be returned directly as `jsonb`.

### Install and Use the Transform

```sql
CREATE EXTENSION plruby;
CREATE EXTENSION jsonb_plruby;

CREATE FUNCTION ruby_mark_processed(jsonb)
RETURNS jsonb
LANGUAGE plruby
TRANSFORM FOR TYPE jsonb
AS $$
  value = args[0]
  value['processed'] = true
  value
$$;

SELECT ruby_mark_processed('{"id": 42}'::jsonb);
```

The transform is used only by functions that declare `TRANSFORM FOR TYPE jsonb`. Other PL/Ruby functions keep the language's ordinary JSONB conversion behavior.

### Objects and Caveats

- `jsonb_to_plruby(internal)` implements SQL-to-Ruby conversion.
- `plruby_to_jsonb(internal)` implements Ruby-to-SQL conversion.
- The extension version is `1.0`, it requires `plruby`, and it is relocatable.
- Ruby `Hash` keys returned to PostgreSQL must be valid JSON object keys, and numeric/special values must be representable by PostgreSQL `jsonb`. Test nested values and numeric limits explicitly.
- PL/Ruby remains untrusted. Installing this transform does not sandbox Ruby code or reduce the privileges required to create PL/Ruby functions.
