package clientLobby;

import packageManaging.LobbyHandler;
import graphicalReference.LobbyGUI;
import clientNetworking.Connection;

public class LobbyShell {
	
	LobbyHandler handler;
	LobbyGUI graphInterface;
	
	public LobbyShell() {
		this.handler = new LobbyHandler();
		this.graphInterface = new LobbyGUI();
	}
	
}
