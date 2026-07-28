## 用法

来源：

- [官方扩展控制文件（bgValidators.control）](https://api.pgxn.org/src/bgvalidators/bgvalidators-0.1.2/bgValidators.control)
- [官方扩展 SQL（bgValidators--1.0.sql）](https://api.pgxn.org/src/bgvalidators/bgvalidators-0.1.2/bgValidators--1.0.sql)

`bgvalidators` — 用于验证 IBAN、保加利亚 EGN 个人身份证、BULSTAT 增值税号码和 LNCh 外国人身份证的函数。当 SQL 需要这些特殊功能或聚合时，请使用此扩展。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION bgvalidators;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `validate_bulstat` 是一个扩展函数。
- `validate_egn` 是一个扩展函数。
- `validate_iban` 是一个扩展函数。
- `validate_lnch` 是一个扩展函数。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
