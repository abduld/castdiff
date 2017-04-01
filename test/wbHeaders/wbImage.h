

#ifndef __IMAGE_H__
#define __IMAGE_H__

typedef struct st_wbImage_t {
  int width;
  int height;
  int channels;
  int pitch;
  float *data;
} *wbImage_t;

#endif /* __IMAGE_H__ */
