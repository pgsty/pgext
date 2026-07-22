## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/tvondra/ccnumber/blob/ec41b073c9d601c164be9798b5e126eda45a41fa/README.md)
- [1.0 版本 SQL 对象](https://github.com/tvondra/ccnumber/blob/ec41b073c9d601c164be9798b5e126eda45a41fa/ccnumber--1.0.sql)
- [后端实现](https://github.com/tvondra/ccnumber/blob/ec41b073c9d601c164be9798b5e126eda45a41fa/ccnumber.c)
- [比较器实现](https://github.com/tvondra/ccnumber/blob/ec41b073c9d601c164be9798b5e126eda45a41fa/ccnumber-comparator.c)

`ccnumber` 是一种实验性的加密字节串类型。PostgreSQL 保存密文，独立的 TCP 比较器负责解密和比较，从而支持等值、排序、B-tree 与哈希操作符类，以及 `min`/`max`。上游明确说明该项目只是功能极少、低效且基本未经测试的概念验证；不要将其用于支付数据或其他秘密。

### 核心流程

项目没有用于明文加密的 SQL 函数。应使用上游单独构建的工具生成密文，启动 `ccnumber-comparator`，再以 bytea 风格的十六进制输入提供密文：

```sql
CREATE EXTENSION ccnumber;
CREATE TABLE payment_token (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    number ccnumber UNIQUE
);

SHOW ccnumber.comparator_host;
SHOW ccnumber.comparator_port;
SHOW ccnumber.optimize_remote_calls;
```

`ccnumber.comparator_host` 默认是 `127.0.0.1`，`ccnumber.comparator_port` 默认是 `9999`，`ccnumber.optimize_remote_calls` 默认是 `true`。每次比较都可能依赖外部服务；比较器不可用时，相关查询会失败。

### 安全与运维边界

已复核的实现把加密密钥硬编码在比较器中。其 TCP 协议既无身份认证也无传输加密，除非进行严格网络隔离，否则该进程会成为可远程访问的解密与比较预言机。排序会泄露顺序；优化和哈希索引会泄露等值相关信息；密文还会增加存储开销。

后端与比较器都会把攻击者可影响的值复制到固定大小缓冲区，且缺少充分的长度检查。因此，任意或超长 `ccnumber` 输入可能破坏内存。只能在一次性环境中将此代码用于设计研究，不能把它当作安全控制或生产扩展。
