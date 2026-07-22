## 用法

来源：

- [官方 README](https://github.com/higuoxing/pg_badapple/blob/92ddbeea6fee794bfaad70f1b1d9c4bfab068c91/README.md)
- [扩展 SQL](https://github.com/higuoxing/pg_badapple/blob/92ddbeea6fee794bfaad70f1b1d9c4bfab068c91/badapple--1.0.0.sql)
- [C 实现](https://github.com/higuoxing/pg_badapple/blob/92ddbeea6fee794bfaad70f1b1d9c4bfab068c91/badapple.c)
- [扩展控制文件](https://github.com/higuoxing/pg_badapple/blob/92ddbeea6fee794bfaad70f1b1d9c4bfab068c91/badapple.control)

`badapple` 是一个趣味扩展，通过 PostgreSQL `NOTICE` 消息播放 Bad Apple 或篮球字符动画。它只适合交互式终端演示，不适合应用或生产负载：读取并播放服务端动画文件期间，调用它的后端会一直被占用。

### 核心流程

在一次性数据库中安装扩展，启用通知输出，然后从能够解释 ANSI/VT100 控制序列的终端调用两个入口之一：

```sql
CREATE EXTENSION badapple;
SET client_min_messages = notice;

SELECT play_badapple();
-- Or:
SELECT play_basketball();
```

C 代码从数据库服务器上的 PostgreSQL 共享扩展目录打开 `badapple.txt` 或 `basketball.txt`。它每帧读取 30 行文本，将整帧发成一条 `NOTICE`，用 ANSI 转义序列清屏，并在帧间休眠。取消语句即可停止播放。

### 用户接口

- `play_badapple()`：播放随扩展安装的 Bad Apple 文本帧。
- `play_basketball()`：播放随扩展安装的篮球文本帧。

两个函数都返回 `void`；可见结果是连续的 `NOTICE` 消息。

### 运维边界

媒体文件必须安装到服务器的扩展共享数据目录。图形化 SQL 客户端、被抑制的通知，或不支持 ANSI 光标控制的终端都无法正确显示动画。

播放会产生大量消息，并有意长期占用一个 PostgreSQL 后端。不要通过连接池、API 或生产监控路径调用它。实现会把源文件行复制到固定长度的帧缓冲区，因此只能使用上游随附且可信的媒体文件；替换为超长或不可信的行可能破坏后端内存。请在隔离数据库中测试取消与日志行为。
