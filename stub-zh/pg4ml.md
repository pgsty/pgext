

## 用法

> [pg4ml](https://gitee.com/guotiecheng/plpgsql_pg4ml)：PostgreSQL 机器学习框架。
> 来源：[README.md](https://gitee.com/guotiecheng/plpgsql_pg4ml)

`pg4ml` 是一个 PostgreSQL 扩展，完全在数据库内使用 PL/pgSQL 和 PL/Python 实现机器学习框架。它通过 SQL 提供矩阵运算、神经网络构建与训练、聚类算法和科学计算功能。


--------

### 前置条件

- PostgreSQL >= 14，支持 Python3
- 依赖扩展：`plpgsql`、`tablefunc`、`cube`、`plpython3u`

### 快速开始

```sql
CREATE EXTENSION pg4ml CASCADE;
-- 这也会创建所需的依赖：plpgsql, tablefunc, cube, plpython3u
```


--------

### 功能特性

#### 矩阵运算

框架在 `sm_sc` schema 下提供了全面的矩阵运算库：

- **逐元素运算**：算术、比较、取整、拼接、布尔、位运算、复数和广播运算
- **矩阵运算**：乘法、转置、翻转、旋转、拼接
- **构造**：采样、替换、填充、字符匹配、随机生成
- **三角函数**：矩阵上的广播运算
- **聚合**：切片级聚合、矩阵级聚合、按切片值排序、定位极值位置

#### 切片聚合示例

按垂直切片求平均（每组 2 行）：

```sql
SELECT sm_sc.fv_aggr_slice_avg(
    array[[1.5, 11.5],
          [2.1, 12.1],
          [3.3, 13.3],
          [4.3, 14.3],
          [5.5, 15.5],
          [6.1, 16.1]],
    array[2, 1]
);
-- 返回: array[[1.8, 11.8],[3.8, 13.8],[5.8, 15.8]]
```

2x3 块上的最大池化：

```sql
SELECT sm_sc.fv_aggr_slice_max(
    array[[2.3, 5.1, 8.2, 2.56, 3.33, -1.9],
          [3.25, 6.4, 6.6, 6.9, -2.65, -4.6],
          [-2.3, 5.1, -8.2, 2.56, -3.33, -1.9],
          [3.25, -6.4, -6.6, 6.9, -2.65, -4.6]],
    array[2, 3]
);
-- 返回: array[[8.2, 6.9],[5.1, 6.9]]
```

#### 神经网络

框架支持深度神经网络的构建和训练：

- **节点和路径表**：`sm_sc.tb_nn_node` / `sm_sc.tb_nn_path` 用于定义网络结构
- **训练输入缓冲区**：`sm_sc.tb_nn_train_input_buff` 用于接收训练数据
- **任务管理**：`sm_sc.tb_classify_task` 用于部署和管理训练任务
- **激活函数**、**卷积**、**池化**、**Lambda 运算**
- **损失函数**、**导数计算**、**反向传播**
- **推理**：`sm_sc.ft_nn_in_out` 用于将测试/验证数据通过训练好的模型运行

#### 聚类

- **K-means++**：通过 `sm_sc.prc_kmeans_pp` 过程
- **DBSCAN**：通过 `sm_sc.prc_dbscan_pp` 过程

两者都使用 `sm_sc.tb_cluster_task` 进行任务部署和管理。

#### 科学计算

- 波形处理
- 计算图 JSON 序列化/反序列化
- 复数运算
- 线性代数


--------

### 性能提示

- 启用调试模式：`SET session pg4ml._v_is_debug_check = '1';`
- 矩阵乘法使用 `plpython3u` 调用 numpy 进行优化
- 调整 PostgreSQL 并行参数以支持多线程训练：
  - `max_parallel_workers_per_gather`
  - `force_parallel_mode`
  - `parallel_setup_cost`, `parallel_tuple_cost`
- 考虑使用 `pg_strom` 扩展进行 GPU 加速
