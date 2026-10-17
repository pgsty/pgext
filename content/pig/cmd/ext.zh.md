---
title: "CMD: ext"
description: 如何使用 pig ext 子命令管理扩展
icon: fa-solid fa-terminal
weight: 620
---

`pig ext` 查询扩展目录，安装、升级或移除主机上的 PostgreSQL 扩展软件包。它也支持 PostgreSQL 内核包。安装软件包后，按扩展要求在数据库中执行 SQL 并修改配置。

| 命令           | 描述             | 备注                |
|:-------------|:---------------|:------------------|
| `ext list`   | 搜索扩展           |                   |
| `ext info`   | 显示扩展详细信息       |                   |
| `ext avail`  | 显示扩展可用性矩阵      |                   |
| `ext status` | 显示主机可见的扩展文件/软件包 | 不会逐库查询            |
| `ext scan`   | 扫描已安装的控制文件/共享库  | 不会逐库查询            |
| `ext add`    | 安装扩展           | 需要 sudo 或 root 权限 |
| `ext rm`     | 移除扩展           | 需要 sudo 或 root 权限 |
| `ext update` | 更新扩展           | 需要 sudo 或 root 权限 |
| `ext import` | 下载扩展以供离线使用     | 需要 sudo 或 root 权限 |
| `ext link`   | 链接 PG 版本到 PATH | 需要 sudo 或 root 权限 |
| `ext reload` | 刷新扩展目录         |                   |
{.full-width}

## 快速入门

以下示例使用已运行的 PostgreSQL 18 实例。安装与启动步骤见[快速上手](/zh/pig/start/)。

添加 PostgreSQL 仓库，查询扩展，再安装软件包：

```bash
pig repo add pgsql -u
pig ext list vector
pig ext info vector
pig ext add vector -v 18
```

在目标数据库中启用扩展，例如：

```sql
CREATE EXTENSION vector;
```

