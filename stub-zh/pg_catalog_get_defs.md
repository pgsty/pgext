## 用法

来源：

- [Pinned official README](https://github.com/aquameta/pg_catalog_get_defs/blob/93c358ae1851913960861bb010923a245e538b6e/README.md)
- [Pinned extension SQL](https://github.com/aquameta/pg_catalog_get_defs/blob/93c358ae1851913960861bb010923a245e538b6e/pg_get_typedef.sql)

`pg_catalog_get_defs` 增加一组 SQL 函数，可依据 PostgreSQL 系统目录重建 `CREATE TYPE` 语句。它为工具和检查场景补足类型定义内省能力；不会执行返回的 DDL，也不会重建依赖对象。

### 核心流程

创建扩展，将类型 OID 传给 `get_typedef(oid)`：

```sql
CREATE EXTENSION pg_catalog_get_defs;

CREATE TYPE deployment_state AS ENUM ('pending', 'ready', 'failed');

SELECT get_typedef('deployment_state'::regtype::oid);
```

结果是类似 `CREATE TYPE deployment_state AS ENUM (...)` 的文本。在其他位置重放前，应把它当作生成输出进行审阅、补全限定名，并与依赖 DDL 正确排序。

### 分类型函数

- `get_typedef_enum(oid)` 按排序顺序重建枚举标签。
- `get_typedef_composite(oid)` 重建未删除的复合属性及非默认排序规则。
- `get_typedef_domain(oid)` 重建基础类型和默认表达式。
- `get_typedef_range(oid)` 重建子类型、可选运算符类/排序规则、规范化函数和子类型差值函数。
- `get_typedef_base(oid)` 重建基础类型的存储与 I/O 属性。
- `get_typedef(oid)` 根据 `pg_type.typtype` 分派；未定义的伪类型会生成壳类型 DDL，已定义的伪类型会报错。

### 边界

生成语句取决于调用者可见的目录状态。它不是完整转储：注释、权限、所有者、扩展成员关系、域约束、数组类型、类型转换、运算符、依赖函数和外围模式设置都不会被重建成迁移包。

名称通过 `regtype` 或目录格式化逻辑渲染，可能受 `search_path` 影响。应用到其他数据库前，应补上模式限定并验证生成的 DDL。该项目规模很小，SQL 遵循较旧的目录布局，因此要在目标 PostgreSQL 版本上测试每种受支持的类型。无需预加载或重启。
