## 用法

来源：

- [官方上游文档](https://github.com/cybertec-postgresql/merge_ips/blob/3cf654cef7c7ce7fafebc9b0799f15da8725da52/README)
- [官方扩展 SQL](https://github.com/cybertec-postgresql/merge_ips/blob/3cf654cef7c7ce7fafebc9b0799f15da8725da52/merge_ips--1.0.sql)
- [官方扩展控制文件](https://github.com/cybertec-postgresql/merge_ips/blob/3cf654cef7c7ce7fafebc9b0799f15da8725da52/merge_ips.control)

`merge_ips` 1.0 将相邻的 IPv4 或 IPv6 地址与网络合并成不填补空隙的最小精确覆盖网络集合。它是安装在固定 `merge_ips` 模式中的纯 PL/pgSQL 扩展。上游仓库已经归档，因此应固定并测试实现，再用于持久工作流。

### 合并网络

将 `inet[]` 数组传给集合返回函数或其数组包装器：

```sql
CREATE EXTENSION merge_ips;

SELECT *
FROM merge_ips.merge_ips(ARRAY[
  '10.0.52.0/30',
  '10.0.52.4/30',
  '10.0.52.16/30',
  '10.0.52.20/30'
]::inet[]);

SELECT merge_ips.merge_ips_array(ARRAY[
  '192.0.2.0/25',
  '192.0.2.128/25'
]::inet[]);
```

第一个查询返回 `10.0.52.0/29` 与 `10.0.52.16/29`；它不会虚构覆盖两者空隙的网络。第二个查询返回包含 `192.0.2.0/24` 的数组。

### 函数行为

- `merge_ips.merge_ips(inet[])` 返回 `inet` 集合，去除重复项与被其他结果包含的网络，并反复合并准确的兄弟网络。
- `merge_ips.merge_ips_array(inet[])` 将相同结果收集为 `inet[]`。

输出顺序不属于 SQL 契约；需要确定性显示时应添加 `ORDER BY`。

### 运维说明

实现会在当前会话中创建、清空、重命名和删除使用固定名称的临时工作表，因此它不是没有副作用的不可变计算，并可能与同一会话中的嵌套或重入调用冲突。不要在不可变表达式、索引或共享同一会话的并行应用代码中调用它。IPv4 与 IPv6 输入应分开，测试主机位和包含场景，并显式排序结果。对于大数组，应测试临时表 I/O 和循环成本；如果这个归档实现过慢或运维侵入性太强，受维护的应用程序库可能更合适。
