---
title: "扩展"
breadcrumbs: false
excludeSearch: true
weight: 1
---

以下列出了归属 16 个分类的可用 PostgreSQL 扩展（424 ext / 356 pkg）:

{{< category "time" >}}
{{< category "gis" >}}
{{< category "rag" >}}
{{< category "fts" >}}
{{< category "olap" >}}
{{< category "feat" >}}
{{< category "lang" >}}
{{< category "type" >}}
{{< category "util" >}}
{{< category "func" >}}
{{< category "admin" >}}
{{< category "stat" >}}
{{< category "sec" >}}
{{< category "fdw" >}}
{{< category "sim" >}}
{{< category "etl" >}}


## TIME

时间时态扩展：时序数据库 TimescaleDB，时态数据库，版本控制表，定时任务，异步后台任务调度扩展。

| ID | 扩展 | 包 | 版本 | 描述 |
|:---:|:---|:---|:---|:---|
| 1000 | {{< ext "timescaledb" >}} | {{< ext "timescaledb" "timescaledb" >}} | 2.22.0 | 时序数据库扩展插件 |
| 1010 | {{< ext "timescaledb_toolkit" >}} | {{< ext "timescaledb_toolkit" "timescaledb_toolkit" >}} | 1.21.0 | 超表分析查询，时间序列流式处理，以及其他SQL工具 |
| 1020 | {{< ext "timeseries" >}} | {{< ext "timeseries" "pg_timeseries" >}} | 0.1.6 | 时序数据API封装 |
| 1030 | {{< ext "periods" >}} | {{< ext "periods" "periods" >}} | 1.2.3 | 为 PERIODs 和 SYSTEM VERSIONING 提供标准 SQL 功能 |
| 1040 | {{< ext "temporal_tables" >}} | {{< ext "temporal_tables" "temporal_tables" >}} | 1.2.2 | 时态表功能支持 |
| 1050 | {{< ext "emaj" >}} | {{< ext "emaj" "emaj" >}} | 4.7.0 | 让数据库的子集具有细粒度日志和时间旅行功能 |
| 1060 | {{< ext "table_version" >}} | {{< ext "table_version" "table_version" >}} | 1.11.1 | PostgreSQL 版本控制表扩展 |
| 1070 | {{< ext "pg_cron" >}} | {{< ext "pg_cron" "pg_cron" >}} | 1.6.7 | 定时任务调度器 |
| 1080 | {{< ext "pg_task" >}} | {{< ext "pg_task" "pg_task" >}} | 1.0.0 | 在特定时间点在后台执行SQL命令 |
| 1090 | {{< ext "pg_later" >}} | {{< ext "pg_later" "pg_later" >}} | 0.3.0 | 执行查询，并在稍后异步获取查询结果 |
| 1100 | {{< ext "pg_background" >}} | {{< ext "pg_background" "pg_background" >}} | 1.3 | 在后台运行 SQL 查询 |

## GIS

地理空间扩展：PostGIS，地理空间数据类型、函数与索引，天空索引 Q3C，OGR FDW， 寻路算法，地理正/逆查询。

| ID | 扩展 | 包 | 版本 | 描述 |
|:---:|:---|:---|:---|:---|
| 1500 | {{< ext "postgis" >}} | {{< ext "postgis" "postgis" >}} | 3.6.0 | PostGIS 几何和地理空间扩展 |
| 1501 | {{< ext "postgis_topology" >}} | {{< ext "postgis" "postgis" >}} | 3.6.0 | PostGIS 拓扑空间类型和函数 |
| 1502 | {{< ext "postgis_raster" >}} | {{< ext "postgis" "postgis" >}} | 3.6.0 | PostGIS 光栅类型和函数 |
| 1503 | {{< ext "postgis_sfcgal" >}} | {{< ext "postgis" "postgis" >}} | 3.6.0 | PostGIS SFCGAL 函数 |
| 1504 | {{< ext "postgis_tiger_geocoder" >}} | {{< ext "postgis" "postgis" >}} | 3.6.0 | PostGIS tiger 地理编码器和反向地理编码器 |
| 1505 | {{< ext "address_standardizer" >}} | {{< ext "postgis" "postgis" >}} | 3.6.0 | 地址标准化函数。 |
| 1506 | {{< ext "address_standardizer_data_us" >}} | {{< ext "postgis" "postgis" >}} | 3.6.0 | 地址标准化函数：美国数据集示例 |
| 1510 | {{< ext "pgrouting" >}} | {{< ext "pgrouting" "pgrouting" >}} | 3.8.0 | 提供寻路能力 |
| 1520 | {{< ext "pointcloud" >}} | {{< ext "pointcloud" "pointcloud" >}} | 1.2.5 | 提供激光雷达点云数据类型支持 |
| 1521 | {{< ext "pointcloud_postgis" >}} | {{< ext "pointcloud" "pointcloud" >}} | 1.2.5 | 将激光雷达点云与PostGIS几何类型相集成 |
| 1530 | {{< ext "h3" >}} | {{< ext "h3" "pg_h3" >}} | 4.2.3 | H3六边形层级索引支持 |
| 1531 | {{< ext "h3_postgis" >}} | {{< ext "h3" "pg_h3" >}} | 4.2.3 | H3与PostGIS集成的扩展插件 |
| 1540 | {{< ext "q3c" >}} | {{< ext "q3c" "q3c" >}} | 2.0.1 | Q3C天空索引插件 |
| 1550 | {{< ext "ogr_fdw" >}} | {{< ext "ogr_fdw" "ogr_fdw" >}} | 1.1.7 | GIS 数据外部数据源包装器 |
| 1560 | {{< ext "geoip" >}} | {{< ext "geoip" "geoip" >}} | 0.3.0 | IP 地理位置扩展（围绕 MaxMind GeoLite 数据集的包装器） |
| 1570 | {{< ext "pg_polyline" >}} | {{< ext "pg_polyline" "pg_polyline" >}} | 0.0.1 | Google快速Polyline编码解码扩展 |
| 1590 | {{< ext "pg_geohash" >}} | {{< ext "pg_geohash" "pg_geohash" >}} | 1.0 | 使用GeoHash处理空间坐标的函数包 |
| 1650 | {{< ext "mobilitydb" >}} | {{< ext "mobilitydb" "mobilitydb" >}} | 1.2.0 | MobilityDB地理空间投影数据管理分析平台 |
| 1680 | {{< ext "tzf" >}} | {{< ext "tzf" "pg_tzf" >}} | 0.2.2 | 快速根据GPS经纬度坐标查找时区 |
| 1690 | {{< ext "earthdistance" >}} | {{< ext "earthdistance" "earthdistance" >}} | 1.2 | 计算地球表面上的大圆距离 |

## RAG

AI与RAG扩展插件：向量数据库，DiskANN 向量索引，相似度度量函数集，库内机器学习与推理 pgml，等等。

| ID | 扩展 | 包 | 版本 | 描述 |
|:---:|:---|:---|:---|:---|
| 1800 | {{< ext "vector" >}} | {{< ext "vector" "pgvector" >}} | 0.8.1 | 向量数据类型和 ivfflat / hnsw 访问方法 |
| 1810 | {{< ext "vchord" >}} | {{< ext "vchord" "vchord" >}} | 0.5.1 | 使用Rust重写的高性能向量扩展 |
| 1820 | {{< ext "vectorscale" >}} | {{< ext "vectorscale" "pgvectorscale" >}} | 0.8.0 | 使用DiskANN算法对向量进行高效索引 |
| 1830 | {{< ext "vectorize" >}} | {{< ext "vectorize" "pg_vectorize" >}} | 0.22.2 | 在PostgreSQL中封装RAG向量检索服务 |
| 1840 | {{< ext "pg_similarity" >}} | {{< ext "pg_similarity" "pg_similarity" >}} | 1.0 | 提供17种距离度量函数 |
| 1850 | {{< ext "smlar" >}} | {{< ext "smlar" "smlar" >}} | 1.0 | 高效的相似度搜索函数 |
| 1860 | {{< ext "pg_summarize" >}} | {{< ext "pg_summarize" "pg_summarize" >}} | 0.0.1 | 使用LLM对文本字段进行总结 |
| 1870 | {{< ext "pg_tiktoken" >}} | {{< ext "pg_tiktoken" "pg_tiktoken" >}} | 0.0.1 | 在PostgreSQL中计算OpenAI使用的Token数 |
| 1880 | {{< ext "pg4ml" >}} | {{< ext "pg4ml" "pg4ml" >}} | 2.0 | PG4ML是一个机器学习框架 |
| 1890 | {{< ext "pgml" >}} | {{< ext "pgml" "pgml" >}} | 2.10.0 | PostgresML：用SQL运行机器学习算法并训练模型 |

## FTS

全文检索扩展：ES 替代 pg_search，BM25，中文分词，欧洲语言分词字典 hunspell，模糊检索，2-gram/3-gram 索引。

