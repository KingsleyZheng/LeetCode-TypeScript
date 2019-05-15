//简单
//21. 合并两个有序链表
//将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
//当然，每次选两个队伍排头的，比较他们的高矮!组成新的的队伍。
/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
//递归
var mergeTwoLists = function(l1, l2) {
    if (l1==null) return l2;
    if (l2==null) return l1;
    if(l1.val<l2.val) {
        l1.next = mergeTwoLists(l1.next,l2);
        return l1;
    } else {
        l2.next = mergeTwoLists(l1,l2.next);
        return l2;
    }
};

//35. 搜索插入位置
//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//你可以假设数组中无重复元素。
/*
示例 1:
输入: [1,3,5,6], 5
输出: 2
*/
//这个比较简单，就是遍历数组
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var searchInsert = function(nums, target) {
    let l = nums.length;
    for(let i=0;i<l;i++) {
        if (target<=nums[i]) {
            return i;
        }
    }
    return l;        
};
//执行用时 : 68 ms, 在Search Insert Position的JavaScript提交中击败了99.14% 的用户