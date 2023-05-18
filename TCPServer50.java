package sumaNred;


import java.io.BufferedReader;
import java.io.PrintWriter;
import java.net.ServerSocket;
import java.net.Socket;


public class TCPServer50 {
    private String message;
    
    int nrcli = 0;

    public static final int SERVERPORT = 4444;
    private OnMessageReceived messageListener = null;
    private boolean running = false;
    TCPServerThread50[] sendclis = new TCPServerThread50[10];

    PrintWriter mOut;
    BufferedReader in;
    
    ServerSocket serverSocket;

    //el constructor pide una interface OnMessageReceived
    public TCPServer50(OnMessageReceived messageListener) {
        this.messageListener = messageListener;
    }
    
    public OnMessageReceived getMessageListener(){/////¨
        return this.messageListener;
    }
    
    public void sendMessageTCPServerRango(String message, int a, int b, int n){    
        int Rango = b -a;
        int d = (int) ((Rango) / nrcli);
        //n = n/nrcli;
        for (int i = 1; i < nrcli; i++) {
            sendclis[i].sendMessage("evalua " + message + " " + (((i-1)*d+a)) + " " + ((i-1)*d+d+a)+" "+ n);
            System.out.println("ENVIANDO A ESCLAVO " + (i));
            //System.out.println("evalua " + message + " " + (i-1)*d+a + " " + (i-1) * d + d+a);
        }

        sendclis[nrcli].sendMessage("evalua " + message +" "+ ((d * (nrcli - 1))+a) + " " + b+" "+ n);
        System.out.println("ENVIANDO A ESCLAVO " + (nrcli));
        //System.out.println("evalua " + message +" "+ ((d * (nrcli - 1))+a) + " " + b);
    }    
    public void run(){
        running = true;
        try{
            System.out.println("TCP Server"+"S : Connecting...");
            serverSocket = new ServerSocket(SERVERPORT);
            
            while(running){
                Socket client = serverSocket.accept();
                System.out.println("TCP Server"+"S: Receiving...");
                nrcli++;
                System.out.println("Engendrado " + nrcli);
                sendclis[nrcli] = new TCPServerThread50(client,this,nrcli,sendclis);
                Thread t = new Thread(sendclis[nrcli]);
                t.start();
                System.out.println("Nuevo conectado:"+ nrcli+" jugadores conectados");
                
            }
            
        }catch( Exception e){
            System.out.println("Error"+e.getMessage());
        }finally{

        }
    }
    public  TCPServerThread50[] getClients(){
        return sendclis;
    } 

    public  interface OnMessageReceived {
        public void messageReceived(String message);
    }
}
