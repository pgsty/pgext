## 用法

来源：

- [官方上游 README](https://github.com/jac30b/pgv/blob/75fc2437b607af80726ac0b2489d8ea30ab39c52/README.md)
- [官方扩展控制文件 (pgv.control)](https://github.com/jac30b/pgv/blob/75fc2437b607af80726ac0b2489d8ea30ab39c52/pgv.control)
- [官方扩展 SQL (pgv--0.0.1.sql)](https://github.com/jac30b/pgv/blob/75fc2437b607af80726ac0b2489d8ea30ab39c52/sql/pgv--0.0.1.sql)

`pgv` — 本扩展基于 pgvector 和 PostgreSQL 扩展进行学习。使用它来进行相应的向量、模型或检索工作流。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pgv;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `vec_cmp(vec, vec)` 是一个扩展函数，返回 `integer`。
- `vec_cosine_distance(vec, vec)` 是一个扩展函数，返回 `float4`。
- `vec_eq(vec, vec)` 是一个扩展函数，返回 `boolean`。
- `vec_ge(vec, vec)` 是一个扩展函数，返回 `boolean`。
- `vec_gt(vec, vec)` 是一个扩展函数，返回 `boolean`。
- `vec_input(cstring, oid, integer)` 是一个扩展函数，返回 `vec`。
- `vec_le(vec, vec)` 是一个扩展函数，返回 `boolean`。
- `vec_lt(vec, vec)` 是一个扩展函数，返回 `boolean`。
- `vec_ne(vec, vec)` 是一个扩展函数，返回 `boolean`。
- `vec_output(vec)` 是一个扩展函数，返回 `cstring`。
- `vec_typemodifier_in(cstring[])` 是一个扩展函数，返回 `integer`。
- `vec_typemodifier_out(integer)` 是一个扩展函数，返回 `cstring`。
- `vec` 是一个扩展定义的类型。
- `vec_ops` 是一个扩展定义的操作类。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.0.1`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
