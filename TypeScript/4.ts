// 寻找两个正序数组的中位数
// 困难

export function findMedianSortedArrays(nums1: number[], nums2: number[]): number {
  let length1 = nums1.length,
    length2 = nums2.length;
  let totalLength = length1 + length2;
  if (totalLength % 2 === 1) {
    let midIndex = Math.floor(totalLength / 2);
    let median = getKthElement(nums1, nums2, midIndex + 1);
    return median;
  } else {
    let midIndex1 = Math.floor(totalLength / 2) - 1,
      midIndex2 = Math.floor(totalLength / 2);
    let median =
      (getKthElement(nums1, nums2, midIndex1 + 1) +
        getKthElement(nums1, nums2, midIndex2 + 1)) /
      2.0;
    return median;
  }
}

function getKthElement(nums1: number[], nums2: number[], k: number): number {
  let length1 = nums1.length,
    length2 = nums2.length;
  let index1 = 0,
    index2 = 0;
  while (true) {
    // 边界情况
    if (index1 == length1) {
      return nums2[index2 + k - 1];
    }
    if (index2 == length2) {
      return nums1[index1 + k - 1];
    }
    if (k == 1) {
      return Math.min(nums1[index1], nums2[index2]);
    }

    // 正常情况
    let half = Math.floor(k / 2);
    let newIndex1 = Math.min(index1 + half, length1) - 1;
    let newIndex2 = Math.min(index2 + half, length2) - 1;
    let pivot1 = nums1[newIndex1],
      pivot2 = nums2[newIndex2];
    if (pivot1 <= pivot2) {
      k -= newIndex1 - index1 + 1;
      index1 = newIndex1 + 1;
    } else {
      k -= newIndex2 - index2 + 1;
      index2 = newIndex2 + 1;
    }
  }
}

console.log(findMedianSortedArrays([1, 3], [2, 4]));
