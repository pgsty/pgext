---
title: "pgcrypto"
linkTitle: "pgcrypto"
description: "cryptographic functions"
weight: 7980
categories: ["Sec"]
width: full
---

cryptographic functions

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7980** | {{< badge content="pgcrypto" link="https://www.postgresql.org/docs/current/pgcrypto.html" >}} | {{< ext "pgcrypto" "pgcrypto" >}} | `1.3` | {{< category "SEC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-dt-" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "omni_auth" >}} {{< ext "omni_aws" >}} {{< ext "omni_credentials" >}} {{< ext "omni_kube" >}} {{< ext "omni_rest" >}} {{< ext "pgcryptokey" >}} |
|   **See Also**    | {{< ext "pgsodium" >}} {{< ext "pgsmcrypto" >}} {{< ext "lo" >}} {{< ext "anon" >}} {{< ext "pg_tde" >}} {{< ext "sslutils" >}} {{< ext "faker" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< badge content="1.3" color="green" >}} | {{< badge content="1.3" color="green" >}} | {{< badge content="1.3" color="green" >}} | {{< badge content="1.3" color="green" >}} | {{< badge content="1.3" color="green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgcrypto;
```
