#include <cuda.h>

__global__ void vecadd(int * a, int * b, int len) {
	int idx = threadIdx.x + blockDim.x * blockIdx.x;
	a[idx] += b[idx];
}

