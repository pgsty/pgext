## 用法

来源：

- [README.md](https://github.com/FranckPachot/pg-mm-dd-yyyy/blob/6d68f142d9777c2d6ee851d181746e77c6a71d6f/README.md)
- [mmddyyyy.control](https://github.com/FranckPachot/pg-mm-dd-yyyy/blob/6d68f142d9777c2d6ee851d181746e77c6a71d6f/mmddyyyy.control)
- [mmddyyyy--0.1.0.sql](https://github.com/FranckPachot/pg-mm-dd-yyyy/blob/6d68f142d9777c2d6ee851d181746e77c6a71d6f/sql/mmddyyyy--0.1.0.sql)
- [mmddyyyy.c](https://github.com/FranckPachot/pg-mm-dd-yyyy/blob/6d68f142d9777c2d6ee851d181746e77c6a71d6f/src/mmddyyyy.c)

`mmddyyyy` 是教学扩展，将日期保存为十个 MM/DD/YYYY 字面字节，用来演示 C 基础类型、运算符、B-tree 排序与 GiST 运算符类。上游明确不建议用于生产；普通应用日期应使用 PostgreSQL 内置日期类型。

### 核心工作流

```sql
CREATE EXTENSION mmddyyyy;
CREATE TABLE events (id integer PRIMARY KEY, happened_on mmddyyyy NOT NULL);
INSERT INTO events VALUES (1, '09/15/2026'), (2, '12/31/2025');
CREATE INDEX events_date_btree ON events (happened_on);
CREATE INDEX events_date_gist ON events USING gist (happened_on);
SELECT * FROM events WHERE happened_on <@ '09/*/*'::mmddyyyy_pattern;
SELECT * FROM events
ORDER BY happened_on <-> '09/15/2026'::mmddyyyy_pattern LIMIT 5;
```

### 类型与索引语义

`mmddyyyy` 验证真实公历日期及规范的补零文本，可与 `date` 相互转换。`mmddyyyy_month`、`mmddyyyy_day` 和 `mmddyyyy_year` 返回各日期分量。虽然存储文本以月份开头，普通比较与默认 B-tree 仍按时间先后排序。

`mmddyyyy_pattern` 允许任意日期分量使用通配符，`<@` 检查分量匹配。`mmddyyyy_gist_ops` 支持这些模式及 `<->` 最近邻搜索。距离优先比较循环月份差，再比较日和年，并非相隔的实际天数。索引效果取决于查询形态。

### 要求与限制

容器与正确性测试以 PostgreSQL 17 为目标。C 扩展需要超级用户安装，可重定位，无需预加载。年份范围为 0001–9999，不支持公元前、无限日期或时区。未提供二进制收发函数及升级脚本，审核源码中也没有许可证文件；仓库公开本身不构成再分发许可。
