## 用法

来源：

- [0.2 版本 SQL 对象](https://gitlab.com/nicklasaven/pg_gunzip/-/blob/c46688ee9dd6e23702b631a8e0b08964df418b94/gunzip--0.2.sql)
- [C 实现](https://gitlab.com/nicklasaven/pg_gunzip/-/blob/c46688ee9dd6e23702b631a8e0b08964df418b94/gunzip.c)
- [扩展 control 文件](https://gitlab.com/nicklasaven/pg_gunzip/-/blob/c46688ee9dd6e23702b631a8e0b08964df418b94/gunzip.control)

`gunzip` 在数据库内解压 gzip 压缩的字节串。0.2 版本定义了返回文本的 `gunzip_text(bytea)`、返回二进制的 `gunzip_bytea(bytea)`，以及作为文本函数别名的 `gunzip(bytea)`。

### 核心流程

```sql
CREATE EXTENSION gunzip;
SELECT gunzip_text(compressed_payload),
       gunzip_bytea(compressed_payload)
FROM archive_item
WHERE id = 42;
```

### 数据与后端安全

两个输出函数都不具备二进制安全性。它们都把 C `strlen` 的结果作为 PostgreSQL datum 长度，因此遇到第一个零字节就会截断；`gunzip_bytea(bytea)` 也无法保留任意解压后字节。只能接受已知不含 NUL 的负载，并在扩展之外验证结果。

SQL 函数没有声明为 strict，而 C 代码也不检查 SQL `NULL` 参数。空输入或畸形输入会在 notice 后返回 `NULL`，分配失败会调用 `exit`，解压还会在后端本地无上限地扩张缓冲区。因此，一次错误调用可能终止对应后端，小型压缩值也可能耗尽内存。不要向不可信输入开放这些函数。0.2 已较陈旧，目录也只记录 PostgreSQL 10 覆盖；生产使用前必须修复这些缺陷，并在精确的 PostgreSQL 与 zlib 版本上测试。
