package clientLobby;

import packageManaging.LobbyHandler;
import graphicalReference.LobbyGUI;
import clientNetworking.Connection;

public class LobbyShell {
	
	Connection conn;
	LobbyHandler handler;
	LobbyGUI graphInterface;
	
	public LobbyShell(Connection conn) {
		this.conn = conn;
		this.handler = new LobbyHandler();
		this.graphInterface = new LobbyGUI();
	}
	
}
