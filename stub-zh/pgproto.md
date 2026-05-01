## 用法

来源：[README](https://github.com/Apaezmx/pgproto/blob/v0.5.0/README.md), [release 0.5.0](https://github.com/Apaezmx/pgproto/releases/tag/v0.5.0), [PGXN 0.5.0](https://pgxn.org/dist/pgproto/0.5.0/), [SQL definitions](https://github.com/Apaezmx/pgproto/blob/v0.5.0/sql/pgproto--1.0.sql), [Makefile](https://github.com/Apaezmx/pgproto/blob/v0.5.0/Makefile), [pgproto.control](https://github.com/Apaezmx/pgproto/blob/v0.5.0/pgproto.control)

`pgproto` 在 PostgreSQL 中以原生 `protobuf` 类型存储 Protocol Buffers `proto3` payload，提供 schema 感知提取、更新辅助函数、包含/索引支持，以及文本/整数路径操作符。上游包版本为 `0.5.0`；扩展 SQL/control 默认版本仍为 `1.0`。

当前上游源码是 C/PGXS 扩展：官方 `Makefile` 设置 `MODULE_big = pgproto`，从 `src/*.o` 构建 C object，并包含 `$(PGXS)`。README 描述该实现为纯 C，不依赖外部 Protobuf 库。

```sql
CREATE EXTENSION pgproto;
```

### Schema 注册与存储

`pgproto` 需要运行时 protobuf descriptor，才能用名称/路径提取解释二进制 payload。可以把序列化后的 `FileDescriptorSet` 注册到 `pb_schemas`，也可以在适合的工作流中调用 SQL 注册辅助函数：

```sql
INSERT INTO pb_schemas (name, data)
VALUES ('MySchema', '\x...');

SELECT pb_register_schema('MySchema', '\x...');
```

把序列化后的 protobuf bytes 存入 `protobuf` 列：

```sql
CREATE TABLE items (
  id serial PRIMARY KEY,
  data protobuf
);

INSERT INTO items (data) VALUES ('\x0a02082a');
```

0.5.0 SQL 还安装了从 `protobuf` 到 `bytea` 的便捷 cast，因此需要时可以使用 `length(data::bytea)` 这类面向 byte 的函数。

### 查询

对嵌套、repeated 和 map 字段使用路径操作符：

```sql
-- Integer accessor: returns int4
SELECT data #> '{Outer, inner, id}'::text[] FROM items;

-- Text accessor: returns text
SELECT data #>> '{Outer, tags, mykey}'::text[] FROM items;

-- Array index lookup
SELECT data #> '{Outer, scores, 0}'::text[] FROM items;
```

扩展定义的其他面向用户的提取辅助函数和操作符包括：

- `pb_get_int32(protobuf, int4)`：按 tag 提取 `int4`。
- `pb_get_int32_by_name(protobuf, text, text)` 与 `pb_get_int32_by_name_dot(protobuf, text)`：按名称提取整数。
- `->`：通过 `pb_get_int32_by_name_dot` 进行 dot-path 整数查找的简写。
- `pb_get_int32_by_path(protobuf, text[])`：支撑 `#>`。
- `pb_get_text_by_path(protobuf, text[])`：支撑 `#>>`。
- `pb_to_json(protobuf, text)`：在提供 message name 时转换为 text JSON。

### 更新与合并

`pb_set`、`pb_insert` 和 `pb_delete` 都是纯函数：它们返回新的 `protobuf` 值，因此需要用 `UPDATE ... SET` 持久化修改。上游 0.5.0 文档说明这些 mutation 会自动 compaction，以移除 stale tags。

```sql
UPDATE items
SET data = pb_set(data, ARRAY['Outer', 'a'], '42');

UPDATE items
SET data = pb_insert(data, ARRAY['Outer', 'scores', '0'], '100');

UPDATE items
SET data = pb_insert(data, ARRAY['Outer', 'tags', 'key1'], 'value1');

UPDATE items
SET data = pb_delete(data, ARRAY['Outer', 'a']);
```

使用 `||` 操作符合并两个 protobuf 值，它会调用 `pb_merge`：

```sql
UPDATE items
SET data = data || other.data
FROM other
WHERE items.id = other.id;
```

### 索引与包含

可以对提取字段建立普通表达式索引：

```sql
CREATE INDEX idx_items_pb_id
ON items ((data #> '{Outer, inner, id}'::text[]));

SELECT *
FROM items
WHERE (data #> '{Outer, inner, id}'::text[]) = 42;
```

SQL definitions 还暴露了 protobuf 包含操作符 `@>`，以及用于 GIN 索引的默认 `protobuf_gin_ops` operator class：

```sql
CREATE INDEX idx_items_data_gin
ON items USING gin (data protobuf_gin_ops);

SELECT * FROM items WHERE data @> '\x0a02082a'::protobuf;
```

### Schema 演进

README 将 schema 演进描述为常规用例：新增字段从旧消息读取为 `NULL`，deprecated 或 unknown 字段会在遍历时跳过，enum 按标准 varint 读取，未设置的 `oneof` 字段返回 `NULL`。

### 注意事项

- schema 感知的路径导航需要运行时 schema；没有已注册 descriptor 时，扩展无法解析 message field name。
- `#>` 返回 `int4`，`#>>` 返回 `text`；应按预期字段类型选择操作符/函数。
- mutation 辅助函数不会原地更新行；必须把返回值赋回列。
- README 中的 benchmark 数字是上游项目 benchmark，不是独立性能保证。
