

## 用法

> [pglogical_origin: 从 Postgres 9.4 升级时的兼容性虚拟扩展](https://github.com/2ndQuadrant/pglogical)

`pglogical_origin` 扩展是随 pglogical 提供的兼容性填充。它的存在仅仅是为了方便从 PostgreSQL 9.4 升级，在该版本中复制源追踪由 pglogical 扩展本身处理而非 PostgreSQL 核心。

### 启用

```sql
CREATE EXTENSION pglogical_origin;
```

### 概述

从 PostgreSQL 9.5 开始，复制源追踪成为 PostgreSQL 内置功能（`pg_replication_origin`）。`pglogical_origin` 扩展是一个空/虚拟扩展：

- 防止升级之前依赖它的数据库时出错
- 提供从 PostgreSQL 9.4 上 pglogical 到新版本的平滑迁移路径
- 不包含实际功能 —— 所有源追踪由 PostgreSQL 核心处理

### 何时使用

仅在以下情况下需要此扩展：

- 从使用了 pglogical 的 PostgreSQL 9.4 数据库升级
- 数据库中有对 `pglogical_origin` 扩展的现有引用

对于新安装，不需要此扩展。直接使用 pglogical，它利用了 PostgreSQL 内置的复制源支持。
