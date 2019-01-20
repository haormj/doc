#! /bin/bash
genpac --format "pac" \
--pac-proxy "PROXY 127.0.0.1:39307" \
--gfwlist-proxy "PROXY 127.0.0.1:39307" \
--gfwlist-local "./gfwlist.txt" \
--user-rule-from "./user-rules.txt" \
--output "autoproxy.pac"