## Usage

Sources:

- [Official README](https://github.com/b4oshany/livewire/blob/547e59a951fe5d088d89039a2d078186e73c7862/README.md)
- [Version 0.2.0 SQL implementation](https://github.com/b4oshany/livewire/blob/547e59a951fe5d088d89039a2d078186e73c7862/livewire--0.2.0.sql)
- [Extension control file](https://github.com/b4oshany/livewire/blob/547e59a951fe5d088d89039a2d078186e73c7862/livewire.control)

`livewire` builds and caches a routable shadow network for electrical-distribution GIS data. It combines PostGIS geometry with pgRouting-style graph traversal so engineers can trace upstream, downstream, source, blocked, and end nodes across configured line and device layers.

### Core Workflow

Install `postgis`, `postgis_topology`, and pgRouting support first. Initialize a schema with its SRID, register edge and node participants through JSON configuration, generate the shadow network, then cache source traces.

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION postgis_topology;
CREATE EXTENSION livewire;

SELECT lw_initialise('powerflow', 3448);
SELECT lw_generate('powerflow');
SELECT lw_traceall('powerflow');
```

`lw_addedgeparticipant()` registers linear layers. `lw_addnodeparticipant()` additionally accepts source and blocking predicates. `lw_tracesource()`, `lw_traceupstream()`, and `lw_tracednstream()` operate on the prepared network; `lw_sourcenodes()`, `lw_blocknodes()`, `lw_endnodes()`, and `lw_nodesnearnode()` expose supporting topology checks.

### Data and Maintenance Boundaries

Participant JSON must identify schema, table, primary-key, geometry, label, phase, and phase-map fields. A node participant can also define `sourcequery` and `blockquery`; these are SQL conditions and must come only from trusted configuration. When source data is outside the LiveWire schema, initialization may copy it with table-like semantics, so confirm ownership, privileges, indexes, and refresh behavior.

`lw_generate()` rebuilds the combined network and `lw_traceall()` can be expensive because it traces every source. The documented model treats switching devices as all-phases on or all-phases off even though phase data can be stored. The reviewed upstream targets PostgreSQL 10-era PostGIS/pgRouting; validate function compatibility and run generation on a disposable copy before using it with current production GIS data.
