#include <wb.h>


int main(int argc, char **argv) {

  wbTime_start(Generic, "Importing data and creating memory on host");
  wbTime_stop(Generic, "Importing data and creating memory on host");
  wbLog(TRACE, "The input length is ", inputLength);

  return 0;
}
