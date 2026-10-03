## Usage

Sources:

- [Pgpool-II 4.7.3 control](https://github.com/pgpool/pgpool2/blob/V4_7_3/src/sql/pgpool_adm/pgpool_adm.control)
- [SQL API 1.6](https://github.com/pgpool/pgpool2/blob/V4_7_3/src/sql/pgpool_adm/pgpool_adm--1.6.sql)
- [PCP node reference](https://github.com/pgpool/pgpool2/blob/V4_7_3/doc/src/sgml/ref/pgpool_adm_pcp_node_info.sgml)
- [Foreign-server setup](https://github.com/pgpool/pgpool2/blob/V4_7_3/src/sql/pgpool_adm/pgpool_adm.c)

`pgpool_adm` exposes SQL wrappers for Pgpool-II PCP administration. A reachable Pgpool-II PCP service and appropriate PCP credentials are required; installing the extension does not deploy Pgpool-II.

### Install and Inspect a Node

```sql
CREATE EXTENSION pgpool_adm;

SELECT * FROM pcp_node_info(
  node_id => 0, host => 'localhost', port => 9898,
  username => 'pcp_admin', password => 'example-password');
SELECT pcp_node_count(
  host => 'localhost', port => 9898,
  username => 'pcp_admin', password => 'example-password');
```

The actual SQL names are `pcp_node_info`, `pcp_health_check_stats`, `pcp_pool_status`, `pcp_node_count`, `pcp_attach_node`, `pcp_detach_node`, and `pcp_proc_info`. The documentation page names may carry a package prefix; that prefix is not part of these SQL functions.

### Credential and Server References

The direct overload accepts host, port, username and password, with node ID first where required. Other overloads accept a foreign-server reference through the `pcp_server` argument. Configure that server and user mapping according to the source implementation; do not assume a client-side `.pcppass` file supplies credentials to the database backend.

```sql
SELECT * FROM pcp_node_info(node_id => 0, pcp_server => 'pgpool_server');
SELECT pcp_node_count(pcp_server => 'pgpool_server');
```

The usual PCP port is 9898. Credential handling and network access occur on the PostgreSQL server. Literal passwords can enter SQL logs and activity views; prefer the configured server-reference path where suitable.

### Manage Nodes

```sql
SELECT pcp_detach_node(
  node_id => 1, gracefully => true,
  host => 'localhost', port => 9898,
  username => 'pcp_admin', password => 'example-password');
SELECT pcp_attach_node(
  node_id => 1, host => 'localhost', port => 9898,
  username => 'pcp_admin', password => 'example-password');
```

Detaching or attaching nodes changes Pgpool-II routing and can affect application traffic. Restrict function access and PCP credentials to administrators. A database transaction rollback does not undo an external PCP operation.

### Version and Loading

Have a superuser install the extension. It needs no preload or PostgreSQL restart. Pgpool-II package release 4.7.3 ships SQL extension version 1.6; update existing databases using `ALTER EXTENSION pgpool_adm UPDATE` after installing matching files, rather than specifying 4.7.3 as an SQL version.
