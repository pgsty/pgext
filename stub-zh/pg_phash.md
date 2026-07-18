## 用法

来源：

- [固定提交的上游 README](https://github.com/PixNyanNyan/postgres-phash-hamming/blob/4b9b69c41acd75be1fe4e6d91f1117622d2c77f2/README.md)
- [1.0 版安装 SQL](https://github.com/PixNyanNyan/postgres-phash-hamming/blob/4b9b69c41acd75be1fe4e6d91f1117622d2c77f2/pg_phash--1.0.sql)
- [固定提交的 C 实现](https://github.com/PixNyanNyan/postgres-phash-hamming/blob/4b9b69c41acd75be1fe4e6d91f1117622d2c77f2/pg_phash.c)
- [固定提交的扩展 control 文件](https://github.com/PixNyanNyan/postgres-phash-hamming/blob/4b9b69c41acd75be1fe4e6d91f1117622d2c77f2/pg_phash.control)

pg_phash 1.0 版比较两个以十进制字符串表示的 64 位感知图像 hash。它对解析出的整数做 XOR，并返回不同 bit 的数量。

### 示例

```sql
CREATE EXTENSION pg_phash;

SELECT phash_hamming(
    '13121266429874464083',
    '10869537466045227287'
);
```

上游示例返回 16。两个相同的合法 hash 返回零。SQL 函数标记为 strict、immutable 和 parallel safe，因此任一输入为 null 时结果为 null，常量调用也可能被折叠。

### 解析与范围限制

README 声称结果范围是 0 至 255，但实现只比较一个 64 位值，实际只能返回 0 至 64。输入通过 C strtoull 转换，代码却不检查 conversion end pointer 或 errno。空字符串、负数、畸形文本、尾随字符和溢出输入可能被静默解释成零、前缀值或最大无符号整数，而不是被拒绝。

输入长度在分配前还会存入一字节 unsigned 变量，因此超过 255 字节的字符串可能回绕并被截断。只应接受无符号 64 位范围内的规范十进制字符串，在调用前完成验证，并为零、最大值、畸形文本和溢出添加回归测试。

项目只记录了 PostgreSQL 9.6 CI 示例，自 2021 年后没有变化，也没有当前 PostgreSQL 主版本发行矩阵。仓库没有归档，但这个小型 C 函数仍应针对确切服务器版本重新构建和测试。
