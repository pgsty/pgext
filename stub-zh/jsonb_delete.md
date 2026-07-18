## 用法

来源：

- [已复核 commit 的 jsonb_delete README](https://github.com/glynastill/jsonb_delete/blob/f87d296d725ae1ba207a3b65335967c4444c8d6d/README.md)
- [已复核 commit 的 jsonb_delete 1.0 安装 SQL](https://github.com/glynastill/jsonb_delete/blob/f87d296d725ae1ba207a3b65335967c4444c8d6d/jsonb_delete--1.0.sql)
- [已复核 commit 的 jsonb_delete 回归 SQL](https://github.com/glynastill/jsonb_delete/blob/f87d296d725ae1ba207a3b65335967c4444c8d6d/sql/jsonb_delete.sql)

1.0 版 `jsonb_delete` 增加 `jsonb_delete(jsonb, jsonb)` 及对应的二元 `-` 运算符。它只删除左值顶层中那些键和值都与右值条目完整匹配的项目。

```sql
CREATE EXTENSION jsonb_delete;

SELECT '{"a": 1, "b": 2, "nested": {"x": 3}}'::jsonb
       - '{"a": 9, "b": 2, "nested": {"x": 4}}'::jsonb;

SELECT jsonb_delete(
  '{"a": 1, "b": 2}'::jsonb,
  '{"b": 2}'::jsonb
);
```

第一个结果保留 `a` 和 `nested`，因为它们的值不匹配；`b` 的键和值都匹配，所以被删除。

### 注意事项

- 删除不会递归。只有嵌套值在顶层完整匹配时，整个嵌套对象才会被删除。
- 该 API 早于多数 PostgreSQL 内置 JSONB 修改运算符，开发背景是 PostgreSQL 9.4/9.5；应在目标大版本上验证编译与行为。
- 此运算符比较 JSONB 值，并非只按键名删除；没有核对语义前，不要把它直接替换成内置的按键删除表达式。
