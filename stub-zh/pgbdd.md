## 用法

来源：

- [官方上游 README](https://github.com/utwente-db/dubio/blob/b5210d373e24752a5d7e48cb2abdeb1a132df210/pgbdd/README.txt)
- [官方扩展控制文件 (pgbdd.control)](https://github.com/utwente-db/dubio/blob/b5210d373e24752a5d7e48cb2abdeb1a132df210/pgbdd/pgbdd.control)
- [官方扩展 SQL (pgbdd--0.0.1.sql)](https://github.com/utwente-db/dubio/blob/b5210d373e24752a5d7e48cb2abdeb1a132df210/pgbdd/pgbdd--0.0.1.sql)

`pgbdd` — 请参见文件 example-dictionary.sql(.log) 以了解如何使用 'dictionary' 类型的示例。当应用程序需要此类型、域或其操作符时，请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pgbdd;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `add(d dictionary, vardefs text)` 是一个扩展函数，返回 `dictionary`。
- `alg_bdd(alg cstring, expression cstring)` 是一个扩展函数，返回 `bdd`。
- `alternatives(dict dictionary, var cstring)` 是一个扩展函数，返回 `text`。
- `and_accum(internal bdd, next bdd)` 是一个扩展函数，返回 `bdd`。
- `bdd_bytea_in(expression bytea)` 是一个扩展函数，返回 `bdd`。
- `bdd_equal(lhs_bdd bdd,rhs_bdd bdd)` 是一个扩展函数，返回 `BOOLEAN`。
- `bdd_equiv(lhs_bdd bdd,rhs_bdd bdd)` 是一个扩展函数，返回 `BOOLEAN`。
- `bdd_fast_equiv(lhs_bdd bdd,rhs_bdd bdd)` 是一个扩展函数，返回 `BOOLEAN`。
- `bdd_in(expression cstring)` 是一个扩展函数，返回 `bdd`。
- `bdd_out(dict bdd)` 是一个扩展函数，返回 `cstring`。
- `contains(bdd bdd, var cstring, val integer)` 是一个扩展函数，返回 `BOOLEAN`。
- `debug(dict dictionary)` 是一个扩展函数，返回 `text`。
- `del(d dictionary, vardefs text)` 是一个扩展函数，返回 `dictionary`。
- `dictionary_in(dictname cstring)` 是一个扩展函数，返回 `dictionary`。

### 要求与注意事项

- 控制文件声明默认版本为 `0.0.1`。
- 控制文件标记该扩展为不可重定位。
- 控制文件要求超级用户进行安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
