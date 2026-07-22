## 用法

来源：

- [目录版本对应的官方仓库 README](https://github.com/a5221985/PhaseZero_Scripts/blob/443e3a5e59d2816b7b8d7ee0129c8685f3f11d8a/README.md)
- [目录版本对应的构建规则](https://github.com/a5221985/PhaseZero_Scripts/blob/443e3a5e59d2816b7b8d7ee0129c8685f3f11d8a/Makefile)
- [目录版本对应的表 DDL](https://github.com/a5221985/PhaseZero_Scripts/blob/443e3a5e59d2816b7b8d7ee0129c8685f3f11d8a/sql/create_table_ddl.sql)
- [目录版本对应的索引 DDL](https://github.com/a5221985/PhaseZero_Scripts/blob/443e3a5e59d2816b7b8d7ee0129c8685f3f11d8a/sql/create_index.sql)
- [目录版本对应的约束 DDL](https://github.com/a5221985/PhaseZero_Scripts/blob/443e3a5e59d2816b7b8d7ee0129c8685f3f11d8a/sql/create_constraint.sql)

`scripts` 1.4 不是通用函数库，而是打包了 PhaseZero 应用 DDL：创建 pzv_aftermarket schema、固定的业务与暂存表、索引和外键。只有在审核完整生成 SQL 后，才应为该遗留应用安装它。

### 核心流程

Makefile 通过串接表、索引和约束文件生成扩展脚本。安装前应检查生成产物，并先在一次性数据库中试运行。

```sql
CREATE EXTENSION scripts;

SELECT n.nspname, c.relname, c.relkind
FROM pg_class AS c
JOIN pg_namespace AS n ON n.oid = c.relnamespace
WHERE n.nspname = 'pzv_aftermarket'
ORDER BY c.relkind, c.relname;
```

### 安装对象

- 核心表包括 company、category、attribute、part、scope_part、display_metadata、orders 和 order_parts。
- 暂存/辅助表包括 stage_category、stage_attribute、stage_part 和 temp_part_child。
- 脚本为其中若干关系添加应用特定的索引与外键，不提供面向用户的 SQL 函数。

### 关键限制

- 对象名、列和约束均为硬编码。安装可能与已有 pzv_aftermarket schema 或部分同名对象冲突，而且这些 DDL 不是按可重复 schema 迁移编写的。
- 仓库中的独立数据库创建文件不在扩展构建规则内；CREATE EXTENSION 不会创建数据库。
- control 文件注释描述了无关的键值数据类型，而实际生成 SQL 是应用 schema DDL。应信任审核过的 SQL，而不是该元数据注释。
- 上游已停止维护并指向老旧 PGXS 路径。应先备份并测试安装/卸载行为，新部署则应选用仍维护的迁移系统。
