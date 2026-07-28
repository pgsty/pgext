## 用法

来源：

- [官方上游 README](https://github.com/zombodb/postgresconf/blob/f72a1f01f408b19076d20e182568908ee1f6f873/README.md)
- [官方扩展控制文件 (postgresconf.control)](https://github.com/zombodb/postgresconf/blob/f72a1f01f408b19076d20e182568908ee1f6f873/postgresconf.control)
- [官方实现源代码](https://github.com/zombodb/postgresconf/blob/f72a1f01f408b19076d20e182568908ee1f6f873/src/lib.rs)

`postgresconf` — 首先，你需要以下软件来构建此扩展。当应用程序需要此特定数据库功能时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION postgresconf;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `array_of_names()` 是一个扩展函数。
- `array_of_names_with_null()` 是一个扩展函数。
- `hello_postgresconf()` 是一个扩展函数。
- `my_generate_series` 是一个扩展函数。
- `rust_tuple` 是一个扩展函数。
- `set_of_animals()` 是一个扩展函数。
- `sum_array` 是一个扩展函数。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
