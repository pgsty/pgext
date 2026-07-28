## 用法

来源：

- [官方上游 README](https://github.com/supabase/pg_crdt/blob/22109c27c481a62476295d7c5c14ccb8cf654b8a/README.md)
- [官方扩展控制文件 (automerge.control)](https://github.com/supabase/pg_crdt/blob/22109c27c481a62476295d7c5c14ccb8cf654b8a/automerge.control)
- [官方实现源代码](https://github.com/supabase/pg_crdt/blob/22109c27c481a62476295d7c5c14ccb8cf654b8a/src/automerge/automerge.c)

`automerge` — pg_crdt 是一个实验性扩展，为 PostgreSQL 添加了冲突自由复制数据类型 (CRDT) 的支持。当应用程序数据需要这种类型、领域或其操作符时，请使用此扩展。上游将此功能描述为实验性。

### 核心工作流

```sql
CREATE EXTENSION automerge;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.0.1`。
- 控制文件标记该扩展为不可重定位。
- 控制文件不要求超级用户安装。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源代码。
