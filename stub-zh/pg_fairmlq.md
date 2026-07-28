## 用法

来源：

- [官方上游 README](https://github.com/sam-harri/pg_fairmlq/blob/a1df8e9f4bc51282f87e4144973716c33826c136/README.md)
- [官方扩展控制文件 (pg_fairmlq.control)](https://github.com/sam-harri/pg_fairmlq/blob/a1df8e9f4bc51282f87e4144973716c33826c136/pg_fairmlq.control)
- [官方扩展 SQL (pg_fairmlq--0.1.0.sql)](https://github.com/sam-harri/pg_fairmlq/blob/a1df8e9f4bc51282f87e4144973716c33826c136/sql/pg_fairmlq--0.1.0.sql)

`pg_fairmlq` — 测试需要运行 pgtap 扩展。当应用程序需要此特定数据库功能时，请使用它。在安装扩展之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION pg_fairmlq;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.1.0`。
- 请先安装确认的扩展依赖项：`plpgsql`。
- 控制文件将该扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源保持一致。
