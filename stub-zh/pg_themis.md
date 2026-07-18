## 用法

来源：

- [上游固定版本 README](https://github.com/cossacklabs/pg_themis/blob/609525bdd564d90c0067398fa0debe99882b12e8/README.md)
- [1.0 版本安装 SQL](https://github.com/cossacklabs/pg_themis/blob/609525bdd564d90c0067398fa0debe99882b12e8/pg_themis--1.0.sql)
- [固定版本 C 实现](https://github.com/cossacklabs/pg_themis/blob/609525bdd564d90c0067398fa0debe99882b12e8/pg_themis.c)
- [固定版本扩展控制文件](https://github.com/cossacklabs/pg_themis/blob/609525bdd564d90c0067398fa0debe99882b12e8/pg_themis.control)
- [当前 Themis 产品文档](https://www.cossacklabs.com/themis/)

pg_themis 1.0 是 2016 年编写的外部 Themis 密码库 PostgreSQL 绑定。它暴露使用对称密钥的 Secure Cell Seal 加解密，以及使用接收方公钥加密、对应私钥解密的 Secure Message。

### Secure Cell 往返示例

必须先安装兼容的系统 Themis 库。这个隔离示例只借助 pgcrypto 生成临时密钥：

```sql
CREATE EXTENSION pgcrypto;
CREATE EXTENSION pg_themis;

WITH key_material AS MATERIALIZED (
    SELECT gen_random_bytes(32) AS key
)
SELECT convert_from(
    pg_themis_scell_decrypt_seal(
        pg_themis_scell_encrypt_seal(
            convert_to('payload', 'UTF8'),
            key
        ),
        key
    ),
    'UTF8'
)
FROM key_material;
```

扩展不负责生成、存储、轮换、授权或审计应用密钥，这些控制仍完全由调用方承担。

### 不安全的遗留 Secure Message 路径

Secure Message 辅助函数在调用 Themis 前，把明文或密文长度收窄为一个无符号字节。超过 255 字节的输入长度因此会按 256 回绕截断。解密路径还会在确认密文至少有四字节之前就读取内嵌的 32 位公钥长度，随后又信任该长度进行指针运算。恶意密文可触达越界读取。不要使用这个快照中的 Secure Message SQL 函数。

四个函数都把明文或密钥作为 SQL bytea 参数。查询文本、参数、活动视图、日志、追踪、错误、备份和内存都可能暴露秘密。安装 SQL 保留默认 PUBLIC 执行权；应撤销它，避免在 SQL 字面量中放密钥，并采用经过审查的密钥管理设计。

该 PostgreSQL 绑定自 2016 年后没有更新，其 PostgreSQL 9.1+ 声明并不是当前兼容保证；现有 Themis 本身可能已独立演进。必须固定并测试精确库 ABI，使用认证测试向量，对畸形和边界长度输入做模糊测试；生产密码功能优先采用受维护的集成。
