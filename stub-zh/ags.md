## 用法

来源：

- [扩展控制文件](https://github.com/x4m/ags/blob/ceaf4c375e5142db193eae616e10d6704ad28c7e/ags.control)
- [1.0 版安装 SQL](https://github.com/x4m/ags/blob/ceaf4c375e5142db193eae616e10d6704ad28c7e/ags--1.0.sql)
- [访问方法实现](https://github.com/x4m/ags/blob/ceaf4c375e5142db193eae616e10d6704ad28c7e/gist.c)

ags 安装一个派生自 PostgreSQL GiST 代码的实验性索引访问方法。1.0 版公开名为 ags 的访问方法，以及针对 cube 类型的默认操作符类 ags_cube_ops。它用于 cube 的重叠、包含、相等与距离检索。

### 基本用法

该扩展未声明 cube 依赖，因此应先显式安装 cube：

```sql
CREATE EXTENSION cube;
CREATE EXTENSION ags;

CREATE TABLE ags_demo (
    id bigint GENERATED ALWAYS AS IDENTITY,
    point cube NOT NULL
);

INSERT INTO ags_demo (point)
VALUES (cube(ARRAY[0.0, 0.0])),
       (cube(ARRAY[1.0, 2.0])),
       (cube(ARRAY[5.0, 5.0]));

CREATE INDEX ags_demo_point_idx ON ags_demo USING ags (point);

SELECT id, point
FROM ags_demo
ORDER BY point <-> cube(ARRAY[0.0, 0.0])
LIMIT 2;
```

该操作符类还注册了 cube 重叠与包含操作符。在依赖此索引前，应在代表性数据上用 EXPLAIN 核对执行计划。

### 注意事项

- 仓库没有 README、发布历史、许可证文件或 PostgreSQL 兼容性矩阵。
- 大部分实现是 2019 年 PostgreSQL GiST 内部代码的私有副本。这些内部 API 对大版本敏感，因此必须针对每个 PostgreSQL 大版本单独构建并做回归测试。
- 1.0 版只提供 cube 操作符类，未在控制文件中声明 cube 依赖，且禁用并行索引扫描。
- 在针对目标服务器验证正确性、崩溃安全、恢复行为与升级路径之前，应将其视为研究代码。
