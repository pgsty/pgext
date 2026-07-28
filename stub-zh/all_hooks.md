## 用法

来源：

- [官方上游 README](https://github.com/frbn/all_hooks/blob/c66e7b101f6fd7b216ea06aed91d9acb438447d5/README.md)
- [官方扩展控制文件 (all_hooks.control)](https://github.com/frbn/all_hooks/blob/c66e7b101f6fd7b216ea06aed91d9acb438447d5/all_hooks.control)
- [官方扩展 SQL (all_hooks--0.1.sql)](https://github.com/frbn/all_hooks/blob/c66e7b101f6fd7b216ea06aed91d9acb438447d5/all_hooks--0.1.sql)

`all_hooks` — 扩展用于在 PostgreSQL 中为每个可用的钩子发出一条消息。当应用程序需要此特定数据库功能时，请使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION all_hooks;
```

在目标数据库中安装扩展，运行可用的最小上游示例，并在将其集成到应用程序 SQL 之前验证安装版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.1`。
- 控制文件将扩展标记为可重定位。
- 控制文件将扩展标记为可信。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源中的信息一致。
