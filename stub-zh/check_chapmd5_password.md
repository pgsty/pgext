## 用法

来源：

- [官方 README](https://github.com/sshutdownow/pg_check-chap-md5/blob/5371628350775b7d51ac097bee24913705856fb1/README.md)
- [0.0.1 版本扩展 SQL](https://github.com/sshutdownow/pg_check-chap-md5/blob/5371628350775b7d51ac097bee24913705856fb1/check_chapmd5_password--0.0.1.sql)
- [C 实现](https://github.com/sshutdownow/pg_check-chap-md5/blob/5371628350775b7d51ac097bee24913705856fb1/check_chapmd5_password.c)
- [RFC 1994：PPP 挑战握手认证协议](https://www.rfc-editor.org/rfc/rfc1994)

`check_chapmd5_password` 使用质询值和明文密钥校验 RFC 1994 CHAP-MD5 响应。软件包名为 `pg_check_chap_md5`，但规范的 PostgreSQL 扩展名和 SQL 函数名均为 `check_chapmd5_password`。

### 核心流程

```sql
CREATE EXTENSION check_chapmd5_password;

SELECT check_chapmd5_password(
    '00777f2a3f6a2e661947b520c6777e0b25',
    '45c915d82d67257209048420a31292d3',
    'password'
);
```

第一个参数是包含一字节标识符的十六进制 CHAP 响应，第二个参数是十六进制质询值，第三个参数是明文共享密钥。响应匹配时函数返回 `true`，否则返回 `false`。

### 对象索引

- `check_chapmd5_password(chap_password text, chap_challenge text, clear_password text) RETURNS boolean` 执行校验。

### 运维说明

版本 `0.0.1` 可重定位，且未声明扩展依赖或预加载要求。该函数为 `STRICT` 和 `IMMUTABLE`。

第三个参数属于凭据：应限制 `EXECUTE` 权限，避免在查询文本或监控中暴露调用，并且不要启用 `DEBUG2` 日志，因为实现会在该级别记录明文密码。输入由宽松的十六进制解析器解码；非法字符只会产生警告，并没有健壮的校验。调用前应校验响应长度、偶数长度的十六进制编码以及质询大小。CHAP-MD5 本身是旧式协议，不能替代传输加密或现代密码存储。