| ID | 扩展 | 包 | 版本 | 描述 |
|:---:|:---|:---|:---|:---|
| 2100 | {{< ext "pg_search" >}} | {{< ext "pg_search" "pg_search" >}} | 0.18.1 | ParadeDB BM25算法全文检索插件，ES全文检索 |
| 2110 | {{< ext "pgroonga" >}} | {{< ext "pgroonga" "pgroonga" >}} | 4.0.0 | 使用Groonga，面向所有语言的高速全文检索平台 |
| 2111 | {{< ext "pgroonga_database" >}} | {{< ext "pgroonga" "pgroonga" >}} | 4.0.0 | PGGroonga 数据库管理模块 |
| 2120 | {{< ext "pg_bigm" >}} | {{< ext "pg_bigm" "pg_bigm" >}} | 1.2 | 基于二字组的多语言全文检索扩展 |
| 2130 | {{< ext "zhparser" >}} | {{< ext "zhparser" "zhparser" >}} | 2.3 | 中文分词，全文搜索解析器 |
| 2140 | {{< ext "pg_bestmatch" >}} | {{< ext "pg_bestmatch" "pg_bestmatch" >}} | 0.0.1 | 在数据库内生成BM25稀疏向量 |
| 2150 | {{< ext "vchord_bm25" >}} | {{< ext "vchord_bm25" "vchord_bm25" >}} | 0.2.1 | BM25排序算法 |
| 2160 | {{< ext "pg_tokenizer" >}} | {{< ext "pg_tokenizer" "pg_tokenizer" >}} | 0.1.0 | 用于全文检索的分词器 |
| 2170 | {{< ext "hunspell_cs_cz" >}} | {{< ext "hunspell_cs_cz" "hunspell_cs_cz" >}} | 1.0 | Hunspell捷克语全文检索词典 |
| 2171 | {{< ext "hunspell_de_de" >}} | {{< ext "hunspell_de_de" "hunspell_de_de" >}} | 1.0 | Hunspell德语全文检索词典 |
| 2172 | {{< ext "hunspell_en_us" >}} | {{< ext "hunspell_en_us" "hunspell_en_us" >}} | 1.0 | Hunspell英语全文检索词典 |
| 2173 | {{< ext "hunspell_fr" >}} | {{< ext "hunspell_fr" "hunspell_fr" >}} | 1.0 | Hunspell法语全文检索词典 |
| 2174 | {{< ext "hunspell_ne_np" >}} | {{< ext "hunspell_ne_np" "hunspell_ne_np" >}} | 1.0 | Hunspell尼泊尔语全文检索词典 |
| 2175 | {{< ext "hunspell_nl_nl" >}} | {{< ext "hunspell_nl_nl" "hunspell_nl_nl" >}} | 1.0 | Hunspell荷兰语全文检索词典 |
| 2176 | {{< ext "hunspell_nn_no" >}} | {{< ext "hunspell_nn_no" "hunspell_nn_no" >}} | 1.0 | Hunspell挪威语全文检索词典 |
| 2177 | {{< ext "hunspell_pt_pt" >}} | {{< ext "hunspell_pt_pt" "hunspell_pt_pt" >}} | 1.0 | Hunspell葡萄牙语全文检索词典 |
| 2178 | {{< ext "hunspell_ru_ru" >}} | {{< ext "hunspell_ru_ru" "hunspell_ru_ru" >}} | 1.0 | Hunspell俄语全文检索词典 |
| 2179 | {{< ext "hunspell_ru_ru_aot" >}} | {{< ext "hunspell_ru_ru_aot" "hunspell_ru_ru_aot" >}} | 1.0 | Hunspell俄语全文检索词典（来自AOT.ru小组） |
| 2180 | {{< ext "fuzzystrmatch" >}} | {{< ext "fuzzystrmatch" "fuzzystrmatch" >}} | 1.2 | 确定字符串之间的相似性和距离 |
| 2190 | {{< ext "pg_trgm" >}} | {{< ext "pg_trgm" "pg_trgm" >}} | 1.6 | 文本相似度测量函数与模糊检索 |

## OLAP

分析能力扩展：列式存储，DuckDB与外部数据源包装器，Parquet S3，数据冷热分级存储，分布式计算，透明分片，GPU加速

| ID | 扩展 | 包 | 版本 | 描述 |
|:---:|:---|:---|:---|:---|
| 2400 | {{< ext "citus" >}} | {{< ext "citus" "citus" >}} | 13.2.0 | Citus 分布式数据库 |
| 2401 | {{< ext "citus_columnar" >}} | {{< ext "citus" "citus" >}} | 13.2.0 | Citus 列式存储引擎 |
| 2410 | {{< ext "columnar" >}} | {{< ext "columnar" "hydra" >}} | 1.1.2 | 开源列式存储扩展 |
| 2420 | {{< ext "pg_analytics" >}} | {{< ext "pg_analytics" "pg_analytics" >}} | 0.3.7 | 由 DuckDB 驱动的数据分析引擎 |
| 2430 | {{< ext "pg_duckdb" >}} | {{< ext "pg_duckdb" "pg_duckdb" >}} | 0.3.1 | 在PostgreSQL中的嵌入式DuckDB扩展 |
| 2440 | {{< ext "pg_mooncake" >}} | {{< ext "pg_mooncake" "pg_mooncake" >}} | 0.1.2 | PostgreSQL列式存储表 |
| 2450 | {{< ext "duckdb_fdw" >}} | {{< ext "duckdb_fdw" "duckdb_fdw" >}} | 1.1.2 | DuckDB 外部数据源包装器 |
| 2460 | {{< ext "pg_parquet" >}} | {{< ext "pg_parquet" "pg_parquet" >}} | 0.4.3 | 在PostgreSQL与本地/S3中的Parquet文件复制数据 |
| 2500 | {{< ext "pg_fkpart" >}} | {{< ext "pg_fkpart" "pg_fkpart" >}} | 1.7.0 | 按外键实用程序进行表分区的扩展 |
| 2510 | {{< ext "pg_partman" >}} | {{< ext "pg_partman" "pg_partman" >}} | 5.2.4 | 用于按时间或 ID 管理分区表的扩展 |
| 2520 | {{< ext "plproxy" >}} | {{< ext "plproxy" "plproxy" >}} | 2.11.0 | 作为过程语言实现的数据库分区 |
| 2530 | {{< ext "pg_strom" >}} | {{< ext "pg_strom" "pg_strom" >}} | 6.0 | 使用GPU与NVMe加速大数据处理 |
| 2590 | {{< ext "tablefunc" >}} | {{< ext "tablefunc" "tablefunc" >}} | 1.0 | 交叉表函数 |

## FEAT

功能特性扩展：图数据库，Hyperloglog，Rum索引，GraphQL，JsonSchema，Hint，虚拟索引，增量物化视图，消息队列等等

