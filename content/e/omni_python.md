---
title: "omni_python"
linkTitle: "omni_python"
description: "First-class Python support"
weight: 2960
categories: ["FEAT"]
width: full
---

[**omnigres**](https://docs.omnigres.org/omni_python/intro/) : First-class Python support


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2960** | {{< badge content="omni_python" link="https://docs.omnigres.org/omni_python/intro/" >}} | {{< ext "omni_python" "omnigres" >}} | `0.1.1` | {{< category "FEAT" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `omni_python` |
|   **Requires**    | {{< ext "plpython3u" >}} |
|    **Siblings**   | {{< ext "omni" >}} {{< ext "omni_auth" >}} {{< ext "omni_aws" >}} {{< ext "omni_cloudevents" >}} {{< ext "omni_containers" >}} {{< ext "omni_credentials" >}} {{< ext "omni_email" >}} {{< ext "omni_http" >}} {{< ext "omni_httpc" >}} {{< ext "omni_httpd" >}} {{< ext "omni_id" >}} {{< ext "omni_json" >}} {{< ext "omni_kube" >}} {{< ext "omni_ledger" >}} {{< ext "omni_manifest" >}} {{< ext "omni_mimetypes" >}} {{< ext "omni_os" >}} {{< ext "omni_polyfill" >}} {{< ext "omni_regex" >}} {{< ext "omni_rest" >}} {{< ext "omni_schema" >}} {{< ext "omni_seq" >}} {{< ext "omni_service" >}} {{< ext "omni_session" >}} {{< ext "omni_sql" >}} {{< ext "omni_sqlite" >}} {{< ext "omni_test" >}} {{< ext "omni_txn" >}} {{< ext "omni_types" >}} {{< ext "omni_var" >}} {{< ext "omni_vfs" >}} {{< ext "omni_vfs_types_v1" >}} {{< ext "omni_web" >}} {{< ext "omni_worker" >}} {{< ext "omni_xml" >}} {{< ext "omni_yaml" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `omnigres` | `plpython3u` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.1` | {{< bg "18" "omnigres_18" "green" >}} {{< bg "17" "omnigres_17" "green" >}} {{< bg "16" "omnigres_16" "green" >}} {{< bg "15" "omnigres_15" "green" >}} {{< bg "14" "omnigres_14" "green" >}} | `omnigres_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.1` | {{< bg "18" "postgresql-18-omnigres" "green" >}} {{< bg "17" "postgresql-17-omnigres" "green" >}} {{< bg "16" "postgresql-16-omnigres" "green" >}} {{< bg "15" "postgresql-15-omnigres" "green" >}} {{< bg "14" "postgresql-14-omnigres" "green" >}} | `postgresql-$v-omnigres` | - |


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


## Source

{{< cards cols=3 >}}
{{< card link="https://docs.omnigres.org/omni_python/intro/" title="Repository" icon="link" subtitle="docs.omnigres.org/omni_python/intro/" >}}
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
pig install omni_python;		# install by extension name, for the current active PG version

pig install omni_python -v 18;   # install for PG 18
pig install omni_python -v 17;   # install for PG 17
pig install omni_python -v 16;   # install for PG 16
pig install omni_python -v 15;   # install for PG 15
pig install omni_python -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION omni_python CASCADE; -- requires plpython3u
```




## Usage

> [omni_python: First-class Python support](https://docs.omnigres.org/omni_python/intro/)

The `omni_python` extension integrates Python code as stored procedures in PostgreSQL.

### Basic Python Function

Create a Python file with the `@pg` decorator:

```python
# hello.py
from omni_python import pg

@pg
def hello() -> str:
    return "Hey there!"

@pg
def add(a: int, b: int) -> int:
    return a + b
```

### Loading Python Files

```sql
CREATE EXTENSION omni_python CASCADE;
CREATE EXTENSION omni_schema CASCADE;
CREATE EXTENSION omni_vfs CASCADE;

CREATE OR REPLACE FUNCTION demo_fs() RETURNS omni_vfs.local_fs
    LANGUAGE sql AS $$ SELECT omni_vfs.local_fs('/python-files') $$;

SELECT omni_schema.load_from_fs(demo_fs());
```

### Flask Integration

```python
from omni_python import pg
from omni_http import omni_httpd
from omni_http.omni_httpd import flask
from flask import Flask, jsonify, request

app = Flask('myapp')

@app.route('/employees', methods=['GET'])
def get_employees():
    employees = plpy.execute(plpy.prepare("select * from employees"))
    return json.dumps([dict(e) for e in employees])

handle = pg(flask.Adapter(app))
```

Configure the HTTP handler:

```sql
UPDATE omni_httpd.handlers SET query = $$select handle(request.*) from request$$;
```
