# Release procedure

As of now we don't have an automated release process. This guide shows how
to do a new release of IMPS.

Steps:
- Update deploy/charts/imagepullsecrets/Chart.yaml with the new target version
- Update deploy/charts/imagepullsecrets/values.yaml with the new target version
  (image name)
- Execute a make build, so that all generated files are updated
- Commit the change to main (please raise a PR)
- Add and push the following *annotated* tags and push them seperately (circleci ignores
  tags when multiple tags are pushed):
  - vX.Y.Z: releases new version from Docker image, main go.mod
  - deploy/charts/vX.Y.Z: releases a new version from the charts module
  - chart/imagepullsecrets/X.Y.Z: release a new version from the chart
