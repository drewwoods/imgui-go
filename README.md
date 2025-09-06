# Dear ImGui for Go

This library is a continuation of the [Go](https://www.golang.org) wrapper for **[Dear ImGui](https://github.com/ocornut/imgui)** originally created at https://github.com/inkyblackness/imgui-go.

This fork has been updated to Dear ImGui v1.90.9 with **docking support** and exposes additional functionality, building upon the excellent foundation of the original project.

This wrapper started as a special-purpose wrapper for use within InkyBlackness.
However, it is self-contained and can be used for other purposes as well.

This wrapper is
* hand-crafted, for Go
* documented
* versioned
* with ported examples in a separate repository (see below)

![Screenshot from example](assets/screenshot.png)

## API naming

Names of types and functions follow closely those of **Dear ImGui**.

For functions that have optional parameters, the following schema is applied:
* There is the "verbose" variant, followed by the letter `V`, such as `ButtonV(id string, size Vec2) bool`
* Next to it there is the "idiomatic" variant, without any optional parameter, such as `Button(id string) bool`.
* The idiomatic variant calls the verbose variant with the default values for the optional parameters.
Functions that don't have optional parameters don't come in a verbose variant.

The **Dear ImGui** functions `IO()` and `Style()` have been renamed to be `CurrentIO()` and `CurrentStyle()`.
This was done because their returned types have the same name, causing a name clash.
With the `Current` prefix, they also better describe what they return.

## API philosophy
This library does not intend to export all the functions of the wrapped **Dear ImGui**. The following filter applies as a rule of thumb:
* Functions marked as "obsolete" are not available. (The corresponding C code isn't even compiled - disabled by define)
* "Shortcut" Functions, which combine language features and/or other **Dear ImGui** functions, are not available. There may be exceptions for this if the shortcut exists in Dear ImGui code base (e.g. `ImGui::Text` which does printf style string formatting)
* Functions that are not needed by InkyBlackness are ignored. This doesn't mean that they can't be in the wrapper, they are simply not a priority. Feel free to propose an implementation or make a pull request, respecting the previous points :)

## Version philosophy
This library does not mirror the versions of the wrapped **Dear ImGui**. The semantic versioning of this wrapper is defined as:
* Major changes: (Breaking) changes in API or behaviour. Typically done through changes in **Dear ImGui**.
* Minor changes: Extensions in API. Typically done through small version increments of **Dear ImGui** and/or exposing further features in a compatible way.
* Patch changes: Bug fixes - either in the wrapper or the wrapped **Dear ImGui**, given that the API & behaviour remains the same.

At the moment, this library uses version [1.90.9](https://github.com/ocornut/imgui/releases/tag/v1.90.9) of **Dear ImGui** with **docking support** enabled.

## Breaking Changes from v1.85

This major version update from Dear ImGui v1.85 to v1.90.9 includes some breaking changes:
* **KeyMap deprecation**: `io.KeyMap[]` and `io.KeysDown[]` are deprecated since v1.87 in favor of `io.AddKeyEvent()`. The legacy KeyMap system still works but will be removed when `IMGUI_DISABLE_OBSOLETE_KEYIO` is defined. 
  - Old: `ImGui::IsKeyPressed(ImGui::GetIO().KeyMap[ImGuiKey_Space])` 
  - New: `ImGui::IsKeyPressed(ImGuiKey_Space)`
* Some function signatures may have changed
* Window flag bit positions have been corrected to match ImGui v1.90.9  
* Applications using docking features will need to enable docking in their ImGui configuration
* New features and improvements may affect existing behavior

For detailed migration information, see [Dear ImGui issue #4921](https://github.com/ocornut/imgui/issues/4921) and the [Dear ImGui changelog](https://github.com/ocornut/imgui/releases). Please test your applications thoroughly when upgrading.

## Acknowledgments

* [ocornut/imgui](https://github.com/ocornut/imgui) - The original Dear ImGui library that makes all of this possible
* [inkyblackness/imgui-go](https://github.com/inkyblackness/imgui-go) - The original excellent Go wrapper that this project builds upon
* [JetSetIlly/imgui-go](https://github.com/JetSetIlly/imgui-go/blob/master/IO.go) for the key identifiers implementation, which was used as a reference for updating the keyboard input system in this version.

## Examples
A separate repository was created to host ported examples and reference implementations.
See repository [inkyblackness/imgui-go-examples](https://github.com/inkyblackness/imgui-go-examples).

It contains reference implementations for libraries such as [GLFW3](https://github.com/go-gl/glfw) and [SDL2](https://github.com/veandco/go-sdl2), using [OpenGL](https://github.com/go-gl/gl).

The screenshot above was created with such an example.

## Docking Support

This version includes full **docking support**, allowing you to create professional IDE-like interfaces with:
* **DockSpace**: Create dockable areas within your application windows  
* **Tabbed windows**: Multiple windows that start docked as tabs but can be undocked
* **Flexible layouts**: Users can drag and rearrange window layouts at runtime
* **Persistent layouts**: Docking configurations are saved automatically

### Key Docking Functions
* `DockSpace(id, size, flags)` - Create a dockspace within an existing window
* `DockSpaceOverViewport(id, viewport, flags)` - Create a full-screen dockspace
* `SetNextWindowDockID(dockID, cond)` - Dock a window to a specific dock node
* `WindowFlagsNoDocking` - Prevent specific windows from being docked

See the `examples/` directory for usage examples and detailed documentation.

## Extras

### FreeType font rendering

If the `FreeType` library is available for your platform, you can enable using it with the build tag `imguifreetype` - as in
```
go build -tags="imguifreetype"
```
This extra is based on the reference implementation from **Dear ImGui**.

If you set the build tag, yet the corresponding support has not been added to the library, you will receive a build error.
Contributions to support more build environments are happily accepted. See file `FreeType.go`.

> If you are trying to do this on MS Windows with MinGW and receive an error like
> `pkg-config: exec: "pkg-config": executable file not found in %PATH%`,
> refer to [online guides](https://stackoverflow.com/questions/1710922/how-to-install-pkg-config-in-windows) on how to add this to your installation.

## Alternatives

Since 2022-08, there is https://github.com/AllenDang/cimgui-go , which is an auto-generated wrapper that
makes it easier to be at the latest version of **Dear ImGui**. It is recommended to use that one instead. 

Before inkyblackness/imgui-go was created, the following alternatives were considered - and ignored:
* `kdrag0n/go-imgui` (no longer available). Reasons for dismissal at time of decision:
  * Auto-generated bloat, which doesn't help
  * Was using old API (1.5x)
  * Did not compile
  * Project appeared to be abandoned
* [Extrawurst/cimgui](https://github.com/Extrawurst/cimgui). Reasons for dismissal at time of decision:
  * Was using old API (1.5x), 1.6x was attempted
  * Apparently semi-exposed the C++ API, especially through the structures
  * Adding this adds another dependency
  * Note: `cimgui` has since switched to an auto-generated method. You can use that instead of this manually curated wrapper here.


## License

The project is available under the terms of the **New BSD License** (see LICENSE file).
The licenses of included sources are stored in the **_licenses** folder.
