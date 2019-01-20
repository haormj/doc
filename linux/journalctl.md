### journalctl

#### help

```shell
journalctl [OPTIONS...] [MATCHES...]

Query the journal.

Flags:
     --system              Show the system journal
     --user                Show the user journal for the current user
  -M --machine=CONTAINER   Operate on local container
  -S --since=DATE          Show entries not older than the specified date
  -U --until=DATE          Show entries not newer than the specified date
  -c --cursor=CURSOR       Show entries starting at the specified cursor
     --after-cursor=CURSOR Show entries after the specified cursor
     --show-cursor         Print the cursor after all the entries
  -b --boot[=ID]           Show current boot or the specified boot
     --list-boots          Show terse information about recorded boots
  -k --dmesg               Show kernel message log from the current boot
  -u --unit=UNIT           Show logs from the specified unit
  -t --identifier=STRING   Show entries with the specified syslog identifier
  -p --priority=RANGE      Show entries with the specified priority
  -e --pager-end           Immediately jump to the end in the pager
  -f --follow              Follow the journal
  -n --lines[=INTEGER]     Number of journal entries to show
     --no-tail             Show all lines, even in follow mode
  -r --reverse             Show the newest entries first
  -o --output=STRING       Change journal output mode (short, short-iso,
                                   short-precise, short-monotonic, verbose,
                                   export, json, json-pretty, json-sse, cat)
     --utc                 Express time in Coordinated Universal Time (UTC)
  -x --catalog             Add message explanations where available
     --no-full             Ellipsize fields
  -a --all                 Show all fields, including long and unprintable
  -q --quiet               Do not show privilege warning
     --no-pager            Do not pipe output into a pager
  -m --merge               Show entries from all available journals
  -D --directory=PATH      Show journal files from directory
     --file=PATH           Show journal file
     --root=ROOT           Operate on catalog files underneath the root ROOT
     --interval=TIME       Time interval for changing the FSS sealing key
     --verify-key=KEY      Specify FSS verification key
     --force               Override of the FSS key pair with --setup-keys
         
Commands:
  -h --help                Show this help text
     --version             Show package version
  -F --field=FIELD         List all values that a specified field takes
     --new-id128           Generate a new 128-bit ID
     --disk-usage          Show total disk usage of all journal files
     --vacuum-size=BYTES   Reduce disk usage below specified size
     --vacuum-time=TIME    Remove journal files older than specified date
     --flush               Flush all journal data from /run into /var
     --header              Show journal header information
     --list-catalog        Show all message IDs in the catalog
     --dump-catalog        Show entries in the message catalog
     --update-catalog      Update the message catalog database
     --setup-keys          Generate a new FSS key pair
     --verify              Verify journal file consistency   
```



#### 实时查看具体service的日志

```shell
journalctl -u xxx -f
```

#### 查看指定日志之后的日志

```shell
journalctl -u xxx -S 2018-09-18 
```

按照从后到前查看日志

```shell
journalctl -u xxx -r
```



