## 用法

来源：

- [官方 README](https://github.com/MasahikoSawada/pg_cgroup/blob/09214a99f5acfced8697ff3c72b0b05dc6bc7bb2/README.md)
- [扩展控制文件](https://github.com/MasahikoSawada/pg_cgroup/blob/09214a99f5acfced8697ff3c72b0b05dc6bc7bb2/pg_cgroup.control)
- [扩展 SQL](https://github.com/MasahikoSawada/pg_cgroup/blob/09214a99f5acfced8697ff3c72b0b05dc6bc7bb2/pg_cgroup--1.0.sql)
- [libcgroup 实现](https://github.com/MasahikoSawada/pg_cgroup/blob/09214a99f5acfced8697ff3c72b0b05dc6bc7bb2/pg_cgroup.c)

`pg_cgroup` 把 Linux libcgroup 操作暴露为 SQL 函数，使 PostgreSQL 后端能够创建子资源组、设置控制器值并把后端进程加入资源组。它面向 cgroup v1 时代的 Linux，文档测试环境是 CentOS 7.4、PostgreSQL 10 与 libcgroup 0.41。

### 核心流程

操作系统管理员必须先创建由 PostgreSQL 系统账号拥有的父 cgroup，然后预加载模块、设置父组、重启并安装 SQL 对象。

```conf
shared_preload_libraries = 'pg_cgroup'
pg_cgroup.cgroup_name = 'pgsql'
```

```sql
CREATE EXTENSION pg_cgroup;

REVOKE ALL ON FUNCTION create_resource_group(text, text) FROM PUBLIC;
REVOKE ALL ON FUNCTION delete_resource_group(text, text) FROM PUBLIC;
REVOKE ALL ON FUNCTION set_resource_value(text, text, text, bigint) FROM PUBLIC;
REVOKE ALL ON FUNCTION attach_resource_group(text, text) FROM PUBLIC;
REVOKE ALL ON FUNCTION attach_resource_group(text, text, integer) FROM PUBLIC;

SELECT create_resource_group('memory', 'reporting');
SELECT set_resource_value('memory', 'reporting', 'memory.limit_in_bytes', 1073741824);
SELECT attach_resource_group('memory', 'reporting');
```

`create_resource_group` 与 `delete_resource_group` 管理已配置父组下面的子组。`set_resource_value` 写入指定控制器参数。`attach_resource_group` 可附加当前后端，三个参数的重载则附加给定操作系统 PID。

### 权限与平台注意事项

这些函数使用 PostgreSQL 服务器账号的 libcgroup 权限运行，且没有 SQL 超级用户检查。默认公众执行权会允许数据库用户修改资源控制或尝试移动其他可访问 PID，因此必须从 `PUBLIC` 撤销所有重载，并只授予严格受控的管理角色。应在带白名单的包装函数中校验控制器和参数名；C 代码会把负数 `bigint` 输入转换成无符号值。

若干删除与附加调用忽略返回码，因此返回 `true` 不能证明内核操作成功。每次变更后都要核对实际 cgroup 文件系统。固定源码早于统一 cgroup v2 与现代 PostgreSQL API，不能假设它兼容 cgroup v2、容器、systemd 委派或当前服务器大版本。只应在匹配的隔离主机上测试，生产资源约束应优先使用操作系统资源管理器。
