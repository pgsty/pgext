## 用法

来源：

- [官方 README](https://github.com/enquo/pg_enquo/blob/9ae68398212302f877631784131ee93e30badf96/README.md)
- [官方数据类型文档](https://github.com/enquo/pg_enquo/blob/9ae68398212302f877631784131ee93e30badf96/doc/data_types/README.md)
- [官方扩展控制文件](https://github.com/enquo/pg_enquo/blob/9ae68398212302f877631784131ee93e30badf96/pg_enquo.control)

`pg_enquo` 为 Enquo 客户端库生成的密文提供 PostgreSQL 类型与运算符，用于存储和比较。加密与查询令牌构造发生在应用端；在正常查询流程中，扩展不会接收明文或加密密钥。

### 核心流程

创建扩展，并用所需 PostgreSQL 类型对应的加密类型定义列：

```sql
CREATE EXTENSION pg_enquo;

CREATE TABLE private_records (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  secret_number enquo_bigint NOT NULL,
  secret_date enquo_date,
  secret_text enquo_text
);
```

使用官方 Enquo 客户端库生成存储密文和查询密文，再作为参数绑定，不要在 SQL 中构造密文 JSON：

```sql
INSERT INTO private_records(secret_number, secret_date, secret_text)
VALUES (
  :'number_ciphertext'::enquo_bigint,
  :'date_ciphertext'::enquo_date,
  :'text_ciphertext'::enquo_text
);

SELECT id
FROM private_records
WHERE secret_text = :'text_query_ciphertext'::enquo_text;
```

客户端必须为已存储密文和查询密文使用相同的根密钥、字段身份、模式上下文、数据类型与存储选项。

### 支持的类型接口

- `enquo_bigint` 存储加密的 64 位有符号整数并支持比较；其文档默认表示每个值约 500 字节。
- `enquo_date` 存储年份范围为 `-32768` 到 `32767` 的加密日期并支持比较；其文档默认表示约 250 字节。
- `enquo_text` 存储任意 UTF-8 文本并支持相等与不等比较；空值表示从约 110 字节起，并按加密块增长。
- 当值从不需要查询时，客户端选项 `no_query` 可以减少存储，但会移除查询令牌。

### 安全取舍

- 应以默认安全模式为基线。客户端选项 `enable_reduced_security_operations` 增加索引或 `ORDER BY` 所需的确定性与顺序信息，但会泄露相等关系，以及有序类型的整列顺序，从而可能引发推断攻击。
- `enquo_bigint` 只有包含所需的降低安全性材料时才能使用 B-tree、hash 索引与排序；`enquo_date` 同样可以启用 B-tree 与排序，`enquo_text` 可以启用 hash 索引。索引必须与客户端存储选项匹配，否则操作会失败。
- 密文不会隐藏表形状、行数、空值、访问模式或查询频率。采用该方案前，应定义威胁模型、轮换并备份根密钥，并防止日志或链路在加密前记录明文。
- 版本 `0.0` 仍是早期接口。应把客户端库、`enquo-core`、扩展构建、模式字段身份和存储选项作为一个兼容单元固定，并测试恢复、密钥丢失、模式变更、重新加密、索引重建、非法密文和升级。
