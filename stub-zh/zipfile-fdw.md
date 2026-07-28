## 用法

来源：

- [官方上游 README](https://github.com/beargiles/zipfile-fdw/blob/dfdfd5f3be8724ef338d7caa72ba1ecb4d4a5754/README.md)
- [官方扩展控制文件 (zipfile-fdw.control)](https://github.com/beargiles/zipfile-fdw/blob/dfdfd5f3be8724ef338d7caa72ba1ecb4d4a5754/zipfile-fdw.control)
- [官方实现源代码](https://github.com/beargiles/zipfile-fdw/blob/dfdfd5f3be8724ef338d7caa72ba1ecb4d4a5754/src/zipfile-fdw.c)

`zipfile-fdw` — 确保已经安装了 pg_config 并且在路径中。如果使用了包管理系统（如 RPM）安装 PostgreSQL，请确保也安装了 -devel 包。如果必要，请告知构建过程 pg_config 的位置：使用它来使 PostgreSQL 能够通过外部数据接口访问相应的外部数据源。在目标 PostgreSQL 构建上使用上述链接的固定上游版本作为 API 边界并进行测试。

### 核心工作流

```sql
CREATE EXTENSION "zipfile-fdw";
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.0.1`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况，以匹配固定源代码。
