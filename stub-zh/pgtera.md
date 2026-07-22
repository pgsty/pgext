## 用法

来源：

- [Official README](https://github.com/f-prime/pgtera/blob/9ca7b37b4448c6583351777a5530f1582f99867d/README.md)
- [Rust implementation](https://github.com/f-prime/pgtera/blob/9ca7b37b4448c6583351777a5530f1582f99867d/src/lib.rs)
- [Extension control file](https://github.com/f-prime/pgtera/blob/9ca7b37b4448c6583351777a5530f1582f99867d/pgtera.control)

pgtera 从 PostgreSQL 渲染 Tera 模板。它可以从服务端本地 glob 加载命名模板，也可用 JSON 上下文渲染模板字符串，使 PostgREST 等系统可以提供数据库驱动的 HTML。扩展会读取数据库服务器文件系统并在后端中渲染应用输出，因此必须严格限定函数权限和内容边界。

### 核心流程

由管理员设置一个受控模板 glob，再用由 name/value 上下文对象组成的数组渲染命名模板：

```sql
CREATE EXTENSION pgtera;

SELECT pgtera_set_render_path('/srv/app/templates/**/*.html');

SELECT pgtera_render(
  'index.html',
  '[{"name":"title","value":"Status"},
    {"name":"items","value":["ready","healthy"]}]'
);
```

若模板直接以文本提供，可使用 `pgtera_render_str`。两条路径解析相同的 JSON 上下文结构，而且字符串渲染路径仍会初始化已配置的模板 glob。

### 重要对象

- `pgtera.render_path` 保存配置的模板 glob；使用最近插入的一行。
- `pgtera_set_render_path` 插入新的 glob。
- `pgtera_render` 加载模板集合并渲染一个命名模板。
- `pgtera_render_str` 使用给定上下文渲染模板文本。

### 安全与运维说明

PostgreSQL 操作系统账户必须能读取所有匹配文件。要限制 path 表及三个函数的权限，防止不可信用户选择任意服务端文件，或用大型模板集消耗后端 CPU 与内存。path setter 拼接 SQL 文本且没有安全参数绑定，因此不能传入用户控制的字符串。Tera 转义规则和 `safe` 过滤器决定输出安全性；提供 HTML 前必须审查模板。若渲染需要沙箱、缓存或文件系统隔离，应优先放在应用层。
