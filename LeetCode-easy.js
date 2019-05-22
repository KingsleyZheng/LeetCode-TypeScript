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

//38. 报数
//报数序列是一个整数序列，按照其中的整数的顺序进行报数，得到下一个数。其前五项如下：
/**
1.     1
2.     11
3.     21
4.     1211
5.     111221
1 被读作  "one 1"  ("一个一") , 即 11。
11 被读作 "two 1s" ("两个一"）, 即 21。
21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。
给定一个正整数 n（1 ≤ n ≤ 30），输出报数序列的第 n 项。
注意：整数顺序将表示为一个字符串。
 */
//使用递归解法
/**
 * @param {number} n
 * @return {string}
 */
var countAndSay = function(n) {
    if (n==1) return "1";
    var strLast = countAndSay(n-1);
    var count = 1;
    var res = "";
    for (let i = 0; i < strLast.length; i++) {
        if (strLast[i] == strLast[i+1]){
            count++;
            continue;
        } else{
            if(strLast[i]!=strLast[i+1]){
                res += count.toString() + strLast[i];
                count=1;
            }
        }
    }
    return res;
};
//执行用时 : 108 ms, 在Count and Say的JavaScript提交中击败了49.16% 的用户
//不用递归
var countAndSay = function(n) {
    if (n<=0) return "";
    let res = "1";
    while (--n) {
        let cur = "";
        let l = res.length;
        for (let i=0;i<l;i++){
            let count = 1;
            while (i+1<l && res[i] == res[i+1]) {
                count++;
                i++;
            }
            cur += count.toString() + res[i];
        }
        res = cur;
    }
    return res;
};
//执行用时 : 84 ms, 在Count and Say的JavaScript提交中击败了98.20% 的用户

//53. 最大子序和
/**
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
示例:
输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶:
如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
 */
/**
 * @param {number[]} nums
 * @return {number}
 */
//动态递归，一个变量记录最大和，一个变量记录当前和，当前和若小于0则重置为下一个数，时间复杂度n
var maxSubArray = function(nums) {
    let sum = 0;
    let res = nums[0];
    let l = nums.length;
    for(let i=0;i<l;i++){
        let num = nums[i];
        sum = sum>0?sum+num: num;
        if (res<sum) {
            res = sum;
        }
    }
    return res;
};

//58. 最后一个单词的长度
/**
给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。
如果不存在最后一个单词，请返回 0 。
说明：一个单词是指由字母组成，但不包含任何空格的字符串。
示例:
输入: "Hello World"
输出: 5
 */
//先去掉两头的空格，再按空格分成列表，返回最后一个单词的长度。
//其实可以只去掉末尾的空格，再从末尾数到空格有几个数，该长度就是最后一个单词的长度
/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLastWord = function(s) {
    let sList = s.trim().split(' ');
    let l = sList.length;
    if (l == 0) return 0;
    return sList[l-1].length;
};
//执行用时 : 72 ms, 在Length of Last Word的JavaScript提交中击败了98.55% 的用户

//67. 二进制求和
/**
给定两个二进制字符串，返回他们的和（用二进制表示）。
输入为非空字符串且只包含数字 1 和 0。
示例 1:
输入: a = "11", b = "1"
输出: "100"
示例 2:
输入: a = "1010", b = "1011"
输出: "10101"
 */
/**
 * @param {string} a
 * @param {string} b
 * @return {string}
 */
var addBinary = function(a, b) {
    let res = '';
    let na = a.length;
    let nb = b.length;
    let n = Math.max(na, nb);
    let carry = false; // 是否需要进位
    // 在字符串左边补全0至ab长度相等
    if (na>nb) {
        for (let i=0;i<na-nb;i++) {
            b = '0' + b
        }
    } else if (na<nb) {
        for (let i=0;i<nb-na;i++) {
            a = '0' + a
        }
    }
    // 诸位相加比较大小，决定是否需要进位
    for (let i=n-1;i>=0;i--) {
        let tmp = 0;
        if (carry) {tmp = (a[i] - '0') + (b[i] - '0') + 1;
        } else {tmp = (a[i] - '0') + (b[i] - '0');}
        if (tmp == 0) {
            res = '0' + res;
            carry = false;
        } else if (tmp == 1) {
            res = '1' + res;
            carry = false;
        } else if (tmp == 2) {
            res = '0' + res;
            carry = true;
        } else if (tmp == 3) {
            res = '1' + res;
            carry = true;
        }
    }
    if (carry) res = '1' + res;
    return res;
};
// 执行用时 : 84 ms, 在Add Binary的JavaScript提交中击败了98.42% 的用户
// 更巧妙的方法是不进行预处理加0，需要时再取出来

//70.爬楼梯
/**
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
注意：给定 n 是一个正整数。
示例 1：
输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶
示例 2：
输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶
 */
//dp[i]=dp[i−1]+dp[i−2]
/**
 * @param {number} n
 * @return {number}
 */
var climbStairs = function(n) {
    if (n==1) {
        return 1;
    }
    let dp = [];
    dp[1] = 1;
    dp[2] = 2;
    for (let i=3;i<=n;i++) {
        dp[i] = dp[i-1] + dp[i-2];
    }
    return dp[n];
};

//83. 删除排序链表中的重复元素
/**
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
示例 1:
输入: 1->1->2
输出: 1->2
示例 2:
输入: 1->1->2->3->3
输出: 1->2->3
 */
/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var deleteDuplicates = function(head) {
    let current = head;
    while (current != null && current.next != null) {
        if (current.next.val == current.val) {
            current.next = current.next.next;
        } else {
            current = current.next;
        }
    }
    return head;
};

//88. 合并两个有序数组
/**
给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。
说明:
初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
示例:
输入:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3
输出: [1,2,2,3,5,6]
 */
/**
 * @param {number[]} nums1
 * @param {number} m
 * @param {number[]} nums2
 * @param {number} n
 * @return {void} Do not return anything, modify nums1 in-place instead.
 */
var merge = function(nums1, m, nums2, n) {
    let i = m - 1, j = n - 1, k = m + n - 1;
    while (i >= 0 && j >= 0) {
        if (nums1[i] > nums2[j]) nums1[k--] = nums1[i--]
        else nums1[k--] = nums2[j--];
    }
    while (j>=0) nums1[k--] = nums2[j--];
};