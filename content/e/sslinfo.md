---
title: "sslinfo"
linkTitle: "sslinfo"
description: "information about SSL certificates"
weight: 6920
categories: ["STAT"]
width: full
---

[**sslinfo**](https://www.postgresql.org/docs/current/sslinfo.html) : information about SSL certificates


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6920** | {{< badge content="sslinfo" link="https://www.postgresql.org/docs/current/sslinfo.html" >}} | {{< ext "sslinfo" >}} | `1.2` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "sslutils" >}} {{< ext "pg_profile" >}} {{< ext "pg_tracing" >}} {{< ext "pg_show_plans" >}} {{< ext "pg_stat_kcache" >}} {{< ext "pg_stat_monitor" >}} {{< ext "pg_qualstats" >}} {{< ext "pg_store_plans" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.2" "PostgreSQL 18: version 1.2" "green" >}} | {{< bg "1.2" "PostgreSQL 17: version 1.2" "green" >}} | {{< bg "1.2" "PostgreSQL 16: version 1.2" "green" >}} | {{< bg "1.2" "PostgreSQL 15: version 1.2" "green" >}} | {{< bg "1.2" "PostgreSQL 14: version 1.2" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION sslinfo;
```



## Usage

> [sslinfo: SSL certificate information functions](https://www.postgresql.org/docs/current/sslinfo.html)

sslinfo provides functions to access information about the SSL certificate used in the current database connection.

### Functions

```sql
-- Check if current connection uses SSL
SELECT ssl_is_used();

-- SSL/TLS protocol version (TLSv1.2, TLSv1.3, etc.)
SELECT ssl_version();

-- Cipher name (e.g., DHE-RSA-AES256-SHA)
SELECT ssl_cipher();

-- Check if client presented a valid certificate
SELECT ssl_client_cert_present();

-- Client certificate serial number
SELECT ssl_client_serial();

-- Client certificate subject (full DN)
SELECT ssl_client_dn();
-- e.g., /CN=Somebody /C=Some country/O=Some organization

-- Certificate issuer (full DN)
SELECT ssl_issuer_dn();

-- Specific field from client certificate subject
SELECT ssl_client_dn_field('CN');
SELECT ssl_client_dn_field('O');

-- Specific field from certificate issuer
SELECT ssl_issuer_field('CN');

-- Client certificate extensions
SELECT * FROM ssl_extension_info();
-- Returns: name, value, critical
```

### Notes

- Most functions return NULL if the connection does not use SSL
- Requires PostgreSQL compiled with OpenSSL support
- The combination of `ssl_client_serial()` and `ssl_issuer_dn()` uniquely identifies a certificate
