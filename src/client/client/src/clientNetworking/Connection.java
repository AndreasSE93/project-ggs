package clientNetworking;

import java.net.Socket;
import java.io.BufferedReader;
import java.io.PrintWriter;
import java.io.InputStreamReader;
import java.io.IOException;

public class Connection {

	final String hostIP;
	final int portNumber;
	private Socket socket;
	public BufferedReader reader;
	public PrintWriter writer;
	
	public Connection(String host, int portnum) {
		this.hostIP = host;
		this.portNumber = portnum;
	}
	
	public void initConnection() throws IOException {
		this.socket = new Socket(this.hostIP, this.portNumber);
		this.reader = new BufferedReader(new InputStreamReader(socket.getInputStream()));
		this.writer = new PrintWriter(socket.getOutputStream(), true);
	}
	
	
}
