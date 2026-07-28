## 用法

来源：

- [官方上游 README](https://github.com/tyrchen/rust-training/blob/dd08bcb50cffdc4f1d715afdaf83af05eb76705c/README.md)
- [官方扩展控制文件 (db_assist.control)](https://github.com/tyrchen/rust-training/blob/dd08bcb50cffdc4f1d715afdaf83af05eb76705c/live_coding/db_assist/db_assist.control)
- [官方实现源代码](https://github.com/tyrchen/rust-training/blob/dd08bcb50cffdc4f1d715afdaf83af05eb76705c/live_coding/db_assist/src/lib.rs)

`db_assist` — 最小化 pgrx 培训扩展，仅暴露 hello_db_assist 函数。当 SQL 需要这些特殊函数或聚合时使用此扩展。在目标 PostgreSQL 构建中使用上述固定上游版本作为 API 边界并进行测试。

### 核心工作流

```sql
CREATE EXTENSION db_assist;
```

在目标数据库中安装扩展，如果有可用的上游最小示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `hello_db_assist()` 是一个扩展函数。

### 要求与注意事项

- 控制文件记录版本 `0.0.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
