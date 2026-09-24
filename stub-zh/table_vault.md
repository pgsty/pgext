## 用法

来源：

- [README](https://github.com/2ndQuadrant/rls-examples/blob/6a5b6d83e87ac682a0e2c7221ae2c2d31e5d9059/table-vault/README.md)
- [Control](https://github.com/2ndQuadrant/rls-examples/blob/6a5b6d83e87ac682a0e2c7221ae2c2d31e5d9059/table-vault/table_vault.control)
- [SQL](https://github.com/2ndQuadrant/rls-examples/blob/6a5b6d83e87ac682a0e2c7221ae2c2d31e5d9059/table-vault/table_vault--1.0.sql)

`table_vault` 是一个历史 SQL 示例，用于演示行级安全场景中的应用身份上下文。1.0 版本来自 2016 年的 rls-examples 源码快照，适合研究和改造，不是生产认证框架。

### 核心用法

安装文件并启用 `pgcrypto` 后，由管理员创建 `table_vault`，同时需要 PL/pgSQL。对象位于固定的 `table_vault` 模式，没有共享库或预加载步骤。以下 psql 示例要求先安全提供 `vault_passphrase` 变量，密钥表只应初始化一次。

```sql
CREATE EXTENSION pgcrypto;
CREATE EXTENSION table_vault;
INSERT INTO table_vault.key_vault
VALUES (crypt(:'vault_passphrase', gen_salt('bf')));
SELECT table_vault.set_username('app_user', :'vault_passphrase');
SELECT table_vault.get_username();
```

设置函数检查共享口令，然后保存`signed_vault.session_id` 中的随机 UUID，并在私有 `sessions` 表中查找对应上下文，并返回上下文令牌。读取函数仅在上下文通过验证且未满一天时返回应用用户名；上下文缺失或过期会报错。

### 对象与运行边界

- `set_username(text,text)` 与 `get_username()` 是允许 PUBLIC 执行的安全定义者函数；持有共享口令即可设置应用用户名。
- `key_vault` 保存受保护的凭据材料。SQL 撤销了公开访问，并为 API 授予模式使用权限。应限制管理表访问，并保护 SQL 传递的凭据。
- 函数没有固定 `search_path`，且调用了未限定模式的辅助函数。允许不可信 SQL 调用者使用前，应审查函数所有者权限与模式信任边界。
- 会话上下文跨事务保留。连接池每次交出连接都必须设置正确的上下文，不能泄露上一调用者的令牌。
- 上游没有提供当前 PostgreSQL 兼容矩阵或明确许可证。不能将这一历史示例等同于现代安全性或支持承诺。

过期行在读取时被拒绝，但扩展没有安排对 `sessions` 的自动清理，需要由所有者制定清理策略。
