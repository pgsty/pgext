---
title: "pgcrypto"
linkTitle: "pgcrypto"
description: "cryptographic functions"
weight: 7980
categories: ["SEC"]
width: full
---

[**pgcrypto**](https://www.postgresql.org/docs/current/pgcrypto.html) : cryptographic functions


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7980** | {{< badge content="pgcrypto" link="https://www.postgresql.org/docs/current/pgcrypto.html" >}} | {{< ext "pgcrypto" >}} | `1.3` | {{< category "SEC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "omni_auth" >}} {{< ext "omni_aws" >}} {{< ext "omni_credentials" >}} {{< ext "omni_rest" >}} {{< ext "pgcryptokey" >}} {{< ext "pgjwt" >}} |
|   **See Also**    | {{< ext "pgsodium" >}} {{< ext "pgsmcrypto" >}} {{< ext "lo" >}} {{< ext "anon" >}} {{< ext "pg_tde" >}} {{< ext "sslutils" >}} {{< ext "faker" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** | **PG13** |
|:--------:|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.3" "PostgreSQL 18: version 1.3" "green" >}} | {{< bg "1.3" "PostgreSQL 17: version 1.3" "green" >}} | {{< bg "1.3" "PostgreSQL 16: version 1.3" "green" >}} | {{< bg "1.3" "PostgreSQL 15: version 1.3" "green" >}} | {{< bg "1.3" "PostgreSQL 14: version 1.3" "green" >}} | {{< bg "1.3" "PostgreSQL 13: version 1.3" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgcrypto;
```
