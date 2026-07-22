## 用法

来源：

- [扩展控制文件](https://github.com/JonM0/hyperspace/blob/99ebb600ef1b5e606a54e79978a6824a33492d3f/hyperspace.control)
- [类型定义](https://github.com/JonM0/hyperspace/blob/99ebb600ef1b5e606a54e79978a6824a33492d3f/sql/type.sql)
- [操作符与 SP-GiST 操作符类](https://github.com/JonM0/hyperspace/blob/99ebb600ef1b5e606a54e79978a6824a33492d3f/sql/kd.sql)

`hyperspace` 版本 `1.0` 增加了四维几何类型与索引。它适合需要在四轴空间内按距离排序或进行包含判断的点数据。

### 核心流程

```sql
CREATE EXTENSION hyperspace;
CREATE TABLE samples (p point4d);
CREATE INDEX samples_p_spgist ON samples USING spgist (p);

INSERT INTO samples VALUES ('(1,2,3,4)'), ('(5,6,7,8)');
SELECT p, p <-> '(0,0,0,0)'::point4d AS distance
FROM samples
WHERE p <@ '((-10,-10,-10,-10),(10,10,10,10))'::box4d
ORDER BY p <-> '(0,0,0,0)'::point4d
LIMIT 5;
```

默认的 `point4d_kd` SP-GiST 操作符类支持相等判断、矩形或圆形包含判断，以及通过 `<->` 进行最近邻排序。

### 对象与注意事项

主要类型是 `point4d`、`box4d` 和 `circle4d`，并提供同名构造函数。点支持相等、加法、`sum` 与欧氏距离。包含关系使用 `<@`，其交换操作符为 `@>`。

`point4d` 还提供默认的字典序 B-tree 操作符类，以及按模长比较的 `point4d_abs_ops` 操作符类。点的文本输入格式是 `(x,y,z,w)`，矩形由两个点组成，圆形由一个点和半径组成。这些都是自定义 C 类型；持久化数据前，应在目标 PostgreSQL 构建上验证二进制兼容性与查询计划。
