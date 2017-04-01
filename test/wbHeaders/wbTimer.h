

#ifndef __WB_TIMER_H__
#define __WB_TIMER_H__

void wbTime_start(const char * kind, ...);
void wbTime_stop(const char * kind, ...);
#define wbTime_start(kind, ...) wbTime_start(#kind, __VA_ARGS__);
#define wbTime_stop(kind, ...) wbTime_stop(#kind, __VA_ARGS__);

#endif /* __WB_TIMER_H__ */
