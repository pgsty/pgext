---
title: "命令：ext"
description: 如何使用 pig ext 子命令管理扩展
icon: SquareTerminal
weight: 620
---

`pig ext` 命令是管理 PostgreSQL 扩展和内核包的综合工具。它提供了统一的接口来搜索、安装、移除、更新和管理 400 多个 PostgreSQL 扩展，支持不同的 PostgreSQL 版本和操作系统。

------

## 子命令

| 命令 | 描述 |
|---------|------------|
| `list`   | 列出和搜索可用扩展 |
| `info`   | 获取扩展信息 |
| `status` | 显示当前 PG 的已安装扩展 |
| `add`    | 安装 postgres 扩展 |
| `rm`     | 移除 postgres 扩展 |
| `update` | 更新当前 PG 版本的已安装扩展 |
| `scan`   | 扫描当前 PG 的已安装扩展 |
| `import` | 导入扩展包到本地仓库 |
| `link`   | 链接 postgres 到活动 PATH |
| `reload` | 重新加载扩展目录到最新版本 |

------

## 概览

```bash
pig ext - 管理 PostgreSQL 扩展

  快速开始: https://pigsty.io/ext/pig/
  pig repo add -ru             # 添加所有仓库并更新缓存（简单粗暴但有效）
  pig ext add pg17             # 安装可选的 postgresql 17 包
  pig ext list duck            # 在目录中搜索扩展
  pig ext scan -v 17           # 扫描 pg 17 的已安装扩展
  pig ext add pg_duckdb        # 安装特定的 postgresql 扩展

Usage:
  pig ext [command]

别名:
  ext, e, ex, pgext, extension

示例:

  pig ext list    [query]      # 列出和搜索扩展
  pig ext info    [ext...]     # 获取特定扩展的信息
  pig ext status  [-v]         # 显示已安装的扩展和 pg 状态
  pig ext add     [ext...]     # 为当前 pg 版本安装扩展
  pig ext rm      [ext...]     # 为当前 pg 版本移除扩展
  pig ext update  [ext...]     # 更新扩展到最新版本
  pig ext import  [ext...]     # 下载扩展到本地仓库
  pig ext link    [ext...]     # 链接 postgres 安装到路径
  pig ext reload               # 重新加载最新的扩展目录数据

可用命令:
  add         安装 postgres 扩展
  import      导入扩展包到本地仓库
  info        获取扩展信息
  link        链接 postgres 到活动 PATH
  list        列出和搜索可用扩展
  reload      重新加载扩展目录到最新版本
  rm          移除 postgres 扩展
  scan        扫描当前 pg 的已安装扩展
  status      显示当前 pg 上的已安装扩展
  update      更新当前 pg 版本的已安装扩展

标志:
  -h, --help          获取帮助信息
  -p, --path string   通过 pg_config 路径指定 postgres
  -v, --version int   通过主版本号指定 postgres

全局标志:
      --debug              启用调试模式
  -H, --home string        pigsty 主目录路径
  -i, --inventory string   配置清单路径
      --log-level string   日志级别: debug, info, warn, error, fatal, panic (默认 "info")
      --log-path string    日志文件路径，默认为终端输出
```

------

## 快速入门

在管理扩展之前，确保仓库已正确配置：

```bash
# 设置仓库（首次需要）
sudo pig repo add -ru

# 搜索扩展
pig ext list                     # 列出所有可用扩展
pig ext list duckdb              # 搜索特定扩展
pig ext list olap                # 按类别列出扩展

# 安装 PostgreSQL 和扩展
sudo pig ext install pg17        # 安装 PostgreSQL 17
sudo pig ext add pg_duckdb       # 安装 DuckDB 扩展
sudo pig ext add postgis timescaledb pgvector  # 安装多个扩展

# 检查已安装的扩展
pig ext status                   # 显示活动的 PostgreSQL 和扩展
pig ext scan                     # 扫描所有已安装的扩展
```

