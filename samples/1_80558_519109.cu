// MP 1
#include	<wb.h>
#include	<cuda.h>

#define wbCheck(stmt) do {                                                \
    cudaError_t err = stmt;                                               \
    if (err != cudaSuccess) {                                             \
        wbLog(ERROR, "Failed to run stmt ", #stmt);                       \
        wbLog(ERROR, "Got CUDA error ...  ", cudaGetErrorString(err));    \
        return -1;                                                        \
    }                                                                     \
} while(0)

__global__ 
void vecAdd(float *in1, float *in2, float *out, int len) {
    
	int i = blockIdx.x * blockDim.x + threadIdx.x;
	if( i<len ) {
		out[i] = in1[i]+in2[i];
	}
}

__host__
int main(int argc, char ** argv) {
    wbArg_t args;
    int inputLength;
	int numBytes;
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
	numBytes = (sizeof(float) * inputLength);
    hostOutput = (float *) malloc(numBytes);
    wbTime_stop(Generic, "Importing data and creating memory on host");

	wbTime_start(GPU, "Allocating GPU memory.");
    wbCheck(cudaMalloc(&deviceInput1, numBytes));
	wbCheck(cudaMalloc(&deviceInput2, numBytes));
	wbCheck(cudaMalloc(&deviceOutput, numBytes));
	
    wbTime_stop(GPU, "Allocating GPU memory.");

    wbTime_start(GPU, "Copying input memory to the GPU.");
    wbCheck(cudaMemcpy(deviceInput1, hostInput1, numBytes, cudaMemcpyHostToDevice));
	wbCheck(cudaMemcpy(deviceInput2, hostInput2, numBytes, cudaMemcpyHostToDevice));

    wbTime_stop(GPU, "Copying input memory to the GPU.");
    
    dim3 DimGrid((inputLength-1)/256 + 1, 1, 1);
	dim3 DimBlock(256, 1, 1);

    wbTime_start(Compute, "Performing CUDA computation");
    vecAdd<<<DimGrid,DimBlock>>>(deviceInput1, deviceInput2, deviceOutput, inputLength);
	wbCheck(cudaGetLastError());

    wbCheck(cudaThreadSynchronize());
    wbTime_stop(Compute, "Performing CUDA computation");
    
    wbTime_start(Copy, "Copying output memory to the CPU");
    wbCheck(cudaMemcpy(hostOutput, deviceOutput, numBytes, cudaMemcpyDeviceToHost));

    wbTime_stop(Copy, "Copying output memory to the CPU");

    wbTime_start(GPU, "Freeing GPU Memory");
    wbCheck(cudaFree(deviceInput1));
	wbCheck(cudaFree(deviceInput2));
	wbCheck(cudaFree(deviceOutput));

    wbTime_stop(GPU, "Freeing GPU Memory");

    wbSolution(args, hostOutput, inputLength);

    free(hostInput1);
    free(hostInput2);
    free(hostOutput);

    return 0;
}

