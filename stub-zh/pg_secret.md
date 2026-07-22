## 用法

来源：

- [官方设计笔记](https://github.com/TypedLambda/pg_secret/blob/eef18ef9f504fa55fb19011e9f4844495f89709a/NOTES)
- [官方扩展 SQL](https://github.com/TypedLambda/pg_secret/blob/eef18ef9f504fa55fb19011e9f4844495f89709a/pg_secret--0.5.sql)
- [官方 C 实现](https://github.com/TypedLambda/pg_secret/blob/eef18ef9f504fa55fb19011e9f4844495f89709a/pgsecret.c)

`pg_secret` 0.5 是一个实验性的保序泄露加密原型。它把整数或文本转换成 `secret` 域，并提供比较运算符和 B-tree 运算符类。它只适合研究如何按密文顺序建立索引，不是通用的机密管理或生产加密系统。

### 核心流程

`make_secret(key, key, int8)` 加密 64 位整数。文本重载先把输入转换成 64 位 SipHash 值，因此密文顺序是哈希顺序，而不是文本字典序。两个密钥都必须恰好为 16 字节，尽管 `key` 域本身并不约束长度。

```sql
CREATE EXTENSION pg_secret;

CREATE TABLE secure_rankings (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    score secret NOT NULL
);

WITH keys AS (
    SELECT decode('00112233445566778899aabbccddeeff', 'hex')::key AS k1,
           decode('ffeeddccbbaa99887766554433221100', 'hex')::key AS k2
)
INSERT INTO secure_rankings (score)
SELECT make_secret(k1, k2, value::int8)
FROM keys, unnest(ARRAY[10, 30, 20]) AS v(value);

CREATE INDEX secure_rankings_score_idx
    ON secure_rankings USING btree (score);

SELECT id
FROM secure_rankings
ORDER BY score;
```

应用必须在该表之外保留并保护两个密钥。若要重新构造比较边界，必须用同一对密钥调用 `make_secret`。

### 对象

- 域：`secret` 是被约束为 224 字节的 `bytea` 值；`key` 是不限制长度的 `bytea` 域。
- 构造函数：`make_secret(key, key, int8)` 与 `make_secret(key, key, text)`。
- 比较：用于 `secret` 的 `<`、`<=`、`=`、`>=` 和 `>`，底层通过 `secret_cmp` 与默认 B-tree 运算符类 `secret_btree_ops` 实现。

0.5 版没有解密函数、认证加密 API、密钥注册表或密钥轮换流程。

### 安全与运维边界

保序泄露加密会有意泄漏顺序，上游笔记也把该实现称为较弱的安全模型。服务器会收到明文和两个密钥；PostgreSQL 日志、错误报告、内存处理或 WAL 相关行为都可能暴露敏感材料。实现中还有尚未解决的密钥清理工作。

该源码是 2019–2020 年的早期实验，依赖内置密码代码以及 OpenSSL/GMP 时代的构建假设。请独立进行密码学审查、限制执行权限，并测试目标 PostgreSQL 构建。不要用 `pg_secret` 替代经过审计的加密、KMS 集成或应用侧密钥托管。