------

## PostgreSQL 版本检测

扩展管理器需要知道要定位哪个 PostgreSQL 版本。可以通过三种方式指定：

### 1. 自动检测（默认）
```bash
pig ext add pg_duckdb            # 使用 PATH 中的 pg_config
```
工具将在系统 PATH 中搜索 `pg_config` 并使用它来检测活动的 PostgreSQL 安装。

### 2. 通过主版本号（`-v`）
```bash
pig ext add pg_duckdb -v 17      # 为 PostgreSQL 17 安装
pig ext add postgis -v 16        # 为 PostgreSQL 16 安装
```
使用默认的 PGDG 安装路径：
- EL：`/usr/pgsql-$version/bin/pg_config`
- Debian/Ubuntu：`/usr/lib/postgresql/$version/bin/pg_config`

### 3. 通过 pg_config 路径（`-p`）
```bash
pig ext add pg_duckdb -p /opt/pgsql/bin/pg_config
pig ext add postgis -p /usr/local/pgsql/bin/pg_config
```
为自定义安装明确指定 pg_config 路径。

------

## 命令参考

### `ext list`

列出并搜索目录中可用的 PostgreSQL 扩展。

```bash
pig ext list                     # 列出所有扩展
pig ext list postgis             # 按名称/描述搜索
pig ext list olap                # 按类别列出
pig ext list gis -v 16           # 列出 PG 16 的 GIS 扩展
pig ext ls duck                  # 使用别名
```

**输出格式：**
```
名称            版本       描述                                        类别
----            -------    -----------                                --------
pg_duckdb       0.1.0      嵌入 PostgreSQL 的 DuckDB                  OLAP
duckdb_fdw      1.1.2      DuckDB 外部数据包装器                       FDW
pgvector        0.7.4      PostgreSQL 的向量相似性搜索                AI/ML
```

**类别：**
- `OLAP`：分析处理扩展
- `GIS`：地理和空间扩展
- `AI/ML`：机器学习和 AI 扩展
- `FDW`：外部数据包装器
- `STAT`：统计和监控
- `TIME`：时间序列扩展
- `SEARCH`：全文搜索扩展
- `ADMIN`：管理工具
- `SEC`：安全扩展
- `SHARD`：分片和分布式
- `JSON`：JSON/JSONB 实用工具

**搜索提示：**
- 搜索不区分大小写
- 同时搜索名称和描述字段
- 类别名称可用作搜索词
- 使用 `-v` 标志按 PostgreSQL 版本过滤

------

### `ext info`

显示特定扩展的详细信息。

```bash
pig ext info pg_duckdb           # 单个扩展
pig ext info postgis pgvector    # 多个扩展
```

**示例输出：**
```yaml
名称        : pg_duckdb
版本        : 0.1.0
类别        : OLAP
许可证      : MIT
主页        : https://github.com/duckdb/pg_duckdb
描述        : 嵌入 PostgreSQL 的 DuckDB
维护者      : DuckDB Labs
仓库        : pigsty
PostgreSQL  : 12,13,14,15,16,17
平台        : el8,el9,d12,u22,u24
包名称:
  - EL: pg_duckdb_$v
  - DEB: postgresql-$v-pg-duckdb
依赖:
  - duckdb (运行时)
  - cmake (构建)
```

------

### `ext status`

显示当前 PostgreSQL 安装状态和已安装的扩展。

```bash
pig ext status                   # 显示活动 PG 和扩展
pig ext status -v 17             # 特定 PG 版本的状态
```

**示例输出：**
```
PostgreSQL 状态:
  活动版本: 17
  安装位置: /usr/pgsql-17
  配置路径: /usr/pgsql-17/bin/pg_config
  数据目录: /var/lib/pgsql/17/data

已安装扩展 (17):
  - postgis (3.4.2)
  - pgvector (0.7.4)
  - pg_duckdb (0.1.0)
  - timescaledb (2.16.1)
  - pg_cron (1.6.4)

可用 PostgreSQL 版本:
  - 14: /usr/pgsql-14
  - 15: /usr/pgsql-15
  - 16: /usr/pgsql-16
  - 17: /usr/pgsql-17 [活动]
```

