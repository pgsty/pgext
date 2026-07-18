## 用法

来源：

- [已复核 commit 的 parquet_fdw README](https://github.com/adjust/parquet_fdw/blob/d15664ebdcb1cbb759a7bb39fd26fb2fa2fff3ea/README.md)
- [已复核 commit 的 parquet_fdw.control](https://github.com/adjust/parquet_fdw/blob/d15664ebdcb1cbb759a7bb39fd26fb2fa2fff3ea/parquet_fdw.control)

`parquet_fdw` 是用于 Apache Parquet 文件的实验性只读外部数据包装器。它通过 Apache Arrow 和 Parquet 库，将单个文件或一组文件呈现为 PostgreSQL 外部表。

### 读取 Parquet 文件

```sql
CREATE EXTENSION parquet_fdw;
CREATE SERVER parquet_srv FOREIGN DATA WRAPPER parquet_fdw;
CREATE USER MAPPING FOR CURRENT_USER
  SERVER parquet_srv
  OPTIONS (user 'postgres');

CREATE FOREIGN TABLE userdata (
  id integer,
  first_name text,
  last_name text
)
SERVER parquet_srv
OPTIONS (filename '/mnt/data/userdata.parquet');

SELECT * FROM userdata LIMIT 20;
```

`filename` 选项可以包含以空格分隔的多个路径。高级选项包括 `sorted`、`files_func`、`max_open_files`、`use_mmap` 和 `use_threads`；上游还说明可用 `IMPORT FOREIGN SCHEMA` 发现服务端目录中的文件。

### 注意事项

- 该包装器只读，且上游将其标为实验性。路径指向 PostgreSQL 服务端文件系统，而不是 SQL 客户端文件系统。
- 构建已复核版本需要系统 Apache Arrow/Parquet 库；README 要求 0.15 或更高版本。
- 支持的映射包括整数、浮点数、时间戳、日期、字符串、二进制、列表和映射值。不支持结构体和嵌套列表。
- control 文件标识扩展版本 0.2，与 catalog 版本一致。请验证部署所用的外部库与 PostgreSQL 组合。
