## 用法

来源：

- [官方上游 README](https://github.com/siose-innova/pg_siose_bench/blob/e28280ad8b97a19568586056b0d0259e35ad1d1b/README.md)
- [官方扩展控制文件 (pg_siose_bench.control)](https://github.com/siose-innova/pg_siose_bench/blob/e28280ad8b97a19568586056b0d0259e35ad1d1b/pg_siose_bench.control)

`pg_siose_bench` — 一个用于基准测试不同 SIOSE 数据库配置（纯关系型、索引型、json、jsonb、xml 等）的 PostgreSQL 扩展。当应用程序需要这种特定的数据库功能时使用它。在安装扩展之前，必须先安装并验证其依赖项。

### 核心工作流

```sql
CREATE EXTENSION pg_siose_bench;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 控制文件声明默认版本为 `0.0.1`。
- 请先安装确认的扩展依赖项：`postgis`。
- 控制文件标记该扩展为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
