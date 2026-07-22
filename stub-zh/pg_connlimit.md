## 用法

来源：

- [官方 C 实现与文档](https://github.com/fdr/pg_connlimit/blob/5ee6221506df3cb572b5b915fe834abbad18b60a/pg_connlimit.c)
- [官方控制文件](https://github.com/fdr/pg_connlimit/blob/5ee6221506df3cb572b5b915fe834abbad18b60a/pg_connlimit.control)

`pg_connlimit` 1.0 是一个无 SQL 对象的预加载模块，使用 PostgreSQL 目录之外的小文件限制每个角色的并发连接数。其设计目的是让主库和热备即使复制数据库目录，也能执行不同的配额。

### 配置

项目没有扩展 SQL 脚本，也不存在 `CREATE EXTENSION` 流程。安装共享库，创建受保护且每个角色一个文件的目录，然后预加载模块并设置目录路径。

```conf
shared_preload_libraries = 'pg_connlimit'
connlimit.directory = '/etc/postgresql/connlimit-db'
```

```text
/etc/postgresql/connlimit-db/
├── app_reader    # file content: 30
└── report_user   # file content: 5
```

修改 `shared_preload_libraries` 后需要重启。认证 hook 会在每次成功登录时读取匹配的角色文件；当该角色现有后端数已达到整数上限时，它会拒绝新连接。

### 精确执行规则

- 限额按整个集群的 PostgreSQL 角色统计，而不是按数据库统计。
- 只有完全由小写字母、数字和下划线组成的角色名符合条件；其他名称会被跳过，以防路径穿越。
- 文件缺失/不可读、整数无效、角色未知或目录未设置时都会 fail open，不执行限额。
- 后续连接尝试会读取文件变化；修改 `connlimit.directory` 属于 SIGHUP 级配置变更。
- 当前连接尚未计入 `CountUserBackends`，因此现有数量大于或等于配置值时 hook 会拒绝连接。

### 安全与运维边界

必须防止数据库用户和无关操作系统账户修改目录与文件。能删除、替换或破坏文件的用户可以关闭限制；非正整数则会阻止该角色的所有新连接。应使用原子文件替换，并在发布前验证内容。

模块有意采用 fail open，不能单独作为安全边界。仍应按需要保留 `max_connections`、管理员保留连接、连接池限制和目录级 `CONNECTION LIMIT`。源码停留在 2013 年，并使用需要在现代 PostgreSQL 上移植及压测的认证/procarray hook。
