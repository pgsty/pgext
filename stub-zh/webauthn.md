## 用法

来源：

- [官方 README 与 API 参考](https://github.com/truthly/pg-webauthn/blob/c1f8099db1cbed3a89cf12a2668ed67bb10ee354/README.md)
- [官方扩展控制文件](https://github.com/truthly/pg-webauthn/blob/c1f8099db1cbed3a89cf12a2668ed67bb10ee354/webauthn.control)
- [WebAuthn Level 2 规范](https://www.w3.org/TR/webauthn-2/)

`webauthn` 是一个用纯 SQL 实现的 PostgreSQL WebAuthn 凭据注册与断言验证扩展。它在固定的 `webauthn` 模式中保存一次性挑战和公钥凭据，而浏览器仪式、HTTPS 来源、依赖方策略、授权和账户恢复设计仍由应用负责。

### 前置条件与注册

版本 `1.6` 依赖 `pgcrypto`、`pg_ecdsa_verify` 和 `cbor`。只有在审查这些扩展及其信任边界后，才应通过依赖级联安装：

```sql
CREATE EXTENSION webauthn CASCADE;

SELECT webauthn.init_credential(
  challenge := gen_random_bytes(32),
  user_name := 'alex@example.com',
  user_id := gen_random_bytes(32),
  user_display_name := 'Alex',
  relying_party_name := 'Example'
);
```

将返回的 `publicKey` 对象交给 `navigator.credentials.create()`。再把浏览器响应传入 `webauthn.store_credential()`；成功时它会保存公钥并返回已注册的 `user_id`，返回 `NULL` 则表示失败。每个挑战只能使用一次。

### 身份验证

使用服务端新生成的挑战开始登录：

```sql
SELECT webauthn.get_credentials(
  challenge := gen_random_bytes(32),
  user_name := 'alex@example.com',
  relying_party_id := 'example.com'
);
```

把该结果交给 `navigator.credentials.get()`，然后将断言字段传入 `webauthn.verify_assertion()`。验证通过时会返回凭据的 `user_id`；必须拒绝 `NULL`。只有凭据创建时设置了 `require_resident_key := TRUE`，才能使用无用户名登录；此时验证过程通过 `user_handle` 识别用户。

### API 索引

- `webauthn.init_credential()` 生成注册选项并保存注册挑战。
- `webauthn.store_credential()` 验证并保存浏览器创建的凭据。
- `webauthn.get_credentials()` 生成断言选项并保存登录挑战。
- `webauthn.verify_assertion()` 验证浏览器断言并返回其 `user_id`。

两个生成选项的函数都默认使用五分钟超时；上游将自定义超时限制在 30 秒至 10 分钟之间。应明确设置依赖方 ID，使用密码学安全的服务端随机源生成挑战，限制对挑战表和凭据表的直接访问，并将返回身份视为认证输入而非应用授权。生产部署前，应测试受支持的认证器、备份与恢复、克隆数据库、挑战清理、失败路径以及完整的来源/RP 策略。
