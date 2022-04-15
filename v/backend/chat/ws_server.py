import json
import time
from channels.generic.websocket import WebsocketConsumer
from chat.ssh import SSH
from chat.models import Host
from threading import Thread


class WebSSHService(WebsocketConsumer):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.chan = None
        self.ssh = None
        self.id = None

    def _init(self):
        self.send(bytes_data=b'\r\33[KConnecting ...\r')
        h = Host.objects.filter(id=self.id).first()
        if not h:
            msg = f'\r\n \x1B[1;3;31m 必须先请求 http://backend:8000/chat/p-host 接口\r\n  method: POST\r\n  data: ip_address,username,portpassword \x1B[0m'
            self.send(bytes_data=msg.encode())
            self.close()
            return
        time.sleep(0.2)
        try:
            self.ssh = SSH(h.ip_address,
                           username=h.username,
                           port=h.port,
                           password=h.pkey).get_client()
        except Exception as e:
            self.send(bytes_data=f'Exception: {e}\r\n'.encode())
            self.close()
            return
        self.chan = self.ssh.invoke_shell(term='xterm')
        self.chan.transport.set_keepalive(30)
        Thread(target=self.loop_read).start()

    def loop_read(self):
        while True:
            data = self.chan.recv(32 * 1024)
            if not data:
                self.close(3333)
                break
            self.send(bytes_data=data)

    def connect(self):
        self.id = self.scope['url_route']['kwargs']['id']
        self.accept()
        self._init()

    def receive(self, text_data=None, bytes_data=None):
        data = text_data or bytes_data
        if data and self.chan:
            data = json.loads(data)
            resize = data.get('resize')
            if resize and len(resize) == 2:
                self.chan.resize_pty(*resize)
            else:
                self.chan.send(data['data'] + '\n')

    def disconnect(self, code):
        self.chan.close()
        self.ssh.close()