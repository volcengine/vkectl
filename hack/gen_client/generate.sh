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
  git clone -b "$branch" "$GITURL/$repo.git"
  kitex -module ${module} "$repo/$thrift_name"
  rm -rf "$repo"
  cd - || exit;
}

function generate_app_client(){
  branch=$1
  # generate app rpc client
  generate_rpc_client pkg/model/app pkg/server/idl/app.thrift vke-app;  # generate app rpc client
  # generate and format http client
  go run hack/gen_client/main.go;
  goimports -w pkg/client;
}

function generate_resource_client(){
  branch=$1
  # generate resource rpc client
  generate_rpc_client pkg/model/resource pkg/server/idl/paastob-vke-resourceserver.thrift vke-resource;  # generate resource rpc client
  # generate and format http client
  go run hack/gen_client/main.go;
  goimports -w pkg/client;
}

function generate_security_client(){
  branch=$1
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
if [ -z $branch ];then
  branch=release;
fi

if [ -z $GITURL ];then
  # git@github.com:volcengine/vke-resource
  echo "Error: please set env GITURL"
  exit
fi

if [ -z $group ];then
  echo "Error: please import module (app, resource, security)"
elif [ $group = app ];then
  echo "updating app client..."
  generate_app_client $branch;
elif [ $group = resource ];then
  echo "updating resource client..."
  generate_resource_client $branch;
elif [ $group = security ];then
  echo "updating security client..."
  generate_security_client $branch;
else
  echo "group should be app, resource or security, now is " $group
fi
