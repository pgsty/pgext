## 用法

来源：

- [官方文档](https://github.com/mrayva/pg_zerialize/blob/2aa98a11c2c51ee468336b702eebb571428f84d2/README.md)
- [扩展控制文件](https://github.com/mrayva/pg_zerialize/blob/2aa98a11c2c51ee468336b702eebb571428f84d2/pg_zerialize.control)
- [官方仓库](https://github.com/mrayva/pg_zerialize)

`pg_zerialize` 提供 MessagePack、CBOR、ZERA、FlexBuffers、Ion、BSON 与 BEVE 二进制序列化。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `pg_zerialize`：

```sql
CREATE EXTENSION pg_zerialize;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
CREATE TYPE employee AS (id int, name text, active boolean);

SELECT msgpack_populate_record(NULL::employee, row_to_msgpack(ROW(1, 'Ada', true)::employee));
SELECT cbor_populate_record(ROW(1, 'Ada', true)::employee, cbor_build_object('active', false));
SELECT zera_populate_record(NULL::employee, zera_from_jsonb('{"id":2,"name":"Grace"}'::jsonb));
```

### 主要对象

官方来源通过动态方式或供应商工具定义扩展接口；授权前应检查实际安装版本。

### 运维与边界

- 已验证的 PostgreSQL 主版本证据覆盖 16, 17, 18；不要推断未列出的主版本。
