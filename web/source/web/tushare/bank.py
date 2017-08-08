'''
tushare银行间同业拆放利率
'''
from django.http import HttpResponse
import tushare as ts
from web.library import vars as vs
from web.library import framework as fw
import json


# 获取银行间同业拆放利率数据，目前只提供2006年以来的数据
def getShiborData(request):
    params = {}
    if request.POST:
        params['year'] = request.POST['year']
    if (not params or not fw.checkParamsEmpty(params)):
        ret = fw.formatResponse(vs.ERRNO_PARAMS, {})
    else:
        result = ts.shibor_data(year=str(params['year']))
        if result is None:
            ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
        else:
            result = result.to_json(orient='records')
            ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 获取银行间报价数据，目前只提供2006年以来的数据
def getShiborQuoteData(request):
    params = {}
    if request.POST:
        params['year'] = request.POST['year']
    if (not params or not fw.checkParamsEmpty(params)):
        ret = fw.formatResponse(vs.ERRNO_PARAMS, {})
    else:
        result = ts.shibor_quote_data(year=str(params['year']))
        if result is None:
            ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
        else:
            result = result.to_json(orient='records')
            ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 获取Shibor均值数据，目前只提供2006年以来的数据
def getShiborMaData(request):
    params = {}
    if request.POST:
        params['year'] = request.POST['year']
    if (not params or not fw.checkParamsEmpty(params)):
        ret = fw.formatResponse(vs.ERRNO_PARAMS, {})
    else:
        result = ts.shibor_ma_data(year=str(params['year']))
        if result is None:
            ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
        else:
            result = result.to_json(orient='records')
            ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 获取贷款基础利率（LPR）数据，目前只提供2013年以来的数据
def getLprData(request):
    params = {}
    if request.POST:
        params['year'] = request.POST['year']
    if (not params or not fw.checkParamsEmpty(params)):
        ret = fw.formatResponse(vs.ERRNO_PARAMS, {})
    else:
        result = ts.lpr_data(year=str(params['year']))
        if result is None:
            ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
        else:
            result = result.to_json(orient='records')
            ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 获取贷款基础利率均值数据，目前只提供2013年以来的数据
def getLprMaData(request):
    params = {}
    if request.POST:
        params['year'] = request.POST['year']
    if (not params or not fw.checkParamsEmpty(params)):
        ret = fw.formatResponse(vs.ERRNO_PARAMS, {})
    else:
        result = ts.lpr_ma_data(year=str(params['year']))
        if result is None:
            ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
        else:
            result = result.to_json(orient='records')
            ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)
