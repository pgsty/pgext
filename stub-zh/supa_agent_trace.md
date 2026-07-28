## 用法

来源：

- [官方上游 README](https://github.com/jessevent/supa-agent/blob/e2a69b5c551d4925358252724788887bcf1cb862/supabase/extensions/supa_agent_trace/README.md)
- [官方扩展控制文件 (supa_agent_trace.control)](https://github.com/jessevent/supa-agent/blob/e2a69b5c551d4925358252724788887bcf1cb862/supabase/extensions/supa_agent_trace/supa_agent_trace.control)
- [官方扩展 SQL (supa_agent_trace--0.1.0.sql)](https://github.com/jessevent/supa-agent/blob/e2a69b5c551d4925358252724788887bcf1cb862/supabase/extensions/supa_agent_trace/supa_agent_trace--0.1.0.sql)

`supa_agent_trace` — 该扩展和 DevTool 均通过 Supabase **管理 API**（账户级 OAuth）进行身份验证，两者均不持有项目 GoTrue 会话。使用它来收集或解释相应的 PostgreSQL 统计信息。在集成到应用程序 SQL 之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION supa_agent_trace;

select dbdev.install('jessevent@supa_agent_trace');
create extension "jessevent@supa_agent_trace";
```

在目标数据库中安装该扩展，当可用时运行最小的上游示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `public.agent_trace_broadcast()` 是一个扩展函数，返回 `trigger`。
- `public.agent_trace_prune(retention interval default interval '7 days')` 是一个扩展函数，返回 `bigint`。
- `public.agent_trace_topic(uid uuid)` 是一个扩展函数，返回 `text`。
- `public.agent_trace_events` 是一个由该扩展安装或管理的表。

### 要求与注意事项

- 控制文件声明默认版本为 `0.1.0`。
- 先安装并验证确认的扩展依赖项：`pgcrypto`。
- 控制文件将该扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源的一致性。
