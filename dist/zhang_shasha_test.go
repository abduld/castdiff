package main

import (
	"testing"

	cc "github.com/abduld/castdiff/cc"
)

var src0 string = `
// MP 1
#include  <wb.h>

__global__ void vecAdd(float * in1, float * in2, float * out, int len) {
    //@@ Insert code to implement vector addition here
  int i = threadIdx.x+blockDim.x*blockIdx.x;
  if (i<len) out[i]=in1[i]+in2[i];
}

int main(int argc, char ** argv) {
    wbArg_t args;
    int inputLength;
    float * hostInput1;
    float * hostInput2;
    float * hostOutput;
    float * deviceInput1;
    float * deviceInput2;
    float * deviceOutput;

    args = wbArg_read(argc, argv);

    wbTime_start(Generic, "Importing data and creating memory on host");
    hostInput1 = (float *) wbImport(wbArg_getInputFile(args, 0), &inputLength);
    hostInput2 = (float *) wbImport(wbArg_getInputFile(args, 1), &inputLength);
    hostOutput = (float *) malloc(inputLength * sizeof(float));
    wbTime_stop(Generic, "Importing data and creating memory on host");

    wbLog(TRACE, "The input length is ", inputLength);

  wbTime_start(GPU, "Allocating GPU memory.");
    //@@ Allocate GPU memory here
  cudaMalloc((void **)&deviceInput1, sizeof(float)*inputLength);
  cudaMalloc((void **)&deviceInput2, sizeof(float)*inputLength);
  cudaMalloc((void **)&deviceOutput, sizeof(float)*inputLength);

    wbTime_stop(GPU, "Allocating GPU memory.");

    wbTime_start(GPU, "Copying input memory to the GPU.");
    //@@ Copy memory to the GPU here
  cudaMemcpy(deviceInput1, hostInput1, sizeof(float)*inputLength, cudaMemcpyHostToDevice);
  cudaMemcpy(deviceInput2, hostInput2, sizeof(float)*inputLength, cudaMemcpyHostToDevice);

    wbTime_stop(GPU, "Copying input memory to the GPU.");
    
    //@@ Initialize the grid and block dimensions here
  const int BLOCK_SIZE=256;
    dim3 dimGrid;
    dimGrid((inputLength-1)/BLOCK_SIZE+1, 1, 1);

    
    wbTime_start(Compute, "Performing CUDA computation");
    //@@ Launch the GPU Kernel here
  vecAdd<<<dimGrid, dimBlock>>>(deviceInput1, deviceInput2, deviceOutput, inputLength);
    cudaThreadSynchronize();
    wbTime_stop(Compute, "Performing CUDA computation");
    
    wbTime_start(Copy, "Copying output memory to the CPU");
    //@@ Copy the GPU memory back to the CPU here
  cudaMemcpy(hostOutput, deviceOutput, sizeof(float)*inputLength, cudaMemcpyDeviceToHost);
    wbTime_stop(Copy, "Copying output memory to the CPU");

    wbTime_start(GPU, "Freeing GPU Memory");
    //@@ Free the GPU memory here
  cudaFree(deviceInput1);
  cudaFree(deviceInput2);
  cudaFree(deviceOutput);

    wbTime_stop(GPU, "Freeing GPU Memory");

    wbSolution(args, hostOutput, inputLength);

    free(hostInput1);
    free(hostInput2);
    free(hostOutput);

    return 0;
}
`

var src1 string = `
int x = 1;
`
var src2 string = `
float x = 2;
float x = 2;
`

func TestDistanceDecl(t *testing.T) {
	p1, err := cc.ParseProg(src2)
	if err != nil {
		t.Errorf("Unable to parse p1 -- %#q", err)
		return
	}
	p2, err := cc.ParseProg(src1)
	if err != nil {
		t.Errorf("Unable to parse p1 -- %#q", err)
		return
	}
	/*
		p2, err := cc.ParseProg("int x;")
		if err != nil {
			t.Errorf("Unable to parse p2 -- %#q", err)
			return
		}
	*/

	if ASTDistance(p1, p2) != 0 {
		t.Errorf("Distance between %#q, %#q was not expected", p1, p1)
	}
}
