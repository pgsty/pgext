
## 用法

> 语法：
>
> ```sql
> CREATE EXTENSION pgbson;
> SELECT bson_get_datetime(bson_column, 'msg.header.event.ts') FROM my_table;
> SELECT (bson_column->'msg'->'header'->'event'->>'ts')::timestamp FROM my_table;
> ```
>
> 来源：[README](https://github.com/buzzm/postgresbson)

`pgbson` 为 PostgreSQL 增加了 BSON 数据类型，以及用于创建、检查和查询 BSON 文档的函数与操作符。上游 README 将其定位为 JSON/JSONB 的二进制、富类型替代方案，具备精确 round-trip、对日期时间和数值子类型的原生支持，以及原始字节支持。

## BSON 的优势

README 强调 BSON 相比 JSON 的几项优势：

- 日期时间是第一类值
- 数值类型保持区分（`int32`、`int64`、`float`、`decimal`）
- 原始字节数组是第一类值
- 往返转换可保留完全一致的二进制表示
- 多种语言的原生 SDK 支持

## 访问方式

扩展提供两种访问风格：

### Dotpath 访问器

这是上游文档中的高性能类型安全访问器：

```sql
SELECT bson_get_datetime(bson_column, 'msg.header.event.ts') FROM my_table;
SELECT bson_get_bson(bson_column, 'msg.header.event') FROM my_table;
```

README 认为它们比反复使用箭头解引用更省内存，因为它们会直接遍历 BSON 结构，只在终点处分配材料化结果。

### 箭头操作符

它也支持类似 JSON 的操作符：

```sql
SELECT (bson_column->'msg'->'header'->'event'->>'ts')::timestamp
FROM my_table;
```

## JSON 互操作

BSON 类型可以通过 Extended JSON（EJSON）转换为 JSON，从而保留类型精度。这使得 BSON 值可以在需要时交给 JSON/JSONB 函数和操作符处理：

```sql
SELECT (bson_get_bson(bson_column, 'msg.header.event')::jsonb) ?& ARRAY['id','type']
FROM my_table;
```

## 说明

README 中给出了跨 Java、Kafka、Python 和 PostgreSQL 的 BSON 往返示例，强调将存储的 BSON 载荷重新转成 `bytea` 后可以做到字节级一致。
