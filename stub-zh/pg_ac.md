## 用法

来源：

- [官方上游 README](https://github.com/enerzed/pg_ac/blob/4f618296d20de7d26f638abc3fc5603a20bffc3d/README.md)
- [官方扩展控制文件 (pg_ac.control)](https://github.com/enerzed/pg_ac/blob/4f618296d20de7d26f638abc3fc5603a20bffc3d/pg_ac.control)
- [官方扩展 SQL (pg_ac--0.1.sql)](https://github.com/enerzed/pg_ac/blob/4f618296d20de7d26f638abc3fc5603a20bffc3d/pg_ac--0.1.sql)

`pg_ac` — PostgreSQL 12 或更高版本的 C 编译器，支持 C99。使用 pg_config 确保路径中可用。用于相应的文本搜索、解析或语言工作流。使用上述链接的固定上游版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_ac;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `ac_add(bigint, text)` 是一个扩展函数，返回 `boolean`。
- `ac_build(text[])` 是一个扩展函数，返回 `bigint`。
- `ac_build(tsvector)` 是一个扩展函数，返回 `bigint`。
- `ac_deserialize(bytea)` 是一个扩展函数，返回 `bigint`。
- `ac_destroy(bigint)` 是一个扩展函数，返回 `boolean`。
- `ac_fini()` 是一个扩展函数，返回 `boolean`。
- `ac_init()` 是一个扩展函数，返回 `boolean`。
- `ac_match(bigint, text)` 是一个扩展函数，返回 `integer[]`。
- `ac_rank_simple(bigint, text)` 是一个扩展函数，返回 `real`。
- `ac_remove(bigint, text)` 是一个扩展函数，返回 `boolean`。
- `ac_search(bigint, text)` 是一个扩展函数，返回 `boolean`。
- `ac_search(bigint, tsquery)` 是一个扩展函数，返回 `boolean`。
- `ac_serialize(bigint)` 是一个扩展函数，返回 `bytea`。

### 要求与注意事项

- 控制文件声明默认版本为 `0.1`。
- 控制文件标记该扩展为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况。