| ID | 扩展 | 包 | 版本 | 描述 |
|:---:|:---|:---|:---|:---|
| 2760 | {{< ext "age" >}} | {{< ext "age" "age" >}} | 1.5.0 | Apache AGE，图数据库扩展 （Deb可用） |
| 2770 | {{< ext "hll" >}} | {{< ext "hll" "hll" >}} | 2.18 | hyperloglog 数据类型 |
| 2780 | {{< ext "rum" >}} | {{< ext "rum" "rum" >}} | 1.3.14 | RUM 索引访问方法 |
| 2790 | {{< ext "pg_graphql" >}} | {{< ext "pg_graphql" "pg_graphql" >}} | 1.5.11 | PG内的GraphQL支持 |
| 2800 | {{< ext "pg_jsonschema" >}} | {{< ext "pg_jsonschema" "pg_jsonschema" >}} | 0.3.3 | 提供JSON Schema校验能力 |
| 2810 | {{< ext "jsquery" >}} | {{< ext "jsquery" "jsquery" >}} | 1.2 | 用于内省 JSONB 数据类型的查询类型 |
| 2820 | {{< ext "pg_hint_plan" >}} | {{< ext "pg_hint_plan" "pg_hint_plan" >}} | 1.7.1 | 添加强制指定执行计划的能力 |
| 2830 | {{< ext "hypopg" >}} | {{< ext "hypopg" "hypopg" >}} | 1.4.2 | 假设索引，用于创建一个虚拟索引检验执行计划 |
| 2840 | {{< ext "index_advisor" >}} | {{< ext "index_advisor" "index_advisor" >}} | 0.2.0 | 查询索引建议器 |
| 2850 | {{< ext "plan_filter" >}} | {{< ext "plan_filter" "pg_plan_filter" >}} | 0.0.1 | 使用执行计划代价过滤阻止特定查询语句 |
| 2860 | {{< ext "imgsmlr" >}} | {{< ext "imgsmlr" "imgsmlr" >}} | 1.0 | 使用Haar小波分析计算图片相似度 |
| 2870 | {{< ext "pg_ivm" >}} | {{< ext "pg_ivm" "pg_ivm" >}} | 1.12 | 增量维护的物化视图 |
| 2880 | {{< ext "pg_incremental" >}} | {{< ext "pg_incremental" "pg_incremental" >}} | 1.2.0 | 增量处理流式事件 |
| 2900 | {{< ext "pgmq" >}} | {{< ext "pgmq" "pgmq" >}} | 1.5.1 | 基于Postgres实现类似AWS SQS/RSMQ的消息队列 |
| 2910 | {{< ext "pgq" >}} | {{< ext "pgq" "pgq" >}} | 3.5.1 | 通用队列的PG实现 |
| 2920 | {{< ext "orioledb" >}} | {{< ext "orioledb" "orioledb" >}} | 1.5 | OrioleDB，下一代事务处理引擎 |
| 2930 | {{< ext "pg_cardano" >}} | {{< ext "pg_cardano" "pg_cardano" >}} | 1.0.5 | Cardano相关工具包：加密函数，地址编解码，区块链处理 |
| 2940 | {{< ext "rdkit" >}} | {{< ext "rdkit" "rdkit" >}} | 202503.1 | 在PostgreSQL化学领域数据管理功能 |
| 2951 | {{< ext "omni" >}} | {{< ext "omni" "omnigres" >}} | 0.2.9 | PostgreSQL即平台，Omnigres主扩展与加载器 |
| 2952 | {{< ext "omni_auth" >}} | {{< ext "omni" "omnigres" >}} | 0.1.3 | Omnigres 基础会话认证管理模块 |
| 2953 | {{< ext "omni_aws" >}} | {{< ext "omni" "omnigres" >}} | 0.1.2 | Omnigres AWS S3 API封装 |
| 2954 | {{< ext "omni_cloudevents" >}} | {{< ext "omni" "omnigres" >}} | 0.1.0 | Omnigres CloudEvents 支持 |
| 2955 | {{< ext "omni_containers" >}} | {{< ext "omni" "omnigres" >}} | 0.2.0 | Omnigres Docker容器管理模块 |
| 2956 | {{< ext "omni_credentials" >}} | {{< ext "omni" "omnigres" >}} | 0.2.0 | Omnigres 应用密钥管理模块 |
| 2958 | {{< ext "omni_email" >}} | {{< ext "omni" "omnigres" >}} | 0 | Omnigres Email 框架 |
| 2959 | {{< ext "omni_http" >}} | {{< ext "omni" "omnigres" >}} | 0.1.0 | Omnigres 基本HTTP类型 |
| 2960 | {{< ext "omni_httpc" >}} | {{< ext "omni" "omnigres" >}} | 0.1.5 | Omnigres HTTP客户端 |
| 2961 | {{< ext "omni_httpd" >}} | {{< ext "omni" "omnigres" >}} | 0.4.6 | Omnigres HTTP服务器 |
| 2962 | {{< ext "omni_id" >}} | {{< ext "omni" "omnigres" >}} | 0.4.2 | Omnigres ID身份数据类型 |
| 2963 | {{< ext "omni_json" >}} | {{< ext "omni" "omnigres" >}} | 0.1.1 | Omnigres JSON工具箱 |
| 2964 | {{< ext "omni_kube" >}} | {{< ext "omni" "omnigres" >}} | 0.1.1 | Omnigres Kubernetes集成模块 |
| 2965 | {{< ext "omni_ledger" >}} | {{< ext "omni" "omnigres" >}} | 0.1.2 | Omnigres 金融账本模块 |
| 2966 | {{< ext "omni_manifest" >}} | {{< ext "omni" "omnigres" >}} | 0.1.2 | Omnigres 包管理清单模块 |
| 2967 | {{< ext "omni_mimetypes" >}} | {{< ext "omni" "omnigres" >}} | 0.1.0 | Omnigres MIME数据类型 |
| 2968 | {{< ext "omni_os" >}} | {{< ext "omni" "omnigres" >}} | 0.1.1 | Omnigres 操作系统集成模块 |
| 2969 | {{< ext "omni_polyfill" >}} | {{< ext "omni" "omnigres" >}} | 0.2.2 | Omnigres Postgres多态API |
| 2970 | {{< ext "omni_python" >}} | {{< ext "omni" "omnigres" >}} | 0.1.1 | Omnigres 第一类Python支持模块 |
| 2971 | {{< ext "omni_regex" >}} | {{< ext "omni" "omnigres" >}} | 0.1.0 | Omnigres PCRE兼容正则表达式模块 |
| 2972 | {{< ext "omni_rest" >}} | {{< ext "omni" "omnigres" >}} | 0.1.1 | Omnigres REST API 工具包 |
| 2973 | {{< ext "omni_schema" >}} | {{< ext "omni" "omnigres" >}} | 0.3.0 | Omnigres 高级模式管理组件 |
| 2974 | {{< ext "omni_seq" >}} | {{< ext "omni" "omnigres" >}} | 0.1.1 | Omnigres 分布式整型序列号 |
| 2975 | {{< ext "omni_service" >}} | {{< ext "omni" "omnigres" >}} | 0.1.0 | Omnigres 服务管理器 |
| 2976 | {{< ext "omni_session" >}} | {{< ext "omni" "omnigres" >}} | 0.2.0 | Omnigres 会话管理器 |
| 2977 | {{< ext "omni_sql" >}} | {{< ext "omni" "omnigres" >}} | 0.5.1 | Omnigres SQL编程组件 |
| 2979 | {{< ext "omni_sqlite" >}} | {{< ext "omni" "omnigres" >}} | 0.1.2 | Omnigres 嵌入的SQLite支持 |
| 2980 | {{< ext "omni_test" >}} | {{< ext "omni" "omnigres" >}} | 0.4.0 | Omnigres 测试框架 |
| 2981 | {{< ext "omni_txn" >}} | {{< ext "omni" "omnigres" >}} | 0.5.0 | Omnigres 事务管理器模块 |
| 2982 | {{< ext "omni_types" >}} | {{< ext "omni" "omnigres" >}} | 0.3.4 | Omnigres 高级数据类型模块 |
| 2983 | {{< ext "omni_var" >}} | {{< ext "omni" "omnigres" >}} | 0.3.0 | Omnigres 局部变量模块 |
| 2984 | {{< ext "omni_vfs" >}} | {{< ext "omni" "omnigres" >}} | 0.2.1 | Omnigres 虚拟文件系统 |
| 2985 | {{< ext "omni_vfs_types_v1" >}} | {{< ext "omni" "omnigres" >}} | 0.1.0 | Omnigres 虚拟文件系统（v1） |
| 2986 | {{< ext "omni_web" >}} | {{< ext "omni" "omnigres" >}} | 0.3.0 | Omnigres Web工具箱 |
| 2987 | {{< ext "omni_worker" >}} | {{< ext "omni" "omnigres" >}} | 0.1.0 | Omnigres 通用Worker池 |
| 2988 | {{< ext "omni_xml" >}} | {{< ext "omni" "omnigres" >}} | 0.1.2 | Omnigres XML工具包 |
| 2989 | {{< ext "omni_yaml" >}} | {{< ext "omni" "omnigres" >}} | 0.1.0 | Omnigres YAML工具包 |
| 2990 | {{< ext "bloom" >}} | {{< ext "bloom" "bloom" >}} | 1.0 | bloom 索引-基于指纹的索引 |

## LANG

存储过程语言扩展：使用各种编程语言开发，调试，打包，分发，测试 PostgreSQL 存储过程：Java，Js，Lua，R，SH，PRQL…

| ID | 扩展 | 包 | 版本 | 描述 |
|:---:|:---|:---|:---|:---|
| 3000 | {{< ext "pg_tle" >}} | {{< ext "pg_tle" "pg_tle" >}} | 1.5.1 | AWS 可信语言扩展 |
| 3010 | {{< ext "plv8" >}} | {{< ext "plv8" "plv8" >}} | 3.2.4 | PL/JavaScript (v8) 可信过程程序语言 |
| 3020 | {{< ext "pllua" >}} | {{< ext "pllua" "pllua" >}} | 2.0.12 | Lua 程序语言 |
| 3021 | {{< ext "hstore_pllua" >}} | {{< ext "pllua" "pllua" >}} | 2.0.12 | Lua 程序语言的Hstore适配扩展 |
| 3030 | {{< ext "plluau" >}} | {{< ext "pllua" "pllua" >}} | 2.0.12 | Lua 程序语言（不受信任的） |
| 3031 | {{< ext "hstore_plluau" >}} | {{< ext "pllua" "pllua" >}} | 2.0.12 | Lua 程序语言的Hstore适配扩展（不受信任的） |
| 3040 | {{< ext "plprql" >}} | {{< ext "plprql" "plprql" >}} | 1.0.0 | 在PostgreSQL使用PRQL——管线式关系查询语言 |
| 3050 | {{< ext "pldbgapi" >}} | {{< ext "pldbgapi" "pldebugger" >}} | 1.8 | 用于调试 PL/pgSQL 函数的服务器端支持 |
| 3060 | {{< ext "plpgsql_check" >}} | {{< ext "plpgsql_check" "plpgsql_check" >}} | 2.8.2 | 对 plpgsql 函数进行扩展检查 |
| 3070 | {{< ext "plprofiler" >}} | {{< ext "plprofiler" "plprofiler" >}} | 4.2.5 | 剖析 PL/pgSQL 函数 |
| 3080 | {{< ext "plsh" >}} | {{< ext "plsh" "plsh" >}} | 1.20220917 | PL/sh 程序语言 |
| 3090 | {{< ext "pljava" >}} | {{< ext "pljava" "pljava" >}} | 1.6.9 | Java 程序语言 |
| 3100 | {{< ext "plr" >}} | {{< ext "plr" "plr" >}} | 8.4.8 | 从数据库中加载R语言解释器并执行R脚本 |
| 3200 | {{< ext "pgtap" >}} | {{< ext "pgtap" "pgtap" >}} | 1.3.3 | PostgreSQL单元测试框架 |
| 3210 | {{< ext "faker" >}} | {{< ext "faker" "faker" >}} | 0.5.3 | 插入生成的测试伪造数据，Python库的包装 |
| 3220 | {{< ext "dbt2" >}} | {{< ext "dbt2" "dbt2" >}} | 0.45.0 | OSDL-DBT-2 测试组件 |
| 3240 | {{< ext "pltcl" >}} | {{< ext "pltcl" "pltcl" >}} | 1.0 | PL/TCL 存储过程语言 |
| 3250 | {{< ext "pltclu" >}} | {{< ext "pltcl" "pltcl" >}} | 1.0 | PL/TCL 存储过程语言（未受信/高权限） |
| 3260 | {{< ext "plperl" >}} | {{< ext "plperl" "plperl" >}} | 1.0 | PL/Perl 存储过程语言 |
| 3261 | {{< ext "bool_plperl" >}} | {{< ext "plperl" "plperl" >}} | 1.0 | 在 bool 和 plperl 之间转换 |
| 3262 | {{< ext "hstore_plperl" >}} | {{< ext "plperl" "plperl" >}} | 1.0 | 在 hstore 和 plperl 之间转换适配类型 |
| 3263 | {{< ext "jsonb_plperl" >}} | {{< ext "plperl" "plperl" >}} | 1.0 | 在 jsonb 和 plperl 之间转换 |
| 3270 | {{< ext "plperlu" >}} | {{< ext "plperlu" "plperlu" >}} | 1.0 | PL/PerlU 存储过程语言（未受信/高权限） |
| 3271 | {{< ext "bool_plperlu" >}} | {{< ext "plperlu" "plperlu" >}} | 1.0 | 在 bool 和 plperlu 之间转换 |
| 3272 | {{< ext "jsonb_plperlu" >}} | {{< ext "plperlu" "plperlu" >}} | 1.0 | 在 jsonb 和 plperlu 之间转换 |
| 3273 | {{< ext "hstore_plperlu" >}} | {{< ext "plperlu" "plperlu" >}} | 1.0 | 在 hstore 和 plperlu 之间转换适配类型 |
| 3280 | {{< ext "plpgsql" >}} | {{< ext "plpgsql" "plpgsql" >}} | 1.0 | PL/pgSQL 程序设计语言 |
| 3290 | {{< ext "plpython3u" >}} | {{< ext "plpython3u" "plpython3u" >}} | 1.0 | PL/Python3 存储过程语言（未受信/高权限） |
| 3291 | {{< ext "jsonb_plpython3u" >}} | {{< ext "plpython3u" "plpython3u" >}} | 1.0 | 在 jsonb 和 plpython3u 之间转换 |
| 3292 | {{< ext "ltree_plpython3u" >}} | {{< ext "plpython3u" "plpython3u" >}} | 1.0 | 在 ltree 和 plpython3u 之间转换 |
| 3293 | {{< ext "hstore_plpython3u" >}} | {{< ext "plpython3u" "plpython3u" >}} | 1.0 | 在 hstore 和 plpython3u 之间转换 |

