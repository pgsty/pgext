## 用法

来源：

- [postgresbson README at the 2.0.4 revision](https://github.com/buzzm/postgresbson/blob/ec71d314511d484a99ed510f480919dd0509fbe9/README.md)
- [META.json version 2.0.4](https://github.com/buzzm/postgresbson/blob/ec71d314511d484a99ed510f480919dd0509fbe9/META.json)
- [pgbson control file](https://github.com/buzzm/postgresbson/blob/ec71d314511d484a99ed510f480919dd0509fbe9/pgbson.control)
- [Version 2.0 SQL API](https://github.com/buzzm/postgresbson/blob/ec71d314511d484a99ed510f480919dd0509fbe9/pgbson--2.0.sql)

pgbson 添加了 BSON 数据类型、带类型的路径访问器、JSON 风格的操作符、转换以及表达式索引支持。当需要存储二进制 BSON 而不先将其每个值转换为 JSONB 时，请使用此扩展，特别是在保持 BSON 类型精度或进行字节级往返传输时。

分发版本是 2.0.4，而扩展控制和 SQL API 版本仍为 2.0。

### 创建扩展

    CREATE EXTENSION pgbson;

该本地模块依赖于 libbson。请安装一个针对兼容 PostgreSQL 和 libbson 版本构建的包。

### 存储和验证 BSON

当写入值时，bytea 到 bson 的转换会验证输入。版本 2.0.4 文档指出，在读取时可以假设存储的 BSON 是有效的。不要通过不安全的低级写操作绕过类型或转换路径。

### 提取值

带类型的点路径访问器可避免生成每个中间对象：

    SELECT bson_get_datetime(payload, 'msg.header.event.ts'),
           bson_get_string(payload, 'data.customer.name')
    FROM events;

使用 bson_get_bson 获取子文档：

    SELECT bson_get_bson(payload, 'msg.header.event')
    FROM events;

JSON 风格的导航也适用：

    SELECT payload->'msg'->'header'->'event'->>'ts'
    FROM events;

### 函数和操作符索引

- bson_get_string, bson_get_int32, bson_get_int64, bson_get_double, bson_get_decimal：带类型的标量访问器。
- bson_get_datetime, bson_get_binary, bson_get_boolean：用于其他 BSON 类型的访问器。
- bson_get_bson：返回嵌入的 BSON 文档。
- bson_get_jsonb_array：将数组端点转换为 PostgreSQL jsonb 数组。
- -> 和 ->>：使用 JSON 风格语法导航值。
- bson 转换到 json 和 jsonb：暴露扩展的 JSON 用于 PostgreSQL 的 JSON 处理。
- bson 和 bytea 转换：保留 BSON 的二进制表示形式。

### 索引和互操作

在频繁查询路径上创建表达式索引：

    CREATE INDEX events_customer_id_idx
    ON events (bson_get_string(payload, 'data.customer.id'));

将子文档转换为 jsonb 以利用 PostgreSQL 的 JSON 操作符：

    SELECT bson_get_bson(payload, 'msg.header')::jsonb ? 'event'
    FROM events;

### 注意事项

- 带类型的获取器仅在端点具有预期的 BSON 类型时才返回有用的数据。在数据摄取代码中明确表示类型期望。
- bson_get_bson 对于标量端点会返回 NULL，因为标量不是一个 BSON 文档。
- 点路径访问器通常比重复提取中的长操作符链更优，因为它避免了中间的 BSON 值。
- BSON 和 JSONB 在类型和排序语义上有所不同。转换可能有用但不是每个 BSON 工作流程的无损替代品。
