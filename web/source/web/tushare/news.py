'''
tushare新闻事件数据
'''
from django.http import HttpResponse
import tushare as ts
from web.library import vars as vs
from web.library import framework as fw
import json


# 获取即时财经新闻，类型包括国内财经、证券、外汇、期货、港股和美股等新闻信息
def getLatestNews(request):
    params = {}
    if request.POST:
        params['top'] = request.POST['top']
        params['content'] = request.POST.get('content', False)
    if (not params or not fw.checkParamsEmpty(params)):
        ret = fw.formatResponse(vs.ERRNO_PARAMS, {})
    else:
        result = ts.get_latest_news(top=int(params['top']), show_content=bool(params['content']))
        if result is None:
            ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
        else:
            result = result.to_json(orient='records')
            ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 获取个股信息地雷数据
def getNotices(request):
    params = {}
    if request.POST:
        params['code'] = request.POST['code']
        params['date'] = request.POST.get('date')
    if (not params or not fw.checkParamsEmpty(params)):
        ret = fw.formatResponse(vs.ERRNO_PARAMS, {})
    else:
        # result = ts.get_notices(code=params['code'], date=params['date'])
        result = ts.get_notices(code=params['code'])
        if result is None:
            ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
        else:
            result = result.to_json(orient='records')
            ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 获取sina财经股吧首页的重点消息
def getGubaSina(request):
    params = {}
    if request.POST:
        params['content'] = request.POST.get('content', False)
    if (not params or not fw.checkParamsEmpty(params)):
        ret = fw.formatResponse(vs.ERRNO_PARAMS, {})
    else:
        result = ts.guba_sina(show_content=bool(params['content']))
        if result is None:
            ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
        else:
            result = result.to_json(orient='records')
            ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)
