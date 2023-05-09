package design

// import (
// 	"fmt"
// 	"path/filepath"

// 	. "goa.design/goa/v3/dsl"
// )

// var _ = Service("Pet Service 2", func() {
// 	path := filepath.Join(workspaceRoot, "proto")

// 	fmt.Printf("Path: %s\n", path)
// 	Meta("protoc:include", path)

// 	Method("SendDogs", func() {
// 		StreamingPayload(dog)

// 		GRPC(func() {
// 			Response(CodeOK)
// 		})
// 	})
// 	Method("ReceiveDogs", func() {
// 		StreamingResult(dog)

// 		GRPC(func() {
// 			Response(CodeOK)
// 		})
// 	})

// 	Method("SendCat", func() {
// 		Payload(dog)

// 		GRPC(func() {
// 			Response(CodeOK)
// 		})
// 	})
// 	Method("ReceiveCat", func() {
// 		Result(dog)

// 		GRPC(func() {
// 			Response(CodeOK)
// 		})
// 	})
// })

// This results in the error `<value>.state: fields of type struct must use pointers`
// var dog = Type("dog", func() {
// 	ConvertTo(pets.Dog{})
// 	CreateFrom(pets.Dog{})

// 	Field(1, "Weight", Int32)
// 	Field(2, "Color", String)
// 	Field(3, "Age", Int32)

// 	Required("Weight")
// 	Required("Color")
// 	Required("Age")
// })
