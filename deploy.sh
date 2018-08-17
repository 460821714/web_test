#! /bin/sh
kill -9 $(pgrep webserver)
cd ~/../myproject/web_test
git pull https://github.com/460821714/web_test.git
cd webserver
./webserver &
