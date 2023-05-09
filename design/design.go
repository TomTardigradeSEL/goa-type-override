package design

import (
	"fmt"
	"path/filepath"

	. "goa.design/goa/v3/dsl"
)

var workspaceRoot = filepath.Join("/", "workspaces", "goa-type-override")

var _ = Service("Pet Service", func() {
	path := filepath.Join(workspaceRoot, "proto")

	fmt.Printf("Path: %s\n", path)
	Meta("protoc:include", path)

	Method("SendDogs", func() {
		StreamingPayload(dogPayload)

		GRPC(func() {
			Response(CodeOK)
		})
	})
	Method("ReceiveDogs", func() {
		StreamingResult(dogPayload)

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("SendCat", func() {
		Payload(catPayload)

		GRPC(func() {
			Response(CodeOK)
		})
	})
	Method("ReceiveCat", func() {
		Result(catPayload)

		GRPC(func() {
			Response(CodeOK)
		})
	})
})

var goPkg = "petsvc/interfaces/pets"

var dogPayload = Type("dogPayload", func() {
	var (
		protoPkg = "pets/dog.proto"
		typeName = "pets.Dog"
	)

	Field(1, "dog", func() {
		Meta("struct:field:proto", typeName, protoPkg, typeName, goPkg)
		Meta("struct:field:type", typeName, goPkg)
	})

	Required("dog")
})

var catPayload = Type("catPayload", func() {
	var (
		protoPkg = "pets/cat.proto"
		typeName = "pets.Cat"
	)

	Field(1, "cat", func() {
		Meta("struct:field:proto", typeName, protoPkg, typeName, goPkg)
		Meta("struct:field:type", typeName, goPkg)
	})

	Required("cat")
})
