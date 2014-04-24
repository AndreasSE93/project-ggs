package clientNetworking;

import java.io.IOException;

public class NetManager {
	
	public Handler handler; 
	public Connection conn;
	
	public NetManager(Handler interactor) {
		this.handler = interactor;
	}
	
	public void connectToServer(String hostIP, int portNum) throws IOException{
		this.conn = new Connection(hostIP, portNum);
		try {
			conn.initConnection();
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
