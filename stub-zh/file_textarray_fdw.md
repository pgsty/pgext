## 用法

来源：

- [官方 README](https://github.com/adunstan/file_text_array_fdw/blob/dde60712f74a249e9d3675f8c33ee18c33d23bd1/README.md)
- [扩展 SQL](https://github.com/adunstan/file_text_array_fdw/blob/dde60712f74a249e9d3675f8c33ee18c33d23bd1/file_textarray_fdw.sql)
- [FDW 实现](https://github.com/adunstan/file_text_array_fdw/blob/dde60712f74a249e9d3675f8c33ee18c33d23bd1/file_textarray_fdw.c)

`file_textarray_fdw` 版本 `1.0.1` 用于读取每行字段数可能不同的服务端文件。每个解析后的行通过单个 `text[]` 列返回，无需为外部表声明固定的字段列表。

### 核心流程

```sql
CREATE EXTENSION file_textarray_fdw;
CREATE SERVER file_server FOREIGN DATA WRAPPER file_textarray_fdw;

CREATE FOREIGN TABLE agg_csv_array (t text[])
SERVER file_server
OPTIONS (
  format 'csv',
  filename '/path/to/agg.csv',
  header 'true',
  delimiter ';',
  quote '@',
  escape '"',
  null ''
);

SELECT t[1]::int2 AS a, t[2]::float4 AS b
FROM agg_csv_array;
```

外部表唯一的列必须是 `text[]`。字段按文件中的顺序进入数组，因此调用方可以只索引并转换所需位置。

### 选项与边界

该包装器接受 PostgreSQL 文件 FDW 机制实现的解析选项，包括 `filename`、`format`、`header`、`delimiter`、`quote`、`escape` 和 `null`。文件名位于 PostgreSQL 服务器而非客户端，服务器操作系统账号必须能读取它。

此 FDW 只实现扫描，不能回写文件。应谨慎授予服务器与外部表权限，因为它可能暴露 PostgreSQL 服务账号可读的文件。不需要用户映射、预加载或重启。
