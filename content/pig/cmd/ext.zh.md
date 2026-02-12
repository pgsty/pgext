---
title: "CMD: ext"
description: 如何使用 pig ext 子命令管理扩展
icon: SquareTerminal
weight: 620
---


`pig ext` 命令是一个用于管理 PostgreSQL 扩展的全能工具。它允许用户搜索、安装、移除、更新和管理 PostgreSQL 扩展，甚至支持内核包的管理。

```bash
pig ext - Manage PostgreSQL Extensions

  pig repo add -ru             # add all repo and update cache (brute but effective)
  pig ext add pg17             # install optional postgresql 17 package
  pig ext list duck            # search extension in catalog
  pig ext scan -v 17           # scan installed extension for pg 17
  pig ext add pg_duckdb        # install certain postgresql extension

Examples:
  pig ext list    [query]      # list & search extension
  pig ext info    [ext...]     # get information of a specific extension
  pig ext status  [-v]         # show installed extension and pg status
  pig ext add     [ext...]     # install extension for current pg version
  pig ext rm      [ext...]     # remove extension for current pg version
  pig ext update  [ext...]     # update extension to the latest version
  pig ext import  [ext...]     # download extension to local repo
  pig ext link    [ext...]     # link postgres installation to path
  pig ext reload               # reload the latest extension catalog data
```

| 命令 | 描述 | 备注 |
|:---|:---|:---|
| `ext list` | 搜索扩展 | |
| `ext info` | 显示扩展详细信息 | |
| `ext avail` | 显示扩展可用性矩阵 | |
| `ext status` | 显示已安装的扩展 | |
| `ext scan` | 扫描已安装的扩展 | |
| `ext add` | 安装扩展 | 需要 sudo 或 root 权限 |
| `ext rm` | 移除扩展 | 需要 sudo 或 root 权限 |
| `ext update` | 更新扩展 | 需要 sudo 或 root 权限 |
| `ext import` | 下载扩展以供离线使用 | 需要 sudo 或 root 权限 |
| `ext link` | 链接 PG 版本到 PATH | 需要 sudo 或 root 权限 |
| `ext reload` | 刷新扩展目录 | |
{.full-width}


## 快速入门

在安装 PostgreSQL 扩展前，你需要先添加 [`pig repo add`](/zh/pig/cmd/repo)：

```bash
pig repo add pgdg pigsty -u    # 温和方式添加 pgdg 和 pigsty 仓库
pig repo set                   # 粗暴方式移除并添加所有所需仓库
```

然后你可以搜索并安装 PostgreSQL 扩展：

```bash
pig ext install pg_duckdb
pig ext install pg_partman
pig ext install pg_cron
pig ext install pg_repack
pig ext install pg_stat_statements
pig ext install pg_stat_kcache
```

