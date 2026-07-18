## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/ketteq-neon/kq_cx/blob/585f574defae0a5aa15fa746814e9adc3719fa5c/README.md)
- [扩展 control 文件](https://github.com/ketteq-neon/kq_cx/blob/585f574defae0a5aa15fa746814e9adc3719fa5c/kq_cx.control)
- [Cargo 特性矩阵](https://github.com/ketteq-neon/kq_cx/blob/585f574defae0a5aa15fa746814e9adc3719fa5c/Cargo.toml)

`kq_cx` 将应用专用的日历切片保存在 PostgreSQL 共享内存中，并基于该缓存执行日期运算。先将 `kq_cx` 加入 `shared_preload_libraries` 并重启 PostgreSQL，然后在包含所需 ketteQ 日历表的 schema 中安装扩展。

```sql
CREATE EXTENSION kq_cx;
SELECT kq_invalidate_calendar_cache();
SELECT kq_add_days('2008-01-15', 1, 'quarter');
```

`kq_add_days_by_id` 按数字切片类型 ID 选择日历，`kq_add_days` 则按名称选择。底层表发生变化后，`kq_invalidate_calendar_cache` 会清空并重新加载共享缓存。

README 记录支持 PostgreSQL 15–17；当前 Cargo 特性矩阵还定义了 PostgreSQL 18。该扩展不能独立使用：源码要求特定应用 schema 和已填充的日历数据，部署前必须验证这些表契约。缓存失效会影响共享状态，应与应用写入协调。
