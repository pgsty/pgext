## 用法

来源：

- [Official extension control file](https://github.com/zilder/pg_lz4/blob/f3259686848a61c544dbe58df5627ed3d77a2144/pg_lz4.control)
- [Official upstream README](https://github.com/zilder/pg_lz4/blob/f3259686848a61c544dbe58df5627ed3d77a2144/README.md)
- [Required PostgreSQL patch discussion](https://www.postgresql.org/message-id/20190315125203.5da43edb%40ildus-work.localdomain)

`pg_lz4` 1.0 将 LZ4 注册为列压缩方法。这个历史实现要求使用带有链接中自定义压缩补丁的 PostgreSQL 开发源码树，并依赖系统 `liblz4`；它不适用于未经修改的受支持服务器。

### 核心流程

编译打过补丁的服务器并安装与其匹配的扩展二进制后，创建扩展并按列选择压缩方法。

```sql
CREATE EXTENSION pg_lz4;

CREATE TABLE messages (
  id bigint GENERATED ALWAYS AS IDENTITY,
  body text COMPRESSION pg_lz4
);

ALTER TABLE messages
  ALTER COLUMN body SET COMPRESSION pg_lz4 WITH (acceleration '8');
```

`acceleration` 选项调整 LZ4 的速度与压缩率权衡。修改列设置只影响新写入值，本身不会重写现有行。

### 兼容边界

上游 README 面向 2019 年的开发补丁，而不是后来进入内核的压缩方法 API。不要假设它兼容原版 PostgreSQL 或内置 LZ4 TOAST 支持。必须使用完全匹配的补丁服务器 ABI，测试备份还原可移植性，并在迁移到缺少该方法的服务器前安排数据重写。
