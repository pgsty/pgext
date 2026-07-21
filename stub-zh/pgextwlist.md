## 用法

来源：

- [pgextwlist v1.20 README](https://github.com/dimitri/pgextwlist/blob/v1.20/README.md)
- [Changes from v1.19 to v1.20](https://github.com/dimitri/pgextwlist/compare/v1.19...v1.20)

pgextwlist 允许选定的非超级用户在明确的允许列表范围内运行扩展生命周期命令。这些命令会临时以启动超级用户的权限执行，因此允许列表和任何自定义脚本是数据库安全边界的组成部分。

该模块预加载且不需要使用 CREATE EXTENSION 命令。

### 配置允许列表

为每个后端加载模块：

    local_preload_libraries = 'pgextwlist'
    extwlist.extensions = 'hstore,cube,pg_stat_statements'

也可以按角色分配此列表：

    ALTER ROLE extension_admin
      SET extwlist.extensions = 'hstore,pg_stat_statements';

更改 local_preload_libraries 后需要重新连接。允许的用户可以运行以下命令：

    CREATE EXTENSION hstore;
    ALTER EXTENSION hstore UPDATE;
    COMMENT ON EXTENSION hstore IS 'approved utility';
    DROP EXTENSION hstore;

未在 extwlist.extensions 中命名的扩展将被拒绝。

### 限制数据库所有权

版本 1.20 添加了：

    extwlist.restrict_to_database_owner = on

启用后，调用者还必须拥有当前数据库。默认情况下此功能关闭以保持兼容性。当需要防止扩展管理跨越数据库所有权边界时，请启用它。

### 自定义生命周期脚本

设置 extwlist.custom_path 为一个可读目录。版本 1.20 如果路径不存在或不可读则会引发错误，而不是静默跳过。

对于扩展 extname，可以在 extname/ 目录下包含以下脚本：

- before--1.0.sql 和 after--1.0.sql 在特定版本创建时使用。
- before-create.sql 和 after-create.sql 作为创建的备用方案。
- before-update.sql 和 after-update.sql 在更新时使用。
- before-drop.sql 和 after-drop.sql 在删除时使用。

模板可以使用 @extschema@, @current_user@ 和 @database_owner@。只有受信任的管理员才能编写此目录，因为脚本以提升权限执行。

### 配置索引

- local_preload_libraries: 将 pgextwlist 加载到新会话中。
- extwlist.extensions: 逗号分隔的允许列表。
- extwlist.custom_path: 生命周期脚本的基础目录。
- extwlist.restrict_to_database_owner: 还需要求数据库所有权。

### 安全性和兼容性注意事项

- 版本 1.20 拒绝包含引号、美元符号、撇号或反斜杠字符的替换名称，解决了 CVE-2023-39417 跟踪的命令注入类问题。
- 支持 CREATE EXTENSION, DROP EXTENSION, ALTER EXTENSION UPDATE 和 COMMENT ON EXTENSION。不支持 ALTER EXTENSION ADD 和 DROP。
- 通过提升路径创建的对象的所有权根据扩展/启动超级用户的行为确定，不一定由请求的角色拥有。
- 在添加名称之前，请审查扩展 SQL 和自定义脚本。将扩展列入白名单会赋予调用者其安装和更新脚本所体现的权力。

文档结束
