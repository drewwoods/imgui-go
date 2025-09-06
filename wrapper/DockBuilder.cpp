#include "ConfiguredImGui.h"
#include "DockBuilder.h"
#include "WrapperConverter.h"

unsigned int iggDockBuilderAddNode(unsigned int node_id, int flags)
{
   return ImGui::DockBuilderAddNode(node_id, flags);
}

void iggDockBuilderRemoveNode(unsigned int node_id)
{
   ImGui::DockBuilderRemoveNode(node_id);
}

void iggDockBuilderSetNodeSize(unsigned int node_id, IggVec2 const *size)
{
   Vec2Wrapper sizeArg(size);
   ImGui::DockBuilderSetNodeSize(node_id, *sizeArg);
}

unsigned int iggDockBuilderSplitNode(unsigned int node_id, int split_dir, float size_ratio, unsigned int *out_id_at_dir, unsigned int *out_id_at_opposite_dir)
{
   ImGuiID id_at_dir, id_at_opposite_dir;
   unsigned int result = ImGui::DockBuilderSplitNode(node_id, static_cast<ImGuiDir>(split_dir), size_ratio, &id_at_dir, &id_at_opposite_dir);
   
   if (out_id_at_dir != nullptr)
      *out_id_at_dir = id_at_dir;
   if (out_id_at_opposite_dir != nullptr)
      *out_id_at_opposite_dir = id_at_opposite_dir;
   
   return result;
}

void iggDockBuilderDockWindow(const char* window_name, unsigned int node_id)
{
   ImGui::DockBuilderDockWindow(window_name, node_id);
}

void iggDockBuilderFinish(unsigned int node_id)
{
   ImGui::DockBuilderFinish(node_id);
}