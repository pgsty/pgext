


## 用法

> [omni_session: 会话管理](https://docs.omnigres.org/omni_session/session_management/)

`omni_session` 扩展提供标准化的会话管理，创建非日志表 `omni_session.sessions` 用于活跃会话。主要与 HTTP 栈配合使用。

### HTTP 会话处理器

**请求处理器** -- 从 `_session` cookie 中提取会话 UUID：

```sql
SELECT omni_session.session_handler(request);
-- 可选: session_handler(request, cookie_name => 'my_session')
```

如果 cookie 缺失或无效则创建新会话，设置 `omni_session.session` 事务变量。

**响应处理器** -- 在响应中设置会话 cookie：

```sql
SELECT omni_session.session_handler(response);
```

Cookie 参数：

| 参数 | 默认值 | 描述 |
|:-----|:-------|:-----|
| `cookie_name` | `_session` | Cookie 标识符 |
| `http_only` | `true` | 阻止 JavaScript 访问 |
| `secure` | `true` | 仅 HTTPS |
| `same_site` | `Lax` | 跨站行为 |
| `domain` | `null` | 域范围 |
| `max_age` | `null` | 生存时间（秒） |
| `expires` | `null` | 过期时间戳 |
| `path` | `null` | URL 路径范围 |

### 直接会话 ID 处理

```sql
-- 创建或验证会话：
SELECT omni_session.session_handler(null::omni_session.session_id);
-- 返回新的会话 ID 并设置 omni_session.session 事务变量
```
