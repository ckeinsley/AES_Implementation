#!/bin/bash
START=$(date +%s)

./aes_encrypt

END=$(date +%s)
DIFF=$(( $END - $START ))
echo "It took $DIFF seconds"