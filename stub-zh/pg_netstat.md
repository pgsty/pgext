## 用法

来源：

- [官方上游 README](https://github.com/supabase/pg_netstat/blob/debce79df4737c5b5f5ea3826b3b4145df5c5c81/README.md)
- [官方扩展控制文件 (pg_netstat.control)](https://github.com/supabase/pg_netstat/blob/debce79df4737c5b5f5ea3826b3b4145df5c5c81/pg_netstat.control)
- [官方实现源代码](https://github.com/supabase/pg_netstat/blob/debce79df4737c5b5f5ea3826b3b4145df5c5c81/src/lib.rs)

`pg_netstat` — pg_netstat 监控你的 PostgreSQL 数据库网络流量。在收集或解释相应的 PostgreSQL 统计信息时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_netstat;

select * from pg_netstat;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `netstat()` 是一个扩展函数。

### 要求与注意事项

- 该目录记录了扩展的版本 `0.1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
