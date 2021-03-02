// 验证二叉搜索树
// 中等

import { TreeNode } from './types';

const helper = (
  root: TreeNode | null,
  lower: number,
  upper: number
): boolean => {
  if (root === null) {
    return true;
  }
  if (root.val <= lower || root.val >= upper) {
    return false;
  }
  return (
    helper(root.left, lower, root.val) && helper(root.right, root.val, upper)
  );
};

function isValidBST(root: TreeNode | null): boolean {
  return helper(root, -Infinity, Infinity);
}

// 中序遍历
function isValidBST2(root: TreeNode | null): boolean {
  let stack = [];
  let inorder = -Infinity;

  while (stack.length || root !== null) {
    while (root !== null) {
      stack.push(root);
      root = root.left;
    }
    root = stack.pop()!;
    // 如果中序遍历得到的节点的值小于等于前一个 inorder，说明不是二叉搜索树
    if (root.val <= inorder) {
      return false;
    }
    inorder = root.val;
    root = root.right;
  }
  return true;
}