## TYPE

自定义类型扩展：前缀树，语义版本号，SI单位，位图，无符号整型，高精度数值，有理数，哈希值，IP地址段，球面，RRULE等

| ID | 扩展 | 包 | 版本 | 描述 |
|:---:|:---|:---|:---|:---|
| 3500 | {{< ext "prefix" >}} | {{< ext "prefix" "pg_prefix" >}} | 1.2.10 | 前缀树数据类型 |
| 3510 | {{< ext "semver" >}} | {{< ext "semver" "pg_semver" >}} | 0.40.0 | 语义版本号数据类型 |
| 3520 | {{< ext "unit" >}} | {{< ext "unit" "pgunit" >}} | 7.10 | SI 国标单位扩展 |
| 3530 | {{< ext "pgpdf" >}} | {{< ext "pgpdf" "pgpdf" >}} | 0.1.0 | PDF数据类型，管理函数与全文检索 |
| 3540 | {{< ext "pglite_fusion" >}} | {{< ext "pglite_fusion" "pglite_fusion" >}} | 0.0.5 | 在PG表中嵌入SQLite数据库作为数据类型 |
| 3550 | {{< ext "md5hash" >}} | {{< ext "md5hash" "md5hash" >}} | 1.0.1 | 提供128位MD5的原生数据类型 |
| 3560 | {{< ext "asn1oid" >}} | {{< ext "asn1oid" "asn1oid" >}} | 1.6 | ASN1OID数据类型支持 |
| 3570 | {{< ext "roaringbitmap" >}} | {{< ext "roaringbitmap" "roaringbitmap" >}} | 0.5.4 | 支持RoaringBitmap数据类型 |
| 3580 | {{< ext "pgfaceting" >}} | {{< ext "pgfaceting" "pgfaceting" >}} | 0.2.0 | 使用倒排索引的高速切面查询 |
| 3590 | {{< ext "pg_sphere" >}} | {{< ext "pg_sphere" "pgsphere" >}} | 1.5.1 | 球面对象函数、运算符与索引支持 |
| 3600 | {{< ext "country" >}} | {{< ext "country" "pg_country" >}} | 0.0.3 | 国家代码数据类型，遵循ISO 3166-1标准 |
| 3610 | {{< ext "pg_xenophile" >}} | {{< ext "pg_xenophile" "pg_xenophile" >}} | 0.8.3 | PostgreSQL i8n与l10n工具包 |
| 3611 | {{< ext "l10n_table_dependent_extension" >}} | {{< ext "pg_xenophile" "pg_xenophile" >}} | 0.8.3 | PostgreSQL l10n 工具包 |
| 3620 | {{< ext "currency" >}} | {{< ext "currency" "pg_currency" >}} | 0.0.3 | 使用1字节表示的货币数据类型 |
| 3630 | {{< ext "collection" >}} | {{< ext "collection" "pg_collection" >}} | 1.0.0 | 在PlPGSQL中使用的内存优化高性能集合数据结构 |
| 3700 | {{< ext "pgmp" >}} | {{< ext "pgmp" "pgmp" >}} | 1.0.5 | 多精度算术扩展 |
| 3710 | {{< ext "numeral" >}} | {{< ext "numeral" "numeral" >}} | 1.3 | 数值类型扩展 |
| 3720 | {{< ext "pg_rational" >}} | {{< ext "pg_rational" "pg_rational" >}} | 0.0.2 | 使用BIGINT表示的有理数数据类型 |
| 3730 | {{< ext "uint" >}} | {{< ext "uint" "pguint" >}} | 1.20231206 | 无符号整型数据类型 |
| 3740 | {{< ext "uint128" >}} | {{< ext "uint128" "pg_uint128" >}} | 1.1.0 | 原生128位无符号整型数据类型 |
| 3750 | {{< ext "hashtypes" >}} | {{< ext "hashtypes" "hashtypes" >}} | 0.1.5 | 包括SHA1，MD5在内的多种哈希数据类型 |
| 3820 | {{< ext "ip4r" >}} | {{< ext "ip4r" "ip4r" >}} | 2.4.2 | PostgreSQL 的 IPv4/v6 和 IPv4/v6 范围索引类型 |
| 3830 | {{< ext "pg_duration" >}} | {{< ext "pg_duration" "pg_duration" >}} | 1.0.2 | 用于表示时间段的强化数据类型 |
| 3840 | {{< ext "uri" >}} | {{< ext "uri" "pg_uri" >}} | 1.20151224 | URI数据类型 |
| 3850 | {{< ext "emailaddr" >}} | {{< ext "emailaddr" "pgemailaddr" >}} | 0 | Email地址数据类型 |
| 3860 | {{< ext "acl" >}} | {{< ext "acl" "pg_acl" >}} | 1.0.4 | ACL数据类型 |
| 3870 | {{< ext "debversion" >}} | {{< ext "debversion" "debversion" >}} | 1.2.0 | Debian版本号数据类型 |
| 3880 | {{< ext "pg_rrule" >}} | {{< ext "pg_rrule" "pg_rrule" >}} | 0.2.0 | 日历重复规则RRULE数据类型 |
| 3890 | {{< ext "timestamp9" >}} | {{< ext "timestamp9" "timestamp9" >}} | 1.4.0 | 纳秒分辨率时间戳 |
| 3920 | {{< ext "chkpass" >}} | {{< ext "chkpass" "chkpass" >}} | 1.0 | 数据类型：自动加密的密码 |
| 3930 | {{< ext "isn" >}} | {{< ext "isn" "isn" >}} | 1.2 | 用于国际产品编号标准的数据类型 |
| 3940 | {{< ext "seg" >}} | {{< ext "seg" "seg" >}} | 1.4 | 表示线段或浮点间隔的数据类型 |
| 3950 | {{< ext "cube" >}} | {{< ext "cube" "cube" >}} | 1.5 | 用于存储多维立方体的数据类型 |
| 3960 | {{< ext "ltree" >}} | {{< ext "ltree" "ltree" >}} | 1.3 | 用于表示分层树状结构的数据类型 |
| 3970 | {{< ext "hstore" >}} | {{< ext "hstore" "hstore" >}} | 1.8 | 用于存储（键，值）对集合的数据类型 |
| 3980 | {{< ext "citext" >}} | {{< ext "citext" "citext" >}} | 1.6 | 提供大小写不敏感的字符串类型 |
| 3990 | {{< ext "xml2" >}} | {{< ext "xml2" "xml2" >}} | 1.1 | XPath 查询和 XSLT |

## UTIL

实用功能扩展：HTTP请求，GZIP压缩，JWT处理，邮件客户端，正则，字符编码，编码解码，加密解密等实用功能

| ID | 扩展 | 包 | 版本 | 描述 |
|:---:|:---|:---|:---|:---|
| 4010 | {{< ext "gzip" >}} | {{< ext "gzip" "pg_gzip" >}} | 1.0.1 | 使用SQL执行Gzip压缩与解压缩 |
| 4020 | {{< ext "bzip" >}} | {{< ext "bzip" "pg_bzip" >}} | 1.0.0 | BZIP压缩解压缩函数包 |
| 4030 | {{< ext "zstd" >}} | {{< ext "zstd" "pg_zstd" >}} | 1.1.2 | ZSTD压缩解压缩函数包 |
| 4070 | {{< ext "http" >}} | {{< ext "http" "pg_http" >}} | 1.7.0 | HTTP客户端，允许在数据库内收发HTTP请求 (supabase) |
| 4080 | {{< ext "pg_net" >}} | {{< ext "pg_net" "pg_net" >}} | 0.9.2 | 用 SQL 进行异步非阻塞HTTP/HTTPS 请求的扩展 (supabase) |
| 4090 | {{< ext "pg_curl" >}} | {{< ext "pg_curl" "pg_curl" >}} | 2.4.5 | 封装CURL，执行各种用URL传输数据的操作 |
| 4150 | {{< ext "pgjq" >}} | {{< ext "pgjq" "pgjq" >}} | 0.1.0 | 在Postgres中使用jq查询JSON |
| 4160 | {{< ext "pgjwt" >}} | {{< ext "pgjwt" "pgjwt" >}} | 0.2.0 | JSON Web Token API 的PG实现 (supabase) |
| 4170 | {{< ext "pg_smtp_client" >}} | {{< ext "pg_smtp_client" "pg_smtp_client" >}} | 0.2.0 | 使用SMTP从PostgreSQL内发送邮件的客户端扩展 |
| 4180 | {{< ext "pg_html5_email_address" >}} | {{< ext "pg_html5_email_address" "pg_html5_email_address" >}} | 1.2.3 | 验证Email是否符合HTML5规范的扩展 |
| 4190 | {{< ext "url_encode" >}} | {{< ext "url_encode" "url_encode" >}} | 1.2.5 | 提供URL编码解码函数 |
| 4200 | {{< ext "pgsql_tweaks" >}} | {{< ext "pgsql_tweaks" "pgsql_tweaks" >}} | 0.11.3 | 一些日常会用到的便利函数与视图 |
| 4220 | {{< ext "pg_extra_time" >}} | {{< ext "pg_extra_time" "pg_extra_time" >}} | 2.0.0 | 一些关于日期与时间的扩展函数 |
| 4230 | {{< ext "pgpcre" >}} | {{< ext "pgpcre" "pgpcre" >}} | 1 | PCRE/Perl风格的正则表达式支持 |
| 4240 | {{< ext "icu_ext" >}} | {{< ext "icu_ext" "icu_ext" >}} | 1.10.0 | 访问ICU库提供的函数 |
| 4250 | {{< ext "pgqr" >}} | {{< ext "pgqr" "pgqr" >}} | 1.0 | 从数据库中直接生成QR二维码 |
| 4260 | {{< ext "pg_protobuf" >}} | {{< ext "pg_protobuf" "pg_protobuf" >}} | 1.0 | 提供Protobuf函数支持 |
| 4270 | {{< ext "envvar" >}} | {{< ext "envvar" "envvar" >}} | 1.0.1 | 获取环境变量的函数 |
| 4280 | {{< ext "floatfile" >}} | {{< ext "floatfile" "floatfile" >}} | 1.3.1 | 将浮点数组存储到文件中而不是堆表中 |
| 4290 | {{< ext "pg_render" >}} | {{< ext "pg_render" "pg_render" >}} | 0.1.2 | 使用SQL渲染HTML页面 |
| 4300 | {{< ext "pg_readme" >}} | {{< ext "pg_readme" "pg_readme" >}} | 0.7.0 | 为模式与扩展生成Markdown文档 |
| 4301 | {{< ext "pg_readme_test_extension" >}} | {{< ext "pg_readme" "pg_readme" >}} | 0.7.0 | 为模式与扩展生成Markdown文档 |
| 4310 | {{< ext "ddl_historization" >}} | {{< ext "ddl_historization" "ddl_historization" >}} | 0.0.7 | 用SQL将所有DDL变更写入到数据库表中 |
| 4320 | {{< ext "data_historization" >}} | {{< ext "data_historization" "data_historization" >}} | 1.1.0 | 用SQL将数据变更历史保存到分区表中 |
| 4330 | {{< ext "schedoc" >}} | {{< ext "schedoc" "pg_schedoc" >}} | 0.0.1 | 在Django与DBT之间通过注释文档交换元数据 |
| 4400 | {{< ext "hashlib" >}} | {{< ext "hashlib" "pg_hashlib" >}} | 1.1 | 稳定哈希函数包 |
| 4430 | {{< ext "xxhash" >}} | {{< ext "xxhash" "pg_xxhash" >}} | 0.0.1 | xxhash哈希函数包 |
| 4440 | {{< ext "shacrypt" >}} | {{< ext "shacrypt" "shacrypt" >}} | 1.1 | 实现SHA256-CRYPT与SHA512-CRYPT密钥加密算法 |
| 4450 | {{< ext "cryptint" >}} | {{< ext "cryptint" "cryptint" >}} | 1.0.0 | 加密INT与BIGINT类型 |
| 4460 | {{< ext "pguecc" >}} | {{< ext "pguecc" "pg_ecdsa" >}} | 1.0 | PostgreSQL的uECC绑定，椭圆曲线加解密函数包 |
| 4470 | {{< ext "sparql" >}} | {{< ext "sparql" "pgsparql" >}} | 1.0 | 使用SQL查询SPARQL数据源 |

