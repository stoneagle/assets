from web.library import vars as vs
import json


def formatResponse(code, data):
    ret = {}
    ret["errno"] = code
    ret["errmsg"] = vs.getMsg(code)
    ret["data"] = data
    ret = json.dumps(ret)
    return ret


def checkParamsEmpty(params):
    ret = True
    for index, key in enumerate(params):
        if not params[key]:
            ret = False
            break
    return ret
