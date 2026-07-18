## 用法

来源：

- [上游固定版本 README](https://github.com/celian-garcia/firefly-database/blob/1dfe6cf49a670b4d3b2c1ace40969cdab58189fa/README.md)
- [0.2.0 版本安装 SQL](https://github.com/celian-garcia/firefly-database/blob/1dfe6cf49a670b4d3b2c1ace40969cdab58189fa/firefly--0.2.0.sql)
- [固定版本扩展控制文件](https://github.com/celian-garcia/firefly-database/blob/1dfe6cf49a670b4d3b2c1ace40969cdab58189fa/firefly.control)

firefly 0.2.0 是某个三维点编辑应用的数据库模式。它创建 task 与 fpoint3d 表、每任务操作序列、增删历史函数、汇总函数、触发器、枚举以及破坏性的重置函数。它依赖 PL/pgSQL 与 PostGIS，并不是通用的变更数据捕获扩展。

### 应用示例

该扩展会创建名称很通用且未限定模式的对象，因此应使用全新数据库：

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION firefly;

INSERT INTO task (name, description)
VALUES ('demo', 'isolated firefly test')
RETURNING id;

SELECT save_add_operation(
    ST_Force3D(ST_Point(1, 2)),
    0
);

SELECT * FROM collect_operations(0, 0);
```

每次插入任务都会按数字任务 ID 动态创建一个序列。点记录使用整数数组保存历史，并按奇偶性解释当前为新增或删除状态。

### 应用与生命周期限制

安装 SQL 硬编码了表、序列、类型和函数，却没有限定模式或保护 search_path。控制文件虽然声明可重定位，但移动扩展或改变 search_path 可能导致名称解析失败，甚至让函数使用错误对象。控制文件还把 module_pathname 指向 pgtap，而安装内容其实是纯 SQL，说明打包并未整理完整。

实现通过距离匹配点，却没有唯一约束；每个任务创建一个序列，并线性扫描操作数组。并发编辑、近似重复点、任务删除、序列所有权、模式迁移和重置行为都需要针对应用验证。empty_firefly_database 会清空全部应用状态并重置序列。

上游示例面向 PostgreSQL 9.6，仓库自 2018 年后没有更新。只能在空白、可丢弃的数据库中评估；撤销公共函数执行权，仅授权应用角色，并在测试前保留独立备份。
