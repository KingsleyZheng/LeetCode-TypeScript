// 二叉树的最大深度
// 简单

import { TreeNode } from "./types";

// 递归，深度优先搜索
function maxDepth(root: TreeNode | null): number {
  if (root === null) {
    return 0;
  } else {
    return Math.max(maxDepth(root.left), maxDepth(root.right)) + 1;
  }
}
