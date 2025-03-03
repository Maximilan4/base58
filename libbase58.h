#ifndef LIB58_H
#define LIB58_H
#include <stdbool.h>

bool encode_base58(const void *data, size_t binsz, char *b58, size_t *b58sz);
bool decode_base58(const char *b58, size_t b58sz, void *bin, size_t *binszp);
int get_errno(void);

#endif // LIB58_H