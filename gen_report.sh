#!/bin/bash

# 需要先安装lcov，打开./unittest/CMakeLists.txt 测试覆盖率的编译条件
# 将该脚本放置到dde-dock-unit_test二进制文件同级目录运行
lcov -c -i -d ./ -o init.info
./dde_dock_unit_test
lcov -c -d ./ -o cover.info
lcov -a init.info -a cover.info -o total.info
lcov --remove total.info '*/usr/include/*' '*/usr/lib/*' '*/usr/lib64/*' '*/usr/local/include/*' '*/usr/local/lib/*' '*/usr/local/lib64/*' '*/third/*' '*/tests/dde_dock_unit_test_autogen/*' '*/dde-dock/frame/dbus/*' '*/dde-dock/interfaces/*' '*/dde-dock/tests/*' -o final.info

# 生成报告
genhtml -o cover_report --legend --title "lcov"  --prefix=./ final.info

#打开报告
nohup x-www-browser ./cover_report/index.html &

exit 0
