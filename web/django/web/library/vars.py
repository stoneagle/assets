# -*- coding:utf-8 -*-

ERRNO_OK = "0000"
ERRNO_PARAMS = "0001"
ERRNO_TUSHARE = "0002"

ERRMSG = {
    ERRNO_OK: "请求正常",
    ERRNO_PARAMS: "参数错误",
    ERRNO_TUSHARE: "tushare错误",
}


def getMsg(errno):
    if not ERRMSG[errno]:
        ret = "无对应信息"
    else:
        ret = ERRMSG[errno]
    return ret
