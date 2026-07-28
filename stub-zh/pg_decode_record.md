## 用法

来源：

- [官方扩展控制文件 (pg_decode_record.control)](https://github.com/myzel394/pg_decode_record/blob/d6053f846d8bda3eb487e24e39cbd5ec59e66aeb/pg_decode_record.control)
- [官方扩展 SQL (pg_decode_record--0.0.1.sql)](https://github.com/myzel394/pg_decode_record/blob/d6053f846d8bda3eb487e24e39cbd5ec59e66aeb/pg_decode_record--0.0.1.sql)

`pg_decode_record` 是一个从 WAL INSERT 记录中提取 SQL 的实验性扩展，最初来自 <https://github.com/rjuju>。它适用于管理或自动化上文所述的数据库行为。请以上方链接的固定上游修订为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_decode_record;
```

在目标数据库中安装扩展；如果上游提供了最小示例，请运行该示例，并在集成到应用 SQL 前验证安装版本和返回值。

### 要求与注意事项

- 审阅的控制文件声明默认版本为 `0.0.1`。
- 控制文件将该扩展标记为可重定位。
- 生产使用前，请根据固定版本源码确认权限、支持的 PostgreSQL 版本、升级行为和失败情形。
