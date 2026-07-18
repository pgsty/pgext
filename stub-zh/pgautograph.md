## 用法

来源：

- [已复核 commit 的 pgautograph Rust 入口](https://github.com/nevindev/pgautograph/blob/a8c59f407f80fb5ed4f1ee8e3ab15745aab03100/src/lib.rs)
- [已复核 commit 的属性图查询生成器](https://github.com/nevindev/pgautograph/blob/a8c59f407f80fb5ed4f1ee8e3ab15745aab03100/src/queries/mod.rs)
- [已复核 commit 的 pgautograph.control](https://github.com/nevindev/pgautograph/blob/a8c59f407f80fb5ed4f1ee8e3ab15745aab03100/pgautograph.control)
- [已复核 commit 的 pgautograph Cargo 清单](https://github.com/nevindev/pgautograph/blob/a8c59f407f80fb5ed4f1ee8e3ab15745aab03100/Cargo.toml)

`pgautograph` 是一个早期 Rust/`pgrx` 扩展，用于从关系型外键推导属性图 DDL。当前的 `list_foreign_keys()` 函数检查 `public` 模式中的外键约束，返回来源和目标元数据，并以 INFO 消息输出生成的 `CREATE PROPERTY GRAPH graph` 语句。

### 检查外键

```sql
CREATE EXTENSION pgautograph;

SELECT *
FROM list_foreign_keys();
```

应检查返回行和生成的 INFO 消息。当前函数只生成文本，不会执行属性图语句。

### 注意事项

- 已复核 crate 版本是 0.0.0，仓库没有 README、许可证声明、标签或发行版。应将其视为开发中代码。
- control 文件设置了 `superuser = true`，因此安装扩展需要具有相应权限的角色。
- 生成器只检查 `public` 模式，并使用简单的首字母大写与单数化规则构造图标签。执行前必须审查并编辑生成的 DDL。
- Cargo 清单以 `pgrx` 0.16.1 声明 PostgreSQL 13 至 18 功能；这表示源码目标，不是已发布的兼容性保证。
