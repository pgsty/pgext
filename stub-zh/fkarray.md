## 用法

来源：

- [README.md](https://github.com/pjungwir/fkarray/blob/434290df52c2163676ee7f7d73888ab1b1ec19c8/README.md)
- [fkarray--1.0.0.sql](https://github.com/pjungwir/fkarray/blob/434290df52c2163676ee7f7d73888ab1b1ec19c8/fkarray--1.0.0.sql)
- [fkarray.c](https://github.com/pjungwir/fkarray/blob/434290df52c2163676ee7f7d73888ab1b1ec19c8/fkarray.c)
- [fkarray.control](https://github.com/pjungwir/fkarray/blob/434290df52c2163676ee7f7d73888ab1b1ec19c8/fkarray.control)

`fkarray` 1.0.0 检查数组列中的每个值是否引用了已有的唯一键或主键。它以触发器和函数实现约束，不是 PostgreSQL 原生数组外键。

### 核心用法

```sql
CREATE EXTENSION fkarray;
CREATE EXTENSION intarray;
CREATE TABLE colors (id integer PRIMARY KEY);
CREATE TABLE products (id integer PRIMARY KEY, color_ids integer[]);
SELECT fkarray_create('public','products','color_ids','public','colors','id');
INSERT INTO colors VALUES (1), (2);
INSERT INTO products VALUES (1, ARRAY[1,2]);
```

### 运行边界

`fkarray_create()` 接收引用方与被引用方的模式、表和列名，以及可选的 `on_delete`、`on_update` 动作。两者默认都是 `no_action`；`cascade` 会从数组中移除被引用的值，更新键值时也是移除旧值。`fkarray_drop()` 删除生成的约束触发器。数组元素与目标列类型必须一致，目标列必须具有唯一索引或主键索引。整数数组需要 `intarray`，大整数数组需要 `bigintarray`，其他类型使用 PostgreSQL 常规运算符。这些是按类型决定的要求，不是 control 文件中的无条件依赖。control 将扩展标记为可信且不可迁移；调用者仍须有权限在相关表上创建触发器。
