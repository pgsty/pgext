## 用法

来源：

- [README](https://api.pgxn.org/src/pg_mustach/pg_mustach-2.0.0/README.md)
- [Control file / 控制文件](https://api.pgxn.org/src/pg_mustach/pg_mustach-2.0.0/pg_mustach.control)
- [SQL](https://api.pgxn.org/src/pg_mustach/pg_mustach-2.0.0/pg_mustach--2.0.sql)

`pg_mustach` 2.0 通过 libmustach 2 或更新版本将 `jsonb` 渲染为 Mustache 模板输出，提供预备模板，并以统一接口替代旧版特定 JSON 库的适配器。

### 核心工作流

```sql
CREATE EXTENSION pg_mustach;
SELECT mustach('{"name":"PostgreSQL"}'::jsonb, 'Hello {{name}}!');
BEGIN;
SELECT mustach_template('Hello {{name}}!', 'greeting');
SELECT mustach_json('{"name":"Ada"}'::jsonb, tplname := 'greeting');
SELECT mustach_free('greeting');
COMMIT;
```

### 模板与标志

`mustach` 直接渲染，`mustach_template` 准备命名模板或未命名槽位，`mustach_json` 使用模板渲染，`mustach_free` 释放模板。应显式使用命名参数 `tplname`，否则第二个位置文本参数可能解析为写文件重载。

`pg_mustach.transaction` 默认为 true，事务结束时会清理预备模板。自动提交模式中，前一语句准备的模板不会保留到下一语句。只有明确管理会话生命周期时才应关闭该选项，连接池尤其需要注意。`pg_mustach.flags` 和 `mustach_set_flags` 控制渲染标志，`mustach_with_*` 辅助函数返回标志值。

### 升级与访问

控制版本为 `2.0`，PGXN 发行版本为 2.0.0。从 `json` 适配器迁移到 `jsonb` 前应检查随附升级 SQL。文件输出重载写入服务器文件系统，需要超级用户权限。扩展可重定位，无需预加载；处理非受信输入时应限制模板和输出大小。

模板中的本地文件片段由仅超级用户可设置的 `pg_mustach.whitelist` 前缀列表单独控制。普通角色需要显式匹配的许可，非空列表也会限制超级用户；该白名单不会授权文件输出重载。
