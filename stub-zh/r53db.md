## 用法

来源：

- [官方上游 README](https://github.com/apparentorder/r53db/blob/8c66a3f954918c2d292700b7494f5315506b169d/README.md)
- [官方扩展控制文件 (r53db.control)](https://github.com/apparentorder/r53db/blob/8c66a3f954918c2d292700b7494f5315506b169d/r53db.control)
- [官方扩展 SQL (r53db--0.1.sql)](https://github.com/apparentorder/r53db/blob/8c66a3f954918c2d292700b7494f5315506b169d/r53db--0.1.sql)

`r53db` — *r53db* 是一个 PostgreSQL 的外部数据封装器，可以让你像访问 SQL 表一样访问 AWS Route53 数据库区域。当 PostgreSQL 需要通过外部数据接口访问相应的外部数据源时，请使用此扩展。在目标 PostgreSQL 构建中测试链接的上游修订版本作为 API 边界。

### 核心工作流

```sql
CREATE EXTENSION r53db;
CREATE SERVER route53 FOREIGN DATA WRAPPER r53db;
CREATE SCHEMA route53;
IMPORT FOREIGN SCHEMA dummy FROM SERVER route53 INTO route53;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `r53db_fdw_handler()` 是一个扩展函数，返回 `fdw_handler`。
- `r53db` 是一个由扩展定义的外部数据封装器。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.1`。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
