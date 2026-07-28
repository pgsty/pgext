## 用法

来源：

- [官方上游 README](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/contrib/interconnect/README.md)
- [官方扩展控制文件 (interconnect.control)](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/contrib/interconnect/interconnect.control)
- [官方扩展 SQL (interconnect--1.0.sql)](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/contrib/interconnect/interconnect--1.0.sql)

`interconnect` — 此扩展提供 Apache Cloudberry 的累积互联统计信息，包括队列大小、缓冲区使用情况、重传次数、数据包错误和其他与 UDPIFC 相关的指标。将其用于相应的分析或存储工作流。在目标 PostgreSQL 构建中测试链接的上游修订版本作为 API 边界。

### 核心工作流

```sql
CREATE EXTENSION interconnect;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

### 重要对象

- `gp_interconnect_stats` 是一个扩展定义的视图。
- `gp_interconnect_stats_per_host` 是一个扩展定义的视图。
- `gp_interconnect_stats_per_segment` 是一个扩展定义的视图。

### 要求与注意事项

- 审查的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为不可重定位。
- 在生产使用之前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与链接的源代码一致。
