
#ifndef __WB_LOGGER_H__
#define __WB_LOGGER_H__

void wbLog(const char * level, ...);
#define wbLog(level, ...)  wbLog(#level, __VA_ARGS__) 


#endif /* __WB_LOGGER_H__ */
