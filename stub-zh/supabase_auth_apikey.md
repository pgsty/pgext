## 用法

来源：

- [0.0.3 版本官方 Database.dev 软件包页面](https://database.dev/martindonadieu/supabase_auth_apikey)
- [官方 Database.dev 软件包集合](https://database.dev/martindonadieu/pg-extensions)

`supabase_auth_apikey` 是一个 Database.dev SQL 软件包，为 Supabase Auth 用户和权限数组增加 API 密钥认证辅助函数。它面向不适合使用邮箱/密码认证的 IoT 设备、命令行工具等非交互客户端。

### 安装与核心流程

已发布的软件包版本为 `0.0.3`。使用官方页面给出的 SQL 软件包标识安装，然后创建规范 PostgreSQL 扩展，最好放在独立模式中：

```sql
SELECT * FROM dbdev('riderx-supabase_auth_apikey');
CREATE EXTENSION supabase_auth_apikey WITH SCHEMA extensions;

SELECT extensions.create_apikey(ARRAY['read', 'write']);

SELECT extensions.is_allowed_apikey(
    (current_setting('request.headers', true)::json ->> 'apikey'),
    ARRAY['read', 'write']
);
```

官方软件包页面记录了以下用户接口：

- `create_apikey(user_id, permissions)` 供后端代表用户创建密钥。
- `create_apikey(permissions)` 供当前已认证用户创建密钥。
- `delete_apikey(user_id, apikey)` 与 `delete_apikey(apikey)` 用于对应删除路径。
- `get_user_id(apikey)` 解析密钥所属用户。
- `is_allowed_apikey(apikey, permissions)` 检查权限，可用于行级安全策略。

### 安全边界

文档中的 RLS 模式从 Supabase/PostgREST 的 `request.headers` 设置取得 API 密钥。缺少头部、JSON 格式错误、密钥已撤销或权限不足时都必须拒绝访问。后端专用重载不得对客户端角色开放，并且每类调用者只应获得所需模式和函数权限。

软件包发布页当前链接的源代码仓库路径无法公开访问，因此公开页面不足以证明密钥如何生成、哈希、比较或存储。安装前必须审计生成的精确迁移，测试轮换与撤销，并避免不必要地记录或持久化返回的明文密钥。该软件包不等同于 Supabase 平台受支持的 Auth API 接口。
