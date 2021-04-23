# waitfor
`waitfor` is a simple software than can be used as Kubernetes initContainer that will wait for some specific criteria

# Why
It starts with when we experience few small intermittent DNS not resolved error on our cluster, most of our microservices use DNS address (`something.namespace.cluster.local`) to refer each-other,
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

# same as above but will only sleep for 1sec before exit
waitfor --env=ENVA,ENVB --sleep=1s
```

```yaml
    apiVersion: v1
    kind: Pod
    metadata:
        name: myapp-pod
        labels:
            app: myapp
    spec:
        containers:
            - name: myapp-container
              image: busybox:1.28
              command: ['sh', '-c', 'echo The app is running! && sleep 3600']
        initContainers:
            - name: init-env-check
              image: lumioid/waitfor
              command: ["/go/bin/waitfor", "--env=SERVICE_HOST,SERVICE2_HOSTS", "--sleep=1s"]
```

# Todo
```
--envHost mode - to check for value of ENV var and ping the value
--service - to ping for services and make sure those are ready
```
