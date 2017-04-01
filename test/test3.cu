#include <cuda.h>

__device__ int outInvariant(int inValue) {
  return inValue * inValue;
}

__host__ __device__ int outDependent(int value, int inIdx, int outIdx) {
  if (inIdx == outIdx) {
    return 2 * value;
  } else if (inIdx > outIdx) {
    return value / (inIdx - outIdx);
  } else {
    return value / (outIdx - inIdx);
  }
}

