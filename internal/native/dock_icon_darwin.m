#import <AppKit/AppKit.h>
#include "dock_icon_darwin.h"

void SetDockIcon(const void* data, int length) {
    NSData* d = [NSData dataWithBytes:data length:(NSUInteger)length];
    NSImage* img = [[NSImage alloc] initWithData:d];
    if (!img) return;
    // NSApp icon must be set on the main thread.
    dispatch_async(dispatch_get_main_queue(), ^{
        [NSApp setApplicationIconImage:img];
    });
}
