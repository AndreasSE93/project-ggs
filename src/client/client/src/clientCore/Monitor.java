package clientCore;

import java.io.IOException;

import clientLobby.LobbyShell;
import clientNetworking.NetManager;
import clientNetworking.Connection;
import clientChat.ChatShell;

public class Monitor {
	
	Connection conn;
	NetManager net;
	ChatShell chatModule;
	LobbyShell lobbyModule;
	
	public Monitor() {
		this.conn = new Connection("127.0.0.1", 8080);
		this.chatModule = new ChatShell();
		this.lobbyModule = new LobbyShell();
		this.net = new NetManager(this.chatModule.handler, this.conn);
		
		try {
			this.net.connectToServer();
		} catch (IOException e) {
			System.out.println("Connection to server failed!\n");
			e.printStackTrace();
		}
		this.lobbyModule.Initiate();
		this.runner();
	}
	
	private void runner() {
		try {
			String firstcall = net.receiveMessage();
		} catch (IOException e) {
			e.printStackTrace();
		}
		
		
		
	}
	
	
	
}
