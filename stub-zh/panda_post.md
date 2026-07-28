## 用法

来源：

- [官方扩展控制文件 (panda_post.control)](https://api.pgxn.org/src/panda_post/panda_post-0.2.1/panda_post.control)
- [官方扩展 SQL (panda_post.sql)](https://api.pgxn.org/src/panda_post/panda_post-0.2.1/sql/panda_post.sql)

`panda_post` — Python Pandas 数据在 Postgres 中。当需要移动、转换或集成相应的数据到 PostgreSQL 时，请使用此扩展。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION panda_post;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `create_cast(data_type text , transform text DEFAULT '' , cast_type text DEFAULT NULL , create_array_cast boolean DEFAULT true)` 是一个扩展函数，返回 `void`。
- `eval(i text)` 是一个扩展函数，返回 `ndarray`。
- `ndall(i ndarray , axis int , keepdims boolean=False)` 是一个扩展函数，返回 `ndarray`。
- `ndall(i ndarray , axis int[] , keepdims boolean=False)` 是一个扩展函数，返回 `ndarray`。
- `ndall(i ndarray , keepdims boolean=False)` 是一个扩展函数，返回 `ndarray`。
- `ndany(i ndarray , axis int , keepdims boolean=False)` 是一个扩展函数，返回 `ndarray`。
- `ndany(i ndarray , axis int[] , keepdims boolean=False)` 是一个扩展函数，返回 `ndarray`。
- `ndany(i ndarray , keepdims boolean=False)` 是一个扩展函数，返回 `ndarray`。
- `ndarray_from_plpython(internal)` 是一个扩展函数，返回 `ndarray`。
- `ndarray_in(cstring)` 是一个扩展函数，返回 `ndarray`。
- `ndarray_out(ndarray)` 是一个扩展函数，返回 `cstring`。
- `ndarray_to_plpython(internal)` 是一个扩展函数，返回 `internal`。
- `ndunique(ar ndarray , return_index boolean = False , return_inverse boolean = False , return_counts boolean = False)` 是一个扩展函数，返回 `ndarray[]`。
- `ndunique1(ar ndarray)` 是一个扩展函数，返回 `ndarray`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.2.1`。
- 控制文件将扩展标记为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
