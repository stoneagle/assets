'''
tushare股票分类数据
'''
from django.http import HttpResponse
import tushare as ts
from web.library import vars as vs
from web.library import framework as fw
import json


# 按照sina财经对沪深股票进行的行业分类，返回所有股票所属行业的信息
def getIndustryClassified(request):
    result = ts.get_industry_classified()
    if result is None:
        ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
    else:
        result = result.to_json(orient='records')
        ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 返回股票概念的分类数据
def getConceptClassified(request):
    result = ts.get_concept_classified()
    if result is None:
        ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
    else:
        result = result.to_json(orient='records')
        ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 按地域对股票进行分类
def getAreaClassified(request):
    params = {}
    params['file_path'] = request.POST.get('file_path', '')
    if not params['file_path']:
        result = ts.get_area_classified()
    else:
        result = ts.get_area_classified(file_path=params['file_path'])

    if result is None:
        ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
    else:
        result = result.to_json(orient='records')
        ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 获取中小板股票数据，即查找所有002开头的股票
def getSMEClassified(request):
    params = {}
    params['file_path'] = request.POST.get('file_path', '')
    if not params['file_path']:
        result = ts.get_sme_classified()
    else:
        result = ts.get_sme_classified(file_path=params['file_path'])

    if result is None:
        ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
    else:
        result = result.to_json(orient='records')
        ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 获取创业板股票数据，即查找所有300开头的股票
def getGEMClassified(request):
    params = {}
    params['file_path'] = request.POST.get('file_path', '')
    if not params['file_path']:
        result = ts.get_gem_classified()
    else:
        result = ts.get_gem_classified(file_path=params['file_path'])

    if result is None:
        ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
    else:
        result = result.to_json(orient='records')
        ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 获取风险警示板股票数据，即查找所有st股票
def getSTClassified(request):
    params = {}
    params['file_path'] = request.POST.get('file_path', '')
    if not params['file_path']:
        result = ts.get_st_classified()
    else:
        result = ts.get_st_classified(file_path=params['file_path'])

    if result is None:
        ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
    else:
        result = result.to_json(orient='records')
        ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 获取沪深300当前成份股及所占权重
def getHS300s(request):
    result = ts.get_hs300s()
    if result is None:
        ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
    else:
        result = result.to_json(orient='records')
        ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 获取上证50成份股
def getSZ50s(request):
    result = ts.get_sz50s()
    if result is None:
        ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
    else:
        result = result.to_json(orient='records')
        ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 获取中证500成份股
def getZZ500s(request):
    result = ts.get_zz500s()
    if result is None:
        ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
    else:
        result = result.to_json(orient='records')
        ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 获取已经被终止上市的股票列表，数据从上交所获取，目前只有在上海证券交易所交易被终止的股票
def getTerminated(request):
    result = ts.get_terminated()
    if result is None:
        ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
    else:
        result = result.to_json(orient='records')
        ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 获取被暂停上市的股票列表，数据从上交所获取，目前只有在上海证券交易所交易被终止的股票
def getSuspended(request):
    result = ts.get_suspended()
    if result is None:
        ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
    else:
        result = result.to_json(orient='records')
        ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)
