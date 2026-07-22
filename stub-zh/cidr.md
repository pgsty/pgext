## 用法

来源：

- [官方 README](https://github.com/ntkme/postgresql-cidr/blob/0a79fac02992d1863b6f98a533e51abc25c3d69c/README.md)
- [1.0 版本扩展 SQL](https://github.com/ntkme/postgresql-cidr/blob/0a79fac02992d1863b6f98a533e51abc25c3d69c/cidr--1.0.sql)
- [上游回归示例](https://github.com/ntkme/postgresql-cidr/blob/0a79fac02992d1863b6f98a533e51abc25c3d69c/sql/cidr.sql)

`cidr` 增加了双参数函数 `pg_catalog.cidr(inet, inet)`，返回精确覆盖一个闭合 IPv4 或 IPv6 地址范围的最小 CIDR 网络集合。与内置 `inet_merge` 不同，它不会把结果扩展成一个可能包含请求范围之外地址的网络。

### 核心流程

```sql
CREATE EXTENSION cidr;

SELECT network
FROM cidr('192.0.2.5'::inet, '192.0.2.20'::inet) AS r(network);
```

可以用它为基于范围的访问规则或报表物化精确网络覆盖：

```sql
INSERT INTO allowed_networks(network)
SELECT network
FROM cidr('2001:db8::10'::inet, '2001:db8::ff'::inet) AS r(network);
```

### 函数契约

- 两个参数必须属于同一地址族。
- 两个参数必须是完整位宽掩码的主机地址，而不是输入网络。
- 第一个地址必须小于或等于第二个地址。
- 集合返回结果由互不重叠的 `cidr` 值组成，并精确覆盖闭合范围。

### 运维说明

版本 `1.0` 使用纯 PL/pgSQL，扩展层面可重定位，且未声明预加载依赖。函数本身被有意创建在 `pg_catalog` 中，因此无需修改 `search_path` 即可访问；审计或移除扩展时应考虑这个全局名称。

对齐条件允许时，大范围会被紧凑表示；边界高度不对齐时仍可能返回很多行。处理用户提供的范围时，应限制结果规模或将其物化。混用 IPv4/IPv6、输入带网络掩码的地址或使用降序范围都会报错，而不会被自动规范化。
