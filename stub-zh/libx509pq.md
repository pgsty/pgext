## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/crtsh/libx509pq/blob/bed05c6e87d79a98d79e42a1fac16386eead0c13/README.md)
- [1.3 版本 SQL 对象](https://github.com/crtsh/libx509pq/blob/bed05c6e87d79a98d79e42a1fac16386eead0c13/libx509pq--1.3.sql)
- [基于 OpenSSL 的实现](https://github.com/crtsh/libx509pq/blob/bed05c6e87d79a98d79e42a1fac16386eead0c13/libx509pq.c)

`libx509pq` 向 SQL 暴露 OpenSSL X.509 证书解析。它接收 DER `bytea` 证书，并提取主题、签发者、有效期、序列号、密钥与签名信息、扩展、备用名称、指纹等字段。

```sql
CREATE EXTENSION libx509pq;
SELECT x509_commonName(der),
       x509_issuerName(der),
       x509_notBefore(der),
       x509_notAfter(der),
       x509_keyAlgorithm(der),
       x509_keySize(der)
FROM certificate_store;
```

提取多个基础字段时，`x509_basic_info(der)` 可避免重复解析同一证书。其他 API 可枚举扩展与名称、提取公钥、识别部分弱密钥模式，并用指定密钥验证单个证书签名。

解析或签名验证并不等于完整 PKIX 路径验证：它本身不能建立信任、主机名有效性、吊销状态、策略或当前可接受性。畸形 DER 会进入原生 OpenSSL 代码，应限制输入大小，用不可信语料做模糊和回归测试，并及时修补 PostgreSQL 与 OpenSSL。构建和运行时应固定 OpenSSL ABI，审计错误哨兵与空值语义，限制高开销函数，并基准批量解析。
