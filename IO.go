package imgui

// #include "wrapper/IO.h"
import "C"

var keyMap = map[int]int{}

// Taken from https://github.com/JetSetIlly/imgui-go/IO.go

// BackendFlags for IO.AddKeyEvent.
type ImguiKey int

// A key identifier (ImguiKey_XXX or ImGuiMod_XXX value): can represent Keyboard, Mouse and Gamepad values.
// All our named keys are >= 512. Keys value 0 to 511 are left unused and were legacy native/opaque key values (< 1.87).
// Support for legacy keys was completely removed in 1.91.5.
// Read details about the 1.87+ transition : https://github.com/ocornut/imgui/issues/4921
// Note that "Keys" related to physical keys and are not the same concept as input "Characters", the later are submitted via io.AddInputCharacter().
// The keyboard key enum values are named after the keys on a standard US keyboard, and on other keyboard types the keys reported may not match the keycaps.
const (
	KeyNone           ImguiKey = 0
	KeyNamedKey_BEGIN ImguiKey = 512 // First valid key value (other than 0)
)


const (
	KeyTab ImguiKey = iota + KeyNamedKey_BEGIN // == ImguiKeyNamedKey_BEGIN
	KeyLeftArrow
	KeyRightArrow
	KeyUpArrow
	KeyDownArrow
	KeyPageUp
	KeyPageDown
	KeyHome
	KeyEnd
	KeyInsert
	KeyDelete
	KeyBackspace
	KeySpace
	KeyEnter
	KeyEscape
	KeyLeftCtrl
	KeyLeftShift
	KeyLeftAlt
	KeyLeftSuper
	KeyRightCtrl
	KeyRightShift
	KeyRightAlt
	KeyRightSuper
	KeyMenu
	Key0
	Key1
	Key2
	Key3
	Key4
	Key5
	Key6
	Key7
	Key8
	Key9
	KeyA
	KeyB
	KeyC
	KeyD
	KeyE
	KeyF
	KeyG
	KeyH
	KeyI
	KeyJ
	KeyK
	KeyL
	KeyM
	KeyN
	KeyO
	KeyP
	KeyQ
	KeyR
	KeyS
	KeyT
	KeyU
	KeyV
	KeyW
	KeyX
	KeyY
	KeyZ
	KeyF1
	KeyF2
	KeyF3
	KeyF4
	KeyF5
	KeyF6
	KeyF7
	KeyF8
	KeyF9
	KeyF10
	KeyF11
	KeyF12
	KeyF13
	KeyF14
	KeyF15
	KeyF16
	KeyF17
	KeyF18
	KeyF19
	KeyF20
	KeyF21
	KeyF22
	KeyF23
	KeyF24
	KeyApostrophe   // '
	KeyComma        //
	KeyMinus        // -
	KeyPeriod       // .
	KeySlash        // /
	KeySemicolon    // ;
	KeyEqual        // =
	KeyLeftBracket  // [
	KeyBackslash    // \ (this text inhibit multiline comment caused by backslash)
	KeyRightBracket // ]
	KeyGraveAccent  // `
	KeyCapsLock
	KeyScrollLock
	KeyNumLock
	KeyPrintScreen
	KeyPause
	KeyKeypad0
	KeyKeypad1
	KeyKeypad2
	KeyKeypad3
	KeyKeypad4
	KeyKeypad5
	KeyKeypad6
	KeyKeypad7
	KeyKeypad8
	KeyKeypad9
	KeyKeypadDecimal
	KeyKeypadDivide
	KeyKeypadMultiply
	KeyKeypadSubtract
	KeyKeypadAdd
	KeyKeypadEnter
	KeyKeypadEqual
	KeyAppBack // Available on some keyboard/mouses. Often referred as "Browser Back"
	KeyAppForward
	KeyOem102 // Non-US backslash.

	// Gamepad (some of those are analog values, 0.0f to 1.0f)                          // NAVIGATION ACTION
	// (download controller mapping PNG/PSD at http://dearimgui.com/controls_sheets)
	KeyGamepadStart       // Menu (Xbox)      + (Switch)   Start/Options (PS)
	KeyGamepadBack        // View (Xbox)      - (Switch)   Share (PS)
	KeyGamepadFaceLeft    // X (Xbox)         Y (Switch)   Square (PS)        // Tap: Toggle Menu. Hold: Windowing mode (Focus/Move/Resize windows)
	KeyGamepadFaceRight   // B (Xbox)         A (Switch)   Circle (PS)        // Cancel / Close / Exit
	KeyGamepadFaceUp      // Y (Xbox)         X (Switch)   Triangle (PS)      // Text Input / On-screen Keyboard
	KeyGamepadFaceDown    // A (Xbox)         B (Switch)   Cross (PS)         // Activate / Open / Toggle / Tweak
	KeyGamepadDpadLeft    // D-pad Left                                       // Move / Tweak / Resize Window (in Windowing mode)
	KeyGamepadDpadRight   // D-pad Right                                      // Move / Tweak / Resize Window (in Windowing mode)
	KeyGamepadDpadUp      // D-pad Up                                         // Move / Tweak / Resize Window (in Windowing mode)
	KeyGamepadDpadDown    // D-pad Down                                       // Move / Tweak / Resize Window (in Windowing mode)
	KeyGamepadL1          // L Bumper (Xbox)  L (Switch)   L1 (PS)            // Tweak Slower / Focus Previous (in Windowing mode)
	KeyGamepadR1          // R Bumper (Xbox)  R (Switch)   R1 (PS)            // Tweak Faster / Focus Next (in Windowing mode)
	KeyGamepadL2          // L Trig. (Xbox)   ZL (Switch)  L2 (PS) [Analog]
	KeyGamepadR2          // R Trig. (Xbox)   ZR (Switch)  R2 (PS) [Analog]
	KeyGamepadL3          // L Stick (Xbox)   L3 (Switch)  L3 (PS)
	KeyGamepadR3          // R Stick (Xbox)   R3 (Switch)  R3 (PS)
	KeyGamepadLStickLeft  // [Analog]                                         // Move Window (in Windowing mode)
	KeyGamepadLStickRight // [Analog]                                         // Move Window (in Windowing mode)
	KeyGamepadLStickUp    // [Analog]                                         // Move Window (in Windowing mode)
	KeyGamepadLStickDown  // [Analog]                                         // Move Window (in Windowing mode)
	KeyGamepadRStickLeft  // [Analog]
	KeyGamepadRStickRight // [Analog]
	KeyGamepadRStickUp    // [Analog]
	KeyGamepadRStickDown  // [Analog]

	// Aliases: Mouse Buttons (auto-submitted from AddMouseButtonEvent() calls)
	// - This is mirroring the data also written to io.MouseDown[], io.MouseWheel, in a format allowing them to be accessed via standard key API.
	KeyMouseLeft
	KeyMouseRight
	KeyMouseMiddle
	KeyMouseX1
	KeyMouseX2
	KeyMouseWheelX
	KeyMouseWheelY

	// [Internal] Reserved for mod storage
	KeyReservedForModCtrl
	KeyReservedForModShift
	KeyReservedForModAlt
	KeyReservedForModSuper
	KeyNamedKey_END
)

