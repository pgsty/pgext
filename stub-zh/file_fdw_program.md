## 用法

来源：

- [已复核 commit 的 file_fdw_program 文档](https://github.com/corlogic-code/file_fdw_program/blob/412b751c6af32882f335515e5af1337000624ce2/doc/file_fdw_program.md)
- [已复核 commit 的 file_fdw_program README](https://github.com/corlogic-code/file_fdw_program/blob/412b751c6af32882f335515e5af1337000624ce2/README.md)
- [当前 PostgreSQL file_fdw 文档](https://www.postgresql.org/docs/current/file-fdw.html)

`file_fdw_program` 是 `file_fdw` 的 `program` 选项的历史回移植版本。外部表不读取指定文件，而是在服务端执行命令，并使用 COPY 风格的格式选项解析命令的标准输出。

### 创建由程序驱动的外部表

```sql
CREATE EXTENSION file_fdw_program;
CREATE SERVER file_program_server
  FOREIGN DATA WRAPPER file_fdw_program;

CREATE FOREIGN TABLE generated_rows (
  one text,
  two text,
  three text
)
SERVER file_program_server
OPTIONS (program 'echo one,two,three', format 'csv');

SELECT * FROM generated_rows;
```

扫描外部表时会执行命令。可用格式选项沿用该扩展提供的类 `file_fdw` 接口。

### 注意事项

- PostgreSQL 10 及更高版本附带的 `file_fdw` 已支持 `program`；现代 PostgreSQL 应使用内置扩展，而不是这个 1.0.1 回移植版本。
- `program` 由 PostgreSQL 服务端操作系统账户执行。应将服务器和外部表所有权视为代码执行权限，只使用固定命令，绝不能插入不可信输入。
- 上游将该包描述为 2016 年面向 PostgreSQL 9.4 及更高版本的回移植；它没有说明对当前 PostgreSQL 版本的支持情况。
