// 二叉树的最小深度
// 简单

import { TreeNode } from './types';

function minDepth(root: TreeNode | null): number {
  if (root == null) {
    return 0;
  }

  if (root.left == null && root.right == null) {
    return 1;
  }

  let min_depth = Number.MAX_VALUE;
  if (root.left != null) {
    min_depth = Math.min(minDepth(root.left), min_depth);
  }
  if (root.right != null) {
    min_depth = Math.min(minDepth(root.right), min_depth);
  }

  return min_depth + 1;
}
