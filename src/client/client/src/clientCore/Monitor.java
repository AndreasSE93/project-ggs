package clientCore;

import java.io.IOException;

import clientNetworking.NetManager;
import clientNetworking.Connection;
import clientChat.ChatShell;

public class Monitor {
	
	Connection conn;
	NetManager net;
	ChatShell chatModule;
	
	public Monitor() {
		this.conn = new Connection("127.0.0.1", 8080);
		this.chatModule = new ChatShell(this.conn);
		this.net = new NetManager(this.chatModule.handler, this.conn);
		
		try {
			this.net.connectToServer();
		} catch (IOException e) {
			System.out.println("Connection to server failed!" + e);
		}
	}
	
	
	
}