## FUNC

标识聚合函数：ID生成器，各类聚合函数，摘要函数，数组处理函数，数学函数，统计量，伪随机，等等

| ID | 扩展 | 包 | 版本 | 描述 |
|:---:|:---|:---|:---|:---|
| 4500 | {{< ext "pg_idkit" >}} | {{< ext "pg_idkit" "pg_idkit" >}} | 0.3.1 | 生成各式各样的唯一标识符：UUIDv6, ULID, KSUID |
| 4510 | {{< ext "pgx_ulid" >}} | {{< ext "pgx_ulid" "pgx_ulid" >}} | 0.2.0 | ULID数据类型与函数 |
| 4540 | {{< ext "pg_uuidv7" >}} | {{< ext "pg_uuidv7" "pg_uuidv7" >}} | 1.6.0 | UUIDv7 支持 |
| 4550 | {{< ext "permuteseq" >}} | {{< ext "permuteseq" "permuteseq" >}} | 1.2.2 | 伪随机数ID置换生成器 |
| 4560 | {{< ext "pg_hashids" >}} | {{< ext "pg_hashids" "pg_hashids" >}} | 1.3 | 加盐将整型ID转为短字符串ID |
| 4570 | {{< ext "sequential_uuids" >}} | {{< ext "sequential_uuids" "sequential_uuids" >}} | 1.0.3 | 生成连续生成的UUID |
| 4600 | {{< ext "topn" >}} | {{< ext "topn" "topn" >}} | 2.7.0 | top-n JSONB 的类型 |
| 4610 | {{< ext "quantile" >}} | {{< ext "quantile" "quantile" >}} | 1.1.8 | Quantile聚合函数 |
| 4620 | {{< ext "lower_quantile" >}} | {{< ext "lower_quantile" "lower_quantile" >}} | 1.0.3 | Lower Quantile 聚合函数 |
| 4630 | {{< ext "count_distinct" >}} | {{< ext "count_distinct" "count_distinct" >}} | 3.0.2 | COUNT(DISTINCT …) 聚合的替代方案 |
| 4640 | {{< ext "omnisketch" >}} | {{< ext "omnisketch" "omnisketch" >}} | 1.0.2 | 实现OmniSketch数据结构，实现近似摘要聚合 |
| 4650 | {{< ext "ddsketch" >}} | {{< ext "ddsketch" "ddsketch" >}} | 1.0.1 | 实现DDSketch数据结构，实现在线的Quantile聚合 |
| 4660 | {{< ext "vasco" >}} | {{< ext "vasco" "vasco" >}} | 0.1.0 | 使用MIC发现数据中隐含的关联 |
| 4670 | {{< ext "xicor" >}} | {{< ext "xicor" "pgxicor" >}} | 0.1.0 | 在PG中计算XI相关系数 |
| 4700 | {{< ext "tdigest" >}} | {{< ext "tdigest" "tdigest" >}} | 1.4.3 | tdigest 聚合函数 |
| 4710 | {{< ext "first_last_agg" >}} | {{< ext "first_last_agg" "first_last_agg" >}} | 0.1.4 | first() 与 last() 聚合函数 |
| 4720 | {{< ext "extra_window_functions" >}} | {{< ext "extra_window_functions" "extra_window_functions" >}} | 1.0 | 额外的窗口函数 |
| 4730 | {{< ext "floatvec" >}} | {{< ext "floatvec" "floatvec" >}} | 1.1.1 | 数组类型数学运算扩展 |
| 4740 | {{< ext "aggs_for_vecs" >}} | {{< ext "aggs_for_vecs" "aggs_for_vecs" >}} | 1.4.0 | 针对数组类型的聚合函数集合扩展 |
| 4750 | {{< ext "aggs_for_arrays" >}} | {{< ext "aggs_for_arrays" "aggs_for_arrays" >}} | 1.3.3 | 计算数组聚合统计值的函数包 |
| 4760 | {{< ext "arraymath" >}} | {{< ext "arraymath" "pg_arraymath" >}} | 1.1 | 数组逐元素数学运算符包 |
| 4770 | {{< ext "pg_math" >}} | {{< ext "pg_math" "pg_math" >}} | 1.0 | 使用GSL库的数学统计函数 |
| 4780 | {{< ext "random" >}} | {{< ext "random" "pg_random" >}} | 2.0.0 | 随机数生成器 |
| 4800 | {{< ext "base36" >}} | {{< ext "base36" "pg_base36" >}} | 1.0.0 | Base36编码解码扩展 |
| 4810 | {{< ext "base62" >}} | {{< ext "base62" "pg_base62" >}} | 0.0.1 | Base62编码解码扩展 |
| 4830 | {{< ext "pg_base58" >}} | {{< ext "pg_base58" "pg_base58" >}} | 0.0.1 | Base58 编码/解码函数 |
| 4840 | {{< ext "financial" >}} | {{< ext "financial" "pg_financial" >}} | 1.0.1 | 金融领域聚合函数 |
| 4850 | {{< ext "convert" >}} | {{< ext "convert" "pg_convert" >}} | 0.0.4 | 用于空间里程等的公英制转换函数 |
| 4880 | {{< ext "refint" >}} | {{< ext "refint" "refint" >}} | 1.0 | 实现引用完整性的函数 |
| 4881 | {{< ext "autoinc" >}} | {{< ext "autoinc" "autoinc" >}} | 1.0 | 用于自动递增字段的函数 |
| 4882 | {{< ext "insert_username" >}} | {{< ext "insert_username" "insert_username" >}} | 1.0 | 用于跟踪谁更改了表的函数 |
| 4883 | {{< ext "moddatetime" >}} | {{< ext "moddatetime" "moddatetime" >}} | 1.0 | 跟踪最后修改时间 |
| 4890 | {{< ext "tsm_system_time" >}} | {{< ext "tsm_system_time" "tsm_system_time" >}} | 1.0 | 接受毫秒数限制的 TABLESAMPLE 方法 |
| 4900 | {{< ext "dict_xsyn" >}} | {{< ext "dict_xsyn" "dict_xsyn" >}} | 1.0 | 用于扩展同义词处理的文本搜索字典模板 |
| 4910 | {{< ext "tsm_system_rows" >}} | {{< ext "tsm_system_rows" "tsm_system_rows" >}} | 1.0 | 接受行数限制的 TABLESAMPLE 方法 |
| 4920 | {{< ext "tcn" >}} | {{< ext "tcn" "tcn" >}} | 1.0 | 用触发器通知变更 |
| 4930 | {{< ext "uuid-ossp" >}} | {{< ext "uuid-ossp" "uuid-ossp" >}} | 1.1 | 生成通用唯一标识符（UUIDs） |
| 4940 | {{< ext "btree_gist" >}} | {{< ext "btree_gist" "btree_gist" >}} | 1.7 | 用GiST索引常见数据类型 |
| 4950 | {{< ext "btree_gin" >}} | {{< ext "btree_gin" "btree_gin" >}} | 1.3 | 用GIN索引常见数据类型 |
| 4960 | {{< ext "intarray" >}} | {{< ext "intarray" "intarray" >}} | 1.5 | 1维整数数组的额外函数、运算符和索引支持 |
| 4970 | {{< ext "intagg" >}} | {{< ext "intagg" "intagg" >}} | 1.1 | 整数聚合器和枚举器（过时） |
| 4980 | {{< ext "dict_int" >}} | {{< ext "dict_int" "dict_int" >}} | 1.0 | 用于整数的文本搜索字典模板 |
| 4990 | {{< ext "unaccent" >}} | {{< ext "unaccent" "unaccent" >}} | 1.1 | 删除重音的文本搜索字典 |

