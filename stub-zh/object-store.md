## 用法

来源：

- [packages/object-store/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/object-store/README.md)
- [packages/object-store/sql/object-store--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/object-store/sql/object-store--0.47.0.sql)
- [packages/object-store/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/object-store/Makefile)
- [packages/object-store/object-store.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/object-store/object-store.control)

`object-store` 保存版本化 JSONB 对象及子节点引用，冻结后按内容寻址，使不同版本能够共享结构。

### 核心用法

```sql
CREATE EXTENSION "object-store" CASCADE;
SELECT * FROM object_store_public.object LIMIT 10;
```

### 运行边界

`object_store_public.get_node_at_path`、`object_store_public.insert_node_at_path`、`object_store_public.update_node_at_path` 与 `object_store_public.freeze_objects` 实现树导航和写时复制操作。冻结的行不能通过普通操作编辑，保留根节点时也须保留引用的子对象；不可变性和内容哈希不能替代权限控制与备份。

0.47.0 是 SQL/PLpgSQL 扩展，没有自己的共享库，也不要求预加载。须先安装配套版本的依赖：`plpgsql`、`pgcrypto`、`uuid-ossp`、`pgpm-verify`。 安装 SQL 要求平台角色 `authenticated` 已存在，须先使用上游角色初始化流程，并审查授权。 控制文件允许非超级用户安装，但没有标记为 trusted；依赖、模式创建和角色授权权限仍须满足。上游未声明当前 PostgreSQL 主版本矩阵。
