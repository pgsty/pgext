## 用法

来源：

- [官方上游 README](https://github.com/willruggiano/pg_jsonpatch/blob/153c3317bee8707cf076761e352ad6fde5083116/README.md)
- [官方扩展控制文件 (pg_jsonpatch.control)](https://github.com/willruggiano/pg_jsonpatch/blob/153c3317bee8707cf076761e352ad6fde5083116/pg_jsonpatch.control)
- [官方扩展 SQL (pg_jsonpatch--1.0.0.sql)](https://github.com/willruggiano/pg_jsonpatch/blob/153c3317bee8707cf076761e352ad6fde5083116/pg_jsonpatch--1.0.0.sql)

`pg_jsonpatch` — 我尚未为 RFC 的 A.11-A.15 编写测试。请参阅 ./test.sql 以获取更多示例。在 SQL 需要这些特殊函数或聚合时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中对其进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_jsonpatch;

select jsonb_patch(
  '{"foo":{"bar":"baz","waldo":"fred"},"qux":{"corge":"grault"}}',
  '[{"op":"move","from":"/foo/waldo","path":"/qux/thud"}]'
);
                             jsonb_patch
---------------------------------------------------------------------
 {"foo": {"bar": "baz"}, "qux": {"thud": "fred", "corge": "grault"}}
(1 row)
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `jsonb_patch(target jsonb, patches jsonb)` 是一个扩展函数，返回 `jsonb`。
- `jsonb_patch_add(target jsonb, path text[], value jsonb)` 是一个扩展函数，返回 `jsonb`。
- `jsonb_patch_apply(target jsonb, patch jsonb)` 是一个扩展函数，返回 `jsonb`。
- `jsonb_patch_copy(target jsonb, _from text[], path text[])` 是一个扩展函数，返回 `jsonb`。
- `jsonb_patch_move(target jsonb, _from text[], path text[])` 是一个扩展函数，返回 `jsonb`。
- `jsonb_patch_remove(target jsonb, path text[])` 是一个扩展函数，返回 `jsonb`。
- `jsonb_patch_replace(target jsonb, path text[], value jsonb)` 是一个扩展函数，返回 `jsonb`。
- `jsonb_patch_split_path(path text)` 是一个扩展函数，返回 `text[]`。
- `jsonb_patch_test(target jsonb, path text[], value jsonb)` 是一个扩展函数，返回 `jsonb`。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0.0`。
- 控制文件标记该扩展为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源中的信息一致。