const (
	// Keyboard Modifiers (explicitly submitted by backend via AddKeyEvent() calls)
	// - This is mirroring the data also written to io.KeyCtrl, io.KeyShift, io.KeyAlt, io.KeySuper, in a format allowing
	//   them to be accessed via standard key API, allowing calls such as IsKeyPressed(), IsKeyReleased(), querying duration etc.
	// - Code polling every key (e.g. an interface to detect a key press for input mapping) might want to ignore those
	//   and prefer using the real keys (e.g. KeyLeftCtrl, ImguiKeyRightCtrl instead of ImGuiMod_Ctrl).
	// - In theory the value of keyboard modifiers should be roughly equivalent to a logical or of the equivalent left/right keys.
	//   In practice: it's complicated; mods are often provided from different sources. Keyboard layout, IME, sticky keys and
	//   backends tend to interfere and break that equivalence. The safer decision is to relay that ambiguity down to the end-user...
	// - On macOS, we swap Cmd(Super) and Ctrl keys at the time of the io.AddKeyEvent() call.
	KeyModNone  ImguiKey = 0
	KeyModCtrl  ImguiKey = 1 << 12 // Ctrl (non-macOS), Cmd (macOS)
	KeyModShift ImguiKey = 1 << 13 // Shift
	KeyModAlt   ImguiKey = 1 << 14 // Option/Menu
	KeyModSuper ImguiKey = 1 << 15 // Windows/Super (non-macOS), Ctrl (macOS)
)

// IO is where your app communicate with ImGui. Access via CurrentIO().
// Read 'Programmer guide' section in imgui.cpp file for general usage.
type IO struct {
	windowHeight int
	windowWidth  int
	handle       C.IggIO
}

