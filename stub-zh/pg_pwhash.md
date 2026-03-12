

## 用法

> [pg_pwhash: PostgreSQL 高级密码哈希](https://github.com/cybertec-postgresql/pg_pwhash)

`pg_pwhash` 为 PostgreSQL 提供现代自适应密码哈希算法，包括 Argon2、scrypt 和 yescrypt。

```sql
CREATE EXTENSION pg_pwhash;
```

### 支持的算法

| 标识符 | 算法 | Salt 模式 |
|------------|-----------|--------------|
| `argon2i` | Argon2i | `$argon2i$v=19$m=4096,t=3,p=1$<salt>` |
| `argon2d` | Argon2d | `$argon2d$v=19$m=4096,t=3,p=1$<salt>` |
| `argon2id` | Argon2id | `$argon2id$v=19$m=4096,t=3,p=1$<salt>` |
| `scrypt` | Scrypt | `$scrypt$ln=16,r=8,p=1$<salt>` |
| `$7$` | Scrypt (crypt) | `$7$BU<salt>` |
| `yescrypt` | yescrypt (crypt) | `$y$j9T$<salt>` |

### 核心函数

#### 生成盐值和哈希

```sql
-- Argon2id（推荐）
SELECT pwhash_crypt('password', pwhash_gen_salt('argon2id'));
-- $argon2id$v=19$m=4096,t=3,p=1$<salt>$<hash>

-- Scrypt
SELECT pwhash_crypt('password', pwhash_gen_salt('scrypt'));

-- Yescrypt
SELECT pwhash_crypt('password', pwhash_gen_salt('yescrypt'));
```

#### 验证密码

```sql
-- 如果输出等于存储的哈希值则匹配
SELECT stored_hash = pwhash_crypt('entered_password', stored_hash) AS valid;
```

#### 直接哈希函数

```sql
SELECT pwhash_argon2('password', pwhash_gen_salt('argon2id'));
SELECT pwhash_scrypt('password', pwhash_gen_salt('scrypt'));
SELECT pwhash_yescrypt_crypt('password', pwhash_gen_salt('yescrypt'));
```

### 自定义盐值参数

```sql
-- 自定义内存/时间/并行度的 Argon2
SELECT pwhash_gen_salt('argon2id', 'm=65536', 't=4', 'p=2');

-- 自定义参数的 Scrypt
SELECT pwhash_gen_salt('scrypt', 'ln=20', 'r=8', 'p=1');
```

### 配置

| 参数 | 描述 |
|-----------|-------------|
| `pg_pwhash.argon2_default_backend` | Argon2 后端：`libargon2` 或 `openssl` |