## ADMIN

管理工具扩展：膨胀治理，脏读，检视缓冲区，数据目录，校验和，腐败检查，优先级管理，权限管理，语句准备，限制批量更新等

| ID | 扩展 | 包 | 版本 | 描述 |
|:---:|:---|:---|:---|:---|
| 5010 | {{< ext "pg_repack" >}} | {{< ext "pg_repack" "pg_repack" >}} | 1.5.2 | 在线垃圾清理与表膨胀治理 |
| 5020 | {{< ext "pg_rewrite" >}} | {{< ext "pg_rewrite" "pg_rewrite" >}} | 2.0.0 | 在线重写整表，不阻塞读写 |
| 5040 | {{< ext "pg_squeeze" >}} | {{< ext "pg_squeeze" "pg_squeeze" >}} | 1.9.0 | 从关系中删除未使用空间 |
| 5050 | {{< ext "pg_dirtyread" >}} | {{< ext "pg_dirtyread" "pg_dirtyread" >}} | 2.7 | 从表中读取尚未垃圾回收的行 |
| 5060 | {{< ext "pgfincore" >}} | {{< ext "pgfincore" "pgfincore" >}} | 1.3.1 | 检查和管理操作系统缓冲区缓存 |
| 5070 | {{< ext "pg_cooldown" >}} | {{< ext "pg_cooldown" "pg_cooldown" >}} | 0.1 | 从缓冲区中移除特定关系的页面 |
| 5080 | {{< ext "ddlx" >}} | {{< ext "ddlx" "pg_ddlx" >}} | 0.30 | 提取数据库对象的DDL |
| 5090 | {{< ext "prioritize" >}} | {{< ext "prioritize" "pg_prioritize" >}} | 1.0.4 | 获取和设置 PostgreSQL 后端的优先级 |
| 5110 | {{< ext "pg_checksums" >}} | {{< ext "pg_checksums" "pg_checksums" >}} | 1.2 | 在离线模式下激活/启用/禁用数据库集群的校验和功能 |
| 5120 | {{< ext "pg_readonly" >}} | {{< ext "pg_readonly" "pg_readonly" >}} | 1.0.3 | 将集群设置为只读 |
| 5130 | {{< ext "pgdd" >}} | {{< ext "pgdd" "pgdd" >}} | 0.6.0 | 提供通过标准SQL查询数据库目录集簇的能力 |
| 5140 | {{< ext "pg_permissions" >}} | {{< ext "pg_permissions" "pg_permissions" >}} | 1.4 | 查看对象权限并将其与期望状态进行比较 |
| 5150 | {{< ext "pgautofailover" >}} | {{< ext "pgautofailover" "pgautofailover" >}} | 2.2 | PG 自动故障迁移 |
| 5160 | {{< ext "pg_catcheck" >}} | {{< ext "pg_catcheck" "pg_catcheck" >}} | 1.6.0 | 用于诊断系统目录是否损坏的工具 |
| 5170 | {{< ext "pre_prepare" >}} | {{< ext "pre_prepare" "preprepare" >}} | 0.9 | 在服务端预先准备好PreparedStatement备用 |
| 5180 | {{< ext "pg_upless" >}} | {{< ext "pg_upless" "pg_upless" >}} | 0.0.3 | 检测表上的无用UPDATE |
| 5190 | {{< ext "pgcozy" >}} | {{< ext "pgcozy" "pgcozy" >}} | 1.0 | 根据先前的pg_buffercache快照预热内存缓冲区 |
| 5200 | {{< ext "pg_orphaned" >}} | {{< ext "pg_orphaned" "pg_orphaned" >}} | 1.0 | 处理孤儿文件的扩展插件 |
| 5210 | {{< ext "pg_crash" >}} | {{< ext "pg_crash" "pg_crash" >}} | 1.0 | 向数据库进程随机发送信号模拟故障 |
| 5220 | {{< ext "pg_cheat_funcs" >}} | {{< ext "pg_cheat_funcs" "pg_cheat_funcs" >}} | 1.0 | 一些超级实用的作弊函数 |
| 5230 | {{< ext "fio" >}} | {{< ext "fio" "pg_fio" >}} | 1.0 | PostgreSQL文件IO函数包 |
| 5810 | {{< ext "pg_savior" >}} | {{< ext "pg_savior" "pg_savior" >}} | 0.0.1 | 阻止不带条件的全表更新以避免意外事故 |
| 5820 | {{< ext "safeupdate" >}} | {{< ext "safeupdate" "safeupdate" >}} | 1.5 | 强制在 UPDATE 和 DELETE 时提供 Where 条件 |
| 5830 | {{< ext "pg_drop_events" >}} | {{< ext "pg_drop_events" "pg_drop_events" >}} | 0.1.0 | 记录删表删列删视图的事务号，辅助PITR确定时间点 |
| 5840 | {{< ext "table_log" >}} | {{< ext "table_log" "table_log" >}} | 0.6.4 | 记录某张表的修改日志并做表/行级时间点恢复 |
| 5880 | {{< ext "pgagent" >}} | {{< ext "pgagent" "pgagent" >}} | 4.2.3 | PostgreSQL任务调度工具，与PGADMIN配合使用 |
| 5890 | {{< ext "pg_prewarm" >}} | {{< ext "pg_prewarm" "pg_prewarm" >}} | 1.2 | 预热关系数据 |
| 5900 | {{< ext "pgpool_adm" >}} | {{< ext "pgpool_adm" "pgpool" >}} | 4.6.3 | PGPool 管理函数 |
| 5910 | {{< ext "pgpool_recovery" >}} | {{< ext "pgpool_adm" "pgpool" >}} | 4.6.3 | PGPool辅助扩展，从v4.3提供的恢复函数 |
| 5920 | {{< ext "pgpool_regclass" >}} | {{< ext "pgpool_adm" "pgpool" >}} | 4.6.3 | PGPool辅助扩展，RegClass替代 |
| 5930 | {{< ext "lo" >}} | {{< ext "lo" "lo" >}} | 1.1 | 大对象维护 |
| 5940 | {{< ext "basic_archive" >}} | {{< ext "basic_archive" "basic_archive" >}} | - | 归档模块样例 |
| 5950 | {{< ext "basebackup_to_shell" >}} | {{< ext "basebackup_to_shell" "basebackup_to_shell" >}} | - | 添加一种备份到Shell终端到基础备份方式 |
| 5960 | {{< ext "old_snapshot" >}} | {{< ext "old_snapshot" "old_snapshot" >}} | 1.0 | 支持 old_snapshot_threshold 的实用程序 |
| 5970 | {{< ext "adminpack" >}} | {{< ext "adminpack" "adminpack" >}} | 2.1 | PostgreSQL 管理函数集合 |
| 5980 | {{< ext "amcheck" >}} | {{< ext "amcheck" "amcheck" >}} | 1.4 | 校验关系完整性 |
| 5990 | {{< ext "pg_surgery" >}} | {{< ext "pg_surgery" "pg_surgery" >}} | 1.0 | 对损坏的关系进行手术 |

## STAT

监控统计扩展：AWR报告，可观测性指标，显示执行计划，查询统计信息，内存使用，配置变更，等待事件采样，慢查询日志，等等

