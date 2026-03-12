

## 用法

> [md5hash: PostgreSQL 原生 MD5 哈希数据类型](https://github.com/tvondra/md5hash)

`md5hash` 扩展提供了高效的 128 位数据类型，以二进制格式（16 字节）而非文本格式（32+ 字节）存储 MD5 哈希值。

```sql
CREATE EXTENSION md5hash;

CREATE TABLE test_table (
    id md5hash PRIMARY KEY
);

INSERT INTO test_table VALUES ('c4ca4238a0b923820dcc509a6f75849b');

SELECT * FROM test_table
WHERE id = 'c4ca4238a0b923820dcc509a6f75849b';
```

### 运算符

支持标准比较运算符：`=`、`<>`、`<`、`>`、`<=`、`>=`。

### 索引支持

内置 Btree 索引运算符类，可在 `md5hash` 列上实现高效查找和主键约束。

### 存储优势

与以 `text` 类型存储 MD5 相比，`md5hash` 类型大约节省 40% 的存储空间，并且索引查找速度更快。
