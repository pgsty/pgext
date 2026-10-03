## 用法

来源：

- [Control 1.0](https://github.com/TimeB30/Diploma/blob/ee6649138162c2a9db2e48b4f5f1a5651741daba/Implementation/ppm.control)
- [SQL 1.0](https://github.com/TimeB30/Diploma/blob/ee6649138162c2a9db2e48b4f5f1a5651741daba/Implementation/ppm--1.0.sql)
- [Type implementation](https://github.com/TimeB30/Diploma/blob/ee6649138162c2a9db2e48b4f5f1a5651741daba/Implementation/ppm_type.c)
- [Official test workflow](https://github.com/TimeB30/Diploma/blob/ee6649138162c2a9db2e48b4f5f1a5651741daba/Implementation/test_ppm.sql)
- [PostgreSQL 15 build recipe](https://github.com/TimeB30/Diploma/blob/ee6649138162c2a9db2e48b4f5f1a5651741daba/Implementation/Dockerfile)

`ppm` 是一个学术原型，以 PPM 压缩算法实现自定义文本类型。control 文件与安装 SQL 定义的扩展版本为 1.0；启动日志中的 4.0 文本不是扩展版本。仓库提供 PostgreSQL 15 容器构建配置，但没有明确许可证或生产支持保证。

### 基本用法

由超级用户安装扩展，即可使用 `ppm_text` 类型：

```sql
CREATE EXTENSION ppm;
SET ppm.max_order = 4;
CREATE TABLE compressed_notes (body ppm_text);
INSERT INTO compressed_notes VALUES ('Repeated text, repeated text.');
SELECT body, pg_column_size(body) FROM compressed_notes;
```

类型输入函数压缩 C 字符串，输出函数将其解压后用于显示。变长内部表示采用 external 存储，避免再经过内置 TOAST 压缩。它是独立类型，不会替换服务器的默认压缩机制。

### 配置与对象

`ppm.max_order` 设置新编码值使用的 PPM 上下文阶数，默认 2，范围 0-8，可在会话中设置。编码器会将所用阶数一并保存；解码时读取值内保存的阶数，而不是当前设置。

`ppm_in` 与 `ppm_out` 是该类型的输入输出函数。安装 SQL 未定义文本运算符、索引或显式文本转换；SQL 能力限于存储和输入输出。

### 使用边界

上游使用流程是比较存储大小和耗时的压缩实验。将其用于重要数据前，应保留原始文本并独立验证往返转换。项目没有文档承诺稳定的磁盘格式、升级路径、复制测试矩阵或生产持久性验证。除仓库提供的 PostgreSQL 15 配置外，其他版本兼容性及再分发许可均未确认。
