## Usage

Sources:

- [Official upstream README](https://github.com/dalibo/hackingpg/blob/a16eacb921fb5560a504b61c539d98b8919617f8/journee6/README)
- [Official extension control file (cpg.control)](https://github.com/dalibo/hackingpg/blob/a16eacb921fb5560a504b61c539d98b8919617f8/journee6/cpg.control)
- [Official implementation source](https://github.com/dalibo/hackingpg/blob/a16eacb921fb5560a504b61c539d98b8919617f8/journee6/cpg.c)

`cpg` is a Corosync-backed PostgreSQL background-worker experiment. It exchanges primary membership information, updates `primary_conninfo`, and can ask systemd to promote or stop a node; it is a workshop module rather than a general-purpose HA product.

### Core Workflow

The upstream provisioning example installs the library and configures it as follows:

```ini
shared_preload_libraries = 'cpg'
include_if_exists = 'cpg.auto.conf'
cpg.interval = 10
cpg.service = 'postgresql-16.service'
cpg.is_user_service = false
```

Restart PostgreSQL after adding the library. The worker joins the Corosync group, writes discovered connection state to `cpg.auto.conf`, and uses the configured systemd service during its failover path.

### Important Settings

- `cpg.interval` is the maximum interval, in seconds, between worker wakeups.
- `cpg.service` names the systemd service the worker is allowed to manage.
- `cpg.is_user_service` chooses the systemd user bus instead of the system bus.

### Requirements and Caveats

- The reviewed control file identifies version `0.1`; there is no extension SQL or `CREATE EXTENSION` workflow.
- The source requires Corosync CPG and systemd libraries. The demo provisioning also installs a polkit rule so the PostgreSQL account can manage its service.
- The worker can alter replication configuration and trigger promotion or shutdown. Test fencing, quorum loss, stale membership, systemd permissions, and split-brain behavior before any non-lab use.
- This code comes from a PostgreSQL hacking workshop and should be treated as an educational prototype.
