## 用法

来源：

- [官方上游 README](https://github.com/optionfactory/olddog/blob/c4195bea2d90e00b98e1d346c5af480e4a5a0d0e/newtricks/README.md)
- [官方扩展控制文件 (newtricks.control)](https://github.com/optionfactory/olddog/blob/c4195bea2d90e00b98e1d346c5af480e4a5a0d0e/newtricks/newtricks.control)
- [官方实现源代码](https://github.com/optionfactory/olddog/blob/c4195bea2d90e00b98e1d346c5af480e4a5a0d0e/newtricks/src/lib.rs)

`newtricks` — 一些使用 Rust 的 postgres 扩展示例。当应用程序需要此特定数据库功能时，请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION newtricks;

SELECT is_valid_fiscal_code('RSSMRA85T10A562S');

CREATE TABLE users (id INT, cf TEXT CHECK(is_valid_fiscal_code(cf)));
INSERT INTO users (id, cf) VALUES (1, 'asd');
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

### 重要对象

- `emojify` 是一个扩展函数。
- `is_secure` 是一个扩展函数。
- `is_valid_fiscal_code` 是一个扩展函数。
- `list_emojis()` 是一个扩展函数。

### 要求与注意事项

- 该目录记录版本为 `0.0.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
