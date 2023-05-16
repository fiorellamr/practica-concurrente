import socket

class TCPClient50:
    def __init__(self, ip, message_received_callback):
        self.SERVERIP = ip
        self.SERVERPORT = 4444
        self.mMessageListener = message_received_callback
        self.mRun = False
        self.out = None
        self.incoming = None

    def send_message(self, message):
        if self.out is not None and not self.out.closed:
            self.out.write(message + "\n")
            self.out.flush()

    def stop_client(self):
        self.mRun = False

    def run(self):
        self.mRun = True
        try:
            print("TCP Client" + "C: Conectando...")
            with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as sock:
                sock.connect((self.SERVERIP, self.SERVERPORT))
                self.out = sock.makefile(mode='w')
                print("TCP Client" + "C: Sent.")
                print("TCP Client" + "C: Done.")
                self.incoming = sock.makefile(mode='r')
                while self.mRun:
                    server_message = self.incoming.readline()
                    if server_message and self.mMessageListener:
                        self.mMessageListener.message_received(server_message.rstrip("\n"))
        except Exception as e:
            print("TCP" + "C: Error", e)


class Cliente50:
    def __init__(self):
        self.mTcpClient = None
        self.sc = None

    def main(self):
        objcli = Cliente50()
        objcli.iniciar()

    def iniciar(self):
        def on_message_received(message):
            self.cliente_recibe(message)

        def tcp_client_thread():
            self.mTcpClient = TCPClient50("192.168.18.5", on_message_received)
            self.mTcpClient.run()

        # Iniciar el hilo del cliente TCP
        thread = threading.Thread(target=tcp_client_thread)
        thread.start()

        salir = "n"
        self.sc = input()
        print("Cliente bandera 01")
        while salir != "s":
            salir = input()
            self.cliente_envia(salir)
        print("Cliente bandera 02")

    def cliente_recibe(self, llego):
        print("CLINTE50 El mensaje::" + llego)

    def cliente_envia(self, envia):
        if self.mTcpClient is not None:
            self.mTcpClient.send_message(envia)


# Llamada al método main
if __name__ == "__main__":
    objcli = Cliente50()
    objcli.iniciar()