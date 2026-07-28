## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/constraint_uniform/constraint_uniform-1.0.0/README.md)
- [官方扩展控制文件 (constraint_uniform.control)](https://api.pgxn.org/src/constraint_uniform/constraint_uniform-1.0.0/constraint_uniform.control)

`constraint_uniform` — PostgreSQL 约束名称统一扩展 ========= 学位论文 -------- 使用此扩展来管理或自动化上述数据库行为。请使用链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION constraint_uniform;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `1.0.0`。
- 控制文件标记扩展为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
