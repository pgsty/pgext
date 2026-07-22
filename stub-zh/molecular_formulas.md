## 用法

来源：

- [官方控制文件](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/molecular_formulas/molecular_formulas.control)
- [官方 Molecular Formula README](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/molecular_formulas/README.md)
- [生成的 PostgreSQL 17 扩展 SQL](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/molecular_formulas/extension/usr/share/postgresql/17/extension/molecular_formulas--0.1.0.sql)

`molecular_formulas` 将分子式 Rust 解析器和工具库打包为 PGRX 扩展产物。在所核验的 0.1.0 产物中，生成的扩展 SQL 没有暴露面向用户的 SQL 解析器、类型或函数。

### 核心流程

```sql
CREATE EXTENSION molecular_formulas;

SELECT extversion
FROM pg_extension
WHERE extname = 'molecular_formulas';
```

这只会注册扩展。`MolecularFormula`、`Ion`、`Atom` 和 `Token` 等 Rust 对象是库 API，在钉住版本的生成 SQL 中不是 PostgreSQL 对象。

### 当前 SQL 接口

- 所核验生成 SQL 没有声明 SQL 类型、函数、操作符、视图或配置参数。
- README 将 Rust 解析器描述为仍在推进的工作，并把分子轨道校验列为未来工作；不要从“可解析”推断出化学有效性保证。

### 运维说明

0.1.0 版不可重定位。控制文件设置 `superuser = true` 和 `trusted = false`，因此必须由超级用户创建；它未声明前置扩展或预加载要求。升级前应重新检查生成 SQL 与发布说明，因为 SQL 边界可能独立于 Rust crate API 变化。
