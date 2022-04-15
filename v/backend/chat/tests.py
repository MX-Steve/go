from django.test import TestCase

# Create your tests here.
import paramiko


class WebSsh(object):
    def client_ssh(self):
        sh = paramiko.SSHClient()  # 1 创建SSH对象
        sh.set_missing_host_key_policy(
            paramiko.AutoAddPolicy())  # 2 允许连接不在know_hosts文件中的主机
        sh.connect("121.43.41.139",
                   username="root",
                   port="53742",
                   password="Cmdb@2021!")  # 3 连接服务器
        stdin, stdout, stderr = sh.exec_command('ls /opt')
        right_info = stdout.read()
        err_info = stderr.read()

        if right_info:
            print(right_info.decode("utf-8"))
        elif err_info:
            print(err_info.decode("utf-8"))
        else:
            print("命令执行成功")


if __name__ == '__main__':
    a = WebSsh()
    a.client_ssh()