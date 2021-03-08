// 翻转二叉树
// 简单

import { TreeNode } from './types';

function invertTree(root: TreeNode | null): TreeNode | null {
  if (root === null) {
    return null;
  }
  const temp = root.left;
  root.left = invertTree(root.right);
  root.right = invertTree(temp);
  return root;
}
