## 用法

来源：

- [固定提交的上游 README](https://github.com/liang7878/login_refuse/blob/f2af59137c10e786b3abf9646fa11fcd8bf3f8d8/README.md)
- [0.1 版安装 SQL](https://github.com/liang7878/login_refuse/blob/f2af59137c10e786b3abf9646fa11fcd8bf3f8d8/login_refuse--0.1.sql)
- [固定提交的认证 hook 实现](https://github.com/liang7878/login_refuse/blob/f2af59137c10e786b3abf9646fa11fcd8bf3f8d8/login_refuse.c)
- [固定提交的扩展 control 文件](https://github.com/liang7878/login_refuse/blob/f2af59137c10e786b3abf9646fa11fcd8bf3f8d8/login_refuse.control)

login_refuse 0.1 版安装 client authentication hook。对于 password、MD5 和 SCRAM 认证，它按用户名记录失败尝试，在达到阈值后的一段配置时间内拒绝后续连接。它还可以让一个角色在某个 Unix timestamp 之后进入 expired 状态。

### 配置

该 hook 必须预加载。两个数值设置都默认为零，因此应在重启 PostgreSQL 前显式设置正值：

```conf
shared_preload_libraries = 'login_refuse'
login_refuse.minutes = 10
login_refuse.threshold = 5
```

```sql
CREATE EXTENSION login_refuse;

SELECT login_refuse_set_expire_time(
    'temporary_user',
    extract(epoch FROM timestamptz '2030-01-01 00:00:00+00')::bigint
);

SELECT login_refuse_reset_expire_time('temporary_user');
```

两个 SQL 函数会在运行时检查超级用户权限。timestamp 函数会让账户从指定时刻之后持续被该 hook 拒绝，直到 reset；它并不是“锁定到此时刻”的临时解锁时间。

### 认证安全

失败和过期状态存放在 PGDATA 目录下的两个普通文件中。每个执行认证的 backend 都会在没有进程间锁和持久事务语义的情况下读取、重写和追加这些文件。并发尝试可能丢失或破坏状态。用户名和文本行通过很小的固定 C buffer 复制，多个文件和内存错误路径也不安全。

项目只有一行 README，没有测试矩阵，自 2018 年后没有更新。authentication hook 和内部 enum 对版本敏感，缺陷可能锁死管理员或让连接进程崩溃。不要把 login_refuse 作为主要暴力破解控制；应优先使用受维护的 proxy、identity provider、firewall 或 PostgreSQL 支持的认证策略。若要评估，必须保留独立本地恢复通道、限制文件权限、测试并发失败和重启行为，并监控状态文件。
