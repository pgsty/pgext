## 用法

来源：

- [已复核 commit 的官方用法文档](https://github.com/theirix/multicdr_fdw/blob/ee1b813e8515a8ef8b21893796ba46b593b7bcc4/doc/multicdr_fdw.md)
- [已复核 commit 的官方 README](https://github.com/theirix/multicdr_fdw/blob/ee1b813e8515a8ef8b21893796ba46b593b7bcc4/README.md)
- [1.2.2 版本 SQL 对象](https://github.com/theirix/multicdr_fdw/blob/ee1b813e8515a8ef8b21893796ba46b593b7bcc4/sql/multicdr_fdw.sql)
- [FDW 实现](https://github.com/theirix/multicdr_fdw/blob/ee1b813e8515a8ef8b21893796ba46b593b7bcc4/src/multicdr_fdw.c)
- [官方 PGXN 发行页](https://pgxn.org/dist/multicdr_fdw/)

`multicdr_fdw` 是只读外部数据包装器，把服务器本地某个目录中的定长话单文件扫描为 SQL 表。它可以用正则表达式选择文件、把不连续字段映射到列、暴露源文件名，并利用文件名中的日期裁剪文件。1.2.2 是面向 PostgreSQL 9.1 与 9.2 文档化的旧代码。

### 核心流程

PostgreSQL 操作系统账户必须能够读取目录与文件：

```sql
CREATE EXTENSION multicdr_fdw;
CREATE SERVER cdr_files FOREIGN DATA WRAPPER multicdr_fdw;

CREATE FOREIGN TABLE call_records (
  calling_number text,
  called_number text,
  duration_seconds integer
) SERVER cdr_files
OPTIONS (
  directory '/var/lib/postgresql/cdr',
  pattern '.*\.cdr$',
  posfields '0,20,40',
  rowminlen '41'
);

SELECT * FROM call_records;
```

`directory`、`pattern` 与 `posfields` 定义非递归文件扫描，以及定宽字段从零开始的起始位置。空字段变为 SQL NULL。已复核的读取器只支持 `text` 与四字节 `integer` 数据列；应在源文件中验证宽度与溢出。

### 表选项

`mapfields` 把表列映射到指定 CDR 字段，并保持表列顺序。`filefield` 指定接收文件路径的文本列。`rowminlen` 与 `rowmaxlen` 丢弃不符合字节长度范围的记录。没有 `mapfields` 时按顺序映射字段。

`dateformat` 把 `pattern` 的捕获组与文件名日期格式关联。`dateminfield` 与 `datemaxfield` 指定合成 timestamp 列；FDW 识别这些列上的等值条件后，可以跳过请求日期边界之外的文件。这是文件名裁剪，并不是从每条记录解析时间戳字段。

### 安全与兼容性边界

服务器进程会打开外部表选项中提供的本地路径。必须限制谁能创建或修改服务器与表，以及谁能查询它们；否则包装器可能成为服务器文件泄露路径。扫描不递归，也没有远程认证、写入支持、事务稳定的文件快照，或防止查询过程中源文件变化的机制。

实现面向已淘汰的 FDW 与后端 API，没有当前大版本兼容矩阵，并自行完成定宽解析和类型转换。使用前必须针对确切 PostgreSQL 构建进行移植、编译与回归测试，包括畸形行、并发文件轮换、日期裁剪、权限、编码及超长记录。
