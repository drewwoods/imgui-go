#include "ConfiguredImGui.h"

#include "Window.h"
#include "WrapperConverter.h"

void iggShowDemoWindow(IggBool *open)
{
   BoolWrapper openArg(open);

   ImGui::ShowDemoWindow(openArg);
}

void iggShowUserGuide(void)
{
   ImGui::ShowUserGuide();
}

IggBool iggBegin(char const *id, IggBool *open, int flags)
{
   BoolWrapper openArg(open);
   return ImGui::Begin(id, openArg, flags) ? 1 : 0;
}

void iggEnd(void)
{
   ImGui::End();
}

IggBool iggBeginChild(char const *id, IggVec2 const *size, IggBool border, int flags)
{
   Vec2Wrapper sizeArg(size);
   return ImGui::BeginChild(id, *sizeArg, border, flags) ? 1 : 0;
}

void iggEndChild(void)
{
   ImGui::EndChild();
}

void iggWindowPos(IggVec2 *pos)
{
   exportValue(*pos, ImGui::GetWindowPos());
}

void iggWindowSize(IggVec2 *size)
{
   exportValue(*size, ImGui::GetWindowSize());
}

float iggWindowWidth(void)
{
   return ImGui::GetWindowWidth();
}

float iggWindowHeight(void)
{
   return ImGui::GetWindowHeight();
}

void iggContentRegionAvail(IggVec2 *size)
{
   exportValue(*size, ImGui::GetContentRegionAvail());
}

void iggGetContentRegionMax(IggVec2 *out)
{
   ImVec2 im_out = ImGui::GetContentRegionMax();
   exportValue(*out, im_out);
}

void iggGetWindowContentRegionMin(IggVec2 *out)
{
   ImVec2 im_out = ImGui::GetWindowContentRegionMin();
   exportValue(*out, im_out);
}

void iggGetWindowContentRegionMax(IggVec2 *out)
{
   ImVec2 im_out = ImGui::GetWindowContentRegionMax();
   exportValue(*out, im_out);
}

void iggSetNextWindowPos(IggVec2 const *pos, int cond, IggVec2 const *pivot)
{
   Vec2Wrapper posArg(pos);
   Vec2Wrapper pivotArg(pivot);
   ImGui::SetNextWindowPos(*posArg, cond, *pivotArg);
}

void iggSetNextWindowCollapsed(IggBool collapsed, int cond)
{
   ImGui::SetNextWindowCollapsed(collapsed, cond);
}

void iggSetNextWindowSize(IggVec2 const *size, int cond)
{
   Vec2Wrapper sizeArg(size);
   ImGui::SetNextWindowSize(*sizeArg, cond);
}

void iggSetNextWindowSizeConstraints(const IggVec2 *size_min, const IggVec2 *size_max)
{
   Vec2Wrapper sizeMinArg(size_min);
   Vec2Wrapper sizeMaxArg(size_max);
   ImGui::SetNextWindowSizeConstraints(*sizeMinArg, *sizeMaxArg);
}

void iggSetNextWindowContentSize(IggVec2 const *size)
{
   Vec2Wrapper sizeArg(size);
   ImGui::SetNextWindowContentSize(*sizeArg);
}

void iggSetNextWindowFocus(void)
{
   ImGui::SetNextWindowFocus();
}

void iggSetNextWindowBgAlpha(float value)
{
   ImGui::SetNextWindowBgAlpha(value);
}

void iggPushItemWidth(float width)
{
   ImGui::PushItemWidth(width);
}

void iggPopItemWidth(void)
{
   ImGui::PopItemWidth();
}

void iggSetNextItemWidth(float width)
{
   ImGui::SetNextItemWidth(width);
}

void iggPushItemFlag(int options, IggBool enabled)
{
   ImGui::PushItemFlag(options, enabled);
}

void iggPopItemFlag(void)
{
   ImGui::PopItemFlag();
}

float iggCalcItemWidth(void)
{
   return ImGui::CalcItemWidth();
}

void iggPushTextWrapPos(float wrapPosX)
{
   ImGui::PushTextWrapPos(wrapPosX);
}

void iggPopTextWrapPos(void)
{
   ImGui::PopTextWrapPos();
}

void iggPushAllowKeyboardFocus(IggBool allow)
{
   ImGui::PushAllowKeyboardFocus(allow);
}

void iggPopAllowKeyboardFocus()
{
   ImGui::PopAllowKeyboardFocus();
}

void iggPushButtonRepeat(IggBool repeat)
{
   ImGui::PushButtonRepeat(repeat);
}

void iggPopButtonRepeat()
{
   ImGui::PopButtonRepeat();
}

IggViewport iggGetMainViewport()
{
   return static_cast<IggViewport>(ImGui::GetMainViewport());
}

unsigned int iggViewportGetID(IggViewport handle)
{
   ImGuiViewport *viewport = reinterpret_cast<ImGuiViewport *>(handle);
   return viewport->ID;
}

int iggViewportGetFlags(IggViewport handle)
{
   ImGuiViewport *viewport = reinterpret_cast<ImGuiViewport *>(handle);
   return viewport->Flags;
}

void iggViewportGetPos(IggViewport handle, IggVec2 *out)
{
   ImGuiViewport *viewport = reinterpret_cast<ImGuiViewport *>(handle);
   exportValue(*out, viewport->Pos);
}

void iggViewportGetSize(IggViewport handle, IggVec2 *out)
{
   ImGuiViewport *viewport = reinterpret_cast<ImGuiViewport *>(handle);
   exportValue(*out, viewport->Size);
}

void iggViewportGetWorkPos(IggViewport handle, IggVec2 *out)
{
   ImGuiViewport *viewport = reinterpret_cast<ImGuiViewport *>(handle);
   exportValue(*out, viewport->WorkPos);
}

void iggViewportGetWorkSize(IggViewport handle, IggVec2 *out)
{
   ImGuiViewport *viewport = reinterpret_cast<ImGuiViewport *>(handle);
   exportValue(*out, viewport->WorkSize);
}

void iggViewportGetCenter(IggViewport handle, IggVec2 *out)
{
   ImGuiViewport *viewport = reinterpret_cast<ImGuiViewport *>(handle);
   exportValue(*out, viewport->GetCenter());
}

void iggViewportGetWorkCenter(IggViewport handle, IggVec2 *out)
{
   ImGuiViewport *viewport = reinterpret_cast<ImGuiViewport *>(handle);
   exportValue(*out, viewport->GetWorkCenter());
}

unsigned int iggDockSpace(unsigned int dockspace_id, IggVec2 const *size, int flags, const void* window_class)
{
   Vec2Wrapper sizeArg(size);
   return ImGui::DockSpace(dockspace_id, *sizeArg, flags, static_cast<const ImGuiWindowClass*>(window_class));
}

unsigned int iggDockSpaceOverViewport(unsigned int dockspace_id, IggViewport viewport, int flags)
{
   ImGuiViewport *viewportPtr = reinterpret_cast<ImGuiViewport *>(viewport);
   return ImGui::DockSpaceOverViewport(dockspace_id, viewportPtr, flags);
}

void iggSetNextWindowDockID(unsigned int dock_id, int cond)
{
   ImGui::SetNextWindowDockID(dock_id, cond);
}

void iggSetNextWindowViewport(unsigned int viewport_id)
{
   ImGui::SetNextWindowViewport(viewport_id);
}

unsigned int iggGetID(const char* str_id)
{
   return ImGui::GetID(str_id);
}
