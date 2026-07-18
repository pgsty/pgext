## 用法

来源：

- [项目 README](https://github.com/singularita/pg_strverscmp/blob/266395a35003dea9763bf25a584c4771f1a21552/README.md)
- [1.0 版安装 SQL](https://github.com/singularita/pg_strverscmp/blob/266395a35003dea9763bf25a584c4771f1a21552/pg_strverscmp--1.0.sql)
- [C 语言比较封装](https://github.com/singularita/pg_strverscmp/blob/266395a35003dea9763bf25a584c4771f1a21552/pg_strverscmp.c)
- [构建定义](https://github.com/singularita/pg_strverscmp/blob/266395a35003dea9763bf25a584c4771f1a21552/Makefile)

`pg_strverscmp` 1.0 为文本提供类似版本号的自然排序，使连续数字按数值而不是逐字符比较。它提供 `pg_strverscmp(text,text)`、六个比较操作符和 `strverscmp_ops` B-tree 操作符类。

### 自然排序

```sql
CREATE EXTENSION pg_strverscmp;

SELECT value
FROM (VALUES ('item2'), ('item10'), ('item1')) AS v(value)
ORDER BY value USING +<;
```

可用操作符是 `+<`、`+<=`、`+=`、`+<>`、`+>=` 和 `+>`。若要让索引查询采用相同语义，应使用 `strverscmp_ops` 声明文本索引；普通文本索引仍使用 PostgreSQL 的常规排序规则和操作符类。

### 重要限制

构建时需要链接 `libunistring`。上游将该比较器描述为支持 Unicode 且通常不区分大小写，但明确警告：它在 NUL 字节处不是二进制安全的，不使用 PostgreSQL 排序规则设施，并且会在字符比较过程中反复调用 Unicode 排序，效率较低。应把它视为专用排序规则，而不是本地化文本相等语义的直接替代品。查询操作符和索引操作符类必须保持一致，库或平台变化后还应重新验证排序结果。
