## 用法

来源：

- [官方扩展控制文件（PandaPost.control）](https://api.pgxn.org/src/pandapost/pandapost-0.2.0/PandaPost.control)
- [官方扩展 SQL（PandaPost.sql）](https://api.pgxn.org/src/pandapost/pandapost-0.2.0/sql/PandaPost.sql)

`pandapost` — Python Pandas 数据在 Postgres 中。当需要移动、转换或集成相应的数据时，请使用此扩展。经过审核的上游材料已将此功能标记为弃用。

### 核心工作流

```sql
CREATE EXTENSION pandapost;
```

在目标数据库中安装扩展，当可用时运行最小的上游示例，并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

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

- 经过审核的控制文件声明默认版本为 `0.2.0`。
- 控制文件将扩展标记为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源。
