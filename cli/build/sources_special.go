// Source bundle compatibility migrated from pgsty/pig (Apache-2.0).
// Fallbacks for builds without a local recipe or catalog entry.
// Aligned with pgsty/rpm and pgsty/deb on 2026-11-09.
package build

var cloudberrySourceBundle = []string{
	"apache-cloudberry-2.2.0-incubating-src.tar.gz",
	"psutil-7.2.2.tar.gz",
	"PyGreSQL-5.2.5.tar.gz",
	"PyYAML-6.0.1.tar.gz",
	"Cython-0.29.37.tar.gz",
	"apache-cloudberry-backup-2.2.0-incubating-src.tar.gz",
	"apache-cloudberry-pxf-2.2.0-incubating-src.tar.gz",
	"gradle-wrapper-8.14.4.jar",
	"gradle-wrapper-8.14.4.jar.sha256",
	"gradle-8.14.4-bin.zip",
}

var cloudberryPXFSourceBundle = []string{
	"apache-cloudberry-pxf-2.2.0-incubating-src.tar.gz",
	"gradle-wrapper-8.14.4.jar",
	"gradle-wrapper-8.14.4.jar.sha256",
	"gradle-8.14.4-bin.zip",
}

var SpecialSourceMapping = map[string][]string{
	"scws":              {"scws-1.2.3.tar.bz2"},
	"openhalodb":        {"openhalodb-1.0.tar.gz"},
	"cloudberry":        cloudberrySourceBundle,
	"cloudberry-backup": {"apache-cloudberry-backup-2.2.0-incubating-src.tar.gz"},
	"cloudberry-pxf":    cloudberryPXFSourceBundle,
	"ivorysql":          {"ivorysql-5.6.tar.gz"},
	"babelfish":         {"babelfish-17-17.11-5.7.0.tar.gz", "babelfish-18-18.6-6.2.0.tar.gz"},
	"babelfish-17":      {"babelfish-17-17.11-5.7.0.tar.gz"},
	"babelfish-18":      {"babelfish-18-18.6-6.2.0.tar.gz"},
	"antlr4":            {"antlr4-cpp-runtime-4.13.2-source.zip"},
	"orioledb":          orioledbSources("postgres-patches16_50.tar.gz", "postgres-patches17_23.tar.gz", "postgres-patches18_4.tar.gz"),
	"orioledb-16":       orioledbSources("postgres-patches16_50.tar.gz"),
	"orioledb-17":       orioledbSources("postgres-patches17_23.tar.gz"),
	"orioledb-18":       orioledbSources("postgres-patches18_4.tar.gz"),
	"agensgraph":        {"agensgraph-2.18.6.1.tar.gz"},
	"polardb":           {"polardb-for-postgresql-17.11.1.0.tar.gz"},
	"polarstore":        {"polarstore-1.2.42-d0c5dc6.tar.gz"},
	"zlog":              {"zlog-1.2.20.tar.gz"},
	"pgedge":            pgedgeSources("postgresql-15.19.tar.gz", "postgresql-16.15.tar.gz", "postgresql-17.11.tar.gz", "postgresql-18.6.tar.gz"),
	"pgedge-15":         pgedgeSources("postgresql-15.19.tar.gz"),
	"pgedge-16":         pgedgeSources("postgresql-16.15.tar.gz"),
	"pgedge-17":         pgedgeSources("postgresql-17.11.tar.gz"),
	"pgedge-18":         pgedgeSources("postgresql-18.6.tar.gz"),
	"hunspell":          {"hunspell-1.0.tar.gz"},
	"libfq":             {"libfq-0.6.2.tar.gz"},
	"pdu":               {"pdu-3.0.25.12+git20260822.345422f.tar.gz"},
	"pgdog":             {"pgdog-0.1.63.tar.gz", "pgdog-0.1.63-vendor.tar.gz"},
	"inchi":             {"inchi_1.07.5+dfsg.orig.tar.xz", "inchi_1.07.5+dfsg-1.dsc", "inchi_1.07.5+dfsg-1.debian.tar.xz"},
	"graphblas":         {"graphblas-10.5.2.tar.gz"},
	"lagraph":           {"lagraph-1.2.2.tar.gz"},
	"rdkit": {
		"rdkit-Release_2026_09_2.tar.gz",
		"better-enums-0.11.3-enum.h",
		"RingDecomposerLib-1.1.3_rdkit.tar.gz",
		"rdkit_202303.3.orig.tar.xz",
		"rdkit_202303.3-3.pgdg120+1.dsc",
		"rdkit_202303.3-3.pgdg120+1.debian.tar.xz",
		"rdkit_202303.3-3.pgdg22.04+1.dsc",
		"rdkit_202303.3-3.pgdg22.04+1.debian.tar.xz",
	},
	"onesparse": {
		"onesparse-1.0.0.tar.gz",
		"graphblas-10.5.2.tar.gz",
		"lagraph-1.2.2.tar.gz",
	},

	// Multi-version PostgreSQL source packages
	"libpgfeutils": {
		"postgresql-14.24.tar.gz",
		"postgresql-15.19.tar.gz",
		"postgresql-16.15.tar.gz",
		"postgresql-17.11.tar.gz",
		"postgresql-18.6.tar.gz",
	},
}

func pgedgeSources(pgSources ...string) []string {
	return append(pgSources, "spock-5.0.13.tar.gz", "lolor-1.2.3.tar.gz", "snowflake-2.6.0.tar.gz")
}

func orioledbSources(pgPatchSources ...string) []string {
	return append(pgPatchSources, "orioledb-1.10.2.tar.gz")
}

// DocumentDB's prepare.sh consumes these in addition to its Makefile inputs.
var documentDBSources = []struct{ env, name string }{
	{"LIBBSON_SOURCE", "mongo-c-driver-1.28.0.tar.gz"},
	{"PCRE2_SOURCE", "pcre2-10.40.tar.gz"},
	{"UNCRUSTIFY_SOURCE", "uncrustify-uncrustify-0.68.1.tar.gz"},
	{"CITUS_TOOLS_SOURCE", "citus-tools-e36e4ea4258989bf527744334f6c633bb67e0686.tar.gz"},
}
