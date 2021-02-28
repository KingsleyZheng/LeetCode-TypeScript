// 对称二叉树
// 简单

import { TreeNode } from "./types";

function isSymmetric(root: TreeNode | null): boolean {
  return check(root, root);
}

function check(p: TreeNode | null, q: TreeNode | null): boolean {
  if (!p && !q) return true;
  if (!p || !q) return false;
  return p.val === q.val && check(p.left, q.right) && check(p.right, q.left);
}
