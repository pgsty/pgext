## 用法

来源：

- [官方上游 README](https://github.com/pgrouting/vrprouting/blob/b34a8c323b8a0e3a81f7b2e5f4261cbd8634c929/README.md)
- [官方扩展控制文件 (vrprouting.control)](https://github.com/pgrouting/vrprouting/blob/b34a8c323b8a0e3a81f7b2e5f4261cbd8634c929/sql/pg_controls/vrprouting.control)

`vrprouting` — 车辆路径问题在数据库中。使用它来处理相应的空间数据或地理空间工作流。在安装和验证其扩展依赖项之前，请勿将其集成到应用程序 SQL 中。

### 核心工作流

```sql
CREATE EXTENSION vrprouting;
```

在目标数据库中安装扩展，当可用时运行最小的上游示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 该目录记录了扩展的版本 `0.5.0`。
- 首先安装确认的扩展依赖项：`plpgsql`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
