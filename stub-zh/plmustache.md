## 用法

来源：

- [README.md](https://github.com/PostgREST/plmustache/blob/024b550d4fe2bbb4e122d2b3216a9ab7f514837a/README.md)
- [LICENSE](https://github.com/PostgREST/plmustache/blob/024b550d4fe2bbb4e122d2b3216a9ab7f514837a/LICENSE)
- [Makefile](https://github.com/PostgREST/plmustache/blob/024b550d4fe2bbb4e122d2b3216a9ab7f514837a/Makefile)
- [plmustache.control.in](https://github.com/PostgREST/plmustache/blob/024b550d4fe2bbb4e122d2b3216a9ab7f514837a/plmustache.control.in)
- [sql/plmustache.sql](https://github.com/PostgREST/plmustache/blob/024b550d4fe2bbb4e122d2b3216a9ab7f514837a/sql/plmustache.sql)
- [src/plmustache.c](https://github.com/PostgREST/plmustache/blob/024b550d4fe2bbb4e122d2b3216a9ab7f514837a/src/plmustache.c)
- [.github/workflows/ci.yaml](https://github.com/PostgREST/plmustache/blob/024b550d4fe2bbb4e122d2b3216a9ab7f514837a/.github/workflows/ci.yaml)

`plmustache` 通过 PostgreSQL 函数渲染 Mustache 模板。当前源码的 Makefile 生成 0.1 版控制文件和安装 SQL；历史发布标签 v0.1 中的 SQL 版本却是 0.0.1。本文对应固定源码快照，不声称这两个版本之间存在受支持的升级路径。C 动态库依赖 libmustach。当前 CI 覆盖 PostgreSQL 14–17，较旧 README 列出 12–16。

### 渲染模板

安装匹配的生成控制文件、SQL 和动态库后，由超级用户创建扩展。普通使用不要求共享预加载或服务器重启：

```sql
CREATE EXTENSION plmustache;
CREATE FUNCTION greeting(person text) RETURNS text
LANGUAGE plmustache AS $$Hello {{person}}!$$;
SELECT greeting('PostgreSQL');
```

扩展控制文件没有把安装标记为受信任，但安装后的过程语言本身为 TRUSTED，这是两个不同的权限边界。命名函数参数提供模板变量，输出作为文本返回。库包含调用、内联及验证处理函数，但内联处理函数会拒绝 DO 块。

### 转义与区块

双大括号执行 HTML 转义，三重大括号原样返回值，例如：

```sql
CREATE FUNCTION display_value(value text) RETURNS text
LANGUAGE plmustache AS $$Escaped: {{value}} Raw: {{{value}}}$$;
CREATE FUNCTION visible_message(show bool) RETURNS text
LANGUAGE plmustache AS $${{#show}}Visible{{/show}}{{^show}}Hidden{{/show}}$$;
SELECT display_value('<script>example</script>');
SELECT visible_message(false);
```

输出到页面时，应把未转义内容视为不可信输入。HTML 转义不是 SQL 引用、JSON 编码或通用注入防护。README 展示布尔区块和数组迭代，但路线图未同步，应按源码测试验证具体参数类型。本快照没有确立对 partials、继承及 lambda 的支持。

### 维护

生成的 SQL/control 版本与安装的 C 动态库须保持一致。已有后端会保留加载的旧库，直到连接断开。源码布局和版本已不同于历史发布，应先在隔离数据库审查安装及迁移，不应假定旧 v0.1 产物安装的就是当前 SQL 版本。
