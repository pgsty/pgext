---
title: "isn"
linkTitle: "isn"
description: "data types for international product numbering standards"
weight: 3930
categories: ["TYPE"]
width: full
---

[**isn**](https://www.postgresql.org/docs/current/isn.html) : data types for international product numbering standards


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3930** | {{< badge content="isn" link="https://www.postgresql.org/docs/current/isn.html" >}} | {{< ext "isn" >}} | `1.2` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} {{< ext "roaringbitmap" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.2" "PostgreSQL 18: version 1.2" "green" >}} | {{< bg "1.2" "PostgreSQL 17: version 1.2" "green" >}} | {{< bg "1.2" "PostgreSQL 16: version 1.2" "green" >}} | {{< bg "1.2" "PostgreSQL 15: version 1.2" "green" >}} | {{< bg "1.2" "PostgreSQL 14: version 1.2" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION isn;
```



## Usage

> [isn: ISBN, ISSN, EAN, UPC product number types](https://www.postgresql.org/docs/current/isn.html)

The `isn` extension provides data types for international product numbering standards with validation and check-digit verification.

```sql
CREATE EXTENSION isn;
```

### Data Types

| Type | Description |
|------|-------------|
| `EAN13` | European Article Numbers (always 13-digit) |
| `ISBN13` | International Standard Book Numbers (13-digit) |
| `ISMN13` | International Standard Music Numbers (13-digit) |
| `ISSN13` | International Standard Serial Numbers (13-digit) |
| `ISBN` | Book numbers (10-digit short format when possible) |
| `ISMN` | Music numbers (10-digit short format when possible) |
| `ISSN` | Serial numbers (10-digit short format when possible) |
| `UPC` | Universal Product Codes |

### Examples

```sql
SELECT isbn('978-0-393-04002-9');
SELECT isbn13('0901690546');
SELECT issn('1436-4522');

-- Create table with ISBN column
CREATE TABLE books (id isbn);
INSERT INTO books VALUES ('9780393040029');

-- Auto-calculate check digit with ?
INSERT INTO books VALUES ('220500896?');

-- Type casting
SELECT ean13('0901690546'::isbn);
```

### Weak Mode

Enable weak mode to accept invalid check digits (useful for OCR/data cleanup):

```sql
SET isn.weak TO true;
INSERT INTO books VALUES ('978-0-11-000533-4');
SET isn.weak TO false;

-- Find invalid entries
SELECT id FROM books WHERE NOT is_valid(id);

-- Fix invalid entries
UPDATE books SET id = make_valid(id) WHERE id = '2-205-00876-X!';
```

### Functions

| Function | Description |
|----------|-------------|
| `is_valid(isn)` | Check if value has valid check digit |
| `make_valid(isn)` | Clear the invalid-check-digit flag |
| `isn_weak()` | Return current weak mode status |
| `isn_weak(boolean)` | Set weak mode |

### Operators and Indexing

Standard comparison operators (`=`, `<`, `>`, `<=`, `>=`, `<>`) with B-tree and hash index support.
