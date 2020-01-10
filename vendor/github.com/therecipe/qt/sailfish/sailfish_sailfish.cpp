// +build sailfish sailfish_emulator

#define protected public
#define private public

#include "sailfish_sailfish.h"
#include "_cgo_export.h"

#include <sailfishapp.h>

int SailfishApp_SailfishApp_Main(int argc, char* argv)
{
	static int argcs = argc;
	static char** argvs = static_cast<char**>(malloc(argcs * sizeof(char*)));

	QList<QByteArray> aList = QByteArray(argv).split('|');
	for (int i = 0; i < argcs; i++)
		argvs[i] = (new QByteArray(aList.at(i)))->data();

	return SailfishApp::main(argcs, argvs);
}

