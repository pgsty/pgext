

## 用法

> [ddlx: DDL 提取函数](https://github.com/lacanoid/pgddl)

ddlx 是一个纯 SQL 扩展，可从 PostgreSQL 系统目录生成 DDL 脚本。它提供三个主要函数，接受 OID 参数，适用于所有 `reg*` 对象标识符类型。

### 核心函数

```sql
-- 为对象生成 CREATE 语句
SELECT ddlx_create('my_table'::regclass);
SELECT ddlx_create('my_type'::regtype);
SELECT ddlx_create('my_function'::regproc);
SELECT ddlx_create(current_role::regrole);

-- 生成 DROP 语句
SELECT ddlx_drop('my_table'::regclass);

-- 生成包含依赖树的完整 DDL 脚本
SELECT ddlx_script('my_table'::regclass);
SELECT ddlx_script('my_enum');
SELECT ddlx_script(current_role::regrole);
```

### 选项

选项以文本数组形式传递（例如 `'{ine,nodcl}'`）：

```sql
SELECT ddlx_create('my_table'::regclass, '{ine}');        -- 添加 IF NOT EXISTS
SELECT ddlx_create('my_type'::regtype, '{noowner}');       -- 省略 ALTER SET OWNER
SELECT ddlx_script('my_table'::regclass, '{drop}');        -- 包含 DROP 语句
```

可用选项：`drop`、`nodrop`、`owner`、`noowner`、`nogrants`、`nodcl`、`noalter`、`ine`（IF NOT EXISTS）、`ie`（IF EXISTS）、`ext`、`lite`、`nowrap`、`nopartitions`、`comments`、`nocomments`、`nostorage`、`noconstraints`、`noindexes`、`nosettings`、`notriggers`、`grantor`、`data`。

### 对于没有 reg* 类型的对象

```sql
SELECT ddlx_create(oid) FROM pg_foreign_data_wrapper WHERE fdwname = 'postgres_fdw';
SELECT ddlx_create(oid) FROM pg_database WHERE datname = current_database();
```

### 其他函数

```sql
-- 通过 OID 识别任意对象
SELECT * FROM ddlx_identify(oid);

-- 描述类的列
SELECT * FROM ddlx_describe('my_table'::regclass);

-- 获取各个定义部分
SELECT * FROM ddlx_definitions(oid);

-- 仅生成数据前创建语句
SELECT ddlx_createonly('my_table'::regclass);

-- 生成数据后修改语句
SELECT ddlx_alter('my_table'::regclass);

-- 通过正则表达式搜索函数/视图内容
SELECT ddlx_create(objid) FROM ddlx_apropos('users');

-- 获取 GRANT 语句
SELECT ddlx_grants('my_table'::regclass);
```
