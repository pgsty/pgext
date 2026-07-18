## 用法

来源：

- [1.0 版本 SQL 声明](https://github.com/RekGRpth/pg_ssl/blob/e0d6397ad7f034085e2bf3972c901c6f73e858ee/pg_ssl--1.0.sql)
- [OpenSSL C 实现](https://github.com/RekGRpth/pg_ssl/blob/e0d6397ad7f034085e2bf3972c901c6f73e858ee/pg_ssl.c)
- [扩展 control 文件](https://github.com/RekGRpth/pg_ssl/blob/e0d6397ad7f034085e2bf3972c901c6f73e858ee/pg_ssl.control)
- [MIT 许可证](https://github.com/RekGRpth/pg_ssl/blob/e0d6397ad7f034085e2bf3972c901c6f73e858ee/LICENSE)

`pg_ssl` 提供一个 C 函数 `sign(cert text, data text)`，使用 OpenSSL 创建 PKCS#7 签名并返回 base64 文本。第一个参数会同时被解析为 X.509 证书及其私钥。

```sql
CREATE EXTENSION pg_ssl;
SELECT oid::regprocedure
FROM pg_proc
WHERE oid = 'sign(text,text)'::regprocedure;
```

不要把生产私钥粘贴到 SQL、查询日志、监控样本、客户端历史或应用参数中。PostgreSQL 文本值不是硬件保护的密钥库，该函数也没有密钥标识符、密码回调、证书链策略或受保护签名服务边界。

上游 README 为空，也没有测试、发行版、使用指南或兼容矩阵。应把 1.0 视为未经审查的原生密码学粘合代码。必须审计内存处理与 OpenSSL 兼容性、撤销 `PUBLIC` 的 `EXECUTE`，真实凭据则应使用外部密钥管理签名服务。
