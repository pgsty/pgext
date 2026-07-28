## 用法

来源：

- [官方上游 README](https://github.com/kolharsam/pg_bm25/blob/2e6a875874e0adb2f1e8d735059c66e208e2ebc8/docs/src/README.md)
- [官方扩展控制文件 (parodydb.control)](https://github.com/kolharsam/pg_bm25/blob/2e6a875874e0adb2f1e8d735059c66e208e2ebc8/parodydb.control)
- [官方实现源代码](https://github.com/kolharsam/pg_bm25/blob/2e6a875874e0adb2f1e8d735059c66e208e2ebc8/src/lib.rs)

`parodydb` — 一个用 Rust 和 pgrx 编写的 PostgreSQL 玩具全文搜索扩展。使用它来进行相应的文本搜索、解析或语言工作流。使用上述链接的上游版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION parodydb;

SELECT parodydb_search('The quick brown fox', 'quick');  -- true
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `hello_parodydb()` 是一个扩展函数。
- `parodydb_index` 是一个扩展函数。
- `parodydb_index_clear()` 是一个扩展函数。
- `parodydb_index_info()` 是一个扩展函数。
- `parodydb_search` 是一个扩展函数。
- `parodydb_tokenize` 是一个扩展函数。

### 要求与注意事项

- 该目录记录版本为 `0.0.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 控制文件将扩展标记为不受信任。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