------

### `ext add`

安装 PostgreSQL 扩展或内核包。

```bash
# 安装扩展
sudo pig ext add pg_duckdb                      # 单个扩展
sudo pig ext add postgis pgvector timescaledb   # 多个扩展
sudo pig ext add pg_duckdb -y                   # 自动确认
sudo pig ext add pg_duckdb -v 16                # 为特定 PG 版本

# 安装 PostgreSQL 内核包
sudo pig ext install pg17                       # 完整 PostgreSQL 17
sudo pig ext install pg16-core                  # 仅核心包
sudo pig ext install pg15-main                  # 核心 + 基本扩展
sudo pig ext install pg14-devel                 # 包含开发包
sudo pig ext install pgsql-common               # 通用实用工具
```

**PostgreSQL 包集：**
- `pg$v` 或 `pgsql`：最新 PostgreSQL 版本
- `pg$v-core`：最小安装（服务器、客户端）
- `pg$v-main`：核心 + 基本扩展（pgvector、pg_repack、wal2json）
- `pg$v-mini`：最小可行安装
- `pg$v-devel`：包括开发包
- `pgsql-common`：实用工具（patroni、pgbouncer、pgbackrest、pg_exporter）

**选项：**
- `-y, --yes`：自动确认安装
- `-v, --version`：目标 PostgreSQL 版本
- `-p, --path`：通过 pg_config 路径定位 PostgreSQL

**常见扩展集：**

```bash
# OLAP 栈
sudo pig ext add pg_duckdb parquet_fdw

# GIS 栈
sudo pig ext add postgis pgrouting h3 ogr_fdw

# AI/ML 栈
sudo pig ext add pgvector pgvectorscale pg_similarity pgml

# 时间序列栈
sudo pig ext add timescaledb pg_partman pg_cron

# 监控栈
sudo pig ext add pg_stat_statements pg_stat_kcache pg_qualstats pg_wait_sampling

# 开发栈
sudo pig ext add pgtap plpgsql_check pg_hint_plan
```

------

### `ext rm`

移除已安装的 PostgreSQL 扩展。

```bash
sudo pig ext rm pg_duckdb                       # 移除单个扩展
sudo pig ext rm postgis pgvector                # 移除多个
sudo pig ext rm pg_duckdb -v 16                 # 从特定 PG 版本
```

**注意：**某些扩展可能有阻止移除的依赖关系。如需强制移除，请直接使用系统包管理器。

------

### `ext update`

更新已安装的扩展到最新可用版本。

```bash
sudo pig ext update                             # 更新所有扩展
sudo pig ext update pg_duckdb                   # 更新特定扩展
sudo pig ext update postgis pgvector -y         # 自动确认更新
```

**选项：**
- `-y, --yes`：自动确认更新
- `-v, --version`：目标 PostgreSQL 版本

------

### `ext scan`

扫描并列出活动或指定 PostgreSQL 的所有已安装扩展。

```bash
pig ext scan                     # 扫描活动 PostgreSQL
pig ext scan -v 17               # 扫描 PostgreSQL 17
pig ext scan -p /opt/pgsql/bin/pg_config  # 扫描自定义安装
```

**输出包括：**
- 扩展名称和版本
- 安装路径
- 控制文件位置
- 共享库状态
- 数据库可用性

------

### `ext import`

下载扩展包到本地仓库以供离线安装。

```bash
pig ext import pg_duckdb                        # 导入单个扩展
pig ext import postgis pgvector timescaledb     # 导入多个
pig ext import pg_duckdb -v 16,17               # 为多个 PG 版本
```

**用例：**
- 准备离线安装包
- 创建本地镜像仓库
- 为离线环境预下载

