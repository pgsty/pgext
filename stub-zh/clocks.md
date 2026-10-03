## 用法

来源：

- [README.md](https://github.com/nuno-faria/crdv/blob/15e58bb9379e77a77de56517ff5bf6a514ff56d4/README.md)
- [schema/sql/01-types.sql](https://github.com/nuno-faria/crdv/blob/15e58bb9379e77a77de56517ff5bf6a514ff56d4/schema/sql/01-types.sql)
- [schema/clocks/clocks.control](https://github.com/nuno-faria/crdv/blob/15e58bb9379e77a77de56517ff5bf6a514ff56d4/schema/clocks/clocks.control)
- [schema/clocks/clocks--1.0.sql](https://github.com/nuno-faria/crdv/blob/15e58bb9379e77a77de56517ff5bf6a514ff56d4/schema/clocks/clocks--1.0.sql)
- [schema/clocks/clocks.c](https://github.com/nuno-faria/crdv/blob/15e58bb9379e77a77de56517ff5bf6a514ff56d4/schema/clocks/clocks.c)

`clocks` 1.0 为 CRDV 向量时钟及混合逻辑时钟提供 C 辅助函数。CRDV 是在 PostgreSQL 16 上测试的研究原型。该组件自身不提供复制功能，也不创建前置类型。

### 基本用法

```sql
-- In a fresh database, matching the CRDV bootstrap:
CREATE TYPE hlc AS (physical_time bigint, logical_time int);
CREATE DOMAIN vclock AS bigint[];
CREATE EXTENSION clocks;
SELECT vclock_lte(ARRAY[1,2]::vclock, ARRAY[2,3]::vclock);
SELECT vclock_max(ARRAY[1,3]::vclock, ARRAY[2,2]::vclock);
SELECT next_hlc(ROW(0,0)::hlc);
```

### 前置条件与函数

初始化 CRDV 数据库时应使用已有的 `hlc` 与 `vclock` 定义，不要重新定义在用类型；创建扩展时须能找到它们。`vclock_lte` 按分量检查先后关系，`vclock_max` 返回各分量的最大值，`next_hlc` 根据当前时间及先前逻辑状态推进混合时钟。时钟维度与排序约定须与应用保持一致。

安装需要超级用户，未声明共享预加载要求。完整 CRDV 集群的初始化及逻辑复制配置属于外层 CRDV 系统，而非这三个函数。备份或恢复扩展时须保留其类型定义。
