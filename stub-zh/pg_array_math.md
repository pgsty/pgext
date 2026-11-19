## 用法

来源：

- [官方资料](https://github.com/idexsts/pg_array_math/blob/3e9877014f567fb23b643db9a5c5db259419ead6/README.md)
- [官方资料](https://github.com/idexsts/pg_array_math/blob/3e9877014f567fb23b643db9a5c5db259419ead6/pg_array_math.control)
- [官方资料](https://github.com/idexsts/pg_array_math/blob/3e9877014f567fb23b643db9a5c5db259419ead6/pg_array_math--1.0.sql)
- [官方资料](https://github.com/idexsts/pg_array_math/blob/3e9877014f567fb23b643db9a5c5db259419ead6/pg_array_math.c)

pg_array_math 为 smallint、integer、bigint、real、double precision 或 numeric 的一维数组提供 immutable、strict 数学函数。已审核源码未声明许可证或支持的 PostgreSQL 主版本矩阵。

### 数组操作

```sql
CREATE EXTENSION pg_array_math;
SELECT array_add(ARRAY[1,2,3], ARRAY[4,5,6]);
SELECT array_mul_scalar(ARRAY[1,2,3], 2);
SELECT array_sqrt(ARRAY[4.0,9.0]::numeric[]);
SELECT array_log(ARRAY[1.0,2.0]::double precision[]);
```

接口包括 array_add、array_sub、array_mul、array_div、array_add_scalar、array_mul_scalar、array_power、array_sqrt、array_abs、array_log 和 array_exp。双数组操作要求元素类型和长度相同。结果保留第一个数组的下界，NULL 元素按位置传播，NULL 输入返回 NULL。

### 限制

多维数组、不支持的类型、除零、定义域错误和数值溢出会报错。array_log 与 array_exp 仅接受浮点或 numeric 数组。结果沿用输入类型：整数除法截断，整数幂与平方根不会自动成为浮点近似值。计算前应明确转换类型。此 C 扩展可迁移，无需预加载；安装遵循普通非 trusted 扩展的管理员权限要求。
