## 用法

来源：

- [官方 README](https://github.com/snoutdata/snout-utils/blob/331d73c5a37060280d7c29e58b6672f7a493e485/README.md)
- [软件包元数据](https://github.com/snoutdata/snout-utils/blob/331d73c5a37060280d7c29e58b6672f7a493e485/Cargo.toml)
- [分发构建流程](https://github.com/snoutdata/snout-utils/blob/331d73c5a37060280d7c29e58b6672f7a493e485/scripts/build-dist.sh)
- [配置实现](https://github.com/snoutdata/snout-utils/blob/331d73c5a37060280d7c29e58b6672f7a493e485/src/settings.rs)

`snout_utils` 0.1.0 是 PostgreSQL 预加载动态库，将部分超级用户操作委托给配置指定的数据库所有者角色。它不创建 SQL 对象，也没有逐库安装命令。上游说明当前已完成构建，但尚未部署到其生产集群；分发构建流程支持 PostgreSQL 17 和 18。

### 启用权限委托

管理员安装匹配的动态库，并将其放在预加载列表末尾，使其钩子先于其他钩子执行：

```ini
shared_preload_libraries = 'snout_utils'
snout_utils.superuser = 'platform_admin'
snout_utils.privileged_role = 'app_owner'
snout_utils.privileged_extensions = 'postgis,vector,pg_cron,hstore'
snout_utils.extension_custom_scripts_path = '/usr/local/share/extension-custom-scripts'
snout_utils.privileged_role_allowed_configs = 'session_replication_role,pgaudit.*'
```

指定角色及允许使用的扩展动态库必须已经存在。重启 PostgreSQL 以加载动态库。五个配置项均使用 `sighup` 上下文，可通过重新加载配置修改，会话不能直接修改。未配置委托角色时，不执行权限委托。

### 委托语句

配置指定的角色及获得其权限的角色，可以管理明确列入允许列表的扩展，还可以创建或修改发布、创建符合条件的外部数据包装器与事件触发器，以及修改配置允许的参数。权限委托仅作用于当前语句，执行结束后恢复原角色，错误路径也会恢复。

对于外部数据包装器，处理函数和验证函数必须属于同一个已安装扩展。事件触发器委托不允许使用超级用户拥有的函数，额外的钩子检查也会防止超级用户执行 DDL 时，以提升后的权限调用被委托角色创建的事件触发器函数。角色修改不能授予超级用户权限，也不能修改已经拥有被委托角色权限的角色。普通角色仍受 PostgreSQL 原有权限检查约束。

### 自定义脚本与安全

`snout_utils.extension_custom_scripts_path` 指定在创建扩展前后以配置的超级用户执行的脚本目录，其中可以包含全局创建前脚本及按扩展划分的创建前后脚本，缺失的脚本会被跳过。任何角色创建扩展时都适用，包括原本不涉及权限委托的语句；应将该目录及内容视为高权限服务器配置进行保护。

脚本替换项 `@extname@`、`@extschema@`、`@extversion@` 和 `@extcascade@` 使用经过校验的 SQL 字面量，不支持的引号或反斜线字符会被拒绝。列表采用 PostgreSQL 风格的标识符语法，允许配置名末尾的通配符表示前缀匹配。扩展和参数允许列表都应保持在必要范围内。

### 运维

分发产物只有共享动态库。仓库中的控制文件用于构建和测试工具，分发构建流程不会安装它。不要尝试以 SQL 扩展的方式创建该模块。

它没有单独的工作进程或健康检查进程。必需的预加载动态库无法加载时，PostgreSQL 会启动失败。升级需替换动态库并重启 PostgreSQL，无需执行数据库迁移。逻辑备份不包含其配置，恢复与大版本升级流程必须单独部署动态库、角色和服务器设置。
