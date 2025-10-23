---
title: 升级
description: 如何将 PostgreSQL 扩展升级到新版本
icon: RefreshCcw
---

要升级已安装的扩展，需先通过操作系统的包管理器（RPM/DEB）升级扩展包，
然后在 PostgreSQL 中执行 `ALTER EXTENSION ... UPDATE` 命令切换至新版本。

你可以使用如下命令升级扩展包：

```bash tab="pig"
pig ext update 扩展名...
```
```bash tab="yum"
yum upgrade 扩展名...
```
```bash tab="apt"
apt upgrade 扩展名...
```
```bash tab="pigsty"
./pgsql.yml -t pg_ext   # -l cls
```

所有在 [`pg_extensions`](/zh/docs/pgsql/param#pg_extensions) 中列出的扩展，
都会在执行 [`pgsql.yml`](/zh/docs/pgsql/playbook#pgsqlyml) 剧本时自动升级。

------

## 升级扩展包

在 [`pg_extensions`](/zh/docs/pgsql/param#pg_extensions) 中指定的扩展（[包别名](/zh/usage/pkg)）
会通过 [`pgsql.yml`](/zh/docs/pgsql/admin/) 的 `pg_ext` 子任务自动升级：

```bash title="~/pigsty"
./pgsql.yml -t pg_ext
```

该剧本会自动安装当前环境下可用的最新扩展 RPM/DEB 包
（无论是本地仓库还是互联网源）。
你也可以直接用 Linux 系统的 `yum/apt upgrade` 命令升级扩展，但需指定完整包名：

```bash
yum upgrade 扩展名...
apt upgrade 扩展名...
```

Pigsty 的 [`pig`](/zh/pig/) CLI 也可简化操作，无需手动指定完整包名：

```bash
pig ext update 扩展名|包别名
```

------

## 升级扩展版本

通过执行 [`ALTER EXTENSION ... UPDATE`](https://www.postgresql.org/docs/current/sql-alterextension.html) SQL 命令，将扩展升级到新版本：

```sql
ALTER EXTENSION name UPDATE [ TO new_version ]
```

如省略 `TO new_version`，则会升级到可用的最新版本。
