package clientNetworking;

import java.io.IOException;



public class NetManager {
	
	private Connection conn;
	
	public NetManager(Connection c) {
			this.conn = c;
	}
	
	public void connectToServer() throws IOException{
		try {
			this.conn.initConnection();
			String pingString = this.receiveMessage();
			this.send(pingString);
		} catch (IOException e) {
			throw(e);
		}
		
	}
	
	public void send(String message) {
		this.conn.writer.println(message);
	}
	
	public String receiveMessage () throws IOException {
		
		String nextLine = "";
		String message = "";
		int c;
		
		while((c = this.conn.reader.read()) != '\n') {
			if(c != 0){
				nextLine = nextLine + (char)c;
			}
		}
		
		message = nextLine;
		System.out.println(message);
		return message;
	}
	
	
}
