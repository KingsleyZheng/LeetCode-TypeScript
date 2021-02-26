// 合并排序的数组
// 简单

// 方法1 直接合并后排序
function merge1(A: number[], m: number, B: number[], n: number): void {
  for (let i = 0; i < n; i++) {
    A[m + i] = B[i];
  }
  A.sort();
}

// 方法2 双指针，创立一个新数组，修改完再逐个复制到A
function merge2(A: number[], m: number, B: number[], n: number): void {
  let pa = 0,
    pb = 0;
  let sorted = new Array<number>(m + n);
  let cur: number;
  while (pa < m || pb < n) {
    if (pa == m) {
      cur = B[pb++];
    } else if (pb == n) {
      cur = A[pa++];
    } else if (A[pa] < B[pb]) {
      cur = A[pa++];
    } else {
      cur = B[pb++];
    }
    sorted[pa + pb - 1] = cur;
  }
  for (let i = 0; i != m + n; ++i) {
    A[i] = sorted[i];
  }
}

// 方法3 从大到小遍历。这才是真的原地修改，但是要先证明A前面的数不会被覆盖
function merge3(A: number[], m: number, B: number[], n: number): void {
  let pa = m - 1,
    pb = n - 1;
  let tail = m + n - 1;
  let cur: number;
  while (pa >= 0 || pb >= 0) {
    if (pa == -1) {
      cur = B[pb--];
    } else if (pb == -1) {
      cur = A[pa--];
    } else if (A[pa] > B[pb]) {
      cur = A[pa--];
    } else {
      cur = B[pb--];
    }
    A[tail--] = cur;
  }
}
