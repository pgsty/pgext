## 用法

来源：

- [Official README](https://github.com/glynastill/jsonb_delete/blob/f87d296d725ae1ba207a3b65335967c4444c8d6d/README.md)
- [Version 1.0 SQL API](https://github.com/glynastill/jsonb_delete/blob/f87d296d725ae1ba207a3b65335967c4444c8d6d/jsonb_delete--1.0.sql)
- [C deletion implementation](https://github.com/glynastill/jsonb_delete/blob/f87d296d725ae1ba207a3b65335967c4444c8d6d/jsonb_delete.c)
- [Official regression examples](https://github.com/glynastill/jsonb_delete/blob/f87d296d725ae1ba207a3b65335967c4444c8d6d/sql/jsonb_delete.sql)

`jsonb_delete` 1.0 添加 `jsonb_delete(jsonb, jsonb)` 和二元 `jsonb - jsonb` 操作符。对于对象，只有键和完整值都与右侧对象匹配时，才删除对应顶层条目；对于数组，则删除与右侧数组元素相等的顶层元素。

### 对象差集

```sql
CREATE EXTENSION jsonb_delete;

SELECT '{"a":1,"b":2,"nested":{"x":3}}'::jsonb
       - '{"a":9,"b":2,"nested":{"x":4}}'::jsonb;
-- {"a": 1, "nested": {"x": 3}}

SELECT jsonb_delete(
  '{"a":1,"b":2}'::jsonb,
  '{"b":2}'::jsonb
);
```

只改变右侧的值并不足以删除键。这与 PostgreSQL 核心的 `jsonb - text` 等按键删除形式不同。

### 数组差集

```sql
SELECT '["a",2,{"x":1}]'::jsonb
       - '["a",{"x":1}]'::jsonb;
-- [2]
```

### 使用边界

删除不会递归。只有整个顶层嵌套对象或数组相等时才会删除；保留值内部的匹配后代不会被修改。函数声明为 `STRICT`，因此 SQL NULL 参数返回 SQL NULL，而 JSON `null` 会正常参与 JSON 相等比较。

该 API 创建于 PostgreSQL 9.4/9.5 时期，早于大多数核心 JSONB 修改操作符，并使用 PostgreSQL 内部 JSONB C 结构。应在每个目标大版本上构建并运行回归测试。由于它添加了 `jsonb - jsonb` 重载，当其他 JSONB 扩展也定义相关重载时，应审计应用中的操作符解析结果。
