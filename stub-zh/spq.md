## 用法

来源：

- [官方上游 README](https://github.com/opengauss-mirror/spq_plugin_v2/blob/cd71e70ff3ee744c4f8a4cf7853774f38fb16af6/README.md)
- [官方扩展控制文件 (spq.control)](https://github.com/opengauss-mirror/spq_plugin_v2/blob/cd71e70ff3ee744c4f8a4cf7853774f38fb16af6/src/backend/distributed/spq.control)
- [官方扩展 SQL (spq--2.0.0.sql)](https://github.com/opengauss-mirror/spq_plugin_v2/blob/cd71e70ff3ee744c4f8a4cf7853774f38fb16af6/src/backend/distributed/sql/spq--2.0.0.sql)

`spq` — Spq 分布式查询。用于相应的分析或存储工作流。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION spq;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 之前验证安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `2.0.0`。
- 控制文件标记该扩展为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况，以匹配固定源代码。
