#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern unsigned int iggDockBuilderAddNode(unsigned int node_id, int flags);
extern void iggDockBuilderRemoveNode(unsigned int node_id);
extern void iggDockBuilderSetNodeSize(unsigned int node_id, IggVec2 const *size);
extern unsigned int iggDockBuilderSplitNode(unsigned int node_id, int split_dir, float size_ratio, unsigned int *out_id_at_dir, unsigned int *out_id_at_opposite_dir);
extern void iggDockBuilderDockWindow(const char* window_name, unsigned int node_id);
extern void iggDockBuilderFinish(unsigned int node_id);

#ifdef __cplusplus
}
#endif