//go:build android
// +build android

package main

import (
	"fmt"
	"time"
)

func runDemoCGO(client *CGOClient) {
	motion := NewMotion(client)
	app := NewApp(client)
	device := NewDevice(client)
	storages := NewStorages(client)

	info, err := device.Info()
	if err != nil {
		fmt.Printf("Device info error: %v\n", err)
		return
	}
	fmt.Printf("Device: %s\n", info)

	size, err := device.GetScreenSize()
	if err == nil {
		fmt.Printf("Screen size: %s\n", size)
	}

	pkg, err := app.CurrentPackage()
	if err == nil {
		fmt.Printf("Current package: %s\n", pkg)
	}

	fmt.Println("Running simple automation...")

	if err := motion.Click(100, 200); err != nil {
		fmt.Printf("Click failed: %v\n", err)
	} else {
		fmt.Println("Click(100,200) OK")
	}

	time.Sleep(500 * time.Millisecond)

	if err := motion.InputText("hello"); err != nil {
		fmt.Printf("InputText failed: %v\n", err)
	} else {
		fmt.Println("InputText OK")
	}

	time.Sleep(300 * time.Millisecond)

	if err := motion.Swipe(500, 800, 500, 400, 300); err != nil {
		fmt.Printf("Swipe failed: %v\n", err)
	} else {
		fmt.Println("Swipe OK")
	}

	if err := device.Screenshot("/sdcard/rustgo_test.png"); err != nil {
		fmt.Printf("Screenshot failed: %v\n", err)
	} else {
		fmt.Println("Screenshot saved to /sdcard/rustgo_test.png")
	}

	_ = storages

	fmt.Println("Demo complete!")
}