// CurrentIO returns access to the ImGui communication struct for the currently active context.
func CurrentIO() IO {
	return IO{handle: C.iggGetCurrentIO()}
}

func (io *IO) SetKey(imguiKey int, nativeKey int) {
	keyMap[imguiKey] = nativeKey
}

func (io IO) GetKey(imguiKey int) int {
	return keyMap[imguiKey]
}

func setKey(imguiKey int, nativeKey int) {
	keyMap[imguiKey] = nativeKey
}

func GetKey(imguiKey int) int {
	return keyMap[imguiKey]
}

// WantCaptureMouse returns true if imgui will use the mouse inputs.
// Do not dispatch them to your main game/application in this case.
// In either case, always pass on mouse inputs to imgui.
//
// e.g. unclicked mouse is hovering over an imgui window, widget is active,
// mouse was clicked over an imgui window, etc.
func (io IO) WantCaptureMouse() bool {
	return C.iggWantCaptureMouse(io.handle) != 0
}

// WantCaptureMouseUnlessPopupClose returns true if imgui will use the mouse inputs.
// Alternative to WantCaptureMouse: (WantCaptureMouse == true &&
// WantCaptureMouseUnlessPopupClose == false) when a click over void is
// expected to close a popup.
func (io IO) WantCaptureMouseUnlessPopupClose() bool {
	return C.iggWantCaptureMouseUnlessPopupClose(io.handle) != 0
}

// WantCaptureKeyboard returns true if imgui will use the keyboard inputs.
// Do not dispatch them to your main game/application (in both cases, always pass keyboard inputs to imgui).
//
// e.g. InputText active, or an imgui window is focused and navigation is enabled, etc.
func (io IO) WantCaptureKeyboard() bool {
	return C.iggWantCaptureKeyboard(io.handle) != 0
}

// WantTextInput is true, you may display an on-screen keyboard.
// This is set by ImGui when it wants textual keyboard input to happen (e.g. when a InputText widget is active).
func (io IO) WantTextInput() bool {
	return C.iggWantTextInput(io.handle) != 0
}

// Framerate application estimation, in frame per second. Solely for convenience.
// Rolling average estimation based on IO.DeltaTime over 120 frames.
func (io IO) Framerate() float32 {
	return float32(C.iggFramerate(io.handle))
}

// MetricsRenderVertices returns vertices output during last call to Render().
func (io IO) MetricsRenderVertices() int {
	return int(C.iggMetricsRenderVertices(io.handle))
}

// MetricsRenderIndices returns indices output during last call to Render() = number of triangles * 3.
func (io IO) MetricsRenderIndices() int {
	return int(C.iggMetricsRenderIndices(io.handle))
}

// MetricsRenderWindows returns number of visible windows.
func (io IO) MetricsRenderWindows() int {
	return int(C.iggMetricsRenderWindows(io.handle))
}

// MetricsActiveWindows returns number of active windows.
func (io IO) MetricsActiveWindows() int {
	return int(C.iggMetricsActiveWindows(io.handle))
}

// MetricsActiveAllocations returns number of active allocations, updated by MemAlloc/MemFree
// based on current context. May be off if you have multiple imgui contexts.
func (io IO) MetricsActiveAllocations() int {
	return int(C.iggMetricsActiveAllocations(io.handle))
}

// MousePosition returns the mouse position.
func (io IO) MousePosition() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggIoGetMousePosition(io.handle, valueArg)
	valueFin()
	return value
}

// MouseDelta returns the mouse delta movement. Note that this is zero if either current or previous position
// are invalid (-math.MaxFloat32,-math.MaxFloat32), so a disappearing/reappearing mouse won't have a huge delta.
func (io IO) MouseDelta() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggMouseDelta(io.handle, valueArg)
	valueFin()
	return value
}

// MouseWheel returns the mouse wheel movement.
func (io IO) MouseWheel() (float32, float32) {
	var mouseWheelH, mouseWheel C.float
	C.iggMouseWheel(io.handle, &mouseWheelH, &mouseWheel)
	return float32(mouseWheelH), float32(mouseWheel)
}

// DisplayFrameBufferScale returns scale factor for HDPI displays.
// It is for retina display or other situations where window coordinates are different from framebuffer coordinates.
func (io IO) DisplayFrameBufferScale() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggDisplayFrameBufferScale(io.handle, valueArg)
	valueFin()
	return value
}

