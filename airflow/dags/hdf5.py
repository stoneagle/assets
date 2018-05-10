import pandas as pd
import sys
import json
import pprint
# import h5py
import re
import time
from pandas.util.testing import _network_error_classes
try:
    from urllib.request import urlopen, Request
except ImportError:
    from urllib2 import urlopen, Request

FOR_CLASSIFY_B_COLS = ['code', 'name']
THE_FIELDS = ['code', 'symbol', 'name', 'changepercent', 'trade', 'open', 'high', 'low', 'settlement', 'volume', 'turnoverratio']


def get_industry_classified():
    "由于tushare的延时会导致ip被封，提取该方法至脚本中"
    url = 'http://vip.stock.finance.sina.com.cn/q/view/newSinaHy.php'
    df = _get_type_data(url)
    return df
    # data = []
    # _write_head()
    # for row in df.values:
    #     rowDf = _get_detail(row[0], retry_count=3, pause=1)
    #     rowDf['c_name'] = row[1]
    #     data.append(rowDf)
    # data = pd.concat(data, ignore_index=True)
    # return data


def get_concept_classified():
    "由于tushare的延时会导致ip被封，提取该方法至脚本中"
    _write_head()
    url = 'http://money.finance.sina.com.cn/q/view/newFLJK.php?param=class'
    df = _get_type_data(url)
    data = []
    for row in df.values:
        rowDf = _get_detail(row[0], retry_count=3, pause=1)
        if rowDf is not None:
            rowDf['c_name'] = row[1]
            data.append(rowDf)
    if len(data) > 0:
        data = pd.concat(data, ignore_index=True)
    return data


def _get_type_data(url):
    try:
        request = Request(url)
        data_str = urlopen(request, timeout=10).read()
        data_str = data_str.decode('GBK')
        data_str = data_str.split('=')[1]
        data_json = json.loads(data_str)
        df = pd.DataFrame(
            [[row.split(',')[0], row.split(',')[1]] for row in data_json.values()],
            columns=['tag', 'name']
        )
        return df
    except Exception as er:
        print(str(er))


def _get_detail(tag, retry_count=3, pause=0.5):
    url = 'http://vip.stock.finance.sina.com.cn/quotes_service/api/json_v2.php/Market_Center.getHQNodeData?page=%s&num=1000&sort=symbol&asc=1&node=%s&symbol=&_s_r_a=page'
    dfc = pd.DataFrame()
    p = 0
    num_limit = 100
    while(True):
        p = p + 1
        for _ in range(retry_count):
            time.sleep(pause)
            try:
                _write_console()
                request = Request(url % (p, tag))
                text = urlopen(request, timeout=10).read()
                text = text.decode('gbk')
            except _network_error_classes:
                pass
            else:
                break
        reg = re.compile(r'\,(.*?)\:')
        text = reg.sub(r',"\1":', text)
        text = text.replace('"{symbol', '{"symbol')
        text = text.replace('{symbol', '{"symbol"')
        jstr = json.dumps(text)
        js = json.loads(jstr)
        df = pd.DataFrame(pd.read_json(js, dtype={'code': object}), columns=THE_FIELDS)
        df = df[FOR_CLASSIFY_B_COLS]
        dfc = pd.concat([dfc, df])
        if df.shape[0] < num_limit:
            return dfc


def _write_head():
    DATA_GETTING_TIPS = '[Getting data:]'
    sys.stdout.write(DATA_GETTING_TIPS)
    sys.stdout.flush()


def _write_console():
    sys.stdout.write('#')
    sys.stdout.flush()


get_industry_classified()
