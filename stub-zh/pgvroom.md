## 用法

来源：

- [官方上游 README](https://github.com/pgrouting/pgvroom/blob/d38a14de6b8b98fafd7949d427b427e51fb65f2d/README.md)
- [官方扩展控制文件 (pgvroom.control)](https://github.com/pgrouting/pgvroom/blob/d38a14de6b8b98fafd7949d427b427e51fb65f2d/sql/pg_controls/pgvroom.control)

`pgvroom` — VROOM 功能可通过数据库访问。使用它来进行相应的空间数据或地理空间工作流。在将其集成到应用程序 SQL 中之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION pgvroom;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 该目录记录了扩展的版本 `0.1.0`。
- 首先安装确认的扩展依赖项：`plpgsql`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源中的信息一致。
