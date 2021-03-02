// 二叉树的中序遍历
// 中等
// 给定一个二叉树的根节点 root ，返回它的 中序 遍历。

import { TreeNode } from './types';

// 递归
function inorderTraversal(root: TreeNode | null): number[] {
  let ans: number[] = [];
  inorder(root, ans);
  return ans;
}

function inorder(root: TreeNode | null, ans: number[]) {
  if (root === null) return;
  inorder(root.left, ans);
  ans.push(root.val);
  inorder(root.right, ans);
}

console.log(inorderTraversal(new TreeNode(1)));

// 迭代
var inorderTraversal2 = function (root: TreeNode | null): number[] {
  const res = [];
  const stk = [];
  while (root || stk.length) {
    while (root) {
      stk.push(root);
      root = root.left;
    }
    root = stk.pop()!;
    res.push(root.val);
    root = root.right;
  }
  return res;
};
