package clientNetworking;

import java.io.IOException;

import packageManaging.Handler;

public class NetManager {
	
	public Handler handler; 
	private Connection conn;
	
	public NetManager(Handler interactor, Connection c) {
		this.handler = interactor;
		this.conn = c;
	}
	
	public void connectToServer() throws IOException{
		try {
			this.conn.initConnection();
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
		return message;
	}
	
	
}
