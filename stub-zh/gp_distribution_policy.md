## 用法

来源：

- [官方上游 README](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/README.md)
- [官方扩展控制文件 (gp_distribution_policy.control)](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/gpcontrib/gp_distribution_policy/gp_distribution_policy.control)
- [官方扩展 SQL (gp_distribution_policy--1.0.sql)](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/gpcontrib/gp_distribution_policy/gp_distribution_policy--1.0.sql)

`gp_distribution_policy` — 检查 Greenplum 家族集群中的表分布策略。在进行数据库管理或自动化上述描述的行为时使用它。请使用上述链接的固定上游版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION gp_distribution_policy;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `gp_distribution_policy_table_check(relid regclass)` 是一个扩展函数，返回 `boolean`。
- `gp_table_distribution_check(regclass)` 是一个扩展函数，返回 `TABLE`。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `1.0`。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
