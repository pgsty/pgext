

## 用法

> [lolor: 逻辑复制友好的 PostgreSQL 大对象替代](https://github.com/pgEdge/lolor)

使 PostgreSQL 大对象与逻辑复制兼容，将其存储在非目录表中。

### 启用

```sql
CREATE EXTENSION lolor;
```

在 `postgresql.conf` 中配置节点标识符：

```ini
lolor.node = 1  -- 唯一节点 ID（1 到 2^28）
```

可选调整搜索路径：

```sql
SET search_path = lolor, "$user", public, pg_catalog;
```

### 大对象操作

安装后，标准 `lo_*` 函数会重定向到使用 lolor 的表：

```sql
-- 创建大对象
SELECT lo_create(0);

-- 将文件导入大对象
SELECT lo_import('/path/to/file.bin');

-- 将大对象导出到文件
SELECT lo_export(oid, '/path/to/output.bin');

-- 打开、读取、写入、定位、关闭
SELECT lo_open(oid, x'40000'::int);  -- INV_WRITE
SELECT lowrite(fd, 'data'::bytea);
SELECT loread(fd, 1024);
SELECT lo_close(fd);

-- 删除大对象
SELECT lo_unlink(oid);
```

### 复制设置

将 lolor 表添加到复制集：

```sql
-- 用于 spock/pgedge 复制
SELECT spock.repset_add_table('default', 'lolor.pg_largeobject');
SELECT spock.repset_add_table('default', 'lolor.pg_largeobject_metadata');
```

### 内部表

该扩展在以下表中管理大对象：

- `lolor.pg_largeobject` - 存储对象数据块
- `lolor.pg_largeobject_metadata` - 存储对象元数据

### 限制

- lolor 激活时不能使用原生 PostgreSQL 大对象功能
- 不支持将现有原生大对象迁移到 lolor
- 不支持 `ALTER LARGE OBJECT`、`GRANT ON LARGE OBJECT`、`COMMENT ON LARGE OBJECT` 和 `REVOKE ON LARGE OBJECT`
- 需要 PostgreSQL 16 或更新版本
