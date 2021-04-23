# waitfor
waitfor is a simple software than can be used as Kubernetes initContainer that will wait for some specific criteria

# Why
It starts with when we experience few small intermittent DNS not resolved error on our cluster, most of our microservices use DNS address (something.namespace.cluster.local) to refer each-other,
it turns out that kube-dns got killed because it's sometimes allocated to memory-heavy nodes.

There's a lot of solution to this, such as rebooting kube-dns when the issue occurs,
make sure the pod is allocated to proper node or create dedicated cluster for kube-system processes, create a
node-dns-caching etc.

We opt in to use service ENV variable instead since with the combination of `initContainer` we can also
make sure our `service` dependencies is created in proper order (in the case if we update or restart our whole kube control plane)

# Usage

```bash
# This will check for all specified system ENVS to be exists, will exit with code 1 if one of it is missing
waitfor --env=ENVA,ENVB,ENVC,ENV_D
```

