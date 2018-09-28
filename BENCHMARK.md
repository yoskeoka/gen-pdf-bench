# Benchmark

```sh
bash ./init.sh
bash ./bench.sh
```

## Result

```txt
node version v8.11.2
puppeteer bench #1

real    0m3.354s
user    0m0.503s
sys     0m0.312s
puppeteer bench #2

real    0m1.636s
user    0m0.514s
sys     0m0.180s
puppeteer bench #3

real    0m1.483s
user    0m0.509s
sys     0m0.175s
puppeteer bench #4

real    0m1.505s
user    0m0.512s
sys     0m0.184s
puppeteer bench #5

real    0m1.472s
user    0m0.506s
sys     0m0.173s
puppeteer bench #6

real    0m1.587s
user    0m0.754s
sys     0m0.324s
puppeteer bench #7

real    0m1.483s
user    0m0.486s
sys     0m0.172s
puppeteer bench #8

real    0m1.687s
user    0m0.542s
sys     0m0.192s
puppeteer bench #9

real    0m1.493s
user    0m0.514s
sys     0m0.173s
puppeteer bench #10

real    0m1.430s
user    0m0.490s
sys     0m0.164s
wkhtmltopdf bench #1

real    0m3.500s
user    0m1.562s
sys     0m0.327s
wkhtmltopdf bench #2

real    0m2.921s
user    0m1.502s
sys     0m0.261s
wkhtmltopdf bench #3

real    0m3.053s
user    0m1.562s
sys     0m0.273s
wkhtmltopdf bench #4

real    0m2.844s
user    0m1.452s
sys     0m0.255s
wkhtmltopdf bench #5

real    0m2.951s
user    0m1.503s
sys     0m0.269s
wkhtmltopdf bench #6

real    0m3.046s
user    0m1.558s
sys     0m0.282s
wkhtmltopdf bench #7

real    0m2.955s
user    0m1.497s
sys     0m0.267s
wkhtmltopdf bench #8

real    0m3.611s
user    0m1.717s
sys     0m0.310s
wkhtmltopdf bench #9

real    0m2.945s
user    0m1.500s
sys     0m0.273s
wkhtmltopdf bench #10

real    0m3.065s
user    0m1.583s
sys     0m0.278s
```

### Memory Usage

```txt
PID    COMMAND      %CPU TIME     #TH   #WQ  #PORT MEM    PURG   CMPRS  PGRP  PPID  STATE    BOOSTS          %CPU_ME %CPU_OTHRS UID  FAULTS   COW    MSGSENT   MSGRECV   SYSBSD    SYSMACH   CSW
12881  VTDecoderXPC 0.0  00:00.04 4     3    57+   3072K+ 0B     0B     12881 1     sleeping *0[6+]          0.00000 0.00000    501  2485+    166+   3976+     1614+     833+      2717+     134+
12880  Chromium Hel 0.0  00:00.39 16/3  3    120+  32M+   4096B+ 0B     12877 12877 running  *0[4+]          0.00000 0.00000    501  20144+   1804+  2492+     866+      7891+     2767+     1852+
12879  Chromium Hel 0.0  00:00.22 10    3    90+   24M+   0B     0B     12877 12877 sleeping *0[1+]          0.00000 0.00000    501  19088+   1839+  9133+     3715+     1755+     6225+     854+
12878  Chromium Hel 0.0  00:00.14 18    6    111+  13M+   4096B+ 0B     12877 12877 sleeping *0[3+]          0.00000 0.00000    501  10796+   1784+  1765+     758+      2724+     1564+     972+
12877  Chromium     0.0  00:00.32 25    4    270+  48M+   0B     0B     12877 12876 sleeping *0[3+]          0.00000 0.00000    501  18593+   1903+  1543+     451+      6561+     6067+     2041+
12876  node         9.4  00:00.25 10    0    28+   16M+   0B     0B     12778 12780 sleeping *0[1]           0.00000 0.00000    501  9053+    166+   48+       23+       2371+     124+      845+

PID    COMMAND      %CPU TIME     #TH   #WQ  #PORT MEM    PURG   CMPRS  PGRP  PPID  STATE    BOOSTS           %CPU_ME %CPU_OTHRS UID  FAULTS   COW     MSGSENT   MSGRECV   SYSBSD    SYSMACH   CSW
15500  VTDecoderXPC 0.0  00:00.06 4     3    57+   3132K+ 0B     0B     15500 1     sleeping *0[6+]           0.00000 0.00000    501  2502+    166+    3976+     1614+     821+      2718+     569+
15499  Chromium Hel 52.1 00:00.77 19/1  5    127+  52M+   4096B  0B     15496 15496 running  *0[3]            5.16091 0.00000    501  29786+   1804+   2813+     958+      10520+    4223+     2838+
15498  Chromium Hel 4.9  00:00.21 10    3    91+   22M+   0B     0B     15496 15496 sleeping *0[1]            3.23331 0.00000    501  18356+   1839+   9135+     3714+     1727+     6228+     664+
15497  Chromium Hel 0.0  00:00.12 17    5    109   13M    4096B  0B     15496 15496 sleeping *0[3]            0.00000 0.00000    501  10788    1784    1752      753       2472+     1597      527+
15496  Chromium     15.2 00:00.41 28    5    285+  50M+   0B     0B     15496 15495 sleeping *0[3]            0.50594 0.00000    501  20363+   1903+   1923+     575+      9156+     8994+     2916+
15495  node         1.7  00:00.27 10    0    28    16M+   0B     0B     15411 15413 sleeping *0[1]            0.00000 0.00000    501  9244+    169     48        23        2471+     124       1094+
```

```txt
PID    COMMAND      %CPU TIME     #TH   #WQ  #PORT MEM    PURG   CMPRS  PGRP  PPID  STATE    BOOSTS           %CPU_ME %CPU_OTHRS UID  FAULTS   COW     MSGSENT   MSGRECV   SYSBSD    SYSMACH   CSW
13912  wkhtmltopdf  76.9 00:01.28 4     3    84    25M+   0B     0B     13677 13911 sleeping *0[2]            19.3168 0.00000    501  14317+   1949+   6961+     255+      5487+     9556+     7537+
13911  testwk       0.0  00:00.00 5     0    14    904K   0B     0B     13677 13677 sleeping *0[1]            0.00000 0.00000    501  773      93      23        11        335       56        111

PID    COMMAND      %CPU TIME     #TH   #WQ  #PORT MEM    PURG   CMPRS  PGRP  PPID  STATE    BOOSTS           %CPU_ME %CPU_OTHRS UID  FAULTS   COW     MSGSENT   MSGRECV   SYSBSD    SYSMACH   CSW
13904  wkhtmltopdf  43.3 00:01.61 6     2    97+   34M+   0B     0B     13677 13903 sleeping *0[2]            19.3271 0.00000    501  19308+   2018+   7978+     322+      8184+     10953+    10153+
13903  testwk       0.2  00:00.00 5     0    14    852K+  0B     0B     13677 13677 sleeping *0[1]            0.00000 0.00000    501  765+     97      23        11        768+      57        228+
```

## Hardware Overview

* Model Name: MacBook Pro
* Model Identifier: MacBookPro13,1
* Processor Name: Intel Core i5
* Processor Speed: 2 GHz
* Number of Processors: 1
* Total Number of Cores: 2
* L2 Cache (per Core): 256 KB
* L3 Cache: 4 MB
* Memory: 8 GB
