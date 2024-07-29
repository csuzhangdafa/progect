import pyshark
import json
from typing import Dict, List
import sys
import csv
import logging
import IPython
import os
import nest_asyncio

sys.setrecursionlimit(10000) # 解决python最大递归深度的问题，但不治本
nest_asyncio.apply()


logging.basicConfig(stream=sys.stdout, level=logging.DEBUG)


def dfs(dct: Dict, index: List = []):
    for k, v in dct.items():
        if k.endswith("_raw") and isinstance(v, list):
            fetch(v, index=index)
        elif k.endswith("_tree") or isinstance(v, dict):
            if isinstance(v, dict):
                dfs(v, index=index)
            elif isinstance(v, list):
                for vv in v:
                    dfs(vv, index=index)
            # else:
            #     assert False  # smb2_100会发生错误


def fetch(lst, index=[]):
    if isinstance(lst[0], list):
        for m in lst:
            fetch(m, index=index)
    else:
        if (
            lst[0] != "field length invalid!"
        ):  # 不处理"field length invalid!"字段，处理smb2_100.pcap中部分包（e.g., 第7个）遇到的问题，https://github.com/wireshark/wireshark/blob/master/epan/print.c
            index.append(int(lst[1]))  # 字段的开始下标
            index.append(int(lst[1]) + int(lst[2]))  # 字段的结束下标(开始下标+长度)

def analyze(cap):
    _ = 0
    for packet in cap:
        # 解析结果
        # print(getattr(packet, sys.argv[2])) # packet.xxx

        # json形式输出
        # js = json.dumps(packet.smb2._all_fields, indent=4, separators=(',', ':'))
        # print(js)

        ans = []
        layers_dict = {}
        _ += 1
        logging.debug(f"Processing packet@{_}")
        for layer in range(int(sys.argv[3]), len(packet.layers)):
            layers_dict = dict(
                **layers_dict, **packet.layers[layer]._all_fields
            )
        # IPython.embed()
        dfs(layers_dict, ans)  # packet.xxx._all_fields

        # 去重并排序
        res = list(set(ans))
        res.sort()

        # 处理，得到分割的index
        base = res[0]
        for i in range(len(res)):
            res[i] -= base

        pkt_layer = packet.get_raw_packet().hex()[base * 2 :]  # 一个字节对应2个数字
        if res[-1] != len(pkt_layer) // 2:
            if res[-1] > len(pkt_layer) // 2:
                raise Exception("字段下标超过数据包长度!")
            else:
                res.append(len(pkt_layer) // 2)

        pkt_layer_split = " ".join(
            pkt_layer[2 * res[i] : 2 * res[i + 1]] for i in range(len(res) - 1)
        )

        writer.writerow([pkt_layer])


if __name__ == "__main__":

    if len(sys.argv) != 5:
        print(f"python {sys.argv[0]} input_file protocol_name layer_num output_file")
        exit(0)

    csvfile = open(sys.argv[4], "w")
    writer = csv.writer(csvfile)
    writer.writerow(["Hexstream"])

    # 导入pcap
    file_path = sys.argv[1]
    if file_path.endswith(".pcap") or file_path.endswith(".pcapng"):
        cap = pyshark.FileCapture(sys.argv[1], display_filter=sys.argv[2], include_raw=True, use_json=True)
        # cap = pyshark.FileCapture('/root/boringssl.pcap', display_filter="ssl", include_raw=True, use_json=True)
        analyze(cap)
    else:
        files = os.listdir(sys.argv[1])
        files.sort()

        for file in files:
            cap = pyshark.FileCapture(
                os.path.join(sys.argv[1], file),
                display_filter=sys.argv[2],
                include_raw=True,
                use_json=True,
            )
            
            analyze(cap)