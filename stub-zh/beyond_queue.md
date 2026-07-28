## 用法

来源：

- [官方上游 README](https://github.com/beyondoss/queue/blob/731d6039952814955a77d1bc065f50291deca781/README.md)
- [官方扩展控制文件 (beyond_queue.control)](https://github.com/beyondoss/queue/blob/731d6039952814955a77d1bc065f50291deca781/beyond-queue-extension/beyond_queue.control)
- [官方实现源代码](https://github.com/beyondoss/queue/blob/731d6039952814955a77d1bc065f50291deca781/beyond-queue-extension/src/lib.rs)

`beyond_queue` — 一个更快的 pgmq 分支，适用于 https://beyond.dev 队列。当应用程序需要此特定数据库功能时，请使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION beyond_queue;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 控制文件声明默认版本为 `0.1.0`。
- 控制文件标记该扩展为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码一致。
