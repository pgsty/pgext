---
title: "test_decoding"
linkTitle: "test_decoding"
description: "SQL-based test/example module for WAL logical decoding"
weight: 9970
categories: ["ETL"]
width: full
---

[**test_decoding**](https://www.postgresql.org/docs/current/test-decoding.html) : SQL-based test/example module for WAL logical decoding


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9970** | {{< badge content="test_decoding" link="https://www.postgresql.org/docs/current/test-decoding.html" >}} | {{< ext "test_decoding" >}} | `-` | {{< category "ETL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s----" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "wal2json" >}} {{< ext "decoderbufs" >}} {{< ext "decoder_raw" >}} {{< ext "pgoutput" >}} {{< ext "pglogical" >}} {{< ext "pg_failover_slots" >}} {{< ext "pgactive" >}} {{< ext "kafka_fdw" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** | **PG13** |
|:--------:|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "-" "PostgreSQL 18: version -" "green" >}} | {{< bg "-" "PostgreSQL 17: version -" "green" >}} | {{< bg "-" "PostgreSQL 16: version -" "green" >}} | {{< bg "-" "PostgreSQL 15: version -" "green" >}} | {{< bg "-" "PostgreSQL 14: version -" "green" >}} | {{< bg "-" "PostgreSQL 13: version -" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


This extension does not need `CREATE EXTENSION` DDL command


