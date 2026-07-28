## 用法

来源：

- [官方上游 README](https://github.com/gabrosys/pg_purr/blob/5064eda546b0b36c9d3ff1877b02b708b18b99b9/README.md)
- [官方扩展控制文件 (pg_purr.control)](https://github.com/gabrosys/pg_purr/blob/5064eda546b0b36c9d3ff1877b02b708b18b99b9/pg_purr.control)

`pg_purr` — 一个通过 PL/Python 实现的 PostgreSQL 扩展，将量子计算带到了你的数据库中。在需要这些特殊函数或聚合时使用它。在安装扩展之前，必须先安装并验证其依赖项。

### 核心工作流

```sql
CREATE EXTENSION pg_purr;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 控制文件声明默认版本为 `0.2.0`。
- 请先安装确认的依赖项：`plpython3u`。
- 控制文件将该扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 控制文件将该扩展标记为不受信任。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源保持一致。
