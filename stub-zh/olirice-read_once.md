## 用法

来源：

- [Official database.dev 包页面](https://database.dev/olirice/read_once)

`olirice-read_once` — 发送只能读取一次的消息。当应用程序需要此特定的数据库功能时使用它。在安装并验证其扩展依赖项之前，请勿集成到应用程序 SQL 中。

### 核心工作流

```sql
CREATE EXTENSION "olirice-read_once";
```

在目标数据库中安装扩展，当可用时运行最小的上游示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `read_message(id uuid)` 是一个扩展函数，返回 `text`。
- `send_message(contents text)` 是一个扩展函数，返回 `uuid`。
- `read_once` 是由扩展创建的模式。

### 要求与注意事项

- 该目录记录版本 `0.3.2`。
- 首先安装已确认的扩展依赖项：`pg_cron`。
- 这是一个 database.dev/pg_tle 包；在发出引用的 `CREATE EXTENSION` 之前，请先注册或生成其包迁移。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
