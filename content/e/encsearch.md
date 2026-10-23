---
title: "encsearch"
linkTitle: "encsearch"
description: "Experimental encrypted text search using AEAD ciphers and keyed blind indexes"
weight: 7035
categories: ["SEC"]
languages: ["Rust"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**encsearch**](https://github.com/lokeshmulpuru/encsearch) : Experimental encrypted text search using AEAD ciphers and keyed blind indexes


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7035** | {{< badge content="encsearch" link="https://github.com/lokeshmulpuru/encsearch" >}} | {{< ext "encsearch" >}} | `0.1.1` | {{< category "SEC" >}} | {{< license "PostgreSQL" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |

> [!Note] Early release 0.1.1 changes PGXN status to stable without code changes from 0.1.0; PG17 tested, other Cargo PG14-18 targets untested. Session preload must precede key settings; blind indexes leak equality and token frequency.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.1` | {{< bg "18" "" "red" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `encsearch` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.1` | {{< bg "18" "encsearch_18" "red" >}} {{< bg "17" "encsearch_17" "green" >}} {{< bg "16" "encsearch_16" "red" >}} {{< bg "15" "encsearch_15" "red" >}} {{< bg "14" "encsearch_14" "red" >}} | `encsearch_$v` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "encsearch_18 : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.1.1" "encsearch_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "encsearch_16 : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch_15 : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "encsearch_18 : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.1.1" "encsearch_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "encsearch_16 : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch_15 : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "N/A" "encsearch_18 : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.1.1" "encsearch_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "encsearch_16 : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch_15 : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "N/A" "encsearch_18 : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.1.1" "encsearch_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "encsearch_16 : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch_15 : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "N/A" "encsearch_18 : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.1.1" "encsearch_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "encsearch_16 : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch_15 : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "N/A" "encsearch_18 : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.1.1" "encsearch_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "encsearch_16 : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch_15 : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} | {{< bg "N/A" "encsearch : N/A 0" "gray" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `encsearch_17` | `0.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 517.1 KiB | [encsearch_17-0.1.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/encsearch_17-0.1.1-1PGSTY.el8.x86_64.rpm) |
| `encsearch_17` | `0.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 414.3 KiB | [encsearch_17-0.1.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/encsearch_17-0.1.1-1PGSTY.el8.aarch64.rpm) |
| `encsearch_17` | `0.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 519.3 KiB | [encsearch_17-0.1.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/encsearch_17-0.1.1-1PGSTY.el9.x86_64.rpm) |
| `encsearch_17` | `0.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 445.6 KiB | [encsearch_17-0.1.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/encsearch_17-0.1.1-1PGSTY.el9.aarch64.rpm) |
| `encsearch_17` | `0.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 473.0 KiB | [encsearch_17-0.1.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/encsearch_17-0.1.1-1PGSTY.el10.x86_64.rpm) |
| `encsearch_17` | `0.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 430.4 KiB | [encsearch_17-0.1.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/encsearch_17-0.1.1-1PGSTY.el10.aarch64.rpm) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/lokeshmulpuru/encsearch" title="Repository" icon="github" subtitle="github.com/lokeshmulpuru/encsearch" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="encsearch-0.1.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg encsearch;		# build rpm
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install encsearch;		# install via package name, for the active PG version

pig install encsearch -v 17;   # install for PG 17

```


[**Config**](https://pigsty.io/docs/pgsql/ext/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'encsearch';
```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION encsearch;
```

## Usage

Sources:

- [README.md](https://github.com/lokeshmulpuru/encsearch/blob/cfa1a2af334ebbd4026a9f467c0ec40bed12a035/README.md)
- [src/lib.rs](https://github.com/lokeshmulpuru/encsearch/blob/cfa1a2af334ebbd4026a9f467c0ec40bed12a035/src/lib.rs)
- [Cargo.toml](https://github.com/lokeshmulpuru/encsearch/blob/cfa1a2af334ebbd4026a9f467c0ec40bed12a035/Cargo.toml)
- [META.json](https://github.com/lokeshmulpuru/encsearch/blob/cfa1a2af334ebbd4026a9f467c0ec40bed12a035/META.json)
- [encsearch.control](https://github.com/lokeshmulpuru/encsearch/blob/cfa1a2af334ebbd4026a9f467c0ec40bed12a035/encsearch.control)

`encsearch` stores encrypted text alongside keyed search tokens in an `int8[]` column. Its experimental 0.1.0 release supports exact, word and pattern searches through a GIN index. Upstream has tested PostgreSQL 17; the other Cargo feature targets are untested.

### Core Workflow

```ini
session_preload_libraries = 'encsearch'
```

```sql
CREATE EXTENSION encsearch;
CREATE TABLE encrypted_notes (id bigint GENERATED ALWAYS AS IDENTITY, body bytea, idx int8[]);
CREATE INDEX encrypted_notes_tokens ON encrypted_notes USING gin (idx);
INSERT INTO encrypted_notes (body, idx)
VALUES (enc_encrypt('Patient has diabetes'), enc_index('Patient has diabetes'));
SELECT enc_decrypt(body) FROM encrypted_notes
WHERE enc_like(body, idx, '%diab%');
```

### Operational Boundaries

An administrator installs this untrusted extension. Configure `session_preload_libraries` before opening client sessions and before setting `encsearch.key`; otherwise an early key setting remains an unregistered setting visible to other roles. Supply a securely retained key of at least 32 characters to the session or explicit-key function overloads before running the data example. `enc_gen_key()` generates a 64-character key. Losing the key loses access to the plaintext.

`enc_encrypt` and `enc_decrypt` handle text and ciphertext. `enc_index` creates stored tokens; `enc_like`, `enc_ilike` and `enc_eq` filter candidates and check decrypted values. `enc_has_words` tests token containment without decrypting; `enc_q_like`, `enc_q_exact` and `enc_q_words` expose raw query tokens. `enc_algorithm` and `enc_algorithms` inspect available ciphers.

`encsearch.algorithm` defaults to `aes-256-gcm`; deterministic `aes-256-siv` reveals duplicate plaintext. Keep `encsearch.token_hash` identical when building and querying tokens: `hmac-sha256` is the default and `blake3` is the alternative. Changing the token hash requires regenerating tokens; changing keys requires decrypting with the old key and rebuilding ciphertext and tokens with the new key.

Blind indexes reveal shared values, words and character chunks and their frequency. Keys reach the database server and may appear in statement or parameter logs. This is an experimental privacy/performance tradeoff, not protection against a server that has the key.
