## 用法

来源：

- [上游来源 1](https://github.com/thejens/pgjinja/blob/8298a381f88437a70ae9d22d6b35fb6db64fd73c/README.md)
- [上游来源 2](https://github.com/thejens/pgjinja/blob/8298a381f88437a70ae9d22d6b35fb6db64fd73c/pgjinja.control)

`pgjinja` 0.1.0 在 PostgreSQL 内直接使用 JSONB 数据渲染 MiniJinja 模板。

### 核心流程

`jinja_render` 渲染内联模板，`jinja_render_named` 读取 `jinja_templates` 并支持继承与包含，`jinja_syntax_error` 返回语法诊断。存储模板的变更通过 PostgreSQL 的事务失效机制清理会话缓存。

```sql
CREATE EXTENSION pgjinja;
SELECT jinja_render('Hello {{ name }}', '{"name":"Ada"}'::jsonb, strict => true);
SELECT jinja_is_valid('{{ name }}');
```

### 运行边界

README 声明支持 Linux/macOS 上的 PostgreSQL 14–18，安装仅允许超级用户。内联模板默认不转义，处理不受信任 HTML 时应启用自动转义。应限制 `pgjinja.fuel`：渲染交回控制权前，语句超时无法中断它。JSON 数值会转换为 64 位整数或浮点数，需要精确保留的大数应以文本传入。模板不获得文件系统或网络访问能力。
