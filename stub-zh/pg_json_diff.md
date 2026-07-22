## 用法

来源：

- [官方 README](https://github.com/KhaledSMQ/pg-jsondiff/blob/95cd668636ce9c88d2f6ac57d92ac5301a81ed3c/README.md)
- [扩展 SQL](https://github.com/KhaledSMQ/pg-jsondiff/blob/95cd668636ce9c88d2f6ac57d92ac5301a81ed3c/pg_json_diff--1.0.sql)
- [C++ 实现](https://github.com/KhaledSMQ/pg-jsondiff/blob/95cd668636ce9c88d2f6ac57d92ac5301a81ed3c/src/json_diff_impl.cpp)

`pg_json_diff` 1.0 提供三个基于 nlohmann/json 的不可变 JSONB 转换函数：生成 RFC 6902 JSON Patch、应用这种补丁，或应用 RFC 7386 合并补丁。它适合明确的文档转换，不是自动冲突解决器，也不是稳定的审计差异格式。

### 核心流程

先生成补丁，保存或检查它，再验证应用补丁后是否得到预期目标：

```sql
CREATE EXTENSION pg_json_diff;

WITH docs AS (
  SELECT '{"name":"Ada","age":30}'::jsonb AS old_doc,
         '{"name":"Ada","age":31,"city":"London"}'::jsonb AS new_doc
), patch AS (
  SELECT old_doc, new_doc, generate_json_diff(old_doc, new_doc) AS operations
  FROM docs
)
SELECT operations, apply_patch(old_doc, operations) = new_doc AS verified
FROM patch;
```

应用外部提供的补丁前应先验证，并在事务及应用级授权保护下执行转换。

### 函数索引

- `generate_json_diff(source, target)` 返回 RFC 6902 操作数组。
- `apply_patch(document, patch)` 支持添加、删除、替换、移动、复制与测试操作。
- `merge_patch(document, patch)` 递归合并对象；空值成员删除键，数组则整体替换而不是合并。

三个函数都接受并返回 `jsonb`。

### 运维说明

每次调用都会把 PostgreSQL JSONB 序列化为文本，解析成 C++ JSON 值，执行操作，再次序列化并重新解析为 JSONB。因此，大型或深层嵌套文档会产生复制与解析开销。nlohmann/json 与 PostgreSQL JSONB 的数值模型不同，应测试超大整数和高精度小数的保真度。补丁错误会以通用内部错误返回；生成的操作顺序与路径也不应当作规范化的人类审计记录。
