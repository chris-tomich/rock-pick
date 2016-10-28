# Rock Pick
Rock Pick is an extensible CLI tool for querying a RocksDB database. Rock Pick can be extended with custom "renderers"
for keys and values. Rock Pick was born out of my own need to be able to easily format query results for my own RocksDB
databases which, unfortunately, the built-in ldb tool doesn't do.

Rock Pick is written in Go but, at the time of this writing (v1.7.3), no releases of Go include the _plugin_ package.
As a result, Rock Pick needs to be compiled with a version of Go built from the master branch. In the future, once the
_plugin_ package has been completed, this will cease to be an issue.