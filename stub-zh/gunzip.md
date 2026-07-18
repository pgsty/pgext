## 用法

来源：

- [0.2 版本 SQL 对象](https://gitlab.com/nicklasaven/pg_gunzip/-/blob/master/gunzip--0.2.sql)
- [C 实现](https://gitlab.com/nicklasaven/pg_gunzip/-/blob/master/gunzip.c)
- [扩展 control 文件](https://gitlab.com/nicklasaven/pg_gunzip/-/blob/master/gunzip.control)

`gunzip` 在数据库内解压 gzip 压缩的字节串。0.2 版本定义了返回文本的 `gunzip_text(bytea)`、返回二进制的 `gunzip_bytea(bytea)`，以及作为文本函数别名的 `gunzip(bytea)`。

```sql
CREATE EXTENSION gunzip;
SELECT gunzip_text(compressed_payload),
       gunzip_bytea(compressed_payload)
FROM archive_item
WHERE id = 42;
```

解压结果可能不是有效数据库文本时，应使用 `gunzip_bytea(bytea)`。扩展在后端进程内完成解压，输入大小并不能限制输出大小：调用前应拒绝不可信或过大的压缩值，并配置语句与资源限制以降低解压炸弹风险。

上游 SQL 把这些函数标记为 immutable，但失败与内存压力仍会影响执行语句。0.2 已较陈旧，目录也只记录 PostgreSQL 10 覆盖。生产使用前，应在精确的 PostgreSQL 与 zlib 版本上验证构建兼容性、畸形流处理、编码行为和最大膨胀率。
