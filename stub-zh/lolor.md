## 用法

来源：

- [v1.2.2 README](https://github.com/pgEdge/lolor/blob/v1.2.2/README.md)
- [Control file](https://github.com/pgEdge/lolor/blob/v1.2.2/lolor.control)
- [Version 1.2.2 migration](https://github.com/pgEdge/lolor/blob/v1.2.2/lolor--1.2.1--1.2.2.sql)
- [Usage reference](https://github.com/pgEdge/lolor/blob/v1.2.2/docs/using_lolor.md)
- [Release notes](https://github.com/pgEdge/lolor/blob/v1.2.2/docs/lolor_release_notes.md)
`lolor` 1.2.2 将大对象的数据块和元数据放到普通表中，使逻辑复制能够包含它们。启用后会替换数据库的原生大对象函数；这是影响整个数据库的行为变化，而不是独立对象存储。

### 配置并启用

在创建大对象前，为每个写入节点分配不同的非零 `lolor.node`；上游文档规定范围为 1 到 2^28。配置服务参数，并在每个参与的数据库中创建扩展：

```conf
lolor.node = 1
```

```sql
CREATE EXTENSION lolor;
SET search_path = lolor, "$user", public, pg_catalog;
```

控制文件固定使用 `lolor` 模式，声明为可信安装，且不允许重定位。上游 README 要求 PostgreSQL 16 或更新版本；当前 Pigsty pgEdge 组合包另外通过了 15–18 的构建测试，但这一打包结果不能证明任意原版 PostgreSQL 15 安装都受支持。

### 大对象工作流

```sql
WITH created AS (
  SELECT lo_from_bytea(0, convert_to('example data', 'UTF8')) AS oid
)
SELECT oid, convert_from(lo_get(oid), 'UTF8') AS contents FROM created;
```

`lo_create()`、`lo_get()`、`lo_put()` 和 `lo_unlink()` 等标准调用使用替换后的函数。通过 `lo_open()`、`loread()`、`lowrite()` 和 `lo_close()` 进行描述符访问时，必须保持在同一个事务内。文件导入和导出操作的是服务端路径，仍受相应函数与文件系统权限限制。

扩展使用 `lolor.pg_largeobject` 和 `lolor.pg_largeobject_metadata` 两张表；原来的系统目录函数会改名保留，使禁用或移除扩展时能够恢复原生函数名称。

### 复制

已有 Spock 复制集时，将两张表都纳入：

```sql
SELECT spock.repset_add_table('default', 'lolor.pg_largeobject');
SELECT spock.repset_add_table('default', 'lolor.pg_largeobject_metadata');
```

在所有节点安装兼容版本，并协调节点标识符与对象所有权。仅创建扩展不会建立逻辑复制拓扑。

### 升级与边界

1.2.2 修复扩展升级和数据库大版本升级行为，并新增 `lolor.disable()`、`lolor.enable()` 和 `lolor.is_enabled()`。应当**先更新 lolor，再运行 pg_upgrade**；该修复不会追溯修复已经用旧扩展文件尝试的大版本升级：

```sql
ALTER EXTENSION lolor UPDATE TO '1.2.2';
SELECT lolor.is_enabled();
```

禁用操作改变当前生效的原生函数名称，不会迁移已存储的大对象。上游不提供原生大对象迁移功能；启用 lolor 时，原生大对象功能和 lolor 存储不能混用。上游不支持 ALTER LARGE OBJECT、GRANT ON LARGE OBJECT、COMMENT ON LARGE OBJECT 和 REVOKE ON LARGE OBJECT。应为两张普通表安排备份、恢复与复制。
