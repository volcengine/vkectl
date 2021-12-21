#!/bin/bash
module="github.com/volcengine/vkectl"

# generate rpc client function
function generate_rpc_client() {
  path=$1
  thrift_name=$2;
  repo=$3
  cd "$path" || exit
  rm -rf kitex_gen
  # download idl
  if [ -z "$branch" ]; then
    git clone "$GITURL/$repo.git"
  else
    git clone -b "$branch" "$GITURL/$repo.git"
  fi
  kitex -module ${module} "$repo/$thrift_name"
  rm -rf "$repo"
  cd - || exit;
}

function generate_resource_client(){
  # generate resource rpc client
  generate_rpc_client pkg/model/resource pkg/server/idl/paastob-vke-resourceserver.thrift vke-resource;  # generate resource rpc client
  # generate and format http client
  go run hack/gen_client/main.go;
  goimports -w pkg/client;
}

function generate_security_client(){
  # generate security rpc client
  generate_rpc_client pkg/model/security pkg/server/idl/security.thrift vke-security;  # generate security rpc client
  # generate and format http client
  go run hack/gen_client/main.go;
  goimports -w pkg/client;
}

# download fmt tools
which goimports;
if [ $? -ne 0 ];then
  go install golang.org/x/tools/cmd/goimports
fi

group=$1;
branch=$2;

if [ -z "$GITURL" ];then
  # for example: git@github.com:volcengine
  read -r -p "please input GITURL (eg. git@github.com:volcengine): " GITURL
fi

if [ -z "$group" ];then
  echo "Error: please import module (resource, security)"
elif [ "$group" = resource ];then
  echo "updating resource client..."
  generate_resource_client;
elif [ "$group" = security ];then
  echo "updating security client..."
  generate_security_client;
else
  echo "group should be resource or security, now is " "$group"
fi
