## 用法

来源：

- [官方 SQL 定义](https://github.com/pgstuff/base32_4b/blob/0518fe2e17a7e634405c79eb51ca5b36f0e1275c/sql/base32_4b.sql)
- [官方 README](https://github.com/pgstuff/base32_4b/blob/0518fe2e17a7e634405c79eb51ca5b36f0e1275c/README.md)
- [官方控制文件](https://github.com/pgstuff/base32_4b/blob/0518fe2e17a7e634405c79eb51ca5b36f0e1275c/base32_4b.control)

`base32_4b` 定义了一个按值传递、内部表示为四字节的 Base32 类型。它只适用于扩展输入函数能够接受的标识符；这是紧凑标量类型，并非任意字节串的通用编码器。

### 核心流程

```sql
CREATE EXTENSION base32_4b;

CREATE TABLE compact_ids (
    id base32_4b PRIMARY KEY,
    payload jsonb NOT NULL
);

INSERT INTO compact_ids (id, payload)
VALUES ($1::text::base32_4b, '{"state":"ready"}');

SELECT id::text, payload
FROM compact_ids
WHERE id = $1::text::base32_4b;
```

批量加载前应通过显式类型转换验证应用输入。文本转换是赋值转换而非隐式转换，有助于保持比较的类型边界。

### 重要对象

- `base32_4b` 是四字节标量类型，支持文本与二进制输入输出。
- 比较运算符 `<`、`<=`、`=`、`<>`、`>=` 与 `>` 提供排序和相等判断。
- `base32_4b_ops` 是默认 B-tree 运算符类，因此支持主键和 B-tree 索引。
- 扩展为该类型定义了 `min` 与 `max` 聚合。
- `base32_4b_to_text` 与 `base32_4b_from_text` 实现与 `text` 之间的赋值转换。

### 限制

SQL 接口没有定义哈希运算符类、算术或通用的二进制到 Base32 转换。索引相等与排序应使用 B-tree，并确认扩展排序符合应用期望的文本顺序。

仓库只有一份 2015 年左右的初始实现，文档也很少。采用前应测试可接受字符表、大小写处理、非法与边界输入、二进制收发、索引排序、转储恢复以及目标 PostgreSQL 大版本。数据本身占四字节并不表示每行只增加四字节，因为元组头、空值位图和对齐也会占空间。
