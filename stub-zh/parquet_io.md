## 用法

来源：

- [postgres/resources/parquet_io.md](https://github.com/minyeamer/linkmerce/blob/a662f8ab7646c677bfd9d2112b0c0a6556ca396b/postgres/resources/parquet_io.md)
- [postgres/extension/parquet_io.control](https://github.com/minyeamer/linkmerce/blob/a662f8ab7646c677bfd9d2112b0c0a6556ca396b/postgres/extension/parquet_io.control)
- [postgres/extension/parquet_io--1.0.sql](https://github.com/minyeamer/linkmerce/blob/a662f8ab7646c677bfd9d2112b0c0a6556ca396b/postgres/extension/parquet_io--1.0.sql)
- [postgres/extension/parquet_io.cpp](https://github.com/minyeamer/linkmerce/blob/a662f8ab7646c677bfd9d2112b0c0a6556ca396b/postgres/extension/parquet_io.cpp)

`parquet_io` 1.0 是 LinkMerce 中借助 Apache Arrow C++ 在 PostgreSQL 与 Parquet 之间传输数据的组件。上游环境面向 PostgreSQL 18。安装需要超级用户，无需共享预加载。

### 基本用法

```sql
CREATE EXTENSION parquet_io;
SELECT parquet_create('/srv/import/input.parquet', 'public.imported_rows');
SELECT parquet_read('/srv/import/input.parquet', 'public.imported_rows');
SELECT octet_length(parquet_write('SELECT * FROM public.imported_rows'));
```

### 接口与边界

`parquet_create` 创建空表并返回列数，不导入数据行。默认冲突模式为 `error`；`replace` 会删除目标表。`parquet_read` 向已有表追加数据并返回行数。两者均接受服务器路径或 `bytea` 值。`parquet_write` 只接收 SQL 时返回 Parquet 字节，同时接收路径时写入服务器文件并返回行数。

路径由数据库服务器按其操作系统权限解析。文件输出不受 PostgreSQL 事务回滚保护。应仅允许可信角色执行这些函数，并审核 SQL 参数及替换目标。安装脚本未撤销函数的默认授权。大字节值及 Arrow 转换会占用服务器内存，类型映射可能简化输入文件的模式。
