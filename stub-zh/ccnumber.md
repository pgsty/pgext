## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/tvondra/ccnumber/blob/ec41b073c9d601c164be9798b5e126eda45a41fa/README.md)
- [1.0 版本 SQL 对象](https://github.com/tvondra/ccnumber/blob/ec41b073c9d601c164be9798b5e126eda45a41fa/ccnumber--1.0.sql)
- [扩展 control 文件](https://github.com/tvondra/ccnumber/blob/ec41b073c9d601c164be9798b5e126eda45a41fa/ccnumber.control)

`ccnumber` 是用于卡号或其他敏感字节串的加密数据类型概念验证。比较操作通过 TCP 发送给独立的 `ccnumber-comparator` 进程，使 PostgreSQL 无需在进程内解密即可使用 B-tree 与哈希索引、排序、等值比较以及 `min`/`max`。

```sql
CREATE EXTENSION ccnumber;
CREATE TABLE payment_token (number ccnumber PRIMARY KEY);

SHOW ccnumber.comparator_host;
SHOW ccnumber.comparator_port;
```

计算数据值前必须先启动单独构建的 comparator。其端点由 `ccnumber.comparator_host` 和 `ccnumber.comparator_port` 选择；`ccnumber.optimize_remote_calls` 控制基于哈希的优化。

### 安全状态

上游明确称该项目只有最小功能、效率低、基本未经测试，且不适合生产。它使用 `libsodium`，但加密密钥硬编码、TCP 传输未认证、排序会泄露信息，而且每个密文都有额外存储开销。它只适合研究 trusted-component 设计，绝不能用于真实支付数据或密钥。
