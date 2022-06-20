# cert-manager



## install

```shell
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.8.0/cert-manager.yaml
```



## CA

```yaml
apiVersion: cert-manager.io/v1
kind: Issuer
metadata:
  name: selfsigned-issuer
  namespace: default
spec:
  selfSigned: {}
---
apiVersion: cert-manager.io/v1
kind: Certificate
metadata:
  name: selfsigned-ca
  namespace: default
spec:
  secretName: root-secret
  duration: 2160h # 90d
  renewBefore: 360h # 15d
  isCA: true
  commonName: selfsigned-ca
  privateKey:
    algorithm: ECDSA
    size: 256
  issuerRef:
    group: cert-manager.io
    kind: Issuer
    name: selfsigned-issuer
---
apiVersion: cert-manager.io/v1
kind: Issuer
metadata:
  name: issuer
  namespace: default
spec:
  ca:
    secretName: root-secret
```



## Certificate

```yaml
apiVersion: cert-manager.io/v1
kind: Certificate
metadata:
  name: example.com
  namespace: default
spec:
  secretName: example-com
  secretTemplate:
    annotations:
      my-secret-annotation-1: "foo"
      my-secret-annotation-2: "bar"
    labels:
      my-secret-label: foo
  duration: 2160h # 90d
  renewBefore: 360h # 15d
  isCA: false
  subject:
    organizations:
    - example
  commonName: example.com
  privateKey:
    algorithm: RSA
    encoding: PKCS1
    size: 2048
  usages:
  - server auth
  - client auth
  dnsNames:
    - example.com
    - www.example.com
  ipAddresses:
    - 192.168.0.5
  issuerRef:
    kind: Issuer
    group: cert-manager.io
    name: issuer
```

