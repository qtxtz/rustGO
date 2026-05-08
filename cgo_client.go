//go:build android
// +build android

package main

/*
#cgo LDFLAGS: -L. -lyunxi_rust -ldl
#include <stdlib.h>

extern int rustgo_init();
extern void rustgo_cleanup();
extern int rustgo_tap(int x, int y);
extern int rustgo_long_press(int x, int y, int duration_ms);
extern int rustgo_swipe(int x1, int y1, int x2, int y2, int duration_ms);
extern int rustgo_input_text(const char* text);
extern int rustgo_key_event(int keycode);
extern int rustgo_touch_down(int x, int y, int finger_id);
extern int rustgo_touch_move(int x, int y, int finger_id);
extern int rustgo_touch_up(int x, int y, int finger_id);
extern int rustgo_screenshot(const char* path);
extern char* rustgo_get_screen_size();
extern int rustgo_launch_app(const char* package_name);
extern int rustgo_close_app(const char* package_name);
extern char* rustgo_shell_exec(const char* cmd);
extern int rustgo_file_exists(const char* path);
extern char* rustgo_file_read(const char* path);
extern int rustgo_file_write(const char* path, const char* content);
extern char* rustgo_device_info();
extern int rustgo_wait(int ms);
extern char* rustgo_get_clipboard();
extern int rustgo_set_clipboard(const char* text);
extern char* rustgo_current_package();
extern char* rustgo_current_activity();
extern int rustgo_is_screen_on();
extern void rustgo_free_string(char* s);
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type CGOClient struct {
	initialized bool
}

func NewCGOClient() *CGOClient {
	c := &CGOClient{}
	if C.rustgo_init() == 1 {
		c.initialized = true
		fmt.Println("[rustGO] Rust SO initialized via CGO")
	} else {
		fmt.Println("[rustGO] Rust SO init failed")
	}
	return c
}

func (c *CGOClient) Close() {
	if c.initialized {
		C.rustgo_cleanup()
		c.initialized = false
	}
}

func (c *CGOClient) HealthCheck() bool {
	return c.initialized
}

func (c *CGOClient) Tap(x, y int) error {
	if !c.initialized {
		return fmt.Errorf("not initialized")
	}
	if C.rustgo_tap(C.int(x), C.int(y)) == 0 {
		return fmt.Errorf("tap(%d,%d) failed", x, y)
	}
	return nil
}

func (c *CGOClient) LongPress(x, y, durationMs int) error {
	if !c.initialized {
		return fmt.Errorf("not initialized")
	}
	if C.rustgo_long_press(C.int(x), C.int(y), C.int(durationMs)) == 0 {
		return fmt.Errorf("long_press(%d,%d,%dms) failed", x, y, durationMs)
	}
	return nil
}

func (c *CGOClient) Swipe(x1, y1, x2, y2, durationMs int) error {
	if !c.initialized {
		return fmt.Errorf("not initialized")
	}
	if C.rustgo_swipe(C.int(x1), C.int(y1), C.int(x2), C.int(y2), C.int(durationMs)) == 0 {
		return fmt.Errorf("swipe failed")
	}
	return nil
}

func (c *CGOClient) TouchDown(x, y, fingerID int) error {
	if !c.initialized {
		return fmt.Errorf("not initialized")
	}
	if C.rustgo_touch_down(C.int(x), C.int(y), C.int(fingerID)) == 0 {
		return fmt.Errorf("touch_down failed")
	}
	return nil
}

func (c *CGOClient) TouchMove(x, y, fingerID int) error {
	if !c.initialized {
		return fmt.Errorf("not initialized")
	}
	if C.rustgo_touch_move(C.int(x), C.int(y), C.int(fingerID)) == 0 {
		return fmt.Errorf("touch_move failed")
	}
	return nil
}

func (c *CGOClient) TouchUp(x, y, fingerID int) error {
	if !c.initialized {
		return fmt.Errorf("not initialized")
	}
	if C.rustgo_touch_up(C.int(x), C.int(y), C.int(fingerID)) == 0 {
		return fmt.Errorf("touch_up failed")
	}
	return nil
}

func (c *CGOClient) InputText(text string) error {
	if !c.initialized {
		return fmt.Errorf("not initialized")
	}
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	if C.rustgo_input_text(ctext) == 0 {
		return fmt.Errorf("input_text failed")
	}
	return nil
}

func (c *CGOClient) KeyEvent(keyCode int) error {
	if !c.initialized {
		return fmt.Errorf("not initialized")
	}
	if C.rustgo_key_event(C.int(keyCode)) == 0 {
		return fmt.Errorf("key_event(%d) failed", keyCode)
	}
	return nil
}

func (c *CGOClient) Screenshot(path string) error {
	if !c.initialized {
		return fmt.Errorf("not initialized")
	}
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	if C.rustgo_screenshot(cpath) == 0 {
		return fmt.Errorf("screenshot failed")
	}
	return nil
}

func (c *CGOClient) GetScreenSize() (string, error) {
	if !c.initialized {
		return "", fmt.Errorf("not initialized")
	}
	cstr := C.rustgo_get_screen_size()
	if cstr == nil {
		return "", fmt.Errorf("get_screen_size failed")
	}
	defer C.rustgo_free_string(cstr)
	return C.GoString(cstr), nil
}

func (c *CGOClient) LaunchApp(pkg string) error {
	if !c.initialized {
		return fmt.Errorf("not initialized")
	}
	cpkg := C.CString(pkg)
	defer C.free(unsafe.Pointer(cpkg))
	if C.rustgo_launch_app(cpkg) == 0 {
		return fmt.Errorf("launch_app(%s) failed", pkg)
	}
	return nil
}

func (c *CGOClient) CloseApp(pkg string) error {
	if !c.initialized {
		return fmt.Errorf("not initialized")
	}
	cpkg := C.CString(pkg)
	defer C.free(unsafe.Pointer(cpkg))
	if C.rustgo_close_app(cpkg) == 0 {
		return fmt.Errorf("close_app(%s) failed", pkg)
	}
	return nil
}

func (c *CGOClient) Shell(cmd string) (string, error) {
	if !c.initialized {
		return "", fmt.Errorf("not initialized")
	}
	ccmd := C.CString(cmd)
	defer C.free(unsafe.Pointer(ccmd))
	cstr := C.rustgo_shell_exec(ccmd)
	if cstr == nil {
		return "", fmt.Errorf("shell_exec failed")
	}
	defer C.rustgo_free_string(cstr)
	return C.GoString(cstr), nil
}

func (c *CGOClient) FileExists(path string) bool {
	if !c.initialized {
		return false
	}
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	return C.rustgo_file_exists(cpath) == 1
}

func (c *CGOClient) FileRead(path string) (string, error) {
	if !c.initialized {
		return "", fmt.Errorf("not initialized")
	}
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	cstr := C.rustgo_file_read(cpath)
	if cstr == nil {
		return "", fmt.Errorf("file_read failed")
	}
	defer C.rustgo_free_string(cstr)
	return C.GoString(cstr), nil
}

func (c *CGOClient) FileWrite(path, content string) error {
	if !c.initialized {
		return fmt.Errorf("not initialized")
	}
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	ccontent := C.CString(content)
	defer C.free(unsafe.Pointer(ccontent))
	if C.rustgo_file_write(cpath, ccontent) == 0 {
		return fmt.Errorf("file_write failed")
	}
	return nil
}

func (c *CGOClient) DeviceInfo() (string, error) {
	if !c.initialized {
		return "", fmt.Errorf("not initialized")
	}
	cstr := C.rustgo_device_info()
	if cstr == nil {
		return "", fmt.Errorf("device_info failed")
	}
	defer C.rustgo_free_string(cstr)
	return C.GoString(cstr), nil
}

func (c *CGOClient) Wait(ms int) error {
	if !c.initialized {
		return fmt.Errorf("not initialized")
	}
	C.rustgo_wait(C.int(ms))
	return nil
}

func (c *CGOClient) GetClipboard() (string, error) {
	if !c.initialized {
		return "", fmt.Errorf("not initialized")
	}
	cstr := C.rustgo_get_clipboard()
	if cstr == nil {
		return "", fmt.Errorf("get_clipboard failed")
	}
	defer C.rustgo_free_string(cstr)
	return C.GoString(cstr), nil
}

func (c *CGOClient) SetClipboard(text string) error {
	if !c.initialized {
		return fmt.Errorf("not initialized")
	}
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	if C.rustgo_set_clipboard(ctext) == 0 {
		return fmt.Errorf("set_clipboard failed")
	}
	return nil
}

func (c *CGOClient) CurrentPackage() (string, error) {
	if !c.initialized {
		return "", fmt.Errorf("not initialized")
	}
	cstr := C.rustgo_current_package()
	if cstr == nil {
		return "", fmt.Errorf("current_package failed")
	}
	defer C.rustgo_free_string(cstr)
	return C.GoString(cstr), nil
}

func (c *CGOClient) CurrentActivity() (string, error) {
	if !c.initialized {
		return "", fmt.Errorf("not initialized")
	}
	cstr := C.rustgo_current_activity()
	if cstr == nil {
		return "", fmt.Errorf("current_activity failed")
	}
	defer C.rustgo_free_string(cstr)
	return C.GoString(cstr), nil
}

func (c *CGOClient) IsScreenOn() bool {
	if !c.initialized {
		return false
	}
	return C.rustgo_is_screen_on() == 1
}
