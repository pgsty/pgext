---
title: Update
description: How to update PostgreSQL extensions to newer versions
icon: RefreshCcw
---

To update an existing extension, you need to first update the RPM/DEB package with your OS's package manager,
then alter the extension to the new version in PostgreSQL with `ALTER EXTENSION ... UPDATE`.

You can upgrade extension packages with the following commands

```bash tab="pig"
pig ext update extname...
```
```bash tab="yum"
yum upgrade extname...
```
```bash tab="apt"
apt upgrade extname...
```
```bash tab="pigsty"
./pgsql.yml -t pg_ext   # -l cls
```

All extensions listed in [`pg_extensions`](/docs/pgsql/param#pg_extensions) will be upgraded using during the [`pgsql.yml`](/docs/pgsql/playbook#pgsqlyml) playbook execution.



------

## Upgrade Packages

Extensions ([Package Alias](/usage/pkg)) listed in [`pg_extensions`](/docs/pgsql/param#pg_extensions) will be upgraded with [`pgsql.yml`](/docs/pgsql/admin/)'s `pg_ext` subtask:

```bash title="~/pigsty"
./pgsql.yml -t pg_ext
```

This playbook will automatically install the latest available version of extension RPM/DEB packages in your current environment.
(from built local repo or via Internet directly).
You can also upgrade extensions with linux system's `yum/apt upgrade` command directly, but you need to specify the full package names:

```bash
yum upgrade extname...
apt upgrade extname...
```

Pigsty’s [`pig`](/pig/) cli can also help you with that, without the burden of specifying full package names:

```bash
pig ext update extname|pkgalias
```


------

## Alter Extension

Execute the [`ALTER EXTENSION ... UPDATE`](https://www.postgresql.org/docs/current/sql-alterextension.html) SQL command to update the extension to the new version:

```sql
ALTER EXTENSION name UPDATE [ TO new_version ]
```

If the `TO new_version` clause is omitted, the extension will be updated to the latest version available.
