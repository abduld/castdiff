// MP 1
#include	<wb.h>

__global__ void vecAdd(float * in1, float * in2, float * out, int len) {
    //@@ Insert code to implement vector addition here
	int i = blockIdx.x*blockDim.x+threadIdx.x;
	if(i<len)
		out[i] = in1[i]+in2[i];
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
	float* d_A;
	float* d_B;
	float* d_C;
    int size = inputLength*sizeof(float);
	cudaMalloc((void**) &d_A,size);
	cudaMalloc((void**) &d_B,size);
	cudaMalloc((void**) &d_C,size);

    wbTime_stop(GPU, "Allocating GPU memory.");

    wbTime_start(GPU, "Copying input memory to the GPU.");
    //@@ Copy memory to the GPU here
	cudaMemcpy(d_A,hostInput1,size,cudaMemcpyHostToDevice);
	cudaMemcpy(d_B,hostInput2,size,cudaMemcpyHostToDevice);


    wbTime_stop(GPU, "Copying input memory to the GPU.");
    
    //@@ Initialize the grid and block dimensions here
	dim3 DimGrid(((inputLength-1)/256 +1.0),1,1);
	dim3 DimBlock(256,1,1);
	


    
    wbTime_start(Compute, "Performing CUDA computation");
    //@@ Launch the GPU Kernel here
	vecAdd<<<DimGrid,DimBlock>>>(d_A,d_B,d_C,inputLength);

    cudaThreadSynchronize();
    wbTime_stop(Compute, "Performing CUDA computation");
    
    wbTime_start(Copy, "Copying output memory to the CPU");
    //@@ Copy the GPU memory back to the CPU here
	cudaMemcpy(hostOutput,d_C,size,cudaMemcpyDeviceToHost);

    wbTime_stop(Copy, "Copying output memory to the CPU");

    wbTime_start(GPU, "Freeing GPU Memory");
    //@@ Free the GPU memory here
	cudaFree(d_A);
	cudaFree(d_B);
	cudaFree(d_C);


    wbTime_stop(GPU, "Freeing GPU Memory");

    wbSolution(args, hostOutput, inputLength);

    free(hostInput1);
    free(hostInput2);
    free(hostOutput);

    return 0;
}