// SetDisplaySize sets the size in pixels.
func (io IO) SetDisplaySize(value Vec2) {
	out, _ := value.wrapped()
	C.iggIoSetDisplaySize(io.handle, out)
}

// SetDisplayFrameBufferScale sets the frame buffer scale factor.
func (io IO) SetDisplayFrameBufferScale(value Vec2) {
	out, _ := value.wrapped()
	C.iggIoSetDisplayFrameBufferScale(io.handle, out)
}

// Fonts returns the font atlas to load and assemble one or more fonts into a single tightly packed texture.
func (io IO) Fonts() FontAtlas {
	return FontAtlas(C.iggIoGetFonts(io.handle))
}

// SetMousePosition sets the mouse position, in pixels.
// Set to Vec2(-math.MaxFloat32,-mathMaxFloat32) if mouse is unavailable (on another screen, etc.).
func (io IO) SetMousePosition(value Vec2) {
	posArg, _ := value.wrapped()
	C.iggIoSetMousePosition(io.handle, posArg)
}

// SetMouseButtonDown sets whether a specific mouse button is currently pressed.
// Mouse buttons: left, right, middle + extras.
// ImGui itself mostly only uses left button (BeginPopupContext** are using right button).
// Other buttons allows us to track if the mouse is being used by your application +
// available to user as a convenience via IsMouse** API.
func (io IO) SetMouseButtonDown(index int, down bool) {
	var downArg C.IggBool
	if down {
		downArg = 1
	}
	C.iggIoSetMouseButtonDown(io.handle, C.int(index), downArg)
}

// AddMouseWheelDelta adds the given offsets to the current mouse wheel values.
// 1 vertical unit scrolls about 5 lines text.
// Most users don't have a mouse with an horizontal wheel, may not be provided by all back-ends.
func (io IO) AddMouseWheelDelta(horizontal, vertical float32) {
	C.iggIoAddMouseWheelDelta(io.handle, C.float(horizontal), C.float(vertical))
}

// SetDeltaTime sets the time elapsed since last frame, in seconds.
func (io IO) SetDeltaTime(value float32) {
	C.iggIoSetDeltaTime(io.handle, C.float(value))
}

// SetFontGlobalScale sets the global scaling factor for all fonts.
func (io IO) SetFontGlobalScale(value float32) {
	C.iggIoSetFontGlobalScale(io.handle, C.float(value))
}

// KeyPress sets the KeysDown flag.
func (io IO) KeyPress(key int) {
	C.iggIoKeyPress(io.handle, C.int(key))
}

// KeyRelease clears the KeysDown flag.
func (io IO) KeyRelease(key int) {
	C.iggIoKeyRelease(io.handle, C.int(key))
}

// Constants to fill IO.KeyMap() lookup with indices into the IO.KeysDown[512] array.
// The mapped indices are then the ones reported to IO.KeyPress() and IO.KeyRelease().
/*
const (
	KeyTab         = 0
	KeyLeftArrow   = 1
	KeyRightArrow  = 2
	KeyUpArrow     = 3
	KeyDownArrow   = 4
	KeyPageUp      = 5
	KeyPageDown    = 6
	KeyHome        = 7
	KeyEnd         = 8
	KeyInsert      = 9
	KeyDelete      = 10
	KeyBackspace   = 11
	KeySpace       = 12
	KeyEnter       = 13
	KeyEscape      = 14
	KeyKeyPadEnter = 15
	KeyA           = 16 // for text edit CTRL+A: select all
	KeyC           = 17 // for text edit CTRL+C: copy
	KeyV           = 18 // for text edit CTRL+V: paste
	KeyX           = 19 // for text edit CTRL+X: cut
	KeyY           = 20 // for text edit CTRL+Y: redo
	KeyZ           = 21 // for text edit CTRL+Z: undo

	// Custom keys, not mapped by imgui
	KeyW = 22
	KeyS = 23
	KeyD = 24
	Key1 = 25
)
*/

// AddKeyEvents adds the key event (up or down) to the event queue.
func (io IO) AddKeyEvent(key ImguiKey, down bool) {
	var downArg C.IggBool
	if down {
		downArg = 1
	}
	C.iggIoAddKeyEvent(io.handle, C.int(int(key)), downArg)
}

// KeyMap maps a key into the KeysDown array which represents your "native" keyboard state.
func (io IO) KeyMap(imguiKey int, nativeKey int) {
	C.iggIoKeyMap(io.handle, C.int(imguiKey), C.int(nativeKey))
}

