## 用法

来源：

- [官方 README](https://github.com/aquameta/meta_triggers/blob/3bc68b67afea896300adf828b52d5dbbf1d63549/README.md)
- [官方扩展控制文件](https://github.com/aquameta/meta_triggers/blob/3bc68b67afea896300adf828b52d5dbbf1d63549/meta_triggers.control)
- [官方触发器定义](https://github.com/aquameta/meta_triggers/blob/3bc68b67afea896300adf828b52d5dbbf1d63549/000-triggers.sql)

`meta_triggers` 让 Aquameta `meta` 扩展中的部分视图变得可写。针对这些目录视图的插入、更新和删除会转换为真实 DDL 或管理命令，使元数据驱动工具能通过行记录创建和修改数据库对象。

### 核心流程

扩展安装到 `meta` 模式，并依赖 `meta`：

```sql
CREATE EXTENSION meta;
CREATE EXTENSION meta_triggers;

INSERT INTO meta.schema(name) VALUES ('app_data');
UPDATE meta.schema
SET name = 'application_data'
WHERE name = 'app_data';
DELETE FROM meta.schema WHERE name = 'application_data';
```

上游 SQL 为大量 `meta` 视图增加 `INSTEAD OF` 触发器与语句构造函数，涵盖模式、类型、序列、表、视图、列、外键、函数、触发器、角色及成员关系、权限、行级安全策略、约束、扩展、外部数据对象、类型转换与事件触发器等。每个视图的可写字段不同，应以对应 `meta` 视图定义作为对象模型。

### 安全与事务行为

写操作使用调用角色的权限执行 DDL，通常参与调用者事务。应将这些视图的写权限视为管理权限：一次行变更就可能创建或删除模式、角色、函数、扩展、策略或其他数据库对象。SQL 会动态构造命令，因此必须校验应用输入并严格控制授权。`meta_triggers` 不是脱离数据库的元数据存储，底层 PostgreSQL 目录会立即改变。
