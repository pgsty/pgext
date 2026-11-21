---
title: "plluau"
linkTitle: "plluau"
description: "Lua as an untrusted procedural language"
weight: 3030
categories: ["LANG"]
languages: ["C"]
licenses: ["MIT"]
repos: ["PIGSTY"]
page_width: full
---

[**pllua**](https://github.com/pllua/pllua) : Lua as an untrusted procedural language


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3030** | {{< badge content="plluau" link="https://github.com/pllua/pllua" >}} | {{< ext "plluau" "pllua" >}} | `2.0` | {{< category "LANG" >}} | {{< license "MIT" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pg_catalog` |
|    **Need By**    | {{< ext "hstore_plluau" >}} |
|   **See Also**    | {{< ext "plperlu" >}} {{< ext "plpython3u" >}} {{< ext "pltclu" >}} {{< ext "plperl" >}} {{< ext "plv8" >}} {{< ext "pljs" >}} {{< ext "pljava" >}} {{< ext "plsh" >}} {{< ext "plr" >}} {{< ext "plxslt" >}} |
|    **Siblings**   | {{< ext "pllua" >}} {{< ext "hstore_pllua" >}} {{< ext "hstore_plluau" >}} |

> [!Note] pkg 2.0.12; SQL 2.0.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0` | {{< bg "19" "" "red" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} | `pllua` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.12` | {{< bg "19" "pllua_19" "red" >}} {{< bg "18" "pllua_18" "green" >}} {{< bg "17" "pllua_17" "green" >}} {{< bg "16" "pllua_16" "green" >}} {{< bg "15" "pllua_15" "green" >}} | `pllua_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.0.12` | {{< bg "19" "postgresql-19-pllua" "red" >}} {{< bg "18" "postgresql-18-pllua" "green" >}} {{< bg "17" "postgresql-17-pllua" "green" >}} {{< bg "16" "postgresql-16-pllua" "green" >}} {{< bg "15" "postgresql-15-pllua" "green" >}} | `postgresql-$v-pllua` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "pllua_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.0.12" "pllua_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 2.0.12" "pllua_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_15 : AVAIL 4" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "pllua_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.0.12" "pllua_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 2.0.12" "pllua_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_15 : AVAIL 2" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "N/A" "pllua_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.0.12" "pllua_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 2.0.12" "pllua_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_15 : AVAIL 2" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "N/A" "pllua_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.0.12" "pllua_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 2.0.12" "pllua_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_15 : AVAIL 2" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "N/A" "pllua_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.0.12" "pllua_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 2.0.12" "pllua_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_15 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "N/A" "pllua_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.0.12" "pllua_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 2.0.12" "pllua_17 : AVAIL 5" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_15 : AVAIL 5" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "N/A" "postgresql-19-pllua : N/A 0" "gray" >}} | {{< bg "PGDG 2.0.12" "postgresql-18-pllua : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-17-pllua : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-16-pllua : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-15-pllua : AVAIL 2" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "N/A" "postgresql-19-pllua : N/A 0" "gray" >}} | {{< bg "PGDG 2.0.12" "postgresql-18-pllua : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-17-pllua : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-16-pllua : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-15-pllua : AVAIL 2" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "N/A" "postgresql-19-pllua : N/A 0" "gray" >}} | {{< bg "PGDG 2.0.12" "postgresql-18-pllua : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-17-pllua : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-16-pllua : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-15-pllua : AVAIL 2" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "N/A" "postgresql-19-pllua : N/A 0" "gray" >}} | {{< bg "PGDG 2.0.12" "postgresql-18-pllua : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-17-pllua : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-16-pllua : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-15-pllua : AVAIL 2" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "N/A" "postgresql-19-pllua : N/A 0" "gray" >}} | {{< bg "PGDG 2.0.12" "postgresql-18-pllua : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-17-pllua : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-16-pllua : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-15-pllua : AVAIL 2" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "N/A" "postgresql-19-pllua : N/A 0" "gray" >}} | {{< bg "PGDG 2.0.12" "postgresql-18-pllua : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-17-pllua : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-16-pllua : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-15-pllua : AVAIL 2" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "N/A" "postgresql-19-pllua : N/A 0" "gray" >}} | {{< bg "PGDG 2.0.12" "postgresql-18-pllua : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-17-pllua : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-16-pllua : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-15-pllua : AVAIL 2" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "N/A" "postgresql-19-pllua : N/A 0" "gray" >}} | {{< bg "PGDG 2.0.12" "postgresql-18-pllua : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-17-pllua : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-16-pllua : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-15-pllua : AVAIL 2" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "N/A" "postgresql-19-pllua : N/A 0" "gray" >}} | {{< bg "PGDG 2.0.12" "postgresql-18-pllua : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-17-pllua : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-16-pllua : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-15-pllua : AVAIL 2" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "N/A" "postgresql-19-pllua : N/A 0" "gray" >}} | {{< bg "PGDG 2.0.12" "postgresql-18-pllua : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-17-pllua : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-16-pllua : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-15-pllua : AVAIL 2" "blue" >}} |
{.matrix .wide .expand}


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pllua/pllua" title="Repository" icon="github" subtitle="github.com/pllua/pllua" />}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pllua;		# install via package name, for the active PG version
pig install plluau;		# install by extension name, for the current active PG version

pig install plluau -v 18;   # install for PG 18
pig install plluau -v 17;   # install for PG 17
pig install plluau -v 16;   # install for PG 16
pig install plluau -v 15;   # install for PG 15

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION plluau;
```




## Usage

> [plluau: Lua as an untrusted procedural language](https://github.com/pllua/pllua)

`plluau` is the untrusted variant of `pllua`, allowing Lua functions to access the filesystem, load arbitrary modules, and perform operations restricted in the trusted version.

```sql
CREATE EXTENSION plluau;
```

### Create Functions

```sql
CREATE FUNCTION read_file(path text) RETURNS text LANGUAGE plluau AS $$
  local f = io.open(path, "r")
  if f then
    local content = f:read("*a")
    f:close()
    return content
  end
  return nil
$$;
```

### Differences from pllua (Trusted)

| Feature | pllua (trusted) | plluau (untrusted) |
|---------|----------------|--------------------|
| File I/O | Restricted | Full access |
| Module loading | Whitelisted only | Unrestricted |
| OS access | No | Yes |
| Suitable for | User-defined functions | Admin/superuser functions |

### Same API as pllua

`plluau` shares the same SPI interface, trigger support, set-returning functions, and data type handling as `pllua`. All SPI functions (`spi.execute`, `spi.prepare`, `spi.rows`), coroutine-based set returns, and trigger functions work identically.

```sql
CREATE FUNCTION run_command(cmd text) RETURNS text LANGUAGE plluau AS $$
  local handle = io.popen(cmd)
  local result = handle:read("*a")
  handle:close()
  return result
$$;
```

### Initialization

Configure via GUC (superuser only):

```sql
SET pllua.on_untrusted_init = 'myvar = {}';
```

Only superusers can create `plluau` functions due to the unrestricted access it provides.
