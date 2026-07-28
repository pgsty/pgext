## 用法

来源：

- [官方上游 README](https://github.com/jwdeitch/pg_suuid/blob/5cab05c4eb9d6989eb6d0372116ec65079ec9cd2/readme)
- [官方扩展控制文件 (suuid.control)](https://github.com/jwdeitch/pg_suuid/blob/5cab05c4eb9d6989eb6d0372116ec65079ec9cd2/suuid.control)

`suuid` — 小型 UUID。当应用程序数据需要这种类型、域或其操作符时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION suuid;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.0.1`。
- 控制文件标记扩展为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源代码。
