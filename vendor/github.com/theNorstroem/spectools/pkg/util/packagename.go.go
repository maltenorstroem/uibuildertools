package util

import (
	"path"
)

// returns a go package packagepb/subpackagepb;subpackagepb
func GetGoPackageName(TargetPath string) string {
	// not useable at the moment, because protoc-gen-grpc-gateway_out does not use the package name option (it uses the file structure)
	// segments := strings.Split(TargetPath,"/")
	// for i,p := range segments{
	// 	segments[i] = p + "pb"
	// }
	// tp := strings.Join(segments,"/")
	return TargetPath + ";" + path.Base(TargetPath) + "pb"
}
