# zkcli

Simple, lightweight, dependable CLI for ZooKeeper

**zkcli** is a non-interactive command line client for [ZooKeeper](http://zookeeper.apache.org/). It provides with:

 * Basic CRUD-like operations: `create`, `set`, `delete`, `exists`, `get`, `ls` (aka `children`).
 * Extended operations: `lsr` (ls recursive), `creater` (create recursively)
 * Well formatted and controlled output: supporting either `txt` or `json` format
 * Single, no-dependencies binary file, based on a native Go ZooKeeper library 
   by [github.com/samuel/go-zookeeper](http://github.com/samuel/go-zookeeper) ([LICENSE](https://github.com/go-zkcli/zkcli/blob/master/go-zookeeper-LICENSE))

### Download & Install

There are [pre built binaries](https://github.com/go-zkcli/zkcli/releases) for download.
You can find `RPM` and `deb` packages, as well as pre-compiled, dependency free `zkcli` executable binary.
In fact, the only file installed by the pre-built `RPM` and `deb` packages is said executable binary file. 

Otherwise the source code is freely available; you will need `git` installed as well as `go`, and you're on your own.
  
### Usage:


```
NAME:
   zkcli - zkcli is a non-interactive command line client for ZooKeeper

USAGE:
   zkcli [global options] command [command options] [arguments...]

VERSION:
   1.1.0

AUTHOR(S):

COMMANDS:
   exists                       zkcli excists <path>
   get                          zkcli get <path>
   set                          zkcli set <path> [data]
   create                       zkcli create [command options] <path> <data>
   list, ls                     zkcli list [command options] [path]
   delete, del, rm, remove
   help, h                      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --servers "localhost:2181"   ZK Server list in format: srv1[:port1][,srv2[:port2]...] [$ZKC_SERVERS]
   --format "txt"               Output format [$ZKC_FORMAT]
   --help, -h                   show help
   --version, -v                print the version
```
    

### Examples:
    
    
```
# Look some path with zkcli from docker 
docker run -it zkcli:latest zkcli --servers zk-server.example.com:2181 ls "/path/"
```

```

$ zkcli --servers srv-1,srv-2,srv-3 -c create /demo_only some_value

# Servers can also be define using environment varaible
$ export ZKC_SERVERS=srv-1,srv-2,srv-3 

# Default port is 2181. The above is equivalent to:
$ zkcli create /demo_only some_value

$ zkcli --format=txt get /demo_only
some_value

# Same as above, JSON format output:
$ zkcli --format=json get /demo_only
"some_value"

# exists exits with exit code 0 when path exists, 1 when path does not exist 
$ zkcli exists /demo_only
true

$ zkcli set /demo_only another_value

$ zkcli --format=json get /demo_only
"another_value"

$ zkcli delete /demo_only

$ zkcli get /demo_only
2014-09-15 04:07:16 FATAL zk: node does not exist

$ zkcli create /demo_only "path placeholder"
$ zkcli create /demo_only/key1 "value1"
$ zkcli create /demo_only/key2 "value2"
$ zkcli create /demo_only/key3 "value3"

$ zkcli ls /demo_only
key3
key2
key1

# Same as above, JSON format output:
$ zkcli --format=json ls /demo_only
["key3","key2","key1"]

$ zkcli delete /demo_only
2014-09-15 08:26:31 FATAL zk: node has children

$ zkcli delete /demo_only/key1
$ zkcli delete /demo_only/key2
$ zkcli delete /demo_only/key3
$ zkcli delete /demo_only

# /demo_only path now does not exist.

# Create recursively a path:
$ zkcli create -f "/demo_only/child/key1" "val1"
$ zkcli create -f "/demo_only/child/key2" "val2"

$ zkcli get "/demo_only/child/key1"
val1

# This path was auto generated due to recursive create:
$ zkcli get "/demo_only" 
zkcli auto-generated

# ls recursively a path and all sub children:
$ zkcli ls -r "/demo_only" 
child
child/key1
child/key2
```
    

The tool was built in order to allow with shell scripting seamless
integration with ZooKeeper. There is another, official command line
tool for ZooKeeper that the orginal author found inadequate in terms of
output format and output control, as well as large footprint. **zkcli**
overcomes those limitations and provides with quick, well formatted
output as well as enhanced functionality.

### License

Release under the [Apache 2.0 license](https://github.com/go-zkcli/zkcli/blob/master/LICENSE)

Authored by [Shlomi Noach](https://github.com/shlomi-noach) at [Outbrain](https://github.com/outbrain)
