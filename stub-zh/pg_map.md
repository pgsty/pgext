## 用法

来源：

- [官方 pg_map README](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_map/README.md)
- [版本 3.4 控制文件](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_map/pg_map.control)
- [基础 SQL 定义](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_map/pg_map--1.2.sql)
- [官方扩展测试和示例](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_map/tests/pytests/extension_test.py)

`pg_map` 从 PostgreSQL 类型生成强类型键值对映射域。生成的映射是一个复合键值对数组，具有特定类型的提取、基数、条目和操作符函数。它被 pg_lake 用于嵌套数据，也可以直接使用。

### 创建并使用一个映射类型

```sql
CREATE EXTENSION pg_map;

-- Requires a privileged role; PUBLIC execution is revoked.
SELECT map_type.create('text', 'integer');
-- map_type.key_text_val_int
```

构造一个值并通过生成的函数或 `->` 操作符读取：

```sql
SELECT map_type.extract(
    '{"(me,1)","(myself,2)","(i,3)"}'::map_type.key_text_val_int,
    'i'
);
-- 3

SELECT
    '{"(me,1)","(myself,2)","(i,3)"}'::map_type.key_text_val_int
    -> 'myself';
-- 2

SELECT key, value
FROM map_type.entries(
    '{"(me,1)","(myself,2)","(i,3)"}'::map_type.key_text_val_int
);
```

### 生成的 API

- `map_type.create(keytype regtype, valtype regtype, typname text default null)`: 以幂等方式创建或返回键/值对的映射类型。
- `map_type.extract(map, key)` 和生成的 `map -> key`：返回指定键的值。
- `map_type.cardinality(map)`: 返回条目数量。
- `map_type.entries(map)`: 将映射扩展为 `key, value` 行。
- 生成的名称通常遵循 `map_type.key_<keytype>_val_<valuetype>`；当需要受控名称时，请提供 `typname`。

### 类型和生命周期注意事项

- 数组类型不能用作映射键。支持数组值和嵌套生成的映射类型。
- 调用 `map_type.create` 会创建 PostgreSQL 类型、函数和操作符。将其视为模式 DDL 并在迁移中运行，而不是每次请求代码中运行。
- 生成的对象作为 `pg_map` 的依赖项进行注册；当使用 `CASCADE` 删除扩展时，它们及其依赖的列可能会被移除。
- 映射值使用 PostgreSQL 复合数组语法。重复键和排序语义应根据应用程序选择的构建路径进行测试，而不是假设来自 JSON 对象。
- 版本 `3.4` 相对于 `3.3` 没有更改映射 SQL API。
