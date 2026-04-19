## 用法

来源：[README](https://github.com/buzzm/postgresbson/blob/master/README.md), [META.json 2.0.2](https://github.com/buzzm/postgresbson/blob/master/META.json), [pgbson.control](https://github.com/buzzm/postgresbson/blob/master/pgbson.control)

`pgbson` 增加了 BSON 数据类型，以及感知 BSON 的访问器与操作符。上游文档将包版本标为 `2.0.2`，而扩展 control 文件暴露的 SQL 默认版本仍是 `2.0`；这与其打包说明一致，即发行包版本领先于扩展 SQL 版本。

```sql
CREATE EXTENSION pgbson;
```

### 核心访问模式

#### 类型化 dotpath 访问器

类型化 dotpath 访问器会直接遍历 BSON 结构，是上游推荐的快速路径：

```sql
SELECT bson_get_datetime(bson_column, 'msg.header.event.ts') FROM my_table;
SELECT bson_get_bson(bson_column, 'msg.header.event') FROM my_table;
SELECT bson_get_string(bson_column, 'data.payload.product.definition.id') FROM my_table;
```

#### JSON 风格操作符

也支持 JSON 风格的操作符：

```sql
SELECT (bson_column->'msg'->'header'->'event'->>'ts')::timestamp
FROM my_table;
```

### 主要函数与操作符

- 类型化 getter，例如 `bson_get_string`、`bson_get_int32`、`bson_get_int64`、`bson_get_double`、`bson_get_decimal`、`bson_get_datetime`、`bson_get_binary` 与 `bson_get_boolean`。
- `bson_get_bson` 用于返回 BSON 子文档。
- `bson_get_jsonb_array` 适合在路径解析到数组后继续使用原生 `jsonb` 数组操作符。
- 箭头操作符 `->` 与 `->>`，用法接近 PostgreSQL JSON 类型。
- 通过 Extended JSON 转成 `json`/`jsonb`，以保留类型精度。

### 互操作与索引

需要 PostgreSQL JSON 操作符时，可先将 BSON 转成 `jsonb`：

```sql
SELECT (bson_get_bson(bson_column, 'msg.header.event')::jsonb) ?& ARRAY['id', 'type']
FROM my_table;
```

也可以在提取路径上建立表达式索引：

```sql
CREATE INDEX ON data_collection (bson_get_string(data, 'd.recordId'));
```

README 还说明 BSON 值可通过 `bytea` cast 实现字节级 round-trip。

### 注意事项

- dotpath 访问器通常比长链式 `->` 访问更快、更省内存，因为它不会物化中间子结构。
- `bson_get_bson()` 在路径终点是 scalar 时会返回 `NULL`，因为简单标量不是 BSON 文档。
- 上游明确指出，数组处理与错误类型访问器行为的人机工学仍有待改进。
