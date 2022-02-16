# Changelog

## 0.4.0 - 2022-02-16

This patch changes the group name of ImagePullSecrets from images.banzaicloud.io to images.cisco.com.

This is a breaking change.

# Upgrading

To upgrade to v0.4.0 the following procedure should be followed:

*Step 1)* Scale down the replicaset of IMPS to 0. This prevents IMPS from updating any pull
secrets.

*Step 2)* Change all Secrets containing ECR login information whose are referenced from any ImagePullSecrets object to
`cisco.com/aws-ecr-login-config` (from `banzaicloud.io/aws-ecr-login-config`)

*Step 3)* Install the new version of the imps helm chart.

*Step 4)* Recreate all ImagePullSecrets objects with the new group of `images.cisco.com` (see [README.md](README.md) for examples)

*Step 5)* Delete the old CustomResourceDefinition
from your cluster as:

```
kubectl get imps
```

will display the CRDs with the old group information and not the new group's contents.

To delete the old CRDs please execute:
```
kubectl delete crd imagepullsecrets.images.banzaicloud.io
```

