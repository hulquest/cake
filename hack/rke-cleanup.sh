#!/usr/bin/env bash
RP="/NetApp-HCI-Datacenter-01/host/NetApp-HCI-Cluster-01/Resources/rancher"
if ! $(govc env 2>/dev/null) ; then
        # Sample govc env file.  
        # $ cat ~/.govc-env
        #/usr/bin/env bash
        # export GOVC_USERNAME=administrator@vsphere.local
        # export GOVC_PASSWORD='NetApp1!!'
        # export GOVC_URL=172.60.0.151
        # export GOVC_INSECURE=true
        # govc env
    echo "Trying to set govc env vars (source ~/.govc-env)"
    source ~/.govc-env
    if [[ -z ${GOVC_URL} ]] ; then
        exit 82
    fi
fi

# Get the pool in json     - extract the opaque vm name               - convert to vm name - poweroff and delete.
COUNT=$(govc pool.info -json ${RP} | jq -r '.ResourcePools[].Vm[]? | join(":")' | wc -l)
echo "Will delete ${COUNT} machines in pool [${RP}]."
if [[ ${COUNT} -gt 0 ]] ; then
    govc pool.info -json ${RP} | jq -r '.ResourcePools[].Vm[] | join(":")' | xargs govc ls -L | xargs govc vm.destroy
fi