// KeyCtrl sets the keyboard modifier control pressed.
func (io IO) KeyCtrl(leftCtrl int, rightCtrl int) {
	C.iggIoKeyCtrl(io.handle, C.int(leftCtrl), C.int(rightCtrl))
}

// KeyCtrlPressed get the keyboard modifier control pressed.
func (io IO) KeyCtrlPressed() bool {
	return C.iggIoKeyCtrlPressed(io.handle) != 0
}

// KeyShift sets the keyboard modifier shift pressed.
func (io IO) KeyShift(leftShift int, rightShift int) {
	C.iggIoKeyShift(io.handle, C.int(leftShift), C.int(rightShift))
}

// KeyShiftPressed get the keyboard modifier shif pressed.
func (io IO) KeyShiftPressed() bool {
	return C.iggIoKeyShiftPressed(io.handle) != 0
}

// KeyAlt sets the keyboard modifier alt pressed.
func (io IO) KeyAlt(leftAlt int, rightAlt int) {
	C.iggIoKeyAlt(io.handle, C.int(leftAlt), C.int(rightAlt))
}

// KeyAltPressed get the keyboard modifier alt pressed.
func (io IO) KeyAltPressed() bool {
	return C.iggIoKeyAltPressed(io.handle) != 0
}

// KeySuper sets the keyboard modifier super pressed.
func (io IO) KeySuper(leftSuper int, rightSuper int) {
	C.iggIoKeySuper(io.handle, C.int(leftSuper), C.int(rightSuper))
}

// KeySuperPressed get the keyboard modifier super pressed.
func (io IO) KeySuperPressed() bool {
	return C.iggIoKeySuperPressed(io.handle) != 0
}

// AddInputCharacters adds a new character into InputCharacters[].
func (io IO) AddInputCharacters(chars string) {
	textArg, textFin := wrapString(chars)
	defer textFin()
	C.iggIoAddInputCharactersUTF8(io.handle, textArg)
}

// SetIniFilename changes the filename for the settings. Default: "imgui.ini".
// Use an empty string to disable the ini from being used.
func (io IO) SetIniFilename(value string) {
	valueArg, valueFin := wrapString(value)
	defer valueFin()
	C.iggIoSetIniFilename(io.handle, valueArg)
}

// ConfigFlags for IO.SetConfigFlags.
type ConfigFlags int

const (
	// ConfigFlagsNone default = 0.
	ConfigFlagsNone ConfigFlags = 0
	// ConfigFlagsNavEnableKeyboard main keyboard navigation enable flag. NewFrame() will automatically fill
	// io.NavInputs[] based on io.KeysDown[].
	ConfigFlagsNavEnableKeyboard ConfigFlags = 1 << 0
	// ConfigFlagsNavEnableGamepad main gamepad navigation enable flag.
	// This is mostly to instruct your imgui back-end to fill io.NavInputs[]. Back-end also needs to set
	// BackendFlagHasGamepad.
	ConfigFlagsNavEnableGamepad ConfigFlags = 1 << 1
	// ConfigFlagsNavEnableSetMousePos instruct navigation to move the mouse cursor. May be useful on TV/console systems
	// where moving a virtual mouse is awkward. Will update io.MousePos and set io.WantSetMousePos=true. If enabled you
	// MUST honor io.WantSetMousePos requests in your binding, otherwise ImGui will react as if the mouse is jumping
	// around back and forth.
	ConfigFlagsNavEnableSetMousePos ConfigFlags = 1 << 2
	// ConfigFlagsNavNoCaptureKeyboard instruct navigation to not set the io.WantCaptureKeyboard flag when io.NavActive
	// is set.
	ConfigFlagsNavNoCaptureKeyboard ConfigFlags = 1 << 3
	// ConfigFlagsNoMouse instruct imgui to clear mouse position/buttons in NewFrame(). This allows ignoring the mouse
	// information set by the back-end.
	ConfigFlagsNoMouse ConfigFlags = 1 << 4
	// ConfigFlagsNoMouseCursorChange instruct back-end to not alter mouse cursor shape and visibility. Use if the
	// back-end cursor changes are interfering with yours and you don't want to use SetMouseCursor() to change mouse
	// cursor. You may want to honor requests from imgui by reading GetMouseCursor() yourself instead.
	ConfigFlagsNoMouseCursorChange ConfigFlags = 1 << 5
	// ConfigFlagsNoKeyboard Instruct dear imgui to disable keyboard inputs and interactions. This is done by ignoring
	// keyboard events and clearing existing states.
	ConfigFlagsNoKeyboard ConfigFlags = 1 << 6
	// ConfigFlagsDockingEnable instruct imgui to enable docking features.
	ConfigFlagsDockingEnable ConfigFlags = 1 << 7

	// User storage (to allow your back-end/engine to communicate to code that may be shared between multiple projects.
	// Those flags are not used by core Dear ImGui).

	// ConfigFlagsIsSRGB application is SRGB-aware.
	ConfigFlagsIsSRGB ConfigFlags = 1 << 20
	// ConfigFlagsIsTouchScreen application is using a touch screen instead of a mouse.
	ConfigFlagsIsTouchScreen ConfigFlags = 1 << 21
)

