package cc

var hdr_u_h = `
typedef signed char schar;
typedef unsigned char uchar;
typedef unsigned short ushort;
typedef unsigned int uint;
typedef unsigned long ulong;
typedef long long vlong;
typedef unsigned long long uvlong;

typedef schar int8;
typedef uchar uint8;
typedef short int16;
typedef ushort uint16;
typedef long int32;
typedef ulong uint32;
typedef vlong int64;
typedef uvlong uint64;
typedef float float32;
typedef double float64;
typedef unsigned long uintptr;

typedef schar s8int;
typedef uchar u8int;
typedef short s16int;
typedef ushort u16int;
typedef long s32int;
typedef ulong u32int;
typedef vlong s64int;
typedef uvlong u64int;

void *nil;

typedef struct va_list *va_list;


`

var hdr_libc_h = `
int nelem(void*);

int memcmp(void*, void*, long);
void *memset(void*, int, long);
int strcmp(char*, char*);
int strncmp(char*, char*, int);
char *strcpy(char*, char*);
char *smprint(char*, ...);
void strcat(char*, char*);

int atoi(char*);
ulong strtoul(char*, char**, int);
long strtol(char*, char**, int);
vlong strtoll(char*, char**, int);
int isspace(int);
int close(int);
int read(int, void*, int);
double atof(char*);
int create(char*, int, int);
int open(char*, int);
uvlong strtoull(char*, char**, int);
char *getenv(char*);
int getwd(char*, int);
double cputime(void);

enum
{
	AEXIST = 0,
};

int errstr(char*, uint);
void werrstr(char*, ...);

void exits(char*);
void sysfatal(char*, ...);
char *strstr(char*, char*);
int strlen(char*);
void memmove(void*, void*, int);
char *strdup(char*);
void *malloc(int);
void *calloc(int, int);
void *realloc(void*, int);
void free(void*);

void va_start(va_list, void*);
void va_end(va_list);
void qsort(void *base, int nmemb, int size, int (*compar)(const void *, const void *));
char *GOEXPERIMENT;
void setfcr(int);
void notify(void*);
void signal(int, void*);
uintptr getcallerpc(void*);

enum
{
	OREAD,
	OWRITE,
	ORDWR,
	SIGBUS,
	SIGSEGV,
	NDFLT,
	FPPDBL,
	FPRNR,
	BOM,
	HEADER_IO,
};

extern	void	flagcount(char*, char*, int*);
extern	void	flagint32(char*, char*, int32*);
extern	void	flagint64(char*, char*, int64*);
extern	void	flagstr(char*, char*, char**);
extern	void	flagparse(int*, char***, void (*usage)(void));
extern	void	flagfn0(char*, char*, void(*fn)(void));
extern	void	flagfn1(char*, char*, void(*fn)(char*));
extern	void	flagfn2(char*, char*, void(*fn)(char*, char*));
extern	void	flagprint(int);
extern	char*	strecpy(char*, char*, char*);
extern	void	abort(void);
extern	int	remove(const char*);
extern	char*	getgoos(void);
extern	char*	getgoarch(void);
extern	char*	getgoroot(void);
extern	char*	getgoversion(void);
extern	char*	getgoarm(void);
extern	char*	getgo386(void);
extern	char*	getgoextlinkenabled(void);
extern	char*	getgohostos(void);
extern	char*	getgohostarch(void);
extern	int	runcmd(char**);
extern	char*	strchr(char*, int);
extern	char*	strrchr(char*, int);
extern	double	floor(double);
extern	double	ldexp(double, int);
extern	double	frexp(double, int*);
extern	double	pow(double, double);

extern	int	access(char*, int);
extern	int	isdigit(int);
extern	int	isalpha(int);
extern	int	isalnum(int);
extern	int	getfields(char*, char**, int, int, char*);
extern	char*	cleanname(char*);
extern	int	noted(int);

`

var hdr_wb_h = `
#include <u.h>
#include <libc.h>
#include <math.h>
typedef int *wbArg_t;
typedef int(*dim3)(int, int, int);
`

var hdr_math_h = `
double      acos(double);
float       acosf(float);
double      acosh(double);
float       acoshf(float);
long double acoshl(long double);
long double acosl(long double);
double      asin(double);
float       asinf(float);
double      asinh(double);
float       asinhf(float);
long double asinhl(long double);
long double asinl(long double);
double      atan(double);
double      atan2(double, double);
float       atan2f(float, float);
long double atan2l(long double, long double);
float       atanf(float);
double      atanh(double);
float       atanhf(float);
long double atanhl(long double);
long double atanl(long double);
double      cbrt(double);
float       cbrtf(float);
long double cbrtl(long double);
double      ceil(double);
float       ceilf(float);
long double ceill(long double);
double      copysign(double, double);
float       copysignf(float, float);
long double copysignl(long double, long double);
double      cos(double);
float       cosf(float);
double      cosh(double);
float       coshf(float);
long double coshl(long double);
long double cosl(long double);
double      erf(double);
double      erfc(double);
float       erfcf(float);
long double erfcl(long double);
float       erff(float);
long double erfl(long double);
double      exp(double);
double      exp2(double);
float       exp2f(float);
long double exp2l(long double);
float       expf(float);
long double expl(long double);
double      expm1(double);
float       expm1f(float);
long double expm1l(long double);
double      fabs(double);
float       fabsf(float);
long double fabsl(long double);
double      fdim(double, double);
float       fdimf(float, float);
long double fdiml(long double, long double);
double      floor(double);
float       floorf(float);
long double floorl(long double);
double      fma(double, double, double);
float       fmaf(float, float, float);
long double fmal(long double, long double, long double);
double      fmax(double, double);
float       fmaxf(float, float);
long double fmaxl(long double, long double);
double      fmin(double, double);
float       fminf(float, float);
long double fminl(long double, long double);
double      fmod(double, double);
float       fmodf(float, float);
long double fmodl(long double, long double);
double      frexp(double, int *);
float       frexpf(float value, int *);
long double frexpl(long double value, int *);
double      hypot(double, double);
float       hypotf(float, float);
long double hypotl(long double, long double);
int         ilogb(double);
int         ilogbf(float);
int         ilogbl(long double);
double      j0(double);
double      j1(double);
double      jn(int, double);

`