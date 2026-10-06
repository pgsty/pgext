## 用法

来源：

- [packages/metaschema-modules/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/metaschema-modules/README.md)
- [packages/metaschema-modules/sql/metaschema-modules--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/metaschema-modules/sql/metaschema-modules--0.47.0.sql)
- [packages/metaschema-modules/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/metaschema-modules/Makefile)
- [packages/metaschema-modules/metaschema-modules.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/metaschema-modules/metaschema-modules.control)

`metaschema-modules` 为 metaschema 模型增加可复用应用模块的配置记录，包括账号、成员关系、消息和认证集成。

### 核心用法

```sql
CREATE EXTENSION "metaschema-modules" CASCADE;
SELECT * FROM metaschema_modules_public.connected_accounts_module LIMIT 10;
```

### 运行边界

这些表描述模块配置及对基础 metaschema 的引用，不会启动认证提供方或网络服务。应通过配套平台工具应用变更，使引用的模式、字段、策略与生成 DDL 保持一致，并将写入限制在管理员范围内。

0.47.0 是 SQL/PLpgSQL 扩展，没有自己的共享库，也不要求预加载。须先安装配套版本的依赖：`plpgsql`、`uuid-ossp`、`metaschema-schema`、`pgpm-verify`。 安装 SQL 要求平台角色 `administrator`, `authenticated` 已存在，须先使用上游角色初始化流程，并审查授权。 控制文件允许非超级用户安装，但没有标记为 trusted；依赖、模式创建和角色授权权限仍须满足。上游未声明当前 PostgreSQL 主版本矩阵。
