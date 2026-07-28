## 用法

来源：

- [官方上游 README](https://github.com/sashaaro/taskboss/blob/fcdc0d35ee6d8a9bffc46d3ba7082e6852242dcf/README.md)
- [官方扩展控制文件 (taskboss.control)](https://github.com/sashaaro/taskboss/blob/fcdc0d35ee6d8a9bffc46d3ba7082e6852242dcf/taskboss.control)
- [官方实现源代码](https://github.com/sashaaro/taskboss/blob/fcdc0d35ee6d8a9bffc46d3ba7082e6852242dcf/src/lib.rs)

`taskboss` — 一个使用 pgrx 编写的原生 PostgreSQL 作业队列扩展，基于 Rust 编写。受 pg-boss 启发。当应用程序需要此特定数据库功能时，请使用它。在目标 PostgreSQL 构建中测试上游链接的固定版本作为 API 边界。

### 核心工作流

```sql
CREATE EXTENSION taskboss;
```

在目标数据库中安装扩展，当可用时运行上游最小示例，并在将其集成到应用程序 SQL 之前验证安装的版本和返回值。

### 重要对象

- `create_queue` 是一个扩展函数。
- `delete_queue` 是一个扩展函数。

### 要求与注意事项

- 该目录记录版本 `0.1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 控制文件将扩展标记为不受信任。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
