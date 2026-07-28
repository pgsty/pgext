## 用法

来源：

- [官方上游 README](https://github.com/ringerc/scrapcode/blob/1d713af60d0b3957771448a685706d35779c4491/postgresql/random_bytea/README)
- [官方扩展控制文件 (random_bytea.control)](https://github.com/ringerc/scrapcode/blob/1d713af60d0b3957771448a685706d35779c4491/postgresql/random_bytea/random_bytea.control)
- [官方扩展 SQL (random_bytea--1.0.sql)](https://github.com/ringerc/scrapcode/blob/1d713af60d0b3957771448a685706d35779c4491/postgresql/random_bytea/random_bytea--1.0.sql)

`random_bytea` — 这是一个我在玩 PostgreSQL 中随机字节字符串生成时编写的玩具扩展。它是由这个 Stack Overflow 问题激发的。在需要这些特殊函数或聚合的 SQL 中使用它。在目标 PostgreSQL 构建中使用上面链接的上游固定版本作为 API 边界并进行测试。

### 核心工作流

```sql
CREATE EXTENSION random_bytea;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `random_bytea(integer)` 是一个扩展函数，返回 `bytea`。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 控制文件标记该扩展为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
