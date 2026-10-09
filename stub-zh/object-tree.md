## 用法

来源：

- [packages/object-tree/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/object-tree/README.md)
- [packages/object-tree/sql/object-tree--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/object-tree/sql/object-tree--0.47.0.sql)
- [packages/object-tree/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/object-tree/Makefile)
- [packages/object-tree/object-tree.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/object-tree/object-tree.control)

`object-tree` 在 PGPM 对象存储上增加仓库、提交和命名引用。`object_tree_public.init_empty_repo` 初始化仓库，`object_tree_public.rev_parse` 解析引用，提交辅助函数同时更新对象树与历史。

### 核心用法

```sql
CREATE EXTENSION "object-tree";
SELECT scope_id, store_id, message, tree_id, date
FROM object_tree_public.commit
ORDER BY date DESC
LIMIT 10;
```

### 运行边界

`object_tree_public.set_many_and_commit` 批量编辑路径，`object_tree_public.set_props_and_commit` 更新节点数据并保留子节点。安装会创建持久化的提交、引用和仓库表，并要求应用角色已就绪，还须提供用于 UUID 默认值的 `uuidv7()`。仓库及作用域 ID 必须属于已配置的应用。历史会占用数据库存储，这些函数不会自动回收历史对象。它实现对象版本管理，不会自动管理任意 SQL 表的版本。 依赖为 `plpgsql`, `pgcrypto`, `object-store`, `pgpm-verify`。 安装脚本会向预先存在的角色 `authenticated` 授权。 这是纯 SQL 模块，无需预加载动态库或重启；该修订未公布 PostgreSQL 主版本支持矩阵。模式布局固定，不可迁移。启用前须部署相匹配的完整依赖。
