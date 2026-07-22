## 用法

来源：

- [官方扩展文档](https://github.com/beargiles/pgopenssltypes/blob/7f89bf03d6c0d7ee7a1fbac03276490d575b0037/doc/pgopenssltypes.md)
- [官方 0.0.1 版本 SQL](https://github.com/beargiles/pgopenssltypes/blob/7f89bf03d6c0d7ee7a1fbac03276490d575b0037/sql/pgopenssltypes--0.0.1.sql)
- [官方函数参考](https://github.com/beargiles/pgopenssltypes/blob/7f89bf03d6c0d7ee7a1fbac03276490d575b0037/FUNCTIONS.md)

`pgopenssltypes` 为 OpenSSL 密钥、证书、请求、吊销列表与加密消息结构添加 PostgreSQL 类型，并提供摘要及 X.509 访问函数。它面向需要在 PostgreSQL 内解析和查询 PEM 对象的证书颁发机构或证书清单原型。

### 核心流程

```sql
CREATE EXTENSION pgopenssltypes;

CREATE TABLE certificates (
    id bigserial PRIMARY KEY,
    certificate x509 NOT NULL
);

SELECT x509_get_subject_name(certificate),
       x509_get_issuer_name(certificate),
       x509_get_not_before(certificate),
       x509_get_not_after(certificate)
FROM certificates;

SELECT dgst_sha256('message to hash');
```

重要类型包括 `X509`、`X509_NAME`、`PKEY`、`PKCS8`、`PKCS7`、`X509_REQ`、`X509_CRL` 与任意精度 `BN`。`RSA`、`DSA` 与 `DSA_PARAMS` 为历史格式保留，但上游不建议使用，因为私钥以未加密形式存储。`PKCS12` 虽有类型声明，但文档说明尚未实现。

X.509 访问函数可提取版本、序列号、有效期、主体、颁发者、公钥、别名、哈希与密钥标识。`x509_check_private_key(X509, PKEY)` 检查密钥是否匹配证书。摘要是否可用取决于链接的 OpenSSL 构建；MD4、MD5、SHA-1 等旧算法不得用于新的安全设计。

### 安全与维护说明

多数值使用 PEM 文本输入输出，可能包含私钥材料。因此 PostgreSQL 表、WAL、副本、日志、转储与备份都可能保存秘密；必须实施严格权限及外部加密/密钥管理控制。SQL 可见解析不能替代证书路径验证或托管 HSM。上游文档明确说明当前没有支持服务，而且 `0.0.1` 版本久远；必须针对精确的 PostgreSQL 与 OpenSSL ABI 验证，并在接受不可信输入前审计解析器。
