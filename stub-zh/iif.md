## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/iif/iif-0.0.1/README.md)
- [官方扩展控制文件 (iif.control)](https://api.pgxn.org/src/iif/iif-0.0.1/iif.control)
- [官方扩展 SQL (iif--0.0.1.sql)](https://api.pgxn.org/src/iif/iif-0.0.1/iif--0.0.1.sql)

`iif` — 一个用于向 Postgres 添加函数 _iif_ 的示例扩展。在移植或模拟相应的数据库 API 时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION iif;

SELECT iif(1<0,1,2);
 iif
-----
   2
(1 row)
```

在目标数据库中安装扩展，如果有可用的上游最小示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `iif(boolean, anyelement, anyelement)` 是一个扩展函数，返回 `anyelement`。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.0.1`。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码一致。
