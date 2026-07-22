## Usage

Sources:

- [Official README](https://github.com/vibhorkum/efm_extension/blob/ecb305cc185b05eaf0e766d6a9c5d23b3072e2fb/README.md)
- [Extension control file](https://github.com/vibhorkum/efm_extension/blob/ecb305cc185b05eaf0e766d6a9c5d23b3072e2fb/efm_extension.control)
- [Version 1.1 installation SQL](https://github.com/vibhorkum/efm_extension/blob/ecb305cc185b05eaf0e766d6a9c5d23b3072e2fb/efm_extension--1.1.sql)

`efm_extension` exposes EDB Failover Manager status and management commands through SQL. It executes the installed EFM command-line program through a restricted `sudo` boundary, adds structured status, cache, history, and monitoring views, and separates read-only monitoring from superuser-only topology changes.

### Configure and Monitor

Install and configure EFM first, create a narrowly scoped sudoers policy for the PostgreSQL operating-system account, and set the cluster, binary, properties, Java, and sudo paths. The optional cache/history worker requires preloading and a restart.

```conf
shared_preload_libraries = 'efm_extension'
efm.cluster_name = 'efm'
efm.command_path = '/usr/edb/efm-4.9/bin/efm'
efm.properties_location = '/etc/edb/efm-4.9'
efm.java_home = '/usr/lib/jvm/java-11-openjdk'
efm.bgw_enabled = true
```

```sql
CREATE EXTENSION dblink;
CREATE EXTENSION pgcrypto;
CREATE EXTENSION efm_extension;

SELECT efm_extension.grant_access_to_user('monitoring_user');
SELECT efm_extension.efm_cluster_status_json();
SELECT * FROM efm_extension.efm_nodes_details;
SELECT * FROM efm_extension.efm_metrics;
```

`efm_cluster_status()`, `efm_cluster_status_json()`, `efm_get_nodes()`, `efm_list_properties()`, and cache/metrics views form the monitoring surface. `efm_allow_node()`, `efm_disallow_node()`, `efm_set_priority()`, `efm_failover()`, `efm_switchover()`, and `efm_resume_monitoring()` remain management operations for superusers.

### Failover and Host-Security Boundaries

Management calls change an external HA cluster and are not rolled back when the surrounding PostgreSQL transaction aborts. A failover or switchover may disconnect the very session that invoked it. Require explicit operational authorization, serialize changes, confirm candidate health and lag, and verify the resulting EFM and PostgreSQL topology independently.

The database backend can execute configured host commands as the EFM account. Pin absolute paths, permit only exact EFM subcommands in sudoers, protect GUCs and properties files, and do not grant configuration or function access to untrusted roles. Cached status can be stale by `efm.cache_ttl`, while background-history persistence adds local writes and retention needs. Match extension 1.1 to the exact EFM CLI/output version and exercise exit-code mapping, timeouts, stderr, worker restart, and loss of EFM connectivity before production use.