**工作流：**
```bash
# 在联网机器上
pig ext import pg_duckdb postgis pgvector
pig repo create                  # 创建本地仓库
pig repo cache                   # 创建离线包

# 传输到离线机器
# 在离线机器上
pig repo boot                    # 解压包
pig repo add local               # 使用本地仓库
pig ext add pg_duckdb postgis pgvector
```

------

### `ext link`

创建符号链接，使特定 PostgreSQL 版本在系统 PATH 中活动。

```bash
sudo pig ext link 17             # 链接 PostgreSQL 17
sudo pig ext link pg17           # 替代语法
sudo pig ext link -p /usr/pgsql-16/bin/pg_config  # 通过路径链接
```

**功能：**
1. 在 `/usr/bin/` 中创建 PostgreSQL 二进制文件的符号链接
2. 更新 alternatives 系统（如果可用）
3. 使指定版本成为默认的 `psql`、`pg_dump` 等

**示例：**
```bash
# 链接前
which psql                       # /usr/bin/psql -> /etc/alternatives/pgsql-psql

# 链接 PostgreSQL 17
sudo pig ext link 17

# 链接后
which psql                       # /usr/bin/psql -> /usr/pgsql-17/bin/psql
psql --version                   # psql (PostgreSQL) 17.0
```

------

### `ext reload`

更新扩展目录到最新版本。

```bash
pig ext reload                   # 重新加载目录
```

**功能：**
- 下载最新的扩展元数据
- 更新本地目录缓存
- 刷新可用扩展列表

**使用时机：**
- 主要 PostgreSQL 版本发布后
- 新扩展可用时
- 扩展元数据看起来过时时

------

## 常见工作流

### 工作流 1：设置新 PostgreSQL 安装

```bash
# 1. 设置仓库
sudo pig repo add -ru

# 2. 安装 PostgreSQL 17
sudo pig ext install pg17

# 3. 链接到 PATH
sudo pig ext link 17

# 4. 安装基本扩展
sudo pig ext add pg_stat_statements pg_repack wal2json

# 5. 安装特定应用扩展
sudo pig ext add pg_duckdb postgis pgvector timescaledb

# 6. 验证安装
pig ext status
```

### 工作流 2：在 PostgreSQL 版本间迁移

```bash
# 在旧版本旁边安装新版本
sudo pig ext install pg17

# 为新版本安装相同的扩展
pig ext scan -v 16 > pg16_extensions.txt
sudo pig ext add $(cat pg16_extensions.txt | awk '{print $1}') -v 17

# 切换活动版本
sudo pig ext link 17
```

### 工作流 3：构建开发环境

```bash
# 安装多个 PostgreSQL 版本
sudo pig ext install pg15-devel pg16-devel pg17-devel

# 安装开发扩展
sudo pig ext add pgtap plpgsql_check pg_hint_plan -v 17
sudo pig ext add pgbench_extended pg_stat_kcache -v 17

# 安装语言扩展
sudo pig ext add plpython3 plperl pltcl -v 17
```

### 工作流 4：设置分析栈

```bash
# 安装专注于分析的 PostgreSQL
sudo pig ext install pg17

# OLAP 扩展
sudo pig ext add pg_duckdb parquet_fdw apache_arrow_fdw

# 时间序列
sudo pig ext add timescaledb pg_partman pg_cron

# 统计
sudo pig ext add pg_stat_monitor pg_qualstats pg_wait_sampling

# ML/AI
sudo pig ext add pgvector pgml
```

### 工作流 5：离线安装

```bash
# 在联网机器上：准备包
sudo pig repo add -ru
pig ext import pg17 pg_duckdb postgis pgvector
sudo pig repo create /www/pigsty
sudo pig repo cache

# 将 /tmp/pkg.tgz 传输到离线机器

# 在离线机器上：从缓存安装
sudo pig repo boot
sudo pig repo add local
sudo pig ext install pg17
sudo pig ext add pg_duckdb postgis pgvector
```

