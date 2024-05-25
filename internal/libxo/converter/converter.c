#include "converter.h"
#include <libxo/xo.h>

void initialize_xo() {
    xo_set_flags(NULL, XOF_PRETTY);  // Enable pretty printing
}

void finalize_xo() {
    xo_finish();
}

void convert_to_xml(const char *data) {
    xo_open_container("root");

    // Print the input string as XML
    xo_emit("%s\n", data);
    
    xo_close_container("root");
}

