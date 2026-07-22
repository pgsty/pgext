## 用法

来源：

- [官方上游文档](https://github.com/corlogic-code/file_fdw_program/blob/412b751c6af32882f335515e5af1337000624ce2/doc/file_fdw_program.md)
- [官方扩展控制文件](https://github.com/corlogic-code/file_fdw_program/blob/412b751c6af32882f335515e5af1337000624ce2/file_fdw_program.control)
- [官方实现](https://github.com/corlogic-code/file_fdw_program/blob/412b751c6af32882f335515e5af1337000624ce2/src/file_fdw_program.c)

`file_fdw_program` 1.0.1 是 PostgreSQL 9.4 时代的只读 FDW，通过 COPY 解析外部命令的标准输出。它回移了后来在 PostgreSQL 10 中进入核心 `file_fdw` 的 `program` 选项。在受支持的现代 PostgreSQL 发布上，应优先使用内置 `file_fdw` 实现。

### 核心流程

只有超级用户可以定义或修改包含命令的外部表选项：

```sql
CREATE EXTENSION file_fdw_program;
CREATE SERVER local_program FOREIGN DATA WRAPPER file_fdw_program;

CREATE FOREIGN TABLE generated_rows (
  one text,
  two text,
  three text
)
SERVER local_program
OPTIONS (
  program 'printf "one,two,three\n"',
  format 'csv'
);

SELECT * FROM generated_rows;
```

`program` 与 `filename` 只能选择一个。其他表和列选项遵循目标服务器版本支持的 COPY FROM 解析选项。

### 执行边界

命令由 PostgreSQL 服务器以运行数据库的操作系统账号执行，而不是由 SQL 客户端执行；它会继承服务器端文件系统、环境、网络和进程权限。验证器只允许超级用户修改外部表选项，但查询该表就会运行已保存的命令；必须严格控制所有权、授权和依赖链，避免不可信角色触发特权命令执行或修改被引用对象。

### 兼容性与运维

包装器只读，而且无法从文件系统元数据估算命令输出，因此计划器使用后备估算。命令必须稳定地产生与声明类型和 COPY 格式一致的行；非零退出、标准错误、编码错误或畸形行都会使扫描失败。命令文本可能暴露在目录元数据和日志中，所以不能放入密钥。应在扩展之外配置操作系统限制、绝对可执行文件路径、最小环境与明确超时。该项目已弃用并面向旧 PostgreSQL API；只能在固定的旧服务器上使用，并测试取消、重复扫描、并行计划、命令失败与权限行为。
