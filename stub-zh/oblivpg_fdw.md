## 用法

来源：

- [Official README](https://github.com/rogerioacp/oblivpg_ftw/blob/5a26f0ffa02242aa187d22bde8ca8dc3bc05dde0/README.md)
- [Extension SQL](https://github.com/rogerioacp/oblivpg_ftw/blob/5a26f0ffa02242aa187d22bde8ca8dc3bc05dde0/oblivpg_fdw--1.0.sql)
- [FDW implementation](https://github.com/rogerioacp/oblivpg_ftw/blob/5a26f0ffa02242aa187d22bde8ca8dc3bc05dde0/oblivpg_fdw.c)

oblivpg_fdw 是一个研究原型，通过 Intel SGX 和使用 Path ORAM 或 Forest ORAM 的安全不经意引擎实现不经意表访问。它与自定义 PostgreSQL 源码树构建、enclave 库、共享内存、镜像表及手工提供的对象标识符紧密耦合。它不是通用或生产可用的 FDW。

### 核心流程

创建扩展会注册包装器和名为 `obliv` 的服务器，并创建全局映射表 `obl_ftw`。README 的测试流程随后手工记录外部表、镜像表和镜像索引的 OID，再打开并初始化 enclave：

```sql
CREATE EXTENSION oblivpg_fdw;

INSERT INTO obl_ftw (
  ftw_table_oid, mirror_table_oid, mirror_index_oid,
  ftw_table_nblocks, ftw_index_nblocks, init
) VALUES (16399, 16392, 16395, 100, 100, false);

SELECT open_enclave();
SELECT init_soe(16399);
```

这里的数字 OID 只是占位符；上游要求从实际测试模式中解析它们。安装 SQL 中的示例外部表定义被注释掉了，因此没有可盲目复制的可移植表契约。

### 重要对象

- `oblivpg_fdw` 是包装器，`obliv` 是安装时创建的服务器。
- `obl_ftw` 存储外部表、镜像表和镜像索引 OID 以及块计数。
- `open_enclave` 和 `close_enclave` 控制 SGX enclave 生命周期。
- `init_soe` 为一条映射初始化安全不经意引擎。
- `load_blocks`、`attach_shmem` 和 `set_nextterm` 暴露底层测试及数据装载钩子。

### 安全与兼容性说明

构建依赖 Intel SGX SDK 和外部 SOE/ORAM 库，而且编译模式可以绕过 enclave 或注入虚拟查询。硬编码 OID 在还原及重建模式后不稳定。安装扩展会删除任何已有的 `obliv` 服务器和 `ftw_users` 外部表，然后创建全局对象。该原型没有提供安全证明、现代兼容性承诺、密钥管理流程或恢复方案；只能在隔离环境中用于复现实验研究。
