## 用法

来源：

- [官方上游 README](https://github.com/myyrakle/oddments/blob/e1a3365fc154e44951b99395144d399fde5572cb/SQL_Boilerplates/PostgreSQL/extensions/hello-with-ffi/README.md)
- [官方扩展控制文件 (simpleext.control)](https://github.com/myyrakle/oddments/blob/e1a3365fc154e44951b99395144d399fde5572cb/SQL_Boilerplates/PostgreSQL/extensions/hello-with-ffi/simpleext.control)
- [官方扩展 SQL (simpleext--1.0.sql)](https://github.com/myyrakle/oddments/blob/e1a3365fc154e44951b99395144d399fde5572cb/SQL_Boilerplates/PostgreSQL/extensions/hello/simpleext--1.0.sql)

`simpleext` — 这是一个可以直接在当前项目中构建的最简单的 C++ PostgreSQL 扩展模板。当应用程序需要这种特定的数据库功能时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION simpleext;
   SELECT simpleext.hello_world(1); -- 2 (입력값 + 1)
   SELECT simpleext.add_ints(2, 3); -- 5
```

在目标数据库中安装扩展，如果有可用的上游最小示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

### 重要对象

- `simpleext.add_ints(a int, b int)` 是一个扩展函数，返回 `int`。
- `simpleext.hello_world()` 是一个扩展函数，返回 `text`。
- `simpleext` 是由扩展创建的一个模式。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况。
