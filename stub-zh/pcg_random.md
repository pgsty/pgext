## 用法

来源：

- [上游 README](https://github.com/cayetanobv/pgsql-pcg-random/blob/cfac9eafc3d34f16f43a398fd2b14b692cc359a3/README.md)
- [扩展 control 文件](https://github.com/cayetanobv/pgsql-pcg-random/blob/cfac9eafc3d34f16f43a398fd2b14b692cc359a3/pcg_random.control)
- [SQL 安装脚本](https://github.com/cayetanobv/pgsql-pcg-random/blob/cfac9eafc3d34f16f43a398fd2b14b692cc359a3/pcg_random--0.0.1.sql)
- [C 实现](https://github.com/cayetanobv/pgsql-pcg-random/blob/cfac9eafc3d34f16f43a398fd2b14b692cc359a3/src/pcg_random.c)

`pcg_random` `0.0.1` 版提供由 PCG 生成器支持的标量、有界值、数组和有界数组函数。可选种子参数用于选择序列；实现还会在每次调用中混入时间和一个地址。

### 示例

```sql
CREATE EXTENSION pcg_random;
SELECT pcg_random();
SELECT pcg_random_bound(100);
SELECT pcg_random_array(5);
SELECT pcg_random_bound_array(5, 100);
```

它不是密码学随机源。边界值必须严格为正，数组长度应为较小的非负数；C 代码没有充分验证零值或负数边界和长度。更重要的是，安装 SQL 把非确定性函数声明成 `IMMUTABLE`，可能导致规划器常量折叠或其他错误复用。源码停留在 2016 年且没有当前兼容矩阵；应优先使用仍在维护的 PostgreSQL 随机数设施。
