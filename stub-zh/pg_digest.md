## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/ratbox/pg-digest/blob/44ebe696700423cdd64ccc8f1d8b9dc18728934a/README.md)
- [1.0 版本 SQL 对象](https://github.com/ratbox/pg-digest/blob/44ebe696700423cdd64ccc8f1d8b9dc18728934a/pg_digest--1.0.sql)
- [类型实现](https://github.com/ratbox/pg-digest/blob/44ebe696700423cdd64ccc8f1d8b9dc18728934a/pg_digest.c)

`pg_digest` 定义了变长 `digest` 类型，用于紧凑存储哈希摘要、校验和、加密密钥和其他二进制块。它的表示与 `bytea` 二进制兼容，并提供双向隐式转换。

```sql
CREATE EXTENSION pg_digest;
CREATE TABLE file_upload (
  path text PRIMARY KEY,
  hash digest NOT NULL UNIQUE
);
INSERT INTO file_upload VALUES ('manual.pdf', '\x0123456789abcdef'::digest);
```

该扩展只存储字节，不会计算、校验、限制长度或标注哈希算法。应另行约束预期字节长度，并用 `pgcrypto` 或应用密码库计算摘要。切勿把校验和或短的非密码学摘要当作身份认证。

由于与 `bytea` 的转换是隐式的，应检查重载表达式和索引是否发生意外强制转换。将此类型用于持久模式前，应固定跨升级的二进制/文本表示并测试转储恢复及逻辑复制；若微小的存储差异并不重要，`bytea` 仍是更可移植的选择。
