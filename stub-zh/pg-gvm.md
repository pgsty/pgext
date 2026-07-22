## 用法

来源：

- [官方 22.6 系列 README](https://github.com/greenbone/pg-gvm/blob/v22.6.18/README.md)
- [官方主机工具 SQL](https://github.com/greenbone/pg-gvm/blob/v22.6.18/sql/hosts.in.sql)
- [官方 iCalendar 工具 SQL](https://github.com/greenbone/pg-gvm/blob/v22.6.18/sql/ical.in.sql)
- [官方正则表达式工具 SQL](https://github.com/greenbone/pg-gvm/blob/v22.6.18/sql/regexp.in.sql)

`pg-gvm` 是 Greenbone 为 GVMd 提供的 PostgreSQL 辅助库。版本/API 系列 `22.6` 提供 C 函数，用于处理 Greenbone 主机范围表达式、iCalendar 重复时间计算与正则匹配。它是应用支撑组件，而不是通用 SQL 工具集。

### 核心流程

创建扩展，然后用 Greenbone 组件生成的相同文本格式调用辅助函数：

```sql
CREATE EXTENSION "pg-gvm";

SELECT hosts_contains(
  '192.168.123.1-192.168.123.20, 192.168.123.30',
  '192.168.123.10'
);

SELECT max_hosts(
  '192.168.123.1-192.168.123.20, 192.168.123.30-192.168.123.34',
  '192.168.123.10'
);

SELECT regexp('abc', '^[a-z]+$');
```

第一个调用测试 Greenbone 主机规格是否包含请求的主机；第二个统计第一个范围表达式中的地址数，并排除第二个参数中的地址。

### API

- `hosts_contains(text, text) RETURNS boolean`：主机表达式是否包含指定主机。
- `max_hosts(text, text) RETURNS integer`：第一个表达式所代表的主机数，扣除第二个表达式中的排除项。
- `regexp(text, text) RETURNS boolean`：执行正则匹配；上游测试中无效模式返回 false。
- `next_time_ical(text, bigint, text) RETURNS integer`：根据 iCalendar 文本、Unix 时间戳和时区计算下一次重复时间。
- `next_time_ical(text, bigint, text, integer) RETURNS integer`：带发生次数偏移量的相同计算。

### 注意事项

扩展链接 `libical`、GLib 与 `libgvm-base`；软件包和库版本必须与 Greenbone 栈匹配。生成的 control 文件使用项目 API 版本和共享库名称，因此应安装同一发布的完整产物，不能混用不同标签的文件。

主机语法和 iCalendar 行为遵循 Greenbone 库，而非 PostgreSQL 原生网络/范围类型。把这些函数用于策略判断前，应测试边界地址、排除项、时区、夏令时转换、畸形日历输入与大型范围表达式。