// SetConfigFlags sets the gamepad/keyboard navigation options, etc.
func (io IO) SetConfigFlags(flags ConfigFlags) {
	C.iggIoSetConfigFlags(io.handle, C.int(flags))
}

// BackendFlags for IO.SetBackendFlags.
type BackendFlags int

const (
	// BackendFlagsNone default = 0.
	BackendFlagsNone BackendFlags = 0
	// BackendFlagsHasGamepad back-end Platform supports gamepad and currently has one connected.
	BackendFlagsHasGamepad BackendFlags = 1 << 0
	// BackendFlagsHasMouseCursors back-end Platform supports honoring GetMouseCursor() value to change the OS cursor
	// shape.
	BackendFlagsHasMouseCursors BackendFlags = 1 << 1
	// BackendFlagsHasSetMousePos back-end Platform supports io.WantSetMousePos requests to reposition the OS mouse
	// position (only used if ImGuiConfigFlags_NavEnableSetMousePos is set).
	BackendFlagsHasSetMousePos BackendFlags = 1 << 2
	// BackendFlagsRendererHasVtxOffset back-end Renderer supports ImDrawCmd::VtxOffset. This enables output of large
	// meshes (64K+ vertices) while still using 16-bits indices.
	BackendFlagsRendererHasVtxOffset BackendFlags = 1 << 3
)

// SetBackendFlags sets back-end capabilities.
func (io IO) SetBackendFlags(flags BackendFlags) {
	C.iggIoSetBackendFlags(io.handle, C.int(flags))
}

// GetBackendFlags gets the current backend flags.
func (io IO) GetBackendFlags() BackendFlags {
	return BackendFlags(C.iggIoGetBackendFlags(io.handle))
}

// SetMouseDrawCursor request ImGui to draw a mouse cursor for you (if you are on a platform without a mouse cursor).
func (io IO) SetMouseDrawCursor(show bool) {
	C.iggIoSetMouseDrawCursor(io.handle, castBool(show))
}

// Clipboard describes the access to the text clipboard of the window manager.
type Clipboard interface {
	// Text returns the current text from the clipboard, if available.
	Text() (string, error)
	// SetText sets the text as the current text on the clipboard.
	SetText(value string)
}

var clipboards = map[C.IggIO]Clipboard{}
var dropLastClipboardText = func() {}

// SetClipboard registers a clipboard for text copy/paste actions.
// If no clipboard is set, then a fallback implementation may be used, if available for the OS.
// To disable clipboard handling overall, pass nil as the Clipboard.
//
// Since ImGui queries the clipboard text via a return value, the wrapper has to hold the
// current clipboard text as a copy in memory. This memory will be freed at the next clipboard operation.
func (io IO) SetClipboard(board Clipboard) {
	dropLastClipboardText()

	if board != nil {
		clipboards[io.handle] = board
		C.iggIoRegisterClipboardFunctions(io.handle)
	} else {
		C.iggIoClearClipboardFunctions(io.handle)
		delete(clipboards, io.handle)
	}
}

func (io IO) Clipboard() Clipboard {
	return clipboards[io.handle]
}

//export iggIoGetClipboardText
func iggIoGetClipboardText(handle C.IggIO) *C.char {
	dropLastClipboardText()

	board := clipboards[handle]
	if board == nil {
		return nil
	}

	text, err := board.Text()
	if err != nil {
		return nil
	}
	textPtr, textFin := wrapString(text)
	dropLastClipboardText = func() {
		dropLastClipboardText = func() {}
		textFin()
	}
	return textPtr
}

//export iggIoSetClipboardText
func iggIoSetClipboardText(handle C.IggIO, text *C.char) {
	dropLastClipboardText()

	board := clipboards[handle]
	if board == nil {
		return
	}
	board.SetText(C.GoString(text))
}
