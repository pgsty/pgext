


## 用法

来源：

- [PGXN biscuit 2.4.1](https://pgxn.org/dist/biscuit/2.4.1/)
- [Biscuit README](https://github.com/CrystallineCore/Biscuit)
- [Biscuit CHANGELOG](https://github.com/CrystallineCore/Biscuit/blob/main/CHANGELOG.md)
- [Biscuit documentation](https://biscuit.readthedocs.io/)

`biscuit` 是 PostgreSQL 的索引访问方法，用于加速文本上的 `LIKE`、`NOT LIKE`、`ILIKE` 和 `NOT ILIKE` 模式匹配。它使用类似位图的位置索引，避免 trigram 搜索中常见的 heap recheck 开销，并支持多列索引，适合通配符查询较多的负载。

PGXN 2.4.1 包里携带的 SQL/control 版本是 `2.4.0`，因此扩展可见的 `default_version` 仍是 `2.4.0`。Pigsty 本地扩展名为 `biscuit`，旧包元数据中可能仍能看到 `pg_biscuit`。

> [!WARNING]
> 上游说明 Biscuit 仍处于活跃开发阶段，建议在生产使用前进行充分的 staging 验证。应覆盖代表性数据集、查询模式、升级、备份恢复和性能行为，再用于关键负载。

### 快速开始

```sql
CREATE EXTENSION IF NOT EXISTS biscuit;

CREATE TABLE users (
  id bigserial PRIMARY KEY,
  name text,
  email text,
  bio text
);

CREATE INDEX users_name_biscuit
ON users USING biscuit (name);

SELECT *
FROM users
WHERE name LIKE '%john%';
```

`biscuit` 支持带 `%` 和 `_` 的普通通配符模式：

```sql
SELECT * FROM users WHERE name LIKE 'john%';
SELECT * FROM users WHERE name LIKE '%smith';
SELECT * FROM users WHERE name LIKE '%oh_';
SELECT * FROM users WHERE name ILIKE '%john%';
SELECT * FROM users WHERE name NOT LIKE '%test%';
```

### 多列索引

```sql
CREATE INDEX users_search_biscuit
ON users USING biscuit (name, email, bio);

SELECT *
FROM users
WHERE name ILIKE '%john%'
  AND email LIKE '%example.com'
  AND bio NOT LIKE '%inactive%';
```

Biscuit 可以合并多个索引列上的 bitmap 匹配，并可能按估计选择性重排谓词。

### 表达式索引

2.4.0 增加了 expression index 支持：

```sql
CREATE INDEX users_lower_name_biscuit
ON users USING biscuit (lower(name));

SELECT *
FROM users
WHERE lower(name) LIKE '%john%';
```

对于 `char(n)` / `bpchar` 列，上游建议使用显式 cast 到 `text` 的表达式索引，因为原生 `bpchar` 操作符类尚不可用：

```sql
CREATE INDEX legacy_code_biscuit
ON legacy_table USING biscuit ((code::text));
```

### 查看信息

```sql
SELECT *
FROM biscuit_operators;

SELECT *
FROM biscuit_version_history;
```

`biscuit_operators` 视图列出 Biscuit 访问方法注册的操作符。2.4.0 修复了该视图，使它在后续增加 operator class / family 时仍能保持正确。

### 运维说明

Biscuit 的设计重点是：

- prefix、suffix、substring 和混合通配符 `LIKE` / `ILIKE` 模式
- 多列谓词中通过 bitmap 交集减少候选集
- 不依赖 trigram false-positive recheck 的精确模式匹配
- 文本模式搜索占主要延迟的查询

它不是通用全文检索引擎，也不替代排序、词干化、分词或短语检索。如果需要这些语义，应使用 PostgreSQL 全文检索、trigram 索引或专门的搜索扩展。

### 注意事项

- 上游要求 PostgreSQL 16 或更高版本及标准构建工具。Pigsty 本地元数据目前为 PostgreSQL 16-18 打包 Biscuit。
- PGXN 包版本 2.4.1 携带的 SQL/control `default_version = '2.4.0'`；这是当前源码包的预期状态。
- Biscuit 只面向 `LIKE` / `ILIKE` 风格的通配符匹配；正则表达式不是它支持的搜索界面。
- 如需索引非 text 列，应使用显式 text 表达式。
- 替换现有生产索引前，应基于真实数据分布与 `pg_trgm` 做基准对比。