可用扩展及其名称请查阅 [**扩展列表**](https://pgext.cloud/zh/list)。

**使用说明：**

1. 未指定 PostgreSQL 版本时，工具会尝试从 `PATH` 中的 `pg_config` 自动检测当前活动的 PostgreSQL 安装。
2. PostgreSQL 可通过主版本号（`-v`）或 pg_config 路径（`-p`）指定。
   - 若指定 `-v`，pig 会使用该版本 PGDG 内核包的默认路径。
     - EL 发行版为 `/usr/pgsql-$v/bin/pg_config`，
     - DEB 发行版为 `/usr/lib/postgresql/$v/bin/pg_config` 等。
   - 若指定 `-p`，则直接用该路径定位 PostgreSQL。
3. 扩展管理器会根据操作系统自动适配不同的包格式：
   - RHEL/CentOS/Rocky Linux/AlmaLinux 使用 RPM 包
   - Debian/Ubuntu 使用 DEB 包
4. 某些扩展可能有依赖项，安装时会自动解决。
5. 谨慎使用 `-y` 参数，它会自动确认所有提示。

Pigsty 假定你已安装官方 PGDG 内核包，如未安装，可用如下命令：

```bash
pig ext install pg18          # 安装 PostgreSQL 18 内核（除 devel 包）
```


## ext list

列出（或搜索）扩展目录中的可用扩展。

```bash
pig ext list                     # 列出所有扩展
pig ext list duck                # 搜索包含 "duck" 的扩展
pig ext list -v 17               # 按 PG 版本筛选
pig ext ls olap                  # 列出 olap 类别扩展
pig ext ls gis -v 16             # 列出 PG 16 的 GIS 类扩展
pig ext ls rag                   # 列出 RAG 类别扩展
```

分类筛选通过查询参数直接指定分类名实现，支持的分类包括：`time`, `gis`, `rag`, `fts`, `olap`, `feat`, `lang`, `type`, `func`, `util`, `admin`, `stat`, `sec`, `fdw`, `sim`, `etl`。

**选项：**

- `-v|--version`：按 PG 版本筛选
- `--pkg`：显示包名而非扩展名，仅列出主导扩展

**Status 列说明：**

- `installed`：扩展已安装（绿色）
- `available`：扩展可用但未安装（黄色）
- `not avail`：扩展在当前系统不可用（红色）

默认扩展目录定义在 [**`cli/ext/assets/extension.csv`**](https://github.com/pgsty/pig/blob/main/cli/ext/assets/extension.csv)

可用 `pig ext reload` 命令更新到最新扩展目录，数据将下载到 `~/.pig/extension.csv`。


## ext info

显示指定扩展的详细信息。

```bash
pig ext info postgis        # 显示 PostGIS 详细信息
pig ext info timescaledb    # 显示 TimescaleDB 信息
pig ext info vector postgis # 显示多个扩展信息
```


## ext avail

显示扩展的可用性矩阵，展示扩展在不同操作系统、架构和 PostgreSQL 版本上的可用情况。

```bash
pig ext avail                     # 显示当前系统上所有包的可用性
pig ext avail timescaledb         # 显示 timescaledb 的可用性矩阵
pig ext avail postgis pg_duckdb   # 显示多个扩展的可用性
pig ext av pgvector               # 显示 pgvector 的可用性
pig ext matrix citus              # avail 命令的别名
```

可用性矩阵会显示扩展在各个操作系统（EL8/9/10, Debian 12/13, Ubuntu 22/24）、架构（x86_64/aarch64）和 PostgreSQL 版本（13-18）上的可用情况。


## ext status

显示当前 PostgreSQL 实例已安装扩展的状态。

```bash
pig ext status              # 显示已安装扩展
pig ext status -c           # 包含 contrib 扩展
pig ext status -v 16        # 显示 PG 16 已安装扩展
```

**选项：**

- `-c|--contrib`: 结果中包含 contrib 扩展


## ext scan

扫描当前 PostgreSQL 实例已安装的扩展。

```bash
pig ext scan [-v version]
```

该命令会扫描 postgres 扩展目录，查找所有实际已安装的扩展。


## ext add

安装一个或多个 PostgreSQL 扩展。也可以使用别名 `pig install`。

```bash
pig ext add pg_duckdb            # 安装 pg_duckdb
pig ext add pg_duckdb -v 17      # 为 PG 17 安装
pig ext add pg_duckdb -y         # 自动确认安装
pig ext add vector postgis       # 安装多个扩展

# 使用别名
pig install pg_duckdb
pig install pg_duckdb -v 17 -y

# 安装 PostgreSQL 内核
pig ext install pgsql            # 安装最新版 postgresql 内核
pig ext a pg17                   # 安装 postgresql 17 内核包
pig ext ins pg16                 # 安装 postgresql 16 内核包
pig ext install pg15-core        # 安装 postgresql 15 核心包
pig ext install pg14-main -y     # 安装 pg 14 + 常用扩展（vector, repack, wal2json）
```

**选项：**

- `-v|--version`：指定 PG 大版本
- `-y|--yes`：自动确认安装


## ext rm

移除一个或多个 PostgreSQL 扩展。

```bash
pig ext rm pg_duckdb             # 移除 pg_duckdb
pig ext rm pg_duckdb -v 17       # 移除 PG 17 版本
pig ext rm pgvector -y           # 自动确认移除
```


## ext update

将已安装扩展更新到最新版。

```bash
pig ext update                   # 更新所有已安装扩展
pig ext update pg_duckdb         # 更新特定扩展
pig ext update postgis timescaledb  # 更新多个扩展
pig ext update -y                # 自动确认更新
```


## ext import

下载扩展包到本地仓库，便于离线安装。

```bash
pig ext import postgis                # 导入 PostGIS 包
pig ext import timescaledb pg_cron    # 导入多个扩展包
pig ext import pg16                   # 导入 PostgreSQL 16 包
pig ext import pgsql-common           # 导入常用工具包
pig ext import -d /www/pigsty postgis # 指定路径导入
```

**选项：**

- `-d|--repo`: 指定仓库目录（默认：`/www/pigsty`）


## ext link

将指定 PG 版本链接到系统 PATH。

```bash
pig ext link 18                  # 链接 PG 18 到 PATH
pig ext link 17                  # 链接 PG 17 到 /usr/pgsql
pig ext link /usr/pgsql-16       # 从指定路径链接到 /usr/pgsql
pig ext link null                # 取消当前 PostgreSQL 链接
```

该命令会创建 `/usr/pgsql` 软链接，并写入 `/etc/profile.d/pgsql.sh`。


## ext reload

从 GitHub 刷新扩展元数据。

```bash
pig ext reload                   # 刷新扩展目录
```

更新后的文件会放置于 `~/.pig/extension.csv` 中。
