## 用法

来源：

- [0.0.1 版本安装 SQL](https://github.com/abris-platform/pg_abris/blob/cca112ac51e23ee130be17a25cb0724a085aa58e/pg_abris--0.0.1.sql)
- [扩展控制文件](https://github.com/abris-platform/pg_abris/blob/cca112ac51e23ee130be17a25cb0724a085aa58e/pg_abris.control)
- [上游构建说明](https://github.com/abris-platform/pg_abris/blob/cca112ac51e23ee130be17a25cb0724a085aa58e/README.md)

`pg_abris` 是面向 Abris 平台的纯 PL/pgSQL 元数据组件。它创建固定的 `meta` schema，其中包含元数据表、视图和辅助函数，并依赖 `uuid-ossp`。

```sql
CREATE EXTENSION "uuid-ossp";
CREATE EXTENSION pg_abris;
SELECT nspname FROM pg_namespace WHERE nspname = 'meta';
```

控制文件虽声明可迁移，SQL 却硬编码了 `meta`，因此移动扩展并不安全。上游几乎没有用户文档，也没有发布历史、PostgreSQL 支持声明或明确许可证文件。应审查安装 SQL 创建的每个对象与权限，把它隔离在 Abris 专用数据库中，不要把版本 `0.0.1` 当作通用元数据框架。
