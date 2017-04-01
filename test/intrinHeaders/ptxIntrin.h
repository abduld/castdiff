#ifndef PTX_INTRIN_H
#define PTX_INTRIN_H

extern __device__  int                    __float_as_int(float x);
extern __device__  void                   __syncthreads(void);
extern __device__  void                   __prof_trigger(int);
extern __device__  void                   __threadfence(void);
extern __device__  void                   __threadfence_block(void);
extern __device__  void                   __trap(void);
// extern __device__  void                   __brkpt(int c = 0);
static __device__ __inline__ float __shfl_xor(float var, int laneMask, int width=warpSize);

#endif // PTX_INTRIN_H
