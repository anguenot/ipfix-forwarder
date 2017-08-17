# ipfix-forwarder

[![alt text](https://travis-ci.org/anguenot/ipfix-forwarder.svg?branch=master "Travis CI build status")](https://travis-ci.org/anguenot/ipfix-forwarder)

[![APACHE2 License](https://img.shields.io/badge/license-Apache2.0-blue.svg?style=flat-square)](https://opensource.org/licenses/Apache-2.0)


`ipfix-forwarder` listens for IPFIX (RFC 5101) streams sent over UDP, parses, 
pre-processes, includes (VMware) vendor fields, converts to JSON and optionally 
can forward JSON string representation to a custom syslog destination.

It knows how to interpret and include the following vendor IPFIX fields:

1. VMware NSX
2. VMware vSphere Distributed Switch (VDS)

If using these vendors above the JSON will include an extra field named 
`nsxSegmentId` which will correspond to the edge `segmentId`. It then becomes 
trivial to bind a flow to corresponding inventory entities.

You can choose to export the JSON to a custom syslog destination.

*This server does not yet directly natively export flows to Apache Kafka. 
If you are looking to export your IPFIX flows to Apache Kafka, you can use 
`ipfix-forwarder` along with [syslog-ng](https://github.com/balabit/syslog-ng/) 
and the [syslog_kafka](https://github.com/ilanddev/syslogng_kafka) destination.*

## Examples

Start `ipfix-forwader` on `udp://0.0.0.0:2055`, interpret and include VMware 
vendor fields, log in console with a verbosity of 1.

```console

$ ./ipfix-forwarder -logtostderr -v 1 -vendor-vmware-vds -vendor-vmware-nsx

```

Start `ipfix-forwader` on `udp://0.0.0.0:2055`, interpret and include VMware 
vendor fields, log in console and file with a verbosity of 1 and export to a 
syslog server on udp://10.10.11.41:2056

```console

$ ./ipfix-forwarder -alsologtostderr -v 1 -vendor-vmware-vds -vendor-vmware-nsx -export-json-to-syslog -export-syslog-host 10.10.11.41 -export-syslog-port 2056 

```

## Usage

```console

$ ./ipfix-forwarder -h

usage: ipfix-forwarder [server-flags] [vendor(s)] [syslog-export-info] [logging-properties]

  -alsologtostderr
        log to standard error as well as files
  -export-json-to-syslog
        export flows to syslog server in JSON format
  -export-syslog-host string
        syslog server address for native flow export. (default "127.0.0.1")
  -export-syslog-port int
        syslog server port for native flows export. (default 514)
  -export-syslog-program string
        syslog message program for native flows export. (default "ipfix-forwarder")
  -export-syslog-proto string
        syslog server proto for native flows export. (default "UDP")
  -log_backtrace_at value
        when logging hits line file:N, emit a stack trace
  -log_dir string
        If non-empty, write log files in this directory
  -logtostderr
        log to standard error instead of files
  -server-address string
        IP the server will be listening to. (default "0.0.0.0")
  -server-port int
        Port we will be listening on. (default 2055)
  -stderrthreshold value
        logs at or above this threshold go to stderr
  -v value
        log level for V logs
  -vendor-vmware-nsx
        Include VMware NSX vendor fields.
  -vendor-vmware-vds
        Include VMware vSphere Distributed Switch (VDS) vendor fields.
  -version
        Version
  -vmodule value
        comma-separated list of pattern=N settings for file-filtered logging
```

## Download

You can find `linux/amd64` binaries XXX

## Build it

You will need Go 1.8.x installed.

```console

$ make build

```
