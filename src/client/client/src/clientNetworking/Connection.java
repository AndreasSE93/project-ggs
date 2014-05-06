package clientNetworking;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.PrintWriter;
import java.net.Socket;

public class Connection {
	public static final String DEFAULT_HOST = "localhost";
	public static final int    DEFAULT_PORT = 8080;

	final String hostIP;
	final int portNumber;
	private Socket socket;
	public BufferedReader reader;
	public PrintWriter writer;

	public Connection(String host) throws NumberFormatException {
		int portSeparator = host.lastIndexOf(':');

		if (portSeparator == -1 && host.length() != 0) {
			this.hostIP = host;
		} else if (portSeparator > 0) {
			this.hostIP = host.substring(0, portSeparator);
		} else {
			this.hostIP = DEFAULT_HOST;
		}

		if (portSeparator != -1 && portSeparator < host.length() - 1) {
			this.portNumber = Integer.parseInt(host.substring(portSeparator + 1));
		} else {
			this.portNumber = DEFAULT_PORT;
		}
	}

	public Connection(String host, int portnum) {
		this.hostIP = host;
		this.portNumber = portnum;
	}
	
	public void initConnection() throws IOException {
		this.socket = new Socket(this.hostIP, this.portNumber);
		this.reader = new BufferedReader(new InputStreamReader(socket.getInputStream()));
		this.writer = new PrintWriter(socket.getOutputStream(), true);
	}

//	public static void main(String[] args) throws IOException {
//		System.out.print("Enter server:  ");
//		String host = new BufferedReader(new InputStreamReader(System.in)).readLine();
//		Connection conn = new Connection(host);
//		System.out.printf("Host=%s%nPort=%d", conn.hostIP, conn.portNumber);
//	}
}
