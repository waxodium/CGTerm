#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include "lib/sout.h"

void printHelp() {
    sout("CGT shell, License GPL-3.0\n\n");

    sout("Standard Commands:\n");
    sout("  host\t\tManage host configuration\n");
    sout("  initscreen\tInitialize terminal screen\n");
    sout("  clear\t\tClear the terminal screen\n");
    sout("  whoami\tPrint current user info\n");
    sout("  lsa\t\tList all files\n");
    sout("  lsd\t\tList directories only\n");
    sout("  lsf\t\tList files only\n");
    sout("  cd\t\tChange directory\n");
    sout("  version\tPrint CGT version\n");

    sout("\nRepository at <https://github.com/MasterArd/CGTerm>\n");
}
