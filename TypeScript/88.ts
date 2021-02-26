// 合并两个有序数组
// 简单

function merge(nums1: number[], m: number, nums2: number[], n: number): void {
  let pa = m - 1,
    pb = n - 1;
  let tail = m + n - 1;
  let cur: number;
  while (pa >= 0 || pb >= 0) {
    if (pa == -1) {
      cur = nums2[pb--];
    } else if (pb == -1) {
      cur = nums1[pa--];
    } else if (nums1[pa] > nums2[pb]) {
      cur = nums1[pa--];
    } else {
      cur = nums2[pb--];
    }
    nums1[tail--] = cur;
  }
};