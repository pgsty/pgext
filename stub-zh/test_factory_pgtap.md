## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/test_factory/test_factory-0.4.2/README.md)
- [官方扩展控制文件 (test_factory_pgtap.control)](https://api.pgxn.org/src/test_factory/test_factory-0.4.2/test_factory_pgtap.control)
- [官方扩展 SQL (test_factory_pgtap.sql)](https://api.pgxn.org/src/test_factory/test_factory-0.4.2/sql/test_factory_pgtap.sql)

`test_factory_pgtap` — 一个用于在 Postgres 中管理单元测试数据的系统。当应用程序需要此特定数据库功能时使用它。在安装扩展及其依赖项并验证它们之前，请勿将其集成到应用程序 SQL 中。

### 核心工作流

```sql
CREATE EXTENSION test_factory_pgtap;
```

在目标数据库中安装扩展，当可用时运行最小的上游示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `tf.tap(table_name text , set_name text DEFAULT 'base')` 是一个扩展函数，返回 `SETOF text`。

### 要求与注意事项

- 控制文件声明默认版本为 `0.1.0`。
- 首先安装并验证确认的扩展依赖项：`pgtap`, `test_factory`。
- 控制文件标记该扩展为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源的一致性。