| ID | 扩展 | 包 | 版本 | 描述 |
|:---:|:---|:---|:---|:---|
| 6000 | {{< ext "pg_profile" >}} | {{< ext "pg_profile" "pg_profile" >}} | 4.10 | PostgreSQL 数据库负载记录与AWR报表工具 |
| 6010 | {{< ext "pg_tracing" >}} | {{< ext "pg_tracing" "pg_tracing" >}} | 0.1.3 | PostgreSQL分布式Tracing |
| 6210 | {{< ext "pg_show_plans" >}} | {{< ext "pg_show_plans" "pg_show_plans" >}} | 2.1.6 | 打印所有当前正在运行查询的执行计划 |
| 6220 | {{< ext "pg_stat_kcache" >}} | {{< ext "pg_stat_kcache" "pg_stat_kcache" >}} | 2.3.0 | 内核统计信息收集 |
| 6230 | {{< ext "pg_stat_monitor" >}} | {{< ext "pg_stat_monitor" "pg_stat_monitor" >}} | 2.2.0 | 提供查询聚合统计、客户端信息、执行计划详细信息和直方图 |
| 6240 | {{< ext "pg_qualstats" >}} | {{< ext "pg_qualstats" "pg_qualstats" >}} | 2.1.2 | 收集有关 quals 的统计信息的扩展 |
| 6250 | {{< ext "pg_store_plans" >}} | {{< ext "pg_store_plans" "pg_store_plans" >}} | 1.8 | 跟踪所有执行的 SQL 语句的计划统计信息 |
| 6260 | {{< ext "pg_track_settings" >}} | {{< ext "pg_track_settings" "pg_track_settings" >}} | 2.1.2 | 跟踪设置更改 |
| 6270 | {{< ext "pg_wait_sampling" >}} | {{< ext "pg_wait_sampling" "pg_wait_sampling" >}} | 1.1.9 | 基于采样的等待事件统计 |
| 6280 | {{< ext "pgsentinel" >}} | {{< ext "pgsentinel" "pgsentinel" >}} | 1.2.0 | 活跃会话历史 |
| 6290 | {{< ext "system_stats" >}} | {{< ext "system_stats" "system_stats" >}} | 3.2 | PostgreSQL 的系统统计函数 |
| 6300 | {{< ext "meta" >}} | {{< ext "meta" "pg_meta" >}} | 0.4.0 | 标准化，更友好的PostgreSQL系统目录视图 |
| 6310 | {{< ext "pgnodemx" >}} | {{< ext "pgnodemx" "pgnodemx" >}} | 1.7 | 使用SQL查询获取操作系统指标 |
| 6320 | {{< ext "pg_proctab" >}} | {{< ext "pgnodemx" "pgnodemx" >}} | 1.7 | 通过SQL接口访问操作系统进程表 |
| 6330 | {{< ext "pg_sqlog" >}} | {{< ext "pg_sqlog" "pg_sqlog" >}} | 1.6 | 提供访问PostgreSQL日志的SQL接口 |
| 6340 | {{< ext "bgw_replstatus" >}} | {{< ext "bgw_replstatus" "bgw_replstatus" >}} | 1.0.8 | 用于汇报本机主从状态的后台工作进程 |
| 6350 | {{< ext "pgmeminfo" >}} | {{< ext "pgmeminfo" "pgmeminfo" >}} | 1.0.0 | 显示内存使用情况 |
| 6360 | {{< ext "toastinfo" >}} | {{< ext "toastinfo" "toastinfo" >}} | 1.5 | 显示TOAST字段的详细信息 |
| 6370 | {{< ext "explain_ui" >}} | {{< ext "explain_ui" "pg_explain_ui" >}} | 0.0.1 | 快速跳转至PEV查阅可视化执行计划 |
| 6380 | {{< ext "pg_relusage" >}} | {{< ext "pg_relusage" "pg_relusage" >}} | 0.0.1 | 打印查询引用的表与列 |
| 6800 | {{< ext "pagevis" >}} | {{< ext "pagevis" "pagevis" >}} | 0.1 | 使用ASCII字符可视化数据库物理页面布局 |
| 6810 | {{< ext "powa" >}} | {{< ext "powa" "powa" >}} | 5.0.1 | PostgreSQL 工作负载分析器-核心 |
| 6880 | {{< ext "pg_overexplain" >}} | {{< ext "pg_overexplain" "pg_overexplain" >}} | 1.0 | 允许 EXPLAIN 转储更多详细 |
| 6890 | {{< ext "pg_logicalinspect" >}} | {{< ext "pg_logicalinspect" "pg_logicalinspect" >}} | 1.0 | 检视逻辑解码组件详情 |
| 6900 | {{< ext "pageinspect" >}} | {{< ext "pageinspect" "pageinspect" >}} | 1.12 | 检查数据库页面二进制内容 |
| 6910 | {{< ext "pgrowlocks" >}} | {{< ext "pgrowlocks" "pgrowlocks" >}} | 1.2 | 显示行级锁信息 |
| 6920 | {{< ext "sslinfo" >}} | {{< ext "sslinfo" "sslinfo" >}} | 1.2 | 关于 SSL 证书的信息 |
| 6930 | {{< ext "pg_buffercache" >}} | {{< ext "pg_buffercache" "pg_buffercache" >}} | 1.5 | 检查共享缓冲区缓存 |
| 6940 | {{< ext "pg_walinspect" >}} | {{< ext "pg_walinspect" "pg_walinspect" >}} | 1.1 | 用于检查 PostgreSQL WAL 日志内容的函数 |
| 6950 | {{< ext "pg_freespacemap" >}} | {{< ext "pg_freespacemap" "pg_freespacemap" >}} | 1.2 | 检查自由空间映射的内容（FSM） |
| 6960 | {{< ext "pg_visibility" >}} | {{< ext "pg_visibility" "pg_visibility" >}} | 1.2 | 检查可见性图（VM）和页面级可见性信息 |
| 6970 | {{< ext "pgstattuple" >}} | {{< ext "pgstattuple" "pgstattuple" >}} | 1.5 | 显示元组级统计信息 |
| 6980 | {{< ext "auto_explain" >}} | {{< ext "auto_explain" "auto_explain" >}} | - | 提供一种自动记录执行计划的手段 |
| 6990 | {{< ext "pg_stat_statements" >}} | {{< ext "pg_stat_statements" "pg_stat_statements" >}} | 1.11 | 跟踪所有执行的 SQL 语句的计划和执行统计信息 |

## SEC

安全功能扩展：强制密码强度，阉割超级用户，密钥管理，商密算法，PII匿名处理，扩展白名单，审计日志，变更追溯，反病毒等等

| ID | 扩展 | 包 | 版本 | 描述 |
|:---:|:---|:---|:---|:---|
| 7000 | {{< ext "passwordcheck_cracklib" >}} | {{< ext "passwordcheck_cracklib" "passwordcheck" >}} | 3.1.0 | 使用cracklib加固PG用户密码 |
| 7010 | {{< ext "supautils" >}} | {{< ext "supautils" "supautils" >}} | 2.10.0 | 用于在云环境中确保数据库集群的安全 |
| 7020 | {{< ext "pgsodium" >}} | {{< ext "pgsodium" "pgsodium" >}} | 3.1.9 | 表数据加密存储 TDE |
| 7030 | {{< ext "supabase_vault" >}} | {{< ext "supabase_vault" "pg_vault" >}} | 0.3.1 | 在 Vault 中存储加密凭证的扩展 (supabase) |
| 7040 | {{< ext "pg_session_jwt" >}} | {{< ext "pg_session_jwt" "pg_session_jwt" >}} | 0.3.1 | 使用JWT进行会话认证 |
| 7050 | {{< ext "anon" >}} | {{< ext "anon" "pg_anon" >}} | 2.3.0 | 数据匿名化处理工具 |
| 7060 | {{< ext "pg_tde" >}} | {{< ext "pg_tde" "pg_tde" >}} | 1.0 | Percona加密存储引擎 |
| 7070 | {{< ext "pgsmcrypto" >}} | {{< ext "pgsmcrypto" "pgsmcrypto" >}} | 0.1.0 | 为PostgreSQL提供商密算法支持：SM2,SM3,SM4 |
| 7080 | {{< ext "pgaudit" >}} | {{< ext "pgaudit" "pgaudit" >}} | 17.1 | 提供审计功能 |
| 7090 | {{< ext "pgauditlogtofile" >}} | {{< ext "pgauditlogtofile" "pgauditlogtofile" >}} | 1.7.1 | pgAudit 子扩展，将审计日志写入单独的文件中 |
| 7100 | {{< ext "pg_auth_mon" >}} | {{< ext "pg_auth_mon" "pg_auth_mon" >}} | 3.0 | 监控每个用户的连接尝试 |
| 7110 | {{< ext "credcheck" >}} | {{< ext "credcheck" "credcheck" >}} | 3.0 | 明文凭证检查器 |
| 7120 | {{< ext "pgcryptokey" >}} | {{< ext "pgcryptokey" "pgcryptokey" >}} | 0.85 | PG密钥管理 |
| 7130 | {{< ext "pg_jobmon" >}} | {{< ext "pg_jobmon" "pg_jobmon" >}} | 1.4.1 | 记录和监控函数 |
| 7140 | {{< ext "logerrors" >}} | {{< ext "logerrors" "logerrors" >}} | 2.1.3 | 用于收集日志文件中消息统计信息的函数 |
| 7150 | {{< ext "login_hook" >}} | {{< ext "login_hook" "login_hook" >}} | 1.7 | 在用户登陆时执行login_hook.login()函数 |
| 7160 | {{< ext "set_user" >}} | {{< ext "set_user" "set_user" >}} | 4.1.0 | 增加了日志记录的 SET ROLE |
| 7170 | {{< ext "pg_snakeoil" >}} | {{< ext "pg_snakeoil" "pg_snakeoil" >}} | 1.4 | PostgreSQL动态链接库反病毒功能 |
| 7180 | {{< ext "pgextwlist" >}} | {{< ext "pgextwlist" "pgextwlist" >}} | 1.19 | PostgreSQL扩展白名单功能 |
| 7190 | {{< ext "pg_auditor" >}} | {{< ext "pg_auditor" "pg_auditor" >}} | 0.2 | 审计数据变更并提供闪回能力 |
| 7200 | {{< ext "sslutils" >}} | {{< ext "sslutils" "sslutils" >}} | 1.4 | 使用SQL管理SSL证书 |
| 7210 | {{< ext "noset" >}} | {{< ext "noset" "pg_noset" >}} | 0.3.0 | 阻止非超级用户使用SET/RESET设置变量 |
| 7960 | {{< ext "sepgsql" >}} | {{< ext "sepgsql" "sepgsql" >}} | - | 基于SELinux标签的强制访问控制 |
| 7970 | {{< ext "auth_delay" >}} | {{< ext "auth_delay" "auth_delay" >}} | - | 在返回认证失败前暂停一会，避免爆破 |
| 7980 | {{< ext "pgcrypto" >}} | {{< ext "pgcrypto" "pgcrypto" >}} | 1.3 | 实用加解密函数 |
| 7990 | {{< ext "passwordcheck" >}} | {{< ext "passwordcheck_cracklib" "passwordcheck" >}} | - | 用于强制拒绝修改弱密码的扩展 |

## FDW

外部数据源包装器：FDW开发框架 Wrappers，Multicorn，访问外部的 Mongo，MySQL，SQLite，HDFS，MSSQL，Oracle，DB2，……

