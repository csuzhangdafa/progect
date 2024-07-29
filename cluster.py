from pcapfile import savefile
from pcapfile.protocols.linklayer import ethernet
from pcapfile.protocols.network import ip
from pcapfile.protocols.transport import tcp, udp
import os

# 创建输出文件夹（如果不存在）
def create_output_folder(folder):
    if not os.path.exists(folder):
        os.makedirs(folder)

# 根据前四个字节的不同将 pcap 包分类并写入不同的 pcap 文件
def classify_and_save_pcap(pcap_file, output_folder):
    with open(pcap_file, 'rb') as f:
        pcap = savefile.load_savefile(f, layers=4, verbose=True)

        for packet in pcap.packets:
            # 解析以太网帧
            eth_frame = ethernet.Ethernet(packet.raw())
            
            # 解析 IP 数据包
            if isinstance(eth_frame.payload, ip.IP):
                ip_packet = eth_frame.payload
                
                # 获取传输层数据
                if isinstance(ip_packet.data, tcp.TCP):
                    transport = ip_packet.data
                elif isinstance(ip_packet.data, udp.UDP):
                    transport = ip_packet.data
                else:
                    continue  # 不处理非 TCP 和 UDP 数据包

                # 获取应用层数据
                if hasattr(transport, 'data'):
                    app_layer_data = transport.data
                    first_four_bytes = app_layer_data[:4]

                    # 构建分类键
                    if len(first_four_bytes) >= 4:
                        classification_key = first_four_bytes.hex()

                        # 创建对应分类的输出文件
                        output_pcap_file = os.path.join(output_folder, f"{classification_key}.pcap")

                        # 写入数据包到对应的 pcap 文件
                        with open(output_pcap_file, 'ab') as fout:
                            fout.write(packet.raw())

# 使用示例
if __name__ == "__main__":
    input_pcap_file = 'your_input_pcap.pcap'  # 输入的 pcap 文件路径
    output_folder = 'classified_pcaps'  # 输出的分类后 pcap 文件夹路径

    create_output_folder(output_folder)
    classify_and_save_pcap(input_pcap_file, output_folder)