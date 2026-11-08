// Source bundle compatibility migrated from pgsty/pig (Apache-2.0).
package build

var cloudberrySourceBundle = []string{
	"apache-cloudberry-2.1.0-incubating-src.tar.gz",
	"psutil-5.7.0.tar.gz",
	"PyGreSQL-5.2.tar.gz",
	"PyYAML-5.4.1.tar.gz",
	"cloudberry-2.1.0-rpm-patches.tar.gz",
	"Cython-0.29.37.tar.gz",
	"apache-cloudberry-backup-2.1.0-incubating-src.tar.gz",
	"apache-cloudberry-pxf-2.1.0-incubating-src.tar.gz",
	"cloudberry-pxf-2.1.0-rpm-patches.tar.gz",
	"gradle-wrapper-6.8.2.jar",
	"gradle-wrapper-6.8.2.jar.sha256",
	"gradle-wrapper-8.5.jar",
	"gradle-wrapper-8.5.jar.sha256",
	"gradle-6.8.2-bin.zip",
	"gradle-8.5-bin.zip",
}

var cloudberryPXFSourceBundle = []string{
	"apache-cloudberry-pxf-2.1.0-incubating-src.tar.gz",
	"cloudberry-pxf-2.1.0-rpm-patches.tar.gz",
	"gradle-wrapper-6.8.2.jar",
	"gradle-wrapper-6.8.2.jar.sha256",
	"gradle-wrapper-8.5.jar",
	"gradle-wrapper-8.5.jar.sha256",
	"gradle-6.8.2-bin.zip",
	"gradle-8.5-bin.zip",
}

var SpecialSourceMapping = map[string][]string{
	"scws":              {"scws-1.2.3.tar.bz2"},
	"openhalodb":        {"openhalodb-1.0.tar.gz"},
	"cloudberry":        cloudberrySourceBundle,
	"cloudberry-backup": {"apache-cloudberry-backup-2.1.0-incubating-src.tar.gz"},
	"cloudberry-pxf":    cloudberryPXFSourceBundle,
	"ivorysql":          {"ivorysql-5.4.tar.gz"},
	"babelfish":         {"babelfish-17-17.7-5.4.0.tar.gz"},
	"babelfish-17":      {"babelfish-17-17.7-5.4.0.tar.gz"},
	"babelfish-18":      {"babelfish-18-18.3-6.0.0.tar.gz"},
	"antlr4":            {"antlr4-cpp-runtime-4.13.2-source.zip"},
	"orioledb":          orioledbSources("postgres-patches18_1.tar.gz"),
	"orioledb-16":       orioledbSources("postgres-patches16_47.tar.gz"),
	"orioledb-17":       orioledbSources("postgres-patches17_20.tar.gz"),
	"orioledb-18":       orioledbSources("postgres-patches18_1.tar.gz"),
	"agensgraph":        {"agensgraph-2.17.0.tar.gz"},
	"polardb":           {"polardb-for-postgresql-17.10.1.0.tar.gz"},
	"polarstore":        {"polarstore-1.2.42-d0c5dc6.tar.gz"},
	"zlog":              {"zlog-1.2.18.tar.gz"},
	"pgedge":            pgedgeSources("postgresql-18.4.tar.gz"),
	"pgedge-15":         pgedgeSources("postgresql-15.18.tar.gz"),
	"pgedge-16":         pgedgeSources("postgresql-16.14.tar.gz"),
	"pgedge-17":         pgedgeSources("postgresql-17.10.tar.gz"),
	"pgedge-18":         pgedgeSources("postgresql-18.4.tar.gz"),
	"hunspell":          {"hunspell-1.0.tar.gz"},
	"libfq":             {"libfq-0.6.2.tar.gz"},
	"pdu":               {"pdu-3.0.25.12.tar.gz"},
	"pgdog":             {"pgdog-0.1.32.tar.gz"},
	"inchi":             {"inchi-1.07.3.tar.gz"},
	"graphblas":         {"graphblas-10.2.0.tar.gz"},
	"lagraph":           {"lagraph-1.2.1.tar.gz"},
	"rdkit": {
		"rdkit_202503.6.orig.tar.xz",
		"better-enums-0.11.3-enum.h",
		"rapidjson-1.1.0.tar.gz",
		"inchi-1.07.3.tar.gz",
	},
	"onesparse": {
		"onesparse-1.0.0.tar.gz",
		"graphblas-10.2.0.tar.gz",
		"lagraph-1.2.1.tar.gz",
	},

	// Multi-version PostgreSQL source packages
	"libpgfeutils": {
		"postgresql-14.23.tar.gz",
		"postgresql-15.18.tar.gz",
		"postgresql-16.14.tar.gz",
		"postgresql-17.10.tar.gz",
		"postgresql-18.4.tar.gz",
	},
}

func pgedgeSources(pgSource string) []string {
	return []string{
		pgSource,
		"spock-5.0.10.tar.gz",
		"lolor-1.2.2.tar.gz",
		"snowflake-2.5.0.tar.gz",
	}
}

func orioledbSources(pgPatchSource string) []string {
	return []string{
		pgPatchSource,
		"orioledb-beta16.tar.gz",
	}
}
