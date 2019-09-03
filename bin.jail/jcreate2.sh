#!/bin/sh

_MYDIR=$( /usr/bin/dirname `/bin/realpath $0` )

${_MYDIR}/../bs_router-client -config ${_MYDIR}/../config.json -body '{"Command":"jcreate","CommandArgs":{
"jname":"jail2",
"jail_profile": "default",
"host_hostname": "jail1.my.domain",
"astart": "1",
"pkglist":"misc/mc devel/git"
}}'
