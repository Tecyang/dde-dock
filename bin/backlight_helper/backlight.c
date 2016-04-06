/**
 * Copyright (C) 2016 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <libudev.h>

#include "backlight.h"

struct udev* udev = NULL;
struct udev_enumerate* enumerate = NULL;

// key range: brightness, max_brightness
static int get_brightness_by_key(char* syspath, char* key);

int
init_udev()
{
    if (udev || enumerate) {
        return 0;
    }

    udev = udev_new();
    enumerate = udev_enumerate_new(udev);
    if (!enumerate) {
        finalize_udev();
        fprintf(stderr, "Get enumerate failed!\n");
        return -1;
    }

    int ret = udev_enumerate_add_match_subsystem(enumerate, "backlight");
    if (ret != 0) {
        finalize_udev();
        fprintf(stderr, "Enumerate match backlight failed!\n");
        return -1;
    }

    return 0;
}

void
finalize_udev()
{
    if (enumerate) {
        udev_enumerate_unref(enumerate);
        enumerate = NULL;
    }

    if (udev) {
        udev_unref(udev);
        udev = NULL;
    }
}

char**
get_syspath_list(int* num)
{
    *num = 0;
    int ret = udev_enumerate_scan_devices(enumerate);
    if (ret != 0) {
        fprintf(stderr, "Enumerate scan device failed!\n");
        return NULL;
    }

    struct udev_list_entry* entries = udev_enumerate_get_list_entry(enumerate);
    if (!entries) {
        fprintf(stderr, "Enumerate list entry failed!\n");
        return NULL;
    }

    char** list = NULL;
    struct udev_list_entry* current = NULL;
    udev_list_entry_foreach(current, entries) {
        const char* name = udev_list_entry_get_name(current);
        char* tmp = strdup(name);
        if (tmp) {
            list = (char**)realloc(list, (*num+1)*sizeof(char*));
            if (!list) {
                fprintf(stderr, "Realloc failed: %d\n", *num);
                continue;
            }
            list[*num] = tmp;
            *num += 1;
        }
    }

    return list;
}

void
free_syspath_list(char** list, int num)
{
    if (!list) {
        return ;
    }

    int i = 0;
    for (; i < num; i++) {
        free(list[i]);
    }
    free(list);
    list = NULL;
}

char*
get_syspath_by_type(char* type)
{
    int ret = udev_enumerate_scan_devices(enumerate);
    if (ret != 0) {
        fprintf(stderr, "Enumerate scan device failed!\n");
        return NULL;
    }

    struct udev_list_entry* entries = udev_enumerate_get_list_entry(enumerate);
    if (!entries) {
        fprintf(stderr, "Enumerate list entry failed!\n");
        return NULL;
    }

    struct udev_list_entry* current = NULL;
    udev_list_entry_foreach(current, entries) {
        const char* name = udev_list_entry_get_name(current);
        struct udev_device* device = udev_device_new_from_syspath(udev, name);
        const char* ty = udev_device_get_sysattr_value(device, "type");
        udev_device_unref(device);
        if (strcmp(ty, type) == 0) {
            return strdup(name);
        }
    }

    return NULL;
}

int
get_brightness(char* syspath)
{
    return get_brightness_by_key(syspath, "brightness");
}

int
get_max_brightness(char* syspath)
{
    return get_brightness_by_key(syspath, "max_brightness");
}

int
set_brightness(char* syspath, int value)
{
    struct udev_device* device = udev_device_new_from_syspath(udev, syspath);
    if (!device) {
        fprintf(stderr, "Invalid device: %s\n", syspath);
        return -1;
    }

	const char* str_max = udev_device_get_sysattr_value(device, "max_brightness");
	int max = atoi(str_max);
	if (value < 0) {
		value  = 0;
	} else if (value > max) {
		value = max;
	}

    char str[1000] = {0};
    sprintf(str, "%d", value);
    int ret = udev_device_set_sysattr_value(device, "brightness", str);
    udev_device_unref(device);
    if (ret != 0) {
        fprintf(stderr, "Set brightness for '%s' to %s failed!\n",
                syspath, str);
    }
    return ret;
}

static int
get_brightness_by_key(char* syspath, char* key)
{
    struct udev_device* device = udev_device_new_from_syspath(udev, syspath);
    if (!device) {
        fprintf(stderr, "Invalid device: %s\n", syspath);
        return -1;
    }

    const char* str = udev_device_get_sysattr_value(device, key);
    udev_device_unref(device);
    return atoi(str);
}
