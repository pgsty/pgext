


## 用法

来源：[README](https://github.com/neondatabase/pg_session_jwt/blob/v0.5.0/README.md)、[v0.5.0 tag](https://github.com/neondatabase/pg_session_jwt/tree/v0.5.0)、[control file](https://github.com/neondatabase/pg_session_jwt/blob/v0.5.0/pg_session_jwt.control)

`pg_session_jwt` 通过 JWT 处理认证会话。配置 JWK 后验证 JWT 真实性；无 JWK 时回退到兼容 PostgREST 的 `request.jwt.claims`。

```sql
CREATE EXTENSION pg_session_jwt;
```

### 模式1：JWK 验证

在连接时通过 libpq 选项设置 JWK：

```bash
export PGOPTIONS="-c pg_session_jwt.jwk=$MY_JWK"
```

在会话中使用：

```sql
SELECT auth.init();                        -- 使用 JWK 初始化
SELECT auth.jwt_session_init('eyJ...');    -- 设置并验证 JWT
SELECT auth.user_id();                     -- 获取 'sub' 声明
SELECT auth.session();                     -- 获取完整 JWT 载荷（JSONB）
```

### 模式2：PostgREST 兼容（无 JWK）

开箱即用适配 PostgREST，无需初始化：

```sql
SELECT auth.user_id();   -- 从 request.jwt.claims 返回 'sub'
SELECT auth.session();   -- 返回完整声明（JSONB）
```

### 函数

| 函数 | 返回类型 | 描述 |
|----------|---------|-------------|
| `auth.init()` | `void` | 使用 JWK 初始化会话 |
| `auth.jwt_session_init(jwt text)` | `void` | 设置并验证 JWT |
| `auth.session()` | `jsonb` | 获取 JWT 载荷或回退声明 |
| `auth.jwt()` | `jsonb` | `auth.session()` 的别名 |
| `auth.user_id()` | `text` | 获取 `sub` 声明 |
| `auth.uid()` | `uuid` | 获取 `sub` 作为 UUID（或 NULL） |
| `auth.organization()` | `jsonb` | Neon Auth organization claim helper |
| `auth.organization_id()` | `uuid` | Neon Auth organization id helper |

### 配置

| 参数 | 描述 |
|-----------|-------------|
| `pg_session_jwt.jwk` | 用于 JWT 验证的 JWK（在启动或连接时设置） |
| `pg_session_jwt.audit_log` | 启用审计日志（`on`/`off`） |

### RLS 示例

```sql
CREATE POLICY user_isolation ON my_table
    USING (user_id = auth.user_id());
```

对 Neon Auth 的 organization-scoped policies，使用 `o` claim helpers：

```sql
CREATE POLICY team_select ON team
  FOR SELECT
  USING (organization_id = auth.organization_id());
```

### 版本说明

v0.5.0 README 增加 Neon Auth organization helpers，并明确区分 `auth.jwt()`、`auth.user_id()`、`auth.uid()` 等 portable helpers，以及 Neon-specific 的 `auth.organization()` 和 `auth.organization_id()`。其他 auth provider 应使用 `auth.jwt()` 并直接提取 provider-specific claims。
