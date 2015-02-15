// MP 1
#include	<wb.h>
#include    <cuda.h>
#include    <stdio.h>
#include    <stdlib.h>

__global__ void vecAdd(float * in1, float * in2, float * out, int len) {
    //@@ Insert code to implement vector addition here
	int i;
	i = blockIdx.x * blockDim.x + threadIdx.x;
	
	if(i<len)
	{
		out[i] = in1[i] + in2[i];
	}
						
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
	cudaError_t err;
	int size;

    wbTime_start(Generic, "Importing data and creating memory on host");
    hostInput1 = (float *) wbImport(wbArg_getInputFile(args, 0), &inputLength);
    hostInput2 = (float *) wbImport(wbArg_getInputFile(args, 1), &inputLength);
    hostOutput = (float *) malloc(inputLength * sizeof(float));
    wbTime_stop(Generic, "Importing data and creating memory on host");

    wbLog(TRACE, "The input length is ", inputLength);
	size = inputLength * sizeof(float);

	wbTime_start(GPU, "Allocating GPU memory.");
    //@@ Allocate GPU memory here
	err = cudaMalloc((void**)&deviceInput1, size);
	if(err != cudaSuccess)
	{
		printf("\nError...! Can't allocate memory for input array1 in device\n");
	}
	err = cudaMalloc((void**)&deviceInput2, size);
	if(err != cudaSuccess)
	{
		printf("\nError...! Can't allocate memory for input array2 in device\n");
	}
	err = cudaMalloc((void**)&deviceOutput, size);
	if(err != cudaSuccess)
	{
		printf("\nError...! Can't allocate memory for output array in device\n");
	}

    wbTime_stop(GPU, "Allocating GPU memory.");

    wbTime_start(GPU, "Copying input memory to the GPU.");
    //@@ Copy memory to the GPU here
	err = cudaMemcpy(deviceInput1, hostInput1, size, cudaMemcpyHostToDevice);
	if(err != cudaSuccess)
		printf("\nError...! Couldn't copy input array1 from host to device\n");
	
	err = cudaMemcpy(deviceInput2, hostInput2, size, cudaMemcpyHostToDevice);
	if(err != cudaSuccess)
		printf("\nError...! Couldn't copy input array2 from host to device\n");

    wbTime_stop(GPU, "Copying input memory to the GPU.");
    
    //@@ Initialize the grid and block dimensions here
	dim3 Grid(((inputLength-1)/256)+1,1,1);
	dim3 Block(256,1,1);

    
    wbTime_start(Compute, "Performing CUDA computation");
    //@@ Launch the GPU Kernel here
    vecAdd<<<Grid, Block>>>(deviceInput1, deviceInput2, deviceOutput, inputLength);
    cudaThreadSynchronize();
    wbTime_stop(Compute, "Performing CUDA computation");
    
    wbTime_start(Copy, "Copying output memory to the CPU");
    //@@ Copy the GPU memory back to the CPU here
	err = cudaMemcpy(hostOutput, deviceOutput, size, cudaMemcpyDeviceToHost);
	if(err != cudaSuccess)
		printf("\nError...! Couldn't copy output array from device to host\n");
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