其他扩展见[扩展列表](https://pigsty.cc/ext/list/)。
扩展需要预加载或额外配置时，先按 `pig ext info NAME` 的提示完成配置，再启用扩展。

### 选择 PostgreSQL 安装

- `-v, --version 18`：指定 PostgreSQL 主版本。
- `-p, --path /path/to/pg_config`：指定一套 PostgreSQL 安装的 `pg_config`。
- 两个参数不能同时使用。省略时，PIG 尝试探测本机 PostgreSQL；自动化脚本建议显式传入 `-v`。

PGDG 默认路径为 EL 的 `/usr/pgsql-18/bin/pg_config`，以及 Debian/Ubuntu 的 `/usr/lib/postgresql/18/bin/pg_config`。包管理器负责解析 RPM/DEB 依赖；`-y` 自动确认软件包操作。

## 软件包与数据库操作 {#软件包层与数据库层}

| 任务 | PIG 操作 | 数据库操作 |
|:---|:---|:---|
| 安装 | `ext add NAME` | 按要求配置预加载并执行 `CREATE EXTENSION` |
| 升级 | `ext update NAME` | 按扩展升级说明执行 `ALTER EXTENSION ... UPDATE` |
| 移除 | `ext rm NAME` | 先在受影响数据库中移除或迁移依赖，再卸载软件包 |
| 检查 | `ext status` / `ext scan` 检查主机文件 | 查询 `pg_extension` 查看数据库已启用扩展 |
| 刷新目录 | `ext reload` 更新扩展元数据 | 无 |

`pg_available_extensions` 显示服务器可用的扩展，`pg_extension` 显示当前数据库已启用的扩展。

## ext list

列出（或搜索）扩展目录中的条目。

```bash
pig ext list                     # 列出所有扩展
pig ext list duck                # 搜索包含 "duck" 的扩展
pig ext list -v 18               # 按 PG 18 判断可用性
pig ext ls olap                  # 列出 olap 类别扩展
pig ext ls gis -v 16             # 列出 GIS 条目及其在 PG 16 上的可用性
pig ext ls rag                   # 列出 RAG 类别扩展
```

分类筛选通过查询参数直接指定分类名实现，支持的分类包括：`time`, `gis`, `rag`, `fts`, `olap`, `feat`, `lang`, `type`, `func`, `util`, `admin`, `stat`, `sec`, `fdw`, `sim`, `etl`。

**选项：**

- `-v|--version`：指定判断可用性时使用的 PG 版本；不可用的条目仍会列出
- `--pkg`：显示包名而非扩展名，仅列出主导扩展

**Status 列说明：**

- `installed`：扩展已安装（绿色）
- `available`：扩展可用但未安装（黄色）
- `not avail`：扩展在当前系统不可用（红色）

可用性判断与 text/JSON/YAML 摘要使用同一平台、架构和 PostgreSQL 版本。平台有包矩阵时，
以矩阵为准，缺少对应单元格即不可用；条目没有矩阵时，沿用 RPM/DEB 版本元数据与平台例外规则。
`amd64`/`x86_64`、`arm64`/`aarch64` 分别等价。macOS 的此状态只表示声明支持的 PostgreSQL 版本，
不表示存在可安装的 RPM/DEB 包。

### 目录计数 {#catalog-count}

官网与 CLI 按全部目录记录统计，包括 `pg_relation_sql` 这类已打包的 SQL 工具，以及已移除扩展的保留记录。当前内嵌目录共 {{< param pig_catalog_count >}} 个条目。这个总数不表示每个条目都支持 `CREATE EXTENSION`，也不表示它们在所有平台上都可安装。

无筛选列表包含全部记录。文本行数与 JSON/YAML 的 `data.count` 都表示实际返回的条目数；搜索统计匹配条目，`--pkg` 统计主导包记录。`ext avail` 的全局文本摘要使用完整目录总数。`ext reload` 的 `data.extension_count` 表示实际下载目录的记录数。

[计数决策](https://pig.pgsty.com/zh/design/catalog-entry-count/)说明了统一口径的原因。默认扩展目录定义在 [**`cli/ext/assets/extension.csv`**](https://github.com/pgsty/pig/blob/main/cli/ext/assets/extension.csv)。

可用 `pig ext reload` 命令更新到最新扩展目录，数据将下载到 `~/.pig/extension.csv`；在线最新版目录同步发布于 [**repo.pigsty.cc/ext/data/extension.csv**](https://repo.pigsty.cc/ext/data/extension.csv)。

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

可用性矩阵会显示扩展在各个操作系统（EL8/9/10, Debian 12/13, Ubuntu 22/24/26）、架构（x86_64/aarch64）和 PostgreSQL 版本（14-18）上的可用情况。

## ext status

显示所选 PostgreSQL 安装在当前主机上可见的扩展文件/软件包状态。

```bash
pig ext status              # 显示已安装扩展
pig ext status -c           # 包含 contrib 扩展
pig ext status -v 16        # 显示 PG 16 已安装扩展
```

**选项：**

- `-c|--contrib`：结果中包含 contrib 扩展

## ext scan

扫描所选 PostgreSQL 安装的扩展目录。

```bash
pig ext scan [-v version]
```

该命令读取所选 PostgreSQL 安装目录中的控制文件及相关文件。
同时按目录元数据匹配共享库名称，因此没有控制文件的条目也可以显示为已安装，
但这不意味着该条目支持 `CREATE EXTENSION`。

## ext add

安装一个或多个 PostgreSQL 扩展。`pig ext add` 的同级别名包括 `pig ext install`、`pig ext ins` 与 `pig ext a`。顶层 [`pig install`](/zh/pig/cmd/#pig-install) 也接受扩展别名和普通系统软件包。

```bash
pig ext add pg_duckdb            # 安装 pg_duckdb
pig ext add pg_duckdb -v 18      # 为 PG 18 安装
pig ext add pg_duckdb -y         # 自动确认安装
pig ext add vector postgis       # 安装多个扩展
pig ext add postgis --plan       # 预览安装计划，不执行

# 使用顶层软件包安装命令
pig install pg_duckdb
pig install pg_duckdb -v 18 -y

# 安装 PostgreSQL 软件包
pig ext add pg18                 # 安装 PostgreSQL 18
pig ext add pg18-client          # 仅安装客户端工具
```

以下别名用于选择 PostgreSQL 软件包组：

| 别名 | 软件包 |
|:---|:---|
| `pg18` | 服务端、客户端、contrib 模块与过程语言 |
| `pg18-server` | 服务端与 contrib 模块 |
| `pg18-client` | 客户端工具 |
| `pg18-devel` | 开发文件 |
| `pg18-basic` | pg_repack、wal2json 和 pgvector |

将 `18` 替换为所需的大版本，也可以用 `pgsql` 配合 `-v 18`。
安装特定软件包版本时，使用 `NAME=VERSION`，版本必须存在于已配置的仓库中。

**选项：**

- `-v|--version`：指定 PG 大版本
- `-y|--yes`：自动确认安装
- `--plan`：预览安装计划，不执行包管理器命令

软件包安装成功，只代表主机上已有扩展文件。请按照 `ext info` 给出的提示完成预加载、重启与 `CREATE EXTENSION`。

## ext rm

移除一个或多个 PostgreSQL 扩展。

```bash
pig ext rm pg_duckdb             # 移除 pg_duckdb
pig ext rm pg_duckdb -v 18       # 移除 PG 18 版本
pig ext rm pgvector -y           # 自动确认移除
pig ext rm pgvector --plan       # 预览移除计划，不执行
```

**选项：**

- `-v|--version`：指定 PG 大版本
- `-y|--yes`：自动确认移除
- `--plan`：预览移除计划，不执行包管理器命令

该命令移除的是操作系统软件包，而不是数据库对象。删除共享文件前，应先在所有受影响数据库中移除或迁移依赖它的扩展。

## ext update

将指定软件包更新到已配置仓库提供的版本。不带参数时，命令不执行更新。

```bash
pig ext update                   # no-op：必须显式指定目标
pig ext update pg_duckdb         # 更新特定扩展
pig ext update postgis timescaledb  # 更新多个扩展
pig ext update pg_duckdb -y      # 自动确认更新
pig ext update pg_duckdb -m      # 先从中国镜像刷新扩展目录
```

**选项：**

- `-v|--version`：指定 PG 大版本
- `-y|--yes`：自动确认更新
- `-m|--mirror`：更新软件包前，优先从 `pigsty.cc` 镜像刷新扩展目录；软件包仍从已配置的系统仓库下载

该命令不会执行 SQL 迁移。软件包升级完成后，请按扩展发布说明，在每个适用数据库中执行 `ALTER EXTENSION ... UPDATE`。

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

- `-d|--repo`：指定仓库目录（默认：`/www/pigsty`）

部分名称无法解析时，已解析出的包仍会下载，但命令以失败退出。结构化结果保留 `requested`、
已解析的 `packages`、确认完成的 `downloaded` 与无法解析的 `failed` 名称。
原生下载命令失败也可能在目标目录留下文件；失败不表示回滚，也不表示下载完整。

Debian/Ubuntu 上，`<virtual-package>` 这样的虚拟依赖标记不会加入下载包名列表，
真实软件包依赖仍参与下载。

JSON/YAML 输出及 `PIG_NON_INTERACTIVE=1` 使用非交互 sudo。提权需要密码时，命令直接失败，
不会等待输入；请先配置适当的 sudo 权限。

## ext link

将指定 PG 版本链接到系统 PATH。

```bash
pig ext link 18                  # 链接 PG 18 到 PATH
pig ext link pg17                # pg 前缀会被剥离，链接 PG 17
pig ext link 16                  # 链接 PG 16 到 /usr/pgsql
pig ext link /usr/pgsql-16       # 从指定路径链接到 /usr/pgsql
pig ext link polar               # 链接 PolarDB / PolarPG 安装
pig ext link /usr/polar-17       # 从指定 PolarDB 路径链接
pig ext link null                # 取消当前 PostgreSQL 链接
pig ext link none                # null / none / nil / nop / no 均可取消链接
```

该命令创建 `/usr/pgsql` 软链接，并写入 `/etc/profile.d/pgsql.sh`。在当前 Shell 中加载配置：

```bash
. /etc/profile.d/pgsql.sh
```

## ext reload

刷新扩展元数据。

```bash
pig ext reload                   # 刷新扩展目录
```

更新后的文件会放置于 `~/.pig/extension.csv` 中。
