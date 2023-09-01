import socket
import struct

def wake_on_lan(mac, ip, port):
    # 创建Magic Packet
    mac = mac.replace(':','')
    data = b'FF' * 6 + (mac * 16).encode()
    send_data = b''

    # 将16进制数据转换为bytes类型
    for i in range(0, len(data), 2):
        send_data += struct.pack(b'B', int(data[i: i + 2], 16))

    # 发送Magic Packet到目标设备
    sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    sock.setsockopt(socket.SOL_SOCKET, socket.SO_BROADCAST, 1)

    sock.sendto(send_data, (ip, port))
    sock.close()

# 使用示例
wake_on_lan('11:22:33:44:55:66', '255.255.255.255', 9)
