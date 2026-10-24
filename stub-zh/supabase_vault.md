## 用法

来源：

- [Official source](https://github.com/supabase/vault/blob/2f05b9be4dcac774a458d5a9e6acae794eb6b8c4/README.md)
- [Official source](https://github.com/supabase/vault/blob/2f05b9be4dcac774a458d5a9e6acae794eb6b8c4/supabase_vault.control.in)
- [Official source](https://github.com/supabase/vault/blob/2f05b9be4dcac774a458d5a9e6acae794eb6b8c4/Makefile)
- [Official source](https://github.com/supabase/vault/blob/2f05b9be4dcac774a458d5a9e6acae794eb6b8c4/sql/supabase_vault--0.3.1--0.3.2.sql)

`supabase_vault` 0.3.2 在 PostgreSQL 中保存加密秘密，并通过受权限控制的视图提供解密结果。固定模式为 `vault`，当前版本不再依赖 `pgsodium`。上游测试 PostgreSQL 14–19。由超级用户安装，无须预加载或重启。

### 基本流程

以下值均为演示字符串，不是真实凭据：

```sql
CREATE EXTENSION supabase_vault;
SELECT vault.create_secret('example-secret', 'api_token', 'Demo only');
SELECT name, decrypted_secret FROM vault.decrypted_secrets WHERE name = 'api_token';
SELECT vault.update_secret(id, 'replacement-example', 'api_token', 'Rotated demo')
FROM vault.secrets WHERE name = 'api_token';
```

`vault.create_secret` 返回秘密 UUID，可接受名称与描述；`vault.update_secret` 更新指定 UUID。`vault.secrets` 保存密文和元数据，`vault.decrypted_secrets` 向获准读取的角色提供明文。应明确限制辅助函数与解密视图的权限。

### 安全与升级

0.3.2 将秘密创建和更新辅助函数搜索路径中的 `pg_temp` 移到最后。安装对应文件后，在每个数据库执行 `ALTER EXTENSION supabase_vault UPDATE TO '0.3.2'`。

静态加密不能阻止有权查询的角色读取明文。秘密字面量和返回值可能进入数据库日志、客户端历史或应用遥测；应使用文档规定的秘密处理流程，并正确配置这些系统。保护加密密钥环境，验证备份恢复行为。不能因为存储表加密，就把解密视图的访问视为安全。
