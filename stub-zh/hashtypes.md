

## 用法

> [hashtypes: 哈希与校验和数据类型（SHA、CRC32）](https://github.com/adjust/hashtypes/)

`hashtypes` 扩展提供了原生数据类型，用于以二进制形式存储哈希值和校验和，相比文本存储更节省空间。

```sql
CREATE EXTENSION hashtypes;
```

### 数据类型

| 类型 | 存储大小 | 描述 |
|------|---------|------|
| `sha1` | 20 字节 | SHA-1 哈希（160 位） |
| `sha224` | 28 字节 | SHA-224 哈希（224 位） |
| `sha256` | 32 字节 | SHA-256 哈希（256 位） |
| `sha384` | 48 字节 | SHA-384 哈希（384 位） |
| `sha512` | 64 字节 | SHA-512 哈希（512 位） |
| `crc32` | 4 字节 | CRC-32 校验和 |

### 用法

```sql
CREATE TABLE objects (
    hash sha256 PRIMARY KEY,
    data bytea
);

INSERT INTO objects VALUES (
    'e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855',
    '\x'
);

SELECT * FROM objects WHERE hash = 'e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855';
```

### 运算符

所有类型均支持比较运算符：`=`、`<>`、`<`、`>`、`<=`、`>=`。

### 索引支持

所有类型均提供 Btree 和 Hash 索引运算符类。

### 类型转换

所有类型均支持与 `text` 和 `bytea` 之间的双向转换。
