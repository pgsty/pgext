## 用法

来源：

- [官方 README](https://github.com/yanliu8/cbtree/blob/871f57358cb297f22be6f419a8e52dc86567ab7c/README.md)
- [扩展 SQL](https://github.com/yanliu8/cbtree/blob/871f57358cb297f22be6f419a8e52dc86567ab7c/cbtree--1.0.sql)
- [访问方法实现](https://github.com/yanliu8/cbtree/blob/871f57358cb297f22be6f419a8e52dc86567ab7c/cbtree.c)

`cbtree` 是一种实验性的计数 B-tree 访问方法，它把索引中的 `integer` 值解释为从 1 开始的插入或查询位置。该项目针对 PostgreSQL 10.1 编写，适合研究位置索引，不应作为通用或仍受支持的 B-tree 替代品。

### 核心流程

创建一个整数哑列和单列 `cbtree` 索引。写入值表示期望插入的位置；大于当前行数的值会将行追加到末尾。查询只支持等值策略。

```sql
CREATE EXTENSION cbtree;

CREATE TABLE demo (
    payload text,
    sequence_position integer NOT NULL
);
CREATE INDEX demo_position_idx
    ON demo USING cbtree (sequence_position);

INSERT INTO demo VALUES ('first', 1);
INSERT INTO demo VALUES ('new first', 1);
INSERT INTO demo VALUES ('append', 999);

SELECT payload
FROM demo
WHERE sequence_position = 1;
```

为已有数据建索引时，初始顺序来自堆表扫描顺序。哑列中保存的值是插入指令，并不是 PostgreSQL 会自动重编号的持久排名。

### 对象与限制

- `cbtree` 是索引访问方法。
- `int4_ops` 是它唯一的操作符类，适用于 `integer`；只注册了 `=`。
- 实现只读取第一个索引值。即使处理器声明支持多列，也只应使用一个非空 `integer` 列。
- 扩展还会安装 `delta`、`delta_actual_pos`、`delta_sel`、`delta_del`、`delta_ins` 和 `auto_vacuum`。这些原型辅助对象包含不安全的动态 SQL，以及无法工作的触发器内 `VACUUM` 路径；不要开放或依赖它们。

### 运维注意事项

只能使用正数位置。生产构建不会安全拒绝零、负数或空值。该访问方法不支持唯一性、排序、反向扫描、空值查询、数组查询或并行索引扫描。

代码使用 PostgreSQL 10 时代的索引 API，且没有维护中的兼容矩阵。任何实验室使用前都要针对目标 PostgreSQL 源码树构建并做破坏性测试；不要用它保存权威顺序或生产数据。
