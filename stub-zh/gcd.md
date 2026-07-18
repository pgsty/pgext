## 用法

来源：

- [固定提交的上游 README](https://github.com/scottjustin5000/pg-gcd/blob/20286b76a2e6d0304d19fcc6083aaea3fc2683df/README.md)
- [0.0.1 版安装 SQL](https://github.com/scottjustin5000/pg-gcd/blob/20286b76a2e6d0304d19fcc6083aaea3fc2683df/gcd--0.0.1.sql)
- [固定提交的 C 实现](https://github.com/scottjustin5000/pg-gcd/blob/20286b76a2e6d0304d19fcc6083aaea3fc2683df/gcd.c)

gcd 0.0.1 版安装一个 C 函数，用来计算整数数组的最大公约数。SQL 声明接受任意数组，但实现只处理 smallint、integer 和 bigint 元素类型。

### 示例

```sql
CREATE EXTENSION gcd;

SELECT gcd(ARRAY[14, 21, 35, 42]::integer[]);
```

结果为 7。该函数被声明为 immutable 和 strict，因此 PostgreSQL 可以折叠常量数组调用，而数组参数为 null 时会直接返回 null。

### 输入与兼容性限制

只能使用非空且不含 null 元素的数组。C 实现从第一个元素开始初始化计算，不能安全处理空数组；它也忽略了数组的 null bitmap。

虽然 bigint 数组可被接受，但计算时每个值都会收窄为 C int，因此超出有符号 32 位范围的值可能得到错误结果。实现返回的内部 datum 还是 64 位，而 SQL 声明为 integer。应优先使用处于有符号 32 位范围内的 smallint 或 integer 值，并验证边界情况。

项目没有持续维护的 PostgreSQL 兼容矩阵或近期发行历史；源码最后更新于 2019 年，仓库并未归档。应将其视为小型遗留辅助扩展，并针对部署所用的确切 PostgreSQL 主版本和编译器进行测试。
