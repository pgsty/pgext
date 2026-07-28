## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/streaming_lag/streaming_lag-0.0.1/README.md)
- [官方扩展控制文件 (streaming_lag.control)](https://api.pgxn.org/src/streaming_lag/streaming_lag-0.0.1/streaming_lag.control)
- [官方扩展 SQL (streaming_lag--0.0.1.sql)](https://api.pgxn.org/src/streaming_lag/streaming_lag-0.0.1/streaming_lag--0.0.1.sql)

`streaming_lag` — streaming_lag 是一个实验性扩展，用于以时间单位而不是字节来衡量流式从库的滞后。在收集或解释相应的 PostgreSQL 统计信息时使用它。上游将此功能描述为实验性功能。

### 核心工作流

```sql
CREATE EXTENSION streaming_lag;
```

在目标数据库中安装该扩展，当可用时运行上游提供的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `streaming_lag` 是一个由扩展定义的视图。
- `streaming_lag_data` 是一个由扩展安装或管理的表。

### 要求与注意事项

- 审查的控制文件声明默认版本 `0.0.1`。
- 控制文件将该扩展标记为可重定位。
- 控制文件要求超级用户权限进行安装。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用之前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源保持一致。
