## 用法

来源：

- [官方扩展文档](https://github.com/beargiles/pg-cert/blob/555b021861e0c8a52b8fcfb08e732945e56f5d82/doc/cert.md)
- [官方扩展 SQL](https://github.com/beargiles/pg-cert/blob/555b021861e0c8a52b8fcfb08e732945e56f5d82/sql/cert.sql)
- [官方 PGXN 分发页](https://pgxn.org/dist/cert/)

`cert` 0.1.1 为 PostgreSQL 添加 OpenSSL 格式的 X.509 证书、证书请求和 DSA 参数类型，并提供提取证书元数据的函数。上游明确说明该实现仍处于早期阶段，不适合生产使用。

### 核心流程

安装 SQL 依赖 `bignum` 扩展。将 PEM 证书存入 `cert` 列后，即可在普通查询中使用访问函数：

```sql
CREATE EXTENSION bignum;
CREATE EXTENSION cert;

CREATE TABLE certificates (
  certificate cert NOT NULL
);

SELECT get_version(certificate),
       get_serial_number(certificate),
       get_not_before(certificate),
       get_not_after(certificate),
       get_issuer(certificate),
       get_subject(certificate)
FROM certificates;
```

值通过 `cert` 输入函数以完整 OpenSSL PEM 证书形式写入。无效或不受支持的证书文本会被 C 输入例程拒绝。

### 重要对象

- `cert` 存储 X.509 证书；`cert_req` 存储证书请求。
- `basic_cert_info` 是由序列号、有效期、颁发者和主题组成的复合类型；`get_basic_info(cert)` 返回该类型。
- `get_version(cert)`、`get_serial_number(cert)`、`get_not_before(cert)`、`get_not_after(cert)`、`get_issuer(cert)` 和 `get_subject(cert)` 暴露基本证书字段。
- `get_signature(cert)`、`get_issuer_hash(cert)`、`get_iands_hash(cert)`、`get_common_name(cert)`、`get_subject_hash(cert)` 和 `get_subject_keyid_hash(cert)` 暴露其他文本属性。
- `dsa_params`、`size(dsa_params)` 和 `generate_dsa_params(integer)` 用于存储与生成 DSA 参数。

### 运维说明

扩展仅支持 OpenSSL 格式输入。SQL 注释指出，颁发者和主题目前返回文本而不是专用的 distinguished-name 类型，若干公钥操作仍未实现。C 模块链接 OpenSSL，安装脚本还会创建 `bignum`；应在目标 PostgreSQL 构建上核验这两项依赖。不要把已存储证书当作自动完成信任验证的对象：文档中的 API 负责解析与检查，但上游没有提供证书链验证或吊销检查说明。
