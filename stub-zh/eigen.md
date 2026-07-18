## 用法

来源：

- [上游 README](https://github.com/tlb-lab/pgeigen/blob/af9c58f84d58e629ed2a2329ffcbe8d6d5c95cff/README.rst)
- [扩展 control 文件](https://github.com/tlb-lab/pgeigen/blob/af9c58f84d58e629ed2a2329ffcbe8d6d5c95cff/eigen.control)
- [1.1 版 SQL](https://github.com/tlb-lab/pgeigen/blob/af9c58f84d58e629ed2a2329ffcbe8d6d5c95cff/sql/eigen--1.1.sql)
- [C++ 矩阵实现](https://github.com/tlb-lab/pgeigen/blob/af9c58f84d58e629ed2a2329ffcbe8d6d5c95cff/src/matrixxd.cpp)

`eigen` `1.1` 版把 Eigen C++ 库封装为 PostgreSQL 域和函数，用于三维向量、整数/双精度数组及双精度矩阵。它提供逐元素算术、向量点积/叉积、范数与距离、集合/相似度运算、矩阵乘积、归约和拼接辅助函数。

### 示例

```sql
CREATE EXTENSION eigen;
SELECT vector3d_dot(
  ARRAY[1.0, 0.0, 0.0]::vector3d,
  ARRAY[0.0, 1.0, 0.0]::vector3d
);
SELECT vector3d_norm(ARRAY[3.0, 4.0, 0.0]::vector3d);
```

输入是由域约束的 PostgreSQL 数组；昂贵运算前应强制维度正确并拒绝空元素。源码标称面向 PostgreSQL 9.1+ 与 Eigen 3，但仓库自 2020 年后没有更新，也没有当前兼容矩阵。它还定义了许多通用函数和运算符名称，因此应安装到规划好的模式，并测试名称冲突和数值边界情况。
