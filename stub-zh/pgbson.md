## 用法

来源：

- [pgbson 2.1.0 README](https://api.pgxn.org/src/bson/bson-2.1.0/README.md)
- [pgbson 2.1 控制文件](https://api.pgxn.org/src/bson/bson-2.1.0/pgbson.control)
- [pgbson 2.1 SQL API](https://api.pgxn.org/src/bson/bson-2.1.0/pgbson--2.1.sql)

`pgbson` 添加了 BSON 数据类型、带类型的点路径访问器、JSON 风格的导航、类型转换、比较操作符，以及 btree/hash 索引。当二进制往返保真度或 BSON 特有的标量类型至关重要时，请使用 BSON；如果主要需求是 PostgreSQL 原生 JSON 索引，请使用 `jsonb`。PGXN 发行版本为 `2.1.0`，而 SQL 扩展版本为 `2.1`。

### 安装并存储 BSON

```sql
CREATE EXTENSION pgbson;
SELECT pgbson_version();

CREATE TABLE events (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  payload bson NOT NULL
);

INSERT INTO events (payload)
VALUES ('{"user":{"name":"Ada"},"attempt":3}'::jsonb::bson);
```

本地模块依赖 `libbson`。隐式的 `bytea` 到 `bson` 转换会验证 BSON 输入，而反向转换会保留二进制表示。

### 提取值

带类型的访问器无需物化每一层中间文档：

```sql
SELECT bson_get_string(payload, 'user.name'),
       bson_get_int32(payload, 'attempt')
FROM events;
```

其他带类型的 getter 覆盖 64 位整数、双精度数、十进制数、日期时间、二进制值、布尔值、嵌套 BSON 文档和 JSONB 数组。路径缺失或类型不匹配时返回 `NULL`；如果必须区分这些情况，请在摄取数据时验证预期的 BSON 模式。

版本 2.1 新增了与类型无关的终端提取器：

```sql
SELECT bson_get_value(payload, 'user.name')
FROM events;
-- { "_" : "Ada" }
```

`bson_get_value` 始终将选中的标量、数组或文档包装在键 `_` 下。调用方应只移除这一层包装。该函数有意不提供可链式使用的 `->` 等价形式。

### 导航、比较与索引

```sql
SELECT payload->'user'->>'name'
FROM events;

CREATE INDEX events_user_name_idx
ON events (bson_get_string(payload, 'user.name'));

CREATE INDEX events_payload_btree_idx ON events (payload);
CREATE INDEX events_payload_hash_idx ON events USING hash (payload);
```

版本 2.1 提供逻辑比较操作符 `=`、`<>`、`<`、`<=`、`>` 和 `>=`；`==` 与 `<<>>` 分别执行二进制相等和不等比较。默认 btree 操作符类使用 BSON 逻辑比较，而 hash 操作符类使用二进制相等。字段顺序或字节完全一致性有影响时，应有意识地选择。

### 升级与注意事项

```sql
ALTER EXTENSION pgbson UPDATE TO '2.1';
```

- 安装 2.1 共享库不会更新已有 2.0 扩展的 SQL 对象；安装文件后应执行扩展更新。
- 2.1 共享库修复了 `bson_get_bson()` 或 `->` 解析到标量端点时导致后端崩溃的问题。即使应用尚未使用新增的 2.1 SQL 函数，也应替换早期二进制文件。
- BSON 到 JSON/JSONB 的转换使用 Extended JSON。BSON 与 JSONB 的类型、相等和排序语义不同，因此这种转换并非对所有工作流都无损。
- 在 2.1 中，BSON 日期时间上的 `->>` 会包含末尾的 `Z`；`bson_get_datetime()` 保持不变。请检查会比较旧文本格式的客户端。
- BSON 顶层值是文档，不能是裸数组或标量。`bson_get_value` 使用 `_` 包装，以便在该限制下返回任意嵌套形态。
