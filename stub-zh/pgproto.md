## 用法

来源：[README](https://github.com/Apaezmx/pgproto/blob/main/README.md), [release 0.3.3](https://github.com/Apaezmx/pgproto/releases/tag/v0.3.3), [pgproto.control](https://github.com/Apaezmx/pgproto/blob/main/pgproto.control)

`pgproto` 增加了 `protobuf` 类型，用于存储二进制 Protocol Buffers，并提供感知 schema 的提取与更新辅助函数。上游最新 release 是 `0.3.3`，而扩展 control 文件声明的 SQL 默认版本为 `1.0`。

```sql
CREATE EXTENSION pgproto;
```

### 基本工作流

先注册 `FileDescriptorSet`，让扩展能够解释消息布局：

```sql
INSERT INTO pb_schemas (name, data) VALUES ('MySchema', '\x...');
```

再把 protobuf payload 存入 `protobuf` 列：

```sql
CREATE TABLE items (
  id serial PRIMARY KEY,
  data protobuf
);
```

### 查询

对嵌套字段使用 PostgreSQL 风格路径操作符：

```sql
SELECT data #> '{Outer, inner, id}'::text[] FROM items;
SELECT data #> '{Outer, tags, mykey}'::text[] FROM items;
```

README 将 `->` 与 `#>` 作为导航嵌套字段、repeated 字段与 map 字段的主要操作符。

### 更新与合并

写辅助函数都是纯函数，会返回新的 protobuf 值：

- `pb_set(...)`
- `pb_insert(...)`
- `pb_delete(...)`
- `||` 用于合并两个同类型消息

```sql
UPDATE items SET data = pb_set(data, ARRAY['Outer', 'a'], '42');
UPDATE items SET data = pb_insert(data, ARRAY['Outer', 'scores', '0'], '100');
UPDATE items SET data = pb_delete(data, ARRAY['Outer', 'a']);
UPDATE items SET data = data || other_data;
```

### 索引与 Schema 演进

可以对提取字段建立表达式索引：

```sql
CREATE INDEX idx_pb_id ON items ((data #> '{Outer, inner, id}'::text[]));
```

README 也把 schema 演进作为一等场景：新增字段向后兼容，已废弃字段在旧 payload 中仍可读取，而用 `ON CONFLICT` 重新注册 schema 是预期更新路径。

### 注意事项

- `pgproto` 依赖已注册的运行时 schema；没有 descriptor set 时，路径提取无法解释 payload。
- 更新辅助函数不会原地修改值，因此必须在 `UPDATE ... SET data = ...` 里使用。
