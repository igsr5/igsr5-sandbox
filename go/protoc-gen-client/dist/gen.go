// Code generated by protoc-gen-client. DO NOT EDIT.

package main

import (
	"encoding/json"
	"log"

	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/jhump/protoreflect/desc"
	"github.com/pkg/errors"
)

var dfs []*descriptor.FileDescriptorProto

func init() {
	dfs = make([]*descriptor.FileDescriptorProto, 0, 1)
	var err error

	byte__FileDescriptorProto__0 := []byte{123, 34, 110, 97, 109, 101, 34, 58, 34, 112, 114, 111, 116, 111, 47, 101, 120, 97, 109, 112, 108, 101, 46, 112, 114, 111, 116, 111, 34, 44, 34, 112, 97, 99, 107, 97, 103, 101, 34, 58, 34, 101, 120, 97, 109, 112, 108, 101, 34, 44, 34, 109, 101, 115, 115, 97, 103, 101, 95, 116, 121, 112, 101, 34, 58, 91, 123, 34, 110, 97, 109, 101, 34, 58, 34, 69, 120, 97, 109, 112, 108, 101, 34, 44, 34, 102, 105, 101, 108, 100, 34, 58, 91, 123, 34, 110, 97, 109, 101, 34, 58, 34, 110, 97, 109, 101, 34, 44, 34, 110, 117, 109, 98, 101, 114, 34, 58, 49, 44, 34, 108, 97, 98, 101, 108, 34, 58, 49, 44, 34, 116, 121, 112, 101, 34, 58, 57, 44, 34, 106, 115, 111, 110, 95, 110, 97, 109, 101, 34, 58, 34, 110, 97, 109, 101, 34, 125, 44, 123, 34, 110, 97, 109, 101, 34, 58, 34, 105, 100, 34, 44, 34, 110, 117, 109, 98, 101, 114, 34, 58, 50, 44, 34, 108, 97, 98, 101, 108, 34, 58, 49, 44, 34, 116, 121, 112, 101, 34, 58, 53, 44, 34, 106, 115, 111, 110, 95, 110, 97, 109, 101, 34, 58, 34, 105, 100, 34, 125, 44, 123, 34, 110, 97, 109, 101, 34, 58, 34, 110, 101, 115, 116, 101, 100, 34, 44, 34, 110, 117, 109, 98, 101, 114, 34, 58, 52, 44, 34, 108, 97, 98, 101, 108, 34, 58, 49, 44, 34, 116, 121, 112, 101, 34, 58, 49, 49, 44, 34, 116, 121, 112, 101, 95, 110, 97, 109, 101, 34, 58, 34, 46, 101, 120, 97, 109, 112, 108, 101, 46, 78, 101, 115, 116, 101, 100, 34, 44, 34, 106, 115, 111, 110, 95, 110, 97, 109, 101, 34, 58, 34, 110, 101, 115, 116, 101, 100, 34, 125, 44, 123, 34, 110, 97, 109, 101, 34, 58, 34, 110, 101, 115, 116, 101, 100, 86, 50, 34, 44, 34, 110, 117, 109, 98, 101, 114, 34, 58, 53, 44, 34, 108, 97, 98, 101, 108, 34, 58, 49, 44, 34, 116, 121, 112, 101, 34, 58, 49, 49, 44, 34, 116, 121, 112, 101, 95, 110, 97, 109, 101, 34, 58, 34, 46, 101, 120, 97, 109, 112, 108, 101, 46, 69, 120, 97, 109, 112, 108, 101, 46, 78, 101, 115, 116, 101, 100, 86, 50, 34, 44, 34, 106, 115, 111, 110, 95, 110, 97, 109, 101, 34, 58, 34, 110, 101, 115, 116, 101, 100, 86, 50, 34, 125, 93, 44, 34, 110, 101, 115, 116, 101, 100, 95, 116, 121, 112, 101, 34, 58, 91, 123, 34, 110, 97, 109, 101, 34, 58, 34, 78, 101, 115, 116, 101, 100, 86, 50, 34, 44, 34, 102, 105, 101, 108, 100, 34, 58, 91, 123, 34, 110, 97, 109, 101, 34, 58, 34, 110, 97, 109, 101, 34, 44, 34, 110, 117, 109, 98, 101, 114, 34, 58, 49, 44, 34, 108, 97, 98, 101, 108, 34, 58, 49, 44, 34, 116, 121, 112, 101, 34, 58, 57, 44, 34, 106, 115, 111, 110, 95, 110, 97, 109, 101, 34, 58, 34, 110, 97, 109, 101, 34, 125, 93, 125, 93, 125, 44, 123, 34, 110, 97, 109, 101, 34, 58, 34, 78, 101, 115, 116, 101, 100, 34, 44, 34, 102, 105, 101, 108, 100, 34, 58, 91, 123, 34, 110, 97, 109, 101, 34, 58, 34, 110, 97, 109, 101, 34, 44, 34, 110, 117, 109, 98, 101, 114, 34, 58, 49, 44, 34, 108, 97, 98, 101, 108, 34, 58, 49, 44, 34, 116, 121, 112, 101, 34, 58, 57, 44, 34, 106, 115, 111, 110, 95, 110, 97, 109, 101, 34, 58, 34, 110, 97, 109, 101, 34, 125, 93, 125, 93, 44, 34, 115, 101, 114, 118, 105, 99, 101, 34, 58, 91, 123, 34, 110, 97, 109, 101, 34, 58, 34, 69, 120, 97, 109, 112, 108, 101, 83, 101, 114, 118, 105, 99, 101, 34, 44, 34, 109, 101, 116, 104, 111, 100, 34, 58, 91, 123, 34, 110, 97, 109, 101, 34, 58, 34, 69, 99, 104, 111, 34, 44, 34, 105, 110, 112, 117, 116, 95, 116, 121, 112, 101, 34, 58, 34, 46, 101, 120, 97, 109, 112, 108, 101, 46, 69, 120, 97, 109, 112, 108, 101, 34, 44, 34, 111, 117, 116, 112, 117, 116, 95, 116, 121, 112, 101, 34, 58, 34, 46, 101, 120, 97, 109, 112, 108, 101, 46, 69, 120, 97, 109, 112, 108, 101, 34, 44, 34, 111, 112, 116, 105, 111, 110, 115, 34, 58, 123, 125, 125, 93, 125, 93, 44, 34, 111, 112, 116, 105, 111, 110, 115, 34, 58, 123, 34, 103, 111, 95, 112, 97, 99, 107, 97, 103, 101, 34, 58, 34, 46, 47, 101, 120, 97, 109, 112, 108, 101, 59, 101, 120, 97, 109, 112, 108, 101, 95, 112, 98, 34, 125, 44, 34, 115, 111, 117, 114, 99, 101, 95, 99, 111, 100, 101, 95, 105, 110, 102, 111, 34, 58, 123, 34, 108, 111, 99, 97, 116, 105, 111, 110, 34, 58, 91, 123, 34, 115, 112, 97, 110, 34, 58, 91, 48, 44, 48, 44, 50, 52, 44, 49, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 49, 50, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 48, 44, 48, 44, 49, 56, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 50, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 50, 44, 48, 44, 49, 54, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 56, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 52, 44, 48, 44, 52, 51, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 56, 44, 49, 49, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 52, 44, 48, 44, 52, 51, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 52, 44, 48, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 54, 44, 48, 44, 49, 54, 44, 49, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 52, 44, 48, 44, 49, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 54, 44, 56, 44, 49, 53, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 52, 44, 48, 44, 50, 44, 48, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 55, 44, 50, 44, 49, 56, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 52, 44, 48, 44, 50, 44, 48, 44, 53, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 55, 44, 50, 44, 56, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 52, 44, 48, 44, 50, 44, 48, 44, 49, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 55, 44, 57, 44, 49, 51, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 52, 44, 48, 44, 50, 44, 48, 44, 51, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 55, 44, 49, 54, 44, 49, 55, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 52, 44, 48, 44, 50, 44, 49, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 56, 44, 50, 44, 49, 53, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 52, 44, 48, 44, 50, 44, 49, 44, 53, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 56, 44, 50, 44, 55, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 52, 44, 48, 44, 50, 44, 49, 44, 49, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 56, 44, 56, 44, 49, 48, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 52, 44, 48, 44, 50, 44, 49, 44, 51, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 56, 44, 49, 51, 44, 49, 52, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 52, 44, 48, 44, 50, 44, 50, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 57, 44, 50, 44, 50, 48, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 52, 44, 48, 44, 50, 44, 50, 44, 54, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 57, 44, 50, 44, 56, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 52, 44, 48, 44, 50, 44, 50, 44, 49, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 57, 44, 57, 44, 49, 53, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 52, 44, 48, 44, 50, 44, 50, 44, 51, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 57, 44, 49, 56, 44, 49, 57, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 52, 44, 48, 44, 51, 44, 48, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 49, 49, 44, 50, 44, 49, 51, 44, 51, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 52, 44, 48, 44, 51, 44, 48, 44, 49, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 49, 49, 44, 49, 48, 44, 49, 56, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 52, 44, 48, 44, 51, 44, 48, 44, 50, 44, 48, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 49, 50, 44, 52, 44, 50, 48, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 52, 44, 48, 44, 51, 44, 48, 44, 50, 44, 48, 44, 53, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 49, 50, 44, 52, 44, 49, 48, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 52, 44, 48, 44, 51, 44, 48, 44, 50, 44, 48, 44, 49, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 49, 50, 44, 49, 49, 44, 49, 53, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 52, 44, 48, 44, 51, 44, 48, 44, 50, 44, 48, 44, 51, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 49, 50, 44, 49, 56, 44, 49, 57, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 52, 44, 48, 44, 50, 44, 51, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 49, 53, 44, 50, 44, 50, 52, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 52, 44, 48, 44, 50, 44, 51, 44, 54, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 49, 53, 44, 50, 44, 49, 48, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 52, 44, 48, 44, 50, 44, 51, 44, 49, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 49, 53, 44, 49, 49, 44, 49, 57, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 52, 44, 48, 44, 50, 44, 51, 44, 51, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 49, 53, 44, 50, 50, 44, 50, 51, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 52, 44, 49, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 49, 56, 44, 48, 44, 50, 48, 44, 49, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 52, 44, 49, 44, 49, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 49, 56, 44, 56, 44, 49, 52, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 52, 44, 49, 44, 50, 44, 48, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 49, 57, 44, 50, 44, 49, 56, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 52, 44, 49, 44, 50, 44, 48, 44, 53, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 49, 57, 44, 50, 44, 56, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 52, 44, 49, 44, 50, 44, 48, 44, 49, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 49, 57, 44, 57, 44, 49, 51, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 52, 44, 49, 44, 50, 44, 48, 44, 51, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 49, 57, 44, 49, 54, 44, 49, 55, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 54, 44, 48, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 50, 50, 44, 48, 44, 50, 52, 44, 49, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 54, 44, 48, 44, 49, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 50, 50, 44, 56, 44, 50, 50, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 54, 44, 48, 44, 50, 44, 48, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 50, 51, 44, 50, 44, 52, 49, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 54, 44, 48, 44, 50, 44, 48, 44, 49, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 50, 51, 44, 54, 44, 49, 48, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 54, 44, 48, 44, 50, 44, 48, 44, 50, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 50, 51, 44, 49, 50, 44, 49, 57, 93, 125, 44, 123, 34, 112, 97, 116, 104, 34, 58, 91, 54, 44, 48, 44, 50, 44, 48, 44, 51, 93, 44, 34, 115, 112, 97, 110, 34, 58, 91, 50, 51, 44, 51, 48, 44, 51, 55, 93, 125, 93, 125, 44, 34, 115, 121, 110, 116, 97, 120, 34, 58, 34, 112, 114, 111, 116, 111, 51, 34, 125}
	f__0 := &descriptor.FileDescriptorProto{}
	err = json.Unmarshal(byte__FileDescriptorProto__0, f__0)
	if err != nil {
		log.Fatalf("failed to unmarshal file: %v", err)
	}

	dfs = append(dfs, f__0)

}

func FileDescriptors() ([]*desc.FileDescriptor, error) {
	var descFiles = make([]*desc.FileDescriptor, 0, len(dfs))
	for _, df := range dfs {
		f, err := desc.CreateFileDescriptor(df, &desc.FileDescriptor{})
		if err != nil {
			return nil, errors.Wrap(err, "failed to create file descriptor")
		}

		descFiles = append(descFiles, f)
	}

	return descFiles, nil
}
