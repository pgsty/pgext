## 用法

来源：

- [官方扩展控制文件 (pgqbw.control)](https://github.com/rekgrpth/pgqbw/blob/a104ef679aa83fc16913ce4b9dea3b4cf98259b3/pgqbw.control)
- [官方实现源代码](https://github.com/rekgrpth/pgqbw/blob/a104ef679aa83fc16913ce4b9dea3b4cf98259b3/src/pgqbw.c)

`pgqbw` — postgres 队列后台工作者。当应用程序需要此特定的数据库功能时使用它。使用链接的上游版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pgqbw;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况，以匹配固定源代码。