| ID | 扩展 | 包 | 版本 | 描述 |
|:---:|:---|:---|:---|:---|
| 8500 | {{< ext "wrappers" >}} | {{< ext "wrappers" "wrappers" >}} | 0.5.4 | Supabase提供的外部数据源包装器捆绑包 |
| 8510 | {{< ext "multicorn" >}} | {{< ext "multicorn" "multicorn" >}} | 3.0 | 用Python编写自定义的外部数据源包装器 |
| 8520 | {{< ext "odbc_fdw" >}} | {{< ext "odbc_fdw" "odbc_fdw" >}} | 0.5.1 | 访问ODBC可访问的任何外部数据源 |
| 8530 | {{< ext "jdbc_fdw" >}} | {{< ext "jdbc_fdw" "jdbc_fdw" >}} | 1.2 | 访问JDBC可访问的任何外部数据源 |
| 8540 | {{< ext "pgspider_ext" >}} | {{< ext "pgspider_ext" "pgspider_ext" >}} | 1.3.0 | 使用多种FDW访问远程数据库服务器 |
| 8600 | {{< ext "mysql_fdw" >}} | {{< ext "mysql_fdw" "mysql_fdw" >}} | 2.9.2 | MySQL外部数据包装器 |
| 8610 | {{< ext "oracle_fdw" >}} | {{< ext "oracle_fdw" "oracle_fdw" >}} | 2.8.0 | 提供对Oracle的外部数据源包装器 |
| 8620 | {{< ext "tds_fdw" >}} | {{< ext "tds_fdw" "tds_fdw" >}} | 2.0.4 | TDS 数据库（Sybase/SQL Server）外部数据包装器 |
| 8630 | {{< ext "db2_fdw" >}} | {{< ext "db2_fdw" "db2_fdw" >}} | 7.0.0 | 提供对DB2的外部数据源包装器 |
| 8640 | {{< ext "sqlite_fdw" >}} | {{< ext "sqlite_fdw" "sqlite_fdw" >}} | 2.5.0 | SQLite 外部数据包装器 |
| 8650 | {{< ext "pgbouncer_fdw" >}} | {{< ext "pgbouncer_fdw" "pgbouncer_fdw" >}} | 1.4.0 | 用SQL查询pgbouncer统计信息，并执行pgbouncer命令 |
| 8700 | {{< ext "mongo_fdw" >}} | {{< ext "mongo_fdw" "mongo_fdw" >}} | 1.1 | MongoDB 外部数据包装器 |
| 8710 | {{< ext "redis_fdw" >}} | {{< ext "redis_fdw" "redis_fdw" >}} | 1.0 | 查询外部Redis数据源 |
| 8720 | {{< ext "redis" >}} | {{< ext "redis" "pg_redis_pubsub" >}} | 0.0.1 | 从PG向Redis发送Pub/Sub消息 |
| 8730 | {{< ext "kafka_fdw" >}} | {{< ext "kafka_fdw" "kafka_fdw" >}} | 0.0.3 | Kafka外部数据源包装器 |
| 8740 | {{< ext "hdfs_fdw" >}} | {{< ext "hdfs_fdw" "hdfs_fdw" >}} | 2.3.2 | hdfs 外部数据包装器 |
| 8750 | {{< ext "firebird_fdw" >}} | {{< ext "firebird_fdw" "firebird_fdw" >}} | 1.4.0 | Firebird外部数据源包装器 |
| 8800 | {{< ext "aws_s3" >}} | {{< ext "aws_s3" "aws_s3" >}} | 0.0.1 | 从S3导入导出数据的外部数据源包装器 |
| 8810 | {{< ext "log_fdw" >}} | {{< ext "log_fdw" "log_fdw" >}} | 1.4 | 访问PostgreSQL日志文件的FDW |
| 8970 | {{< ext "dblink" >}} | {{< ext "dblink" "dblink" >}} | 1.2 | 从数据库内连接到其他 PostgreSQL 数据库 |
| 8980 | {{< ext "file_fdw" >}} | {{< ext "file_fdw" "file_fdw" >}} | 1.0 | 访问外部文件的外部数据包装器 |
| 8990 | {{< ext "postgres_fdw" >}} | {{< ext "postgres_fdw" "postgres_fdw" >}} | 1.1 | 用于远程 PostgreSQL 服务器的外部数据包装器 |

## SIM

数据库兼容扩展：仿真其他 DBMS 的行为：MySQL，Memcache，Mongo，Oracle，Babelfish for Microsoft SQL Server……

| ID | 扩展 | 包 | 版本 | 描述 |
|:---:|:---|:---|:---|:---|
| 9000 | {{< ext "documentdb" >}} | {{< ext "documentdb" "documentdb" >}} | 0.106 | 微软DocumentDB的API层 |
| 9010 | {{< ext "documentdb_core" >}} | {{< ext "documentdb" "documentdb" >}} | 0.106 | 微软DocumentDB的核心API层实现 |
| 9020 | {{< ext "documentdb_distributed" >}} | {{< ext "documentdb" "documentdb" >}} | 0.106 | DocumentDB多节点模式的API层 |
| 9100 | {{< ext "orafce" >}} | {{< ext "orafce" "orafce" >}} | 4.14.4 | 模拟 Oracle RDBMS 的一部分函数和包的函数和运算符 |
| 9110 | {{< ext "pgtt" >}} | {{< ext "pgtt" "pgtt" >}} | 4.4 | 类似Oracle的全局临时表功能 |
| 9120 | {{< ext "session_variable" >}} | {{< ext "session_variable" "session_variable" >}} | 3.4 | Oracle兼容的会话变量/常量操作函数 |
| 9130 | {{< ext "pg_statement_rollback" >}} | {{< ext "pg_statement_rollback" "pg_statement_rollback" >}} | 1.4 | 在服务端提供类似Oracle/DB2的语句级回滚能力 |
| 9240 | {{< ext "pg_dbms_metadata" >}} | {{< ext "pg_dbms_metadata" "pg_dbms_metadata" >}} | 1.0.0 | 添加 Oracle DBMS_METADATA 兼容性支持的扩展 |
| 9250 | {{< ext "pg_dbms_lock" >}} | {{< ext "pg_dbms_lock" "pg_dbms_lock" >}} | 1.0 | 为PG添加对 Oracle DBMS_LOCK 的完整兼容性支持 |
| 9260 | {{< ext "pg_dbms_job" >}} | {{< ext "pg_dbms_job" "pg_dbms_job" >}} | 1.5 | 添加 Oracle DBMS_JOB 兼容性支持的扩展 |
| 9300 | {{< ext "babelfishpg_common" >}} | {{< ext "babelfishpg_common" "babelfishpg_common" >}} | 3.3.3 | SQL Server 数据类型兼容扩展 |
| 9310 | {{< ext "babelfishpg_tsql" >}} | {{< ext "babelfishpg_tsql" "babelfishpg_tsql" >}} | 3.3.1 | SQL Server SQL语法兼容性扩展 |
| 9320 | {{< ext "babelfishpg_tds" >}} | {{< ext "babelfishpg_tds" "babelfishpg_tds" >}} | 1.0.0 | SQL Server TDS线缆协议兼容扩展 |
| 9330 | {{< ext "babelfishpg_money" >}} | {{< ext "babelfishpg_money" "babelfishpg_money" >}} | 1.1.0 | SQL Server 货币数据类型兼容扩展 |
| 9400 | {{< ext "spat" >}} | {{< ext "spat" "spat" >}} | 0.1.0a4 | 在PG中嵌入Redis风格的内存数据库 |
| 9410 | {{< ext "pgmemcache" >}} | {{< ext "pgmemcache" "pgmemcache" >}} | 2.3.0 | 为PG提供memcached兼容接口 |

## ETL

数据复制扩展：逻辑复制，逻辑解码，DDL复制，JSON/BSON/Protobuf 变更抽取，数据迁移，数据导入，数据比对等

| ID | 扩展 | 包 | 版本 | 描述 |
|:---:|:---|:---|:---|:---|
| 9500 | {{< ext "pglogical" >}} | {{< ext "pglogical" "pglogical" >}} | 2.4.5 | PostgreSQL逻辑复制：三方扩展实现 |
| 9501 | {{< ext "pglogical_origin" >}} | {{< ext "pglogical" "pglogical" >}} | 2.4.5 | 用于从 Postgres 9.4 升级时的兼容性虚拟扩展 |
| 9510 | {{< ext "pglogical_ticker" >}} | {{< ext "pglogical_ticker" "pglogical_ticker" >}} | 1.4.1 | pglogical复制延迟以秒计的精确视图 |
| 9520 | {{< ext "pgl_ddl_deploy" >}} | {{< ext "pgl_ddl_deploy" "pgl_ddl_deploy" >}} | 2.2.1 | 使用 pglogical 执行自动 DDL 部署 |
| 9530 | {{< ext "pg_failover_slots" >}} | {{< ext "pg_failover_slots" "pg_failover_slots" >}} | 1.1.0 | 在Failover过程中保留复制槽 |
| 9540 | {{< ext "db_migrator" >}} | {{< ext "db_migrator" "db_migrator" >}} | 1.0.0 | 使用FDW从其他DBMS迁移到PostgreSQL |
| 9550 | {{< ext "pgactive" >}} | {{< ext "pgactive" "pgactive" >}} | 2.1.6 | PostgreSQL多主逻辑复制 |
| 9630 | {{< ext "wal2json" >}} | {{< ext "wal2json" "wal2json" >}} | 2.6 | 用逻辑解码捕获 JSON 格式的 CDC 变更 |
| 9640 | {{< ext "wal2mongo" >}} | {{< ext "wal2mongo" "wal2mongo" >}} | 1.0.7 | 使用逻辑解码捕获MongoDB JSON格式的CDC变更 |
| 9650 | {{< ext "decoderbufs" >}} | {{< ext "decoderbufs" "decoderbufs" >}} | 3.2.0 | 将WAL逻辑解码为ProtocolBuffer协议的消息 |
| 9660 | {{< ext "decoder_raw" >}} | {{< ext "decoder_raw" "decoder_raw" >}} | 1.0 | 逻辑复制解码输出插件：RAW SQL格式 |
| 9700 | {{< ext "mimeo" >}} | {{< ext "mimeo" "mimeo" >}} | 1.5.1 | 在PostgreSQL实例间进行表级复制 |
| 9710 | {{< ext "repmgr" >}} | {{< ext "repmgr" "repmgr" >}} | 5.5.0 | PostgreSQL复制管理组件 |
| 9820 | {{< ext "pg_fact_loader" >}} | {{< ext "pg_fact_loader" "pg_fact_loader" >}} | 2.0.1 | 在 Postgres 中构建事实表 |
| 9830 | {{< ext "pg_bulkload" >}} | {{< ext "pg_bulkload" "pg_bulkload" >}} | 3.1.22 | 向 PostgreSQL 中高速加载数据 |
| 9970 | {{< ext "test_decoding" >}} | {{< ext "test_decoding" "test_decoding" >}} | - | 基于SQL的WAL逻辑解码样例 |
| 9980 | {{< ext "pgoutput" >}} | {{< ext "pgoutput" "pgoutput" >}} | - | PG内置的逻辑解码输出插件 |
