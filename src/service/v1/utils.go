// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"ftx-omx/src/cfg"
	"log"
	"math/rand"
	"time"
)

func generateID(length int) string {
	return randStringBytes(length)
}

const letterBytes = "abcdefghkmnpqrstuvwxyz123456789"

func randStringBytes(n int) string {

	if cfg.SeedRandomNumbGen {
		// https://pkg.go.dev/math/rand#Seed
		// Seed needs to be called first otherwise rand returns always the same value, which comes handy during testing.
		// This flag must be true for production.
		rand.Seed(time.Now().UnixNano())
		// Now we seed again with random-random int
		rand.Seed(time.Now().UnixNano() + rand.Int63n(rand.Int63()))
	}

	// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func PrintInitHeader(serviceName, msg string) {
	DbgPrint("")
	DbgPrint(serviceName)
	DbgPrint("===========================")
	DbgPrint(msg)
	DbgPrint("===========================")
}

func PrintDbgHeader() {
	DbgPrint("")
	DbgPrint("==========================")
	DbgPrint("Start serving:")
	DbgPrint("==========================")
}

func PrintStartHeader(serviceName string, port string, elapsed time.Duration) {
	log.Println()
	log.Println(serviceName)
	log.Printf("Service start time (Milliseconds): %d", elapsed.Milliseconds())
	log.Println("========================================== ")
	log.Println(" Health check at: 	host" + port + "/health")
	log.Println("========================================== ")
	log.Println()
}

func PrintStopHeader(elapsed time.Duration) {
	log.Println()
	log.Printf("Service shutdown took (Milliseconds): %d", elapsed.Milliseconds())
	log.Println("========================================== ")
	log.Println(" Shutdown complete - Switch off now!")
	log.Println("========================================== ")
	log.Println()
}
