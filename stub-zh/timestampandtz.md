## 用法

来源：

- [上游 README](https://github.com/mweber26/timestampandtz/blob/fc3597ec859201d8f1e122c56bbf355d4c79b04f/README.md)
- [安装 SQL](https://github.com/mweber26/timestampandtz/blob/fc3597ec859201d8f1e122c56bbf355d4c79b04f/timestampandtz--1.0.0.sql)
- [C 类型实现](https://github.com/mweber26/timestampandtz/blob/fc3597ec859201d8f1e122c56bbf355d4c79b04f/timestampandtz.c)
- [固定时区标识符表](https://github.com/mweber26/timestampandtz/blob/fc3597ec859201d8f1e122c56bbf355d4c79b04f/zones.c)
- [PGXN 发布元数据](https://pgxn.org/dist/timestampandtz/)

timestampandtz 同时存储时间点与具名来源时区，在保留本地时间显示的同时，以 UTC 时间点进行比较和 B-tree 排序。扩展的 SQL 默认版本是 1.0.0；PGXN 将同一安装脚本标记为 2018 年发布的发行版 1.0.2。

### 保留来源时区

```sql
CREATE EXTENSION timestampandtz;
SELECT '2014-09-01 22:15:00 @ US/Eastern'::timestampandtz;
SELECT tzmove(
  '2014-09-01 22:15:00 @ US/Eastern'::timestampandtz,
  'US/Pacific'
);
SELECT '2014-09-01 22:15:00 @ US/Eastern'::timestampandtz
     = '2014-09-01 19:15:00 @ US/Pacific'::timestampandtz;
```

最后一次比较为真，因为两个值表示同一时间点。间隔运算会保留已存时区，并应用基于墙上时钟的夏令时规则。

### 注意事项

- 两字节时区标识符来自扩展自有固定表，而不是 PostgreSQL 实时时区目录。新增或改名的 IANA 时区可能不可用，不过偏移规则仍来自服务器时区数据。
- 在不同构建间移动二进制数据时必须保留完全相同的标识符表。二进制接收函数不会在后续查表前验证标识符边界。
- 文本值中 @ 分隔符任一侧为空时，可能进入已发布 C 解析器中的不安全缓冲区边界逻辑。强制转换前应拒绝不可信输入，并针对暴露的工作负载测试修复后代码。
- 安装会在 pg_catalog 内添加 timezone 重载，因此需要较高信任级别。安装与删除期间应审查对象所有权和冲突。
- 上游最后一次发布于 2018 年，面向 PostgreSQL 9.2 时代的内部接口。没有当前主版本兼容声明；采用前必须重新构建，并执行回归、夏令时、二进制复制、转储、恢复与索引测试。
