package imgui

// #include "wrapper/DockBuilder.h"
import "C"

// DockBuilderAddNode creates a new dock node and returns its ID.
// Pass 0 for node_id to auto-generate an ID, or pass a specific ID to use.
// Use DockNodeFlags to control the dock node properties.
func DockBuilderAddNode(nodeID uint32, flags DockNodeFlags) uint32 {
	return uint32(C.iggDockBuilderAddNode(C.uint(nodeID), C.int(flags)))
}

// DockBuilderRemoveNode removes a dock node and all its children, undocking all windows.
func DockBuilderRemoveNode(nodeID uint32) {
	C.iggDockBuilderRemoveNode(C.uint(nodeID))
}

// DockBuilderSetNodeSize sets the size of a dock node.
// This should be called before splitting the node if you intend to split it immediately.
func DockBuilderSetNodeSize(nodeID uint32, size Vec2) {
	sizeArg, _ := size.wrapped()
	C.iggDockBuilderSetNodeSize(C.uint(nodeID), sizeArg)
}

// DockBuilderSplitNode splits a dock node into two child nodes.
// Returns the ID of the new node created in the split direction.
// The original node becomes the parent and contains the child nodes.
//
// Parameters:
//   - nodeID: The dock node to split
//   - splitDir: Direction to split (DirLeft, DirRight, DirUp, DirDown)
//   - sizeRatio: Size ratio for the node at the split direction (0.0 to 1.0)
//   - outIDAtDir: Receives the ID of the new node in the split direction (can be nil)
//   - outIDAtOppositeDir: Receives the ID of the node opposite the split direction (can be nil)
func DockBuilderSplitNode(nodeID uint32, splitDir Dir, sizeRatio float32, outIDAtDir, outIDAtOppositeDir *uint32) uint32 {
	var cOutIDAtDir, cOutIDAtOppositeDir *C.uint
	if outIDAtDir != nil {
		cOutIDAtDir = (*C.uint)(outIDAtDir)
	}
	if outIDAtOppositeDir != nil {
		cOutIDAtOppositeDir = (*C.uint)(outIDAtOppositeDir)
	}
	
	return uint32(C.iggDockBuilderSplitNode(C.uint(nodeID), C.int(splitDir), C.float(sizeRatio), cOutIDAtDir, cOutIDAtOppositeDir))
}

// DockBuilderDockWindow docks a window to a specific dock node.
// The window doesn't need to exist yet - this sets up the docking for when it's created.
func DockBuilderDockWindow(windowName string, nodeID uint32) {
	windowNameArg, windowNameFin := wrapString(windowName)
	defer windowNameFin()
	C.iggDockBuilderDockWindow(windowNameArg, C.uint(nodeID))
}

// DockBuilderFinish finalizes the dock builder operations.
// Call this after you've finished setting up your dock layout.
func DockBuilderFinish(nodeID uint32) {
	C.iggDockBuilderFinish(C.uint(nodeID))
}