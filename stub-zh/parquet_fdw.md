## 用法

来源：

- [官方上游文档](https://github.com/adjust/parquet_fdw/blob/d15664ebdcb1cbb759a7bb39fd26fb2fa2fff3ea/README.md)
- [官方扩展控制文件](https://github.com/adjust/parquet_fdw/blob/d15664ebdcb1cbb759a7bb39fd26fb2fa2fff3ea/parquet_fdw.control)
- [官方选项与导入实现](https://github.com/adjust/parquet_fdw/blob/d15664ebdcb1cbb759a7bb39fd26fb2fa2fff3ea/src/parquet_impl.cpp)

`parquet_fdw` 0.2 是面向本地 Apache Parquet 文件的实验性只读外部数据包装器。它使用 Apache Arrow 与 Parquet C++ 库，支持多文件、并行扫描、有序文件合并策略和模式导入。路径在 PostgreSQL 服务器而不是客户端上解析。

### 核心流程

创建服务器并将 Parquet 文件映射为外部表：

```sql
CREATE EXTENSION parquet_fdw;
CREATE SERVER parquet_srv FOREIGN DATA WRAPPER parquet_fdw;

CREATE FOREIGN TABLE public.userdata (
  id integer,
  first_name text,
  last_name text
)
SERVER parquet_srv
OPTIONS (filename '/srv/parquet/userdata.parquet');

SELECT * FROM public.userdata WHERE id >= 100;
```

PostgreSQL 操作系统账号必须能读取每个选定路径。不支持通过外部表写入。

### 文件、类型与导入

`filename` 接受以空格分隔的文件列表。也可以用 `files_func` 指定一个接受 JSONB 参数并返回 `text[]` 的函数，再由 `files_func_arg` 提供该 JSONB 值。其他选项包括 `sorted`、`files_in_order`、`use_mmap`、`use_threads` 和 `max_open_files`。全局开关包括 `parquet_fdw.use_threads`、`parquet_fdw.enable_multifile` 与 `parquet_fdw.enable_multifile_merge`。

支持将 Arrow 整数、浮点、时间戳、日期、字符串、二进制、列表和映射转换为 PostgreSQL 数值类型、timestamp、date、text、bytea、数组与 JSONB。不支持 struct 和嵌套列表。`IMPORT FOREIGN SCHEMA` 将远端模式名解释为带双引号的本地目录路径，并根据目录中的文件创建外部表。

### 安全与正确性

已复核的验证器会检查路径存在性和选项，但没有执行类似 `file_fdw` 的超级用户限制。能够创建或修改这些外部表的角色，可以尝试读取 PostgreSQL 服务器账号可访问的文件，或调用获准的 `files_func`。应只向可信管理员授予外部服务器使用权、模式创建权、表所有权与函数所有权；使用专用只读目录下的绝对路径，并对不可变文件列表函数使用模式限定名。

`sorted` 与 `files_in_order` 选项是对计划器的承诺。错误声明可能产生错误排序结果，因此启用前必须验证文件范围。应固定匹配的 PostgreSQL、Arrow 与 Parquet 库版本，并在代表性数据上测试类型转换、时间戳单位、NULL、模式漂移、内存映射、文件替换、打开文件上限、并行扫描和畸形输入。
