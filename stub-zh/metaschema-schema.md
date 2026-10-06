## 用法

来源：

- [packages/metaschema-schema/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/metaschema-schema/README.md)
- [packages/metaschema-schema/sql/metaschema-schema--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/metaschema-schema/sql/metaschema-schema--0.47.0.sql)
- [packages/metaschema-schema/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/metaschema-schema/Makefile)
- [packages/metaschema-schema/metaschema-schema.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/metaschema-schema/metaschema-schema.control)

`metaschema-schema` 用元数据表保存声明式数据库模型，为外围平台描述数据库、模式、表、字段、约束、索引和授权。

### 核心用法

```sql
CREATE EXTENSION "metaschema-schema" CASCADE;
SELECT * FROM metaschema_public.database LIMIT 10;
SELECT * FROM metaschema_public.schema LIMIT 10;
```

### 运行边界

固定的 `metaschema_public` 与 `metaschema_private` 模式包含验证辅助函数，以及变更和任务集成。元数据行并不证明物理 DDL 已生效，须由平台的协调程序处理模型。修改记录前应审查生成标识符、权限、版本配套及迁移影响。

0.47.0 是 SQL/PLpgSQL 扩展，没有自己的共享库，也不要求预加载。须先安装配套版本的依赖：`citext`、`hstore`、`pgpm-inflection`、`pgpm-database-jobs`、`pgpm-types`、`pgcrypto`、`plpgsql`、`postgis`、`uuid-ossp`、`pgpm-verify`。 安装 SQL 要求平台角色 `authenticated` 已存在，须先使用上游角色初始化流程，并审查授权。 控制文件允许非超级用户安装，但没有标记为 trusted；依赖、模式创建和角色授权权限仍须满足。上游未声明当前 PostgreSQL 主版本矩阵。
