package main

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/volcengine/vkectl/hack/gen_client/template"
)

func main() {
	_, dir, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(dir), "../../")
	client := template.GetClientInfo(root)
	for _, c := range client {
		service, funcparts, imports, err := template.GetFuncPart(c.Path)
		if err != nil {
			fmt.Printf("get function part failed, client info %v, err %v\n", c, err)
			continue
		}
		path := template.ProducePath(root, c, service)

		// generate OpenAPI client
		dir, file := filepath.Split(path)
		apiPath := filepath.Join(dir, file)
		err = template.ProduceBaseFile(apiPath, c.Group, service, imports)
		if err != nil {
			fmt.Printf("[error] template base file failed, api client info %v, err %v\n", c, err)
			continue
		}
		functions := template.FunctionJointForOpenAPI(service, funcparts)
		err = template.WriteFuncs(apiPath, functions)
		if err != nil {
			fmt.Printf("[error] write functions failed, api client info %v, err %v\n", c, err)
			continue
		}
	}
}
