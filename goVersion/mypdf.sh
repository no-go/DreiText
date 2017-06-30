#!/bin/bash

## it converts a pdf to a readable textfile (maybe)

if [ $# -lt 1 ]; then
	exit 1
fi
INFILE=$1
OUTFILE=out.txt
pdftotext "$INFILE" $OUTFILE
cat $OUTFILE | sed -e 's/^$/CASTLEOFAAAAA/g' | tr '\n' ' ' > ${OUTFILE}.2
cat ${OUTFILE}.2 | sed -e 's/CASTLEOFAAAAA/\n\n/g' > ${OUTFILE}
less ${OUTFILE}
rm -f ${OUTFILE}.2
