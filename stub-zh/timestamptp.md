## 用法

来源：

- [README.md](https://github.com/brownan/timestamptp/blob/6663d316cef644075624be31602add416f83060c/README.md)
- [timestamptp--1.0.sql](https://github.com/brownan/timestamptp/blob/6663d316cef644075624be31602add416f83060c/timestamptp--1.0.sql)
- [timestamptp.c](https://github.com/brownan/timestamptp/blob/6663d316cef644075624be31602add416f83060c/timestamptp.c)
- [timestamptp.control](https://github.com/brownan/timestamptp/blob/6663d316cef644075624be31602add416f83060c/timestamptp.control)

`timestamptp` 同时保存绝对时间戳与原始 UTC 偏移，并按该偏移输出。它保留的是偏移值，不是包含未来夏令时规则的具名时区。

### 核心用法

```sql
CREATE EXTENSION timestamptp;
SELECT '2026-10-10 09:00:00+08'::timestamptp;
SELECT '2026-10-10 09:00:00+08'::timestamptp + interval '1 hour';
```

### 运行边界

`make_timestamptp(timestamptz, integer)` 与 `make_timestamptp(timestamptz, text)` 是构造函数。转换为 `timestamptz` 会丢弃保存的偏移，反向转换采用会话时区。相等比较忽略偏移、比较绝对时刻，加上时间间隔则保留偏移。此原生类型可迁移模式。上游未公布 PostgreSQL 主版本支持矩阵或明确的许可证声明。安装采用控制文件的默认超级用户要求。
