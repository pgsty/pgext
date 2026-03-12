

## 用法

> [supabase_vault: Supabase 的加密密钥存储](https://github.com/supabase/vault)

Supabase Vault 提供 `vault.secrets` 表来存储静态加密的敏感信息（API 密钥、令牌等）。通过 `vault.decrypted_secrets` 视图即时解密。

```sql
CREATE EXTENSION supabase_vault CASCADE;
```

### 存储密钥

```sql
INSERT INTO vault.secrets (secret) VALUES ('s3kre3t_k3y') RETURNING *;

-- 或使用辅助函数：
SELECT vault.create_secret('another_s3kre3t');

-- 带可选名称和描述：
SELECT vault.create_secret('my_secret', 'unique_name', 'This is the description');
```

### 读取密钥

`vault.secrets` 表以加密方式存储数据。使用 `vault.decrypted_secrets` 视图读取解密后的值：

```sql
SELECT * FROM vault.decrypted_secrets ORDER BY created_at DESC LIMIT 3;
-- 包含 `decrypted_secret` 列，其中为明文值
```

### 更新密钥

```sql
SELECT vault.update_secret(
    '7095d222-efe5-4cd5-b5c6-5755b451e223',
    'n3w_upd@ted_s3kret',
    'updated_unique_name',
    'This is the updated description'
);
```

### 安全提示

关闭语句日志以防止密钥出现在日志中：

```sql
ALTER SYSTEM SET statement_log = 'none';
```
