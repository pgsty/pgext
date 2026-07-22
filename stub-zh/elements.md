## 用法

来源：

- [官方控制文件](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/elements_rs/elements.control)
- [官方 Elements README](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/elements_rs/README.md)
- [生成的 PostgreSQL 17 扩展 SQL](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/elements_rs/extension/usr/share/postgresql/17/extension/elements--0.1.0.sql)

`elements` 将 EMI Rust 化学元素 crate 打包为 PGRX 扩展产物。Rust crate 包含元素、同位素、原子量、氧化态、价电子、成键和电子排布 API，但所核验的 0.1.0 生成 SQL 没有暴露面向用户的 SQL 类型或函数。

### 核心流程

```sql
CREATE EXTENSION elements;

SELECT extversion
FROM pg_extension
WHERE extname = 'elements';
```

这会注册扩展，但钉住版本的生成 SQL 只包含 PGRX 头注释。不要假设 `Element` 或 `Isotope` 等 Rust 名称在该产物中可作为 PostgreSQL 类型使用。

### 当前 SQL 接口

- 所核验生成 SQL 没有声明 SQL 类型、函数、操作符、视图或配置参数。
- README 记录的大量元素和同位素 API 属于 Rust crate，需要使用相应 Rust feature 编译的消费者。

### 运维说明

0.1.0 版不可重定位。控制文件设置 `superuser = true` 和 `trusted = false`，因此必须由超级用户创建；它未声明前置扩展或预加载要求。升级时应重新核验生成 SQL，因为后续 PGRX 构建可能添加该版本不存在的 SQL 接口。
