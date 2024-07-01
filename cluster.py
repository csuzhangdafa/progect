import csv
from argparse import ArgumentParser
import pandas as pd
from scapy.all import *
import os
from netzob.Import.PCAPImporter.PCAPImporter import PCAPImporter


def add_suffix_to_filename(filename, suffix):
    # 获取文件所在的目录和原文件名
    directory = os.path.dirname(filename)
    base = os.path.basename(filename)
    
    # 分割原文件名为文件名和后缀
    name, ext = os.path.splitext(base)
    
    # 创建新的文件名，把suffix添加到文件名的后面
    new_name = name + "_" + suffix + ext
    
    # 如果原文件名包含目录，把目录添加回新的文件名
    if directory:
        new_name = os.path.join(directory, new_name)

    if not new_name.endswith(".pcap"):  # input 为文件夹
        new_name = new_name + ".pcap"
        
    return new_name


def getOriginMessages(orig_messages_path):
    orig_packets = []
    if os.path.isfile(orig_messages_path):
        orig_packets = PCAPImporter.readFile(orig_messages_path).values()
    else:
        files = os.listdir(orig_messages_path)
        files.sort()

        for file in files:
            filepath_input = os.path.join(orig_messages_path, file)
            packet = PCAPImporter.readFile(filepath_input).values()
            if not orig_packets:
                orig_packets = packet
            else:
                orig_packets = orig_packets + packet
    
    return orig_packets

if __name__ == '__main__':
    parser = ArgumentParser(
        description='Re-Implementation of FieldHunter.')
    parser.add_argument('pcapfilename', help='Filename of the PCAP to load.')
    parser.add_argument('output', help='Filename of the output')
    parser.add_argument('-i', '--interactive', help='Open ipython prompt after finishing the analysis.',
                        action="store_true")
    parser.add_argument('-d', '--debug', help='Enable debug output.', action="store_true")
    parser.add_argument('-f', '--fuzz_messages_filepath', dest='fuzz_messages_filepath', default=None, help='increase in traffic')
    args = parser.parse_args()

    # 原始流量
    orig_messages = getOriginMessages(args.pcapfilename)

    with open(args.output, "w", newline="") as csvfile:
        writer = csv.writer(csvfile)

        # 写入 CSV 文件的表头
        writer.writerow(["Message", "Category"])

         # 遍历原始消息并写入 CSV 文件
        for message in orig_messages:
            # 提取消息的字节内容
            message_content = bytes(message.data)  # 假设 data 属性包含消息的字节数据
            # 将消息转换为十六进制字符串
            hex_message = message_content.hex()
            # 提取前四个字节作为分类
            category = hex_message[:4]  # 假设每个消息至少有四个字节
            # 写入 CSV 文件的一行
            writer.writerow([hex_message, category])

    print(f"Messages have been written to {args.output}")

     # 读取刚刚写入的 CSV 文件
    df = pd.read_csv(args.output)

    # 创建一个字典来存储每个类别的文件名
    category_files = {}

    # 根据前四个字节对数据进行分类
    categories = df['Category'].unique()
    category_counts = {category: df[df['Category'] == category].shape[0] for category in categories}

    # 打印每个类别的数量
    print(f"Number of messages in each category:")
    for category, count in category_counts.items():
        print(f"{category}: {count}")

     # 遍历每个类别
    for category in categories:
        # 选择特定类别的数据
        category_df = df[df['Category'] == category]
        # 生成每个类别的文件名
        category_file = f"category_{category}.csv"
        # 保存每个类别的数据到不同的 CSV 文件
        category_files[category] = category_file
        category_df.to_csv(category_file, index=False)
    # 打印每个类别的文件路径
    print(f"Data has been categorized and saved to separate CSV files:")
    for category, file in category_files.items():
        print(f"{category}: {file}")
    
    '''
    # 保存聚类结果到新的 CSV 文件
    clustered_output = args.output.replace('.csv', '_clustered.csv')
    df.to_csv(clustered_output, index=False)

    print(f"Clustered results have been written to {clustered_output}")

    '''


    '''
    # 根据前四个字节对数据进行分类
    categories = df['Category'].unique()
    for category in categories:
        category_df = df[df['Category'] == category]
        # 保存每个类别的数据到不同的 CSV 文件
        category_file = args.output.replace('.csv', f'_{category}.csv')
        category_df.to_csv(category_file, index=False)

    print(f"Data has been categorized based on the first four bytes.")
'''