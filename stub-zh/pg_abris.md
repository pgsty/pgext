## 用法

来源：

- [0.0.1 版本安装 SQL](https://github.com/abris-platform/pg_abris/blob/cca112ac51e23ee130be17a25cb0724a085aa58e/pg_abris--0.0.1.sql)
- [扩展控制文件](https://github.com/abris-platform/pg_abris/blob/cca112ac51e23ee130be17a25cb0724a085aa58e/pg_abris.control)
- [官方上游构建说明](https://github.com/abris-platform/pg_abris/blob/cca112ac51e23ee130be17a25cb0724a085aa58e/README.md)

`pg_abris` 是面向 Abris 平台的纯 PL/pgSQL 元数据层。它创建固定的 `meta` schema，其视图把数据库对象暴露为 Abris 实体、属性、关系和投影，而触发器让部分元数据界面可写。

### 核心流程

```sql
CREATE EXTENSION "uuid-ossp";
CREATE EXTENSION pg_abris;

SELECT entity_id, schema_name, table_name, primarykey
FROM meta.entity
ORDER BY schema_name, table_name;
```

该扩展需要 `uuid-ossp`。安装 SQL 会直接创建并填充 `meta` schema；只应在专用于理解 Abris 元数据契约之软件的数据库中使用。

### 安装对象

- `meta.entity`、`meta.property` 和 `meta.relation` 通过基于系统目录的视图展示关系、列与关联。
- 投影、菜单、页面与额外属性表保存应用展示元数据。
- INSTEAD OF 触发器与表触发器把对选定元数据视图的写入转换成底层 DDL 或元数据表变更。
- `meta.clean()` 删除所引用系统目录对象已不存在的过期辅助行。

控制文件称版本 `0.0.1` 可迁移，但 SQL 硬编码了 `meta`，因此移动扩展并不安全。上游没有提供用户指南、受支持的 PostgreSQL 矩阵、发布历史或明确的许可证文件。使用前应审查完整安装 SQL、对象所有权、触发器生成的 DDL 与权限；不要向不可信角色暴露可写元数据视图，也不要把它当成通用元数据框架。
