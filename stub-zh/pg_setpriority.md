## 用法

来源：

- [上游 README](https://github.com/sshutdownow/pg_setpriority/blob/d36af737087d590ff8d94cecbf5b24e02ffdb66a/README.md)
- [扩展 control 文件](https://github.com/sshutdownow/pg_setpriority/blob/d36af737087d590ff8d94cecbf5b24e02ffdb66a/pg_setpriority.control)
- [SQL 安装脚本](https://github.com/sshutdownow/pg_setpriority/blob/d36af737087d590ff8d94cecbf5b24e02ffdb66a/pg_setpriority--1.0.sql)
- [C 实现](https://github.com/sshutdownow/pg_setpriority/blob/d36af737087d590ff8d94cecbf5b24e02ffdb66a/pg_setpriority.c)

`pg_setpriority` `1.0` 版封装操作系统的进程优先级调用。`setpriority(integer)` 修改当前 PostgreSQL 后端的 nice 值，`getpriority(integer)` 读取进程优先级，默认读取当前后端。

### 示例

```sql
CREATE EXTENSION pg_setpriority;
SELECT setpriority(10);
SELECT getpriority();
```

修改会持续到后端进程结束，因此连接池复用该后端时可能影响后续会话。普通操作系统用户通常只能降低优先级，若无额外能力便无法恢复或设置负 nice 值。`setpriority(integer)` 返回原始成功/错误码而不抛出异常，所以必须显式检查。安装脚本默认开放执行权限；应从应用角色撤销 `EXECUTE`，只提供严格受控的管理封装。