------

## 扩展类别指南

### OLAP 和分析
分析工作负载的必需品：
- `pg_duckdb`：用于 OLAP 查询的嵌入式 DuckDB
- `parquet_fdw`：直接查询 Parquet 文件
- `citus`：分布式 PostgreSQL
- `columnar`：列式存储引擎

### GIS 和空间
地理和空间数据处理：
- `postgis`：空间数据库扩展器
- `pgrouting`：路由功能
- `h3`：Uber 的 H3 分层空间索引
- `ogr_fdw`：访问 OGR 数据源

### AI 和机器学习
PostgreSQL 内的 ML 能力：
- `pgvector`：向量相似性搜索
- `pgvectorscale`：可扩展向量搜索
- `pgml`：PostgresML 机器学习
- `pg_similarity`：相似性搜索函数

### 时间序列
时间序列数据管理：
- `timescaledb`：时间序列数据库
- `pg_partman`：分区管理
- `pg_cron`：作业调度器

### 全文搜索
高级文本搜索能力：
- `pgroonga`：使用 Groonga 的全文搜索
- `pg_bigm`：双字母搜索
- `zhparser`：中文文本搜索
- `pg_jieba`：中文分词

### 监控和统计
性能监控和分析：
- `pg_stat_statements`：查询性能跟踪
- `pg_stat_kcache`：内核缓存统计
- `pg_stat_monitor`：高级监控
- `pg_wait_sampling`：等待事件采样
- `pg_qualstats`：查询质量统计

### 安全
安全和审计扩展：
- `pgcrypto`：加密函数
- `pg_audit`：审计日志
- `set_user`：权限提升控制
- `supautils`：Supabase 实用工具

### 开发工具
开发和调试：
- `pgtap`：单元测试框架
- `plpgsql_check`：PL/pgSQL 代码检查器
- `pg_hint_plan`：查询计划提示
- `pgbench`：基准测试工具

------

## 故障排除

### 找不到扩展

```bash
# 更新目录
pig ext reload

# 使用不同的术语搜索
pig ext list <部分名称>

# 检查可用仓库
pig repo status
```

### 版本兼容性问题

```bash
# 检查扩展兼容性
pig ext info <扩展名称>

# 指定正确的 PostgreSQL 版本
sudo pig ext add <扩展> -v <版本>
```

### 缺少依赖

```bash
# 某些扩展需要操作系统包
sudo dnf install <所需包>  # 对于 EL
sudo apt install <所需包>  # 对于 Debian/Ubuntu

# 然后重试扩展安装
sudo pig ext add <扩展>
```

### 找不到 pg_config

```bash
# 首先安装 PostgreSQL
sudo pig ext install pg17

# 或明确指定路径
pig ext add <扩展> -p /usr/pgsql-17/bin/pg_config

# 或链接 PostgreSQL 到 PATH
sudo pig ext link 17
```

### 权限被拒绝

```bash
# 安装需要 sudo
sudo pig ext add <扩展>

# 对于系统级操作
sudo pig ext link 17
```

------

## 最佳实践

1. **总是先设置仓库**：在安装扩展前运行 `pig repo add -ru`

2. **使用版本标志**：管理多个安装时明确指定 PostgreSQL 版本

3. **检查兼容性**：使用 `pig ext info` 验证扩展支持您的 PostgreSQL 版本

4. **组合相关扩展**：一起安装相关扩展以获得更好的依赖解析

5. **定期更新**：定期运行 `pig ext update` 以获取安全和错误修复

6. **目录维护**：偶尔运行 `pig ext reload` 以获取新的扩展列表

7. **重大更改前备份**：更新关键扩展前创建数据库备份

8. **先在开发环境测试**：在生产环境前总是在开发环境测试扩展更新

9. **安装后监控**：安装新扩展后检查日志是否有任何问题

10. **记录您的栈**：保留已安装扩展的记录以用于灾难恢复