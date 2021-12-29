#!/bin/bash

if [ -n "$(git diff)" ]; then
  echo "git repo should be clean"
  exit 1
fi

read -r -p "Enter the deb package version to publish: " DebVersion
read -r -p "Enter the deb package message to publish: " DebMessage
DebDate=$(date -R)

sed -i "s/__DebVersion__/$DebVersion/g" debian/changelog
sed -i "s/__DebMessage__/$DebMessage/g" debian/changelog
sed -i "s/__DebDate__/$DebDate/g" debian/changelog

sed -i "1 i\VERSION := $(git describe --tags --always)" Makefile
sed -i "1 i\\BRANCH := $(git branch | grep \* | cut -d ' ' -f2)" Makefile
sed -i "1 i\\GITCOMMIT := $(git rev-parse HEAD)" Makefile
sed -i "1 i\\GITTREESTATE := clean" Makefile
sed -i "1 i\\BUILDDATE := $(date -u +"%Y-%m-%dT%H:%M:%SZ")" Makefile

debuild -S

dput -f ppa:volcengine/ppa "../vkectl_${DebVersion}_source.changes"

# clean
rm -f "../vkectl_${DebVersion}.*"
rm -f "../vkectl_${DebVersion}_source.*"

git checkout -- .