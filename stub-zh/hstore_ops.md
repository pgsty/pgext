## 用法

来源：

- [官方 README](https://github.com/postgrespro/hstore_ops/blob/a9b00d87ee32171a7f7ef46101cfad9e90b0ff2e/README.md)
- [扩展 control 文件](https://github.com/postgrespro/hstore_ops/blob/a9b00d87ee32171a7f7ef46101cfad9e90b0ff2e/hstore_ops.control)
- [版本 1.1 扩展 SQL](https://github.com/postgrespro/hstore_ops/blob/a9b00d87ee32171a7f7ef46101cfad9e90b0ff2e/hstore_ops--1.1.sql)

`hstore_ops` 为 PostgreSQL `hstore` 类型增加两种替代 GIN 操作符类。当以包含查询为主的工作负载需要比内置 GIN 类更小或更快的索引，或希望通过逐字节比较避免昂贵的排序规则处理时，可以使用它。

### 核心流程

```sql
CREATE EXTENSION hstore;
CREATE EXTENSION hstore_ops;

CREATE TABLE item_attributes (
    item_id bigint PRIMARY KEY,
    attrs hstore NOT NULL
);

CREATE INDEX item_attrs_hash_idx
ON item_attributes USING gin (attrs gin_hstore_hash_ops);

SELECT item_id
FROM item_attributes
WHERE attrs @> 'color=>blue'::hstore;
```

另一种逐字节操作符类的用法如下：

```sql
CREATE INDEX item_attrs_bytea_idx
ON item_attributes USING gin (attrs gin_hstore_bytea_ops);
```

两种操作符类都支持 `@>`、`?`、`?|` 与 `?&`。

### 选择操作符类

- `gin_hstore_hash_ops` 存储键和值的哈希。上游报告称其索引比默认 GIN 类更小，`@>` 搜索更快，但键存在性操作符可能略慢。
- `gin_hstore_bytea_ops` 在版本 `1.1` 中加入，是标准 `hstore` GIN 类的变体：它逐字节比较键，而不使用排序规则键。目标是在排序规则比较较慢时提升索引处理性能，同时不改变支持的操作符。

`gin_hstore_hash_ops` 可能发生哈希冲突，因此 PostgreSQL 必须复查候选行。应在代表性数据上测量索引大小、构建时间、更新成本与各种查询形态，而不要假设某一种类始终更快。

### 依赖与维护

control 文件声明依赖 `hstore` 扩展；不需要预加载或重启。把现有 `hstore_ops` `1.0` 安装升级到 `1.1` 会加入 `gin_hstore_bytea_ops`：

```sql
ALTER EXTENSION hstore_ops UPDATE TO '1.1';
```

将现有索引改为另一种操作符类需要重建索引。锁与磁盘空间影响应按大型 GIN 索引构建来规划。
