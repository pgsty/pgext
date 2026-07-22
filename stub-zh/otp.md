## 用法

来源：

- [官方 README](https://api.pgxn.org/src/otp/otp-1.0.0/README.md)
- [官方扩展 SQL](https://api.pgxn.org/src/otp/otp-1.0.0/sql/otp.sql)
- [官方扩展控制文件](https://api.pgxn.org/src/otp/otp-1.0.0/otp.control)

`otp` 版本 `1.0.0` 使用不受信任的服务端 Perl 实现一个小型 SHA-1 TOTP 流程。它可以生成配置 URL 和当前时间窗口的验证码，但其密钥生成器不具备密码学安全性，应把该扩展视为旧式参考代码，而不是生产认证系统。

### 核心流程

安装过程会创建 `plperlu`，因此需要超级用户权限、服务端 PL/PerlU 支持，以及 Perl 模块 `Digest::HMAC_SHA1`：

```sql
CREATE EXTENSION otp;

SELECT provisioning_url(
  'alice@example.com',
  'jbswy3dpehpk3pxp',
  30,
  'Example'
);

SELECT generate_totp('jbswy3dpehpk3pxp', 30, 6);
SELECT verify_totp('jbswy3dpehpk3pxp', 30, '123456');
```

`generate_totp(text, integer, integer)` 生成当前时间步的验证码。`verify_totp(text, integer, text)` 只检查当前时间步，并且始终生成六位候选值；它不实现时钟漂移窗口或防重放机制。

### 辅助函数与安全注意事项

`random_base32(integer)` 生成小写 Base32 风格密钥；`urlencode(text)`、`provisioning_url(text, text, integer, text)`、`pack(text)`、`unpack(text)` 与 `perl_hmac(text, text)` 为该流程提供支持。

`random_base32(integer)` 使用 PostgreSQL 非密码学安全的 `random()`，绝不能用来生成认证密钥。应使用外部密码学安全生成器创建密钥，对其静态加密，限制 SQL 与日志暴露，并在扩展外实现速率限制、防重放、恢复与时钟偏差策略。`urlencode(text)` 只能处理有限的 UTF-8 序列，因此非 ASCII 标签可能失败或编码错误。
