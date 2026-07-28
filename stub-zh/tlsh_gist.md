## 用法

来源：

- [官方上游 README](https://github.com/jinyyu/tlsh_gist/blob/8b028591e961a9ca8d561b34b02817c15f01bf86/README.MD)
- [官方扩展控制文件 (tlsh_gist.control)](https://github.com/jinyyu/tlsh_gist/blob/8b028591e961a9ca8d561b34b02817c15f01bf86/tlsh_gist.control)
- [官方扩展 SQL (tlsh_gist--1.0.sql)](https://github.com/jinyyu/tlsh_gist/blob/8b028591e961a9ca8d561b34b02817c15f01bf86/tlsh_gist--1.0.sql)

`tlsh_gist` — tlsh_gist --------- 一个 PostgreSQL 插件，用于 tlsh 哈希，这是一种模糊匹配程序和库。相似文件将具有相似的哈希值，允许通过比较它们的哈希值来检测相似对象。使用它来进行相应的向量、模型或检索工作流。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION tlsh_gist;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `gtlsh_in(cstring)` 是一个扩展函数，返回 `gtlsh`。
- `gtlsh_out(gtlsh)` 是一个扩展函数，返回 `cstring`。
- `tlsh_compress(internal)` 是一个扩展函数，返回 `internal`。
- `tlsh_consistent(internal, tlsh, smallint, oid, internal)` 是一个扩展函数。
- `tlsh_decompress(internal)` 是一个扩展函数，返回 `internal`。
- `tlsh_dist(tlsh,tlsh)` 是一个扩展函数，返回 `int4`。
- `tlsh_distance(internal, tlsh, smallint, oid, internal)` 是一个扩展函数，返回 `float8`。
- `tlsh_equal(tlsh, tlsh)` 是一个扩展函数。
- `tlsh_in(cstring)` 是一个扩展函数，返回 `tlsh`。
- `tlsh_mean(tlsh, tlsh)` 是一个扩展函数，返回 `tlsh`。
- `tlsh_out(tlsh)` 是一个扩展函数，返回 `cstring`。
- `tlsh_penalty(internal, internal, internal)` 是一个扩展函数，返回 `internal`。
- `tlsh_picksplit(internal, internal)` 是一个扩展函数，返回 `internal`。
- `tlsh_same(gtlsh, gtlsh, internal)` 是一个扩展函数，返回 `internal`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
